use actix_multipart::Multipart;
use actix_web::{error, web, Error, HttpRequest, HttpResponse};
use futures_util::StreamExt;
use sqlx::PgPool;
use std::io::Write;
use uuid::Uuid;

use crate::config::Config;
use crate::models::{
    AvailabilityRequest, ExpirationRequest, Settings, Upload, UploadResponse,
};
use crate::utils::{
    calculate_expiry_time, check_is_blocked, extract_user_id_from_request,
    is_allowed_file_type, is_validity_allowed, sanitize_filename_safe,
};

pub async fn upload(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: HttpRequest,
    mut payload: Multipart,
) -> Result<HttpResponse, Error> {
    let user_id = extract_user_id_from_request(&req, &config)?;

    // Check if user is blocked
    if check_is_blocked(user_id, &pool).await.unwrap_or(false) {
        return Ok(HttpResponse::Forbidden().json(serde_json::json!({
            "error": "Account blocked - uploads are not allowed"
        })));
    }

    // Get settings
    let settings: Settings = sqlx::query_as("SELECT * FROM settings ORDER BY id LIMIT 1")
        .fetch_optional(pool.as_ref())
        .await
        .map_err(error::ErrorInternalServerError)?
        .unwrap_or_default();

    let max_size = if settings.max_upload_size > 0 {
        settings.max_upload_size
    } else {
        100 * 1024 * 1024 // 100MB default
    };

    let upload_id = Uuid::new_v4().to_string();
    let mut uploaded_files = Vec::new();
    let mut total_size: i64 = 0;
    let mut email = String::new();
    let mut validity = String::from("7days");

    // Create uploads directory
    std::fs::create_dir_all("./uploads").map_err(error::ErrorInternalServerError)?;

    // Process multipart form
    while let Some(item) = payload.next().await {
        let mut field = item.map_err(error::ErrorBadRequest)?;
        
        let content_disposition = field.content_disposition();
        let field_name = content_disposition.get_name().unwrap_or("");

        match field_name {
            "files" => {
                let filename = content_disposition
                    .get_filename()
                    .ok_or_else(|| error::ErrorBadRequest("No filename"))?;

                // Validate file type
                if !is_allowed_file_type(filename) {
                    return Ok(HttpResponse::BadRequest().json(serde_json::json!({
                        "error": format!("File type not allowed: {}", filename)
                    })));
                }

                let sanitized = sanitize_filename_safe(filename);
                if sanitized.is_empty() {
                    return Ok(HttpResponse::BadRequest().json(serde_json::json!({
                        "error": "Invalid filename"
                    })));
                }

                let mut file_data = Vec::new();
                while let Some(chunk) = field.next().await {
                    let data = chunk.map_err(error::ErrorBadRequest)?;
                    file_data.extend_from_slice(&data);
                }

                let file_size = file_data.len() as i64;
                
                // Check file size
                if file_size > max_size {
                    return Ok(HttpResponse::BadRequest().json(serde_json::json!({
                        "error": format!("File {} is too large ({} bytes)", filename, file_size)
                    })));
                }

                total_size += file_size;

                // Check total size
                if total_size > max_size {
                    return Ok(HttpResponse::BadRequest().json(serde_json::json!({
                        "error": format!("Total file size ({} bytes) exceeds maximum allowed size ({} bytes)", total_size, max_size)
                    })));
                }

                // Save file
                let file_path = format!("./uploads/{}_{}", upload_id, sanitized);
                let mut file = std::fs::File::create(&file_path)
                    .map_err(error::ErrorInternalServerError)?;
                file.write_all(&file_data)
                    .map_err(error::ErrorInternalServerError)?;

                uploaded_files.push(filename.to_string());
            }
            "email" => {
                while let Some(chunk) = field.next().await {
                    let data = chunk.map_err(error::ErrorBadRequest)?;
                    email.push_str(&String::from_utf8_lossy(&data));
                }
            }
            "validity" => {
                while let Some(chunk) = field.next().await {
                    let data = chunk.map_err(error::ErrorBadRequest)?;
                    validity = String::from_utf8_lossy(&data).to_string();
                }
            }
            _ => {}
        }
    }

    if uploaded_files.is_empty() {
        return Ok(HttpResponse::BadRequest().json(serde_json::json!({
            "error": "No files uploaded"
        })));
    }

    // Validate validity
    if !is_validity_allowed(&validity, &settings.max_validity) {
        return Ok(HttpResponse::BadRequest().json(serde_json::json!({
            "error": format!("Requested validity '{}' exceeds maximum allowed '{}'", validity, settings.max_validity)
        })));
    }

    let expires_at = calculate_expiry_time(&validity);
    let files_json = serde_json::to_string(&uploaded_files)
        .map_err(error::ErrorInternalServerError)?;
    let download_url = format!("/download/{}", upload_id);

    let email_value = if email.is_empty() { None } else { Some(email) };

    // Save to database
    sqlx::query(
        r#"
        INSERT INTO uploads (user_id, upload_id, files, total_size, email, download_url, expires_at, is_available, is_reverse)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        "#
    )
    .bind(user_id)
    .bind(&upload_id)
    .bind(&files_json)
    .bind(total_size)
    .bind(email_value)
    .bind(&download_url)
    .bind(expires_at)
    .bind(true)
    .bind(false)
    .execute(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    Ok(HttpResponse::Ok().json(UploadResponse {
        download_url,
        files: uploaded_files,
        count: uploaded_files.len(),
        expires_at,
    }))
}

pub async fn get_uploads(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: HttpRequest,
) -> Result<HttpResponse, Error> {
    let user_id = extract_user_id_from_request(&req, &config)?;

    let uploads: Vec<Upload> = sqlx::query_as(
        r#"
        SELECT id, user_id, upload_id, files, total_size, email, download_url, created_at, expires_at,
               is_available, is_reverse, reverse_token, is_deleted, deleted_at, deletion_reason
        FROM uploads
        WHERE user_id = $1
        ORDER BY created_at DESC
        "#
    )
    .bind(user_id)
    .fetch_all(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    Ok(HttpResponse::Ok().json(serde_json::json!({
        "uploads": uploads
    })))
}

pub async fn delete_upload(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: HttpRequest,
    upload_id: web::Path<String>,
) -> Result<HttpResponse, Error> {
    let user_id = extract_user_id_from_request(&req, &config)?;

    // Get upload details
    let upload: Option<(String, String, i64, String, chrono::DateTime<chrono::Utc>, Option<chrono::DateTime<chrono::Utc>>, bool, Option<String>, bool)> = sqlx::query_as(
        r#"
        SELECT u.upload_id, u.files, u.total_size, u.download_url, u.created_at, u.expires_at, u.is_reverse, u.reverse_token, u.is_deleted
        FROM uploads u
        WHERE u.user_id = $1 AND u.upload_id = $2
        "#
    )
    .bind(user_id)
    .bind(upload_id.as_str())
    .fetch_optional(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    let (uid, files_str, total_size, download_url, created_at, expires_at, is_reverse, reverse_token, is_deleted) =
        upload.ok_or_else(|| error::ErrorNotFound("Upload not found"))?;

    if is_deleted {
        return Ok(HttpResponse::BadRequest().json(serde_json::json!({
            "error": "Upload is already deleted"
        })));
    }

    // Get user info
    let user_info: (String, String) = sqlx::query_as(
        "SELECT username, email FROM users WHERE id = $1"
    )
    .bind(user_id)
    .fetch_one(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    let (username, user_email) = user_info;

    // Parse files
    let files: Vec<String> = serde_json::from_str(&files_str)
        .map_err(error::ErrorInternalServerError)?;

    // Begin transaction
    let mut tx = pool.begin().await.map_err(error::ErrorInternalServerError)?;

    // Insert deletion log
    sqlx::query(
        r#"
        INSERT INTO deletion_logs (user_id, username, upload_id, files, total_size, email, download_url, uploaded_at, expires_at, is_reverse, reverse_token, deletion_reason)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, 'User deleted')
        "#
    )
    .bind(user_id)
    .bind(&username)
    .bind(&uid)
    .bind(&files_str)
    .bind(total_size)
    .bind(&user_email)
    .bind(&download_url)
    .bind(created_at)
    .bind(expires_at)
    .bind(is_reverse)
    .bind(&reverse_token)
    .execute(&mut *tx)
    .await
    .map_err(error::ErrorInternalServerError)?;

    // Soft delete
    sqlx::query(
        r#"
        UPDATE uploads
        SET is_deleted = TRUE, deleted_at = CURRENT_TIMESTAMP, deletion_reason = 'User deleted'
        WHERE user_id = $1 AND upload_id = $2
        "#
    )
    .bind(user_id)
    .bind(&uid)
    .execute(&mut *tx)
    .await
    .map_err(error::ErrorInternalServerError)?;

    tx.commit().await.map_err(error::ErrorInternalServerError)?;

    // Delete physical files
    for filename in files {
        let sanitized = sanitize_filename_safe(&filename);
        let file_path = format!("./uploads/{}_{}", uid, sanitized);
        if let Err(e) = std::fs::remove_file(&file_path) {
            log::warn!("Failed to delete file {}: {}", file_path, e);
        } else {
            log::info!("Successfully deleted file: {}", file_path);
        }
    }

    Ok(HttpResponse::Ok().json(serde_json::json!({
        "message": "Upload deleted successfully"
    })))
}

pub async fn toggle_availability(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: HttpRequest,
    upload_id: web::Path<String>,
    body: web::Json<AvailabilityRequest>,
) -> Result<HttpResponse, Error> {
    let user_id = extract_user_id_from_request(&req, &config)?;

    sqlx::query("UPDATE uploads SET is_available = $1 WHERE user_id = $2 AND upload_id = $3")
        .bind(body.is_available)
        .bind(user_id)
        .bind(upload_id.as_str())
        .execute(pool.as_ref())
        .await
        .map_err(error::ErrorInternalServerError)?;

    Ok(HttpResponse::Ok().json(serde_json::json!({
        "message": "Upload availability updated successfully"
    })))
}

pub async fn update_expiration(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: HttpRequest,
    upload_id: web::Path<String>,
    body: web::Json<ExpirationRequest>,
) -> Result<HttpResponse, Error> {
    let user_id = extract_user_id_from_request(&req, &config)?;

    // Check if user is blocked
    if check_is_blocked(user_id, &pool).await.unwrap_or(false) {
        return Ok(HttpResponse::Forbidden().json(serde_json::json!({
            "error": "Account blocked"
        })));
    }

    // Check ownership
    let owner: Option<(i32,)> = sqlx::query_as(
        "SELECT user_id FROM uploads WHERE upload_id = $1"
    )
    .bind(upload_id.as_str())
    .fetch_optional(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    let owner_id = owner
        .ok_or_else(|| error::ErrorNotFound("Upload not found"))?
        .0;

    if owner_id != user_id {
        return Ok(HttpResponse::Forbidden().json(serde_json::json!({
            "error": "You can only update your own uploads"
        })));
    }

    let expiry_time = calculate_expiry_time(&body.validity);

    sqlx::query("UPDATE uploads SET expires_at = $1 WHERE upload_id = $2")
        .bind(expiry_time)
        .bind(upload_id.as_str())
        .execute(pool.as_ref())
        .await
        .map_err(error::ErrorInternalServerError)?;

    Ok(HttpResponse::Ok().json(serde_json::json!({
        "message": "Upload expiration updated successfully"
    })))
}
