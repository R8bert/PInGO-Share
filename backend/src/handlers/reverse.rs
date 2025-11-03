use actix_multipart::Multipart;
use actix_web::{error, web, Error, HttpRequest, HttpResponse};
use chrono::Utc;
use futures_util::StreamExt;
use sqlx::PgPool;
use std::io::Write;
use uuid::Uuid;

use crate::config::Config;
use crate::models::{CreateTokenRequest, ReverseShareToken, Settings, UploadResponse};
use crate::utils::{
    calculate_expiry_time, extract_user_id_from_request, is_validity_allowed,
    sanitize_filename_safe,
};

pub async fn create_token(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: HttpRequest,
    body: web::Json<CreateTokenRequest>,
) -> Result<HttpResponse, Error> {
    let user_id = extract_user_id_from_request(&req, &config)?;

    let token = Uuid::new_v4().to_string();
    let expires_at = body
        .expires_in
        .as_ref()
        .and_then(|v| calculate_expiry_time(v));

    let max_uses = if body.max_uses == 0 {
        -1
    } else {
        body.max_uses
    };

    let token_id: (i32,) = sqlx::query_as(
        r#"
        INSERT INTO reverse_share_tokens (user_id, token, name, max_uses, expires_at)
        VALUES ($1, $2, $3, $4, $5) RETURNING id
        "#,
    )
    .bind(user_id)
    .bind(&token)
    .bind(&body.name)
    .bind(max_uses)
    .bind(expires_at)
    .fetch_one(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    Ok(HttpResponse::Created().json(serde_json::json!({
        "id": token_id.0,
        "token": token,
        "name": body.name,
        "max_uses": max_uses,
        "used_count": 0,
        "expires_at": expires_at,
    })))
}

pub async fn get_tokens(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: HttpRequest,
) -> Result<HttpResponse, Error> {
    let user_id = extract_user_id_from_request(&req, &config)?;

    let tokens: Vec<ReverseShareToken> = sqlx::query_as(
        r#"
        SELECT id, user_id, token, name, used_count, max_uses, created_at, expires_at
        FROM reverse_share_tokens
        WHERE user_id = $1
        ORDER BY created_at DESC
        "#,
    )
    .bind(user_id)
    .fetch_all(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    Ok(HttpResponse::Ok().json(serde_json::json!({
        "tokens": tokens
    })))
}

pub async fn delete_token(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: HttpRequest,
    token_id: web::Path<i32>,
) -> Result<HttpResponse, Error> {
    let user_id = extract_user_id_from_request(&req, &config)?;

    sqlx::query("DELETE FROM reverse_share_tokens WHERE user_id = $1 AND id = $2")
        .bind(user_id)
        .bind(*token_id)
        .execute(pool.as_ref())
        .await
        .map_err(error::ErrorInternalServerError)?;

    Ok(HttpResponse::Ok().json(serde_json::json!({
        "message": "Token deleted successfully"
    })))
}

pub async fn reverse_upload(
    pool: web::Data<PgPool>,
    token: web::Path<String>,
    mut payload: Multipart,
) -> Result<HttpResponse, Error> {
    // Validate token
    let token_data: Option<(i32, i32, String, i32, i32, Option<chrono::DateTime<chrono::Utc>>)> = sqlx::query_as(
        r#"
        SELECT id, user_id, name, used_count, max_uses, expires_at
        FROM reverse_share_tokens
        WHERE token = $1
        "#
    )
    .bind(token.as_str())
    .fetch_optional(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    let (token_id, user_id, _name, used_count, max_uses, expires_at) =
        token_data.ok_or_else(|| error::ErrorNotFound("Invalid token"))?;

    // Check expiration
    if let Some(exp) = expires_at {
        if exp < Utc::now() {
            return Ok(HttpResponse::Forbidden().json(serde_json::json!({
                "error": "Token has expired"
            })));
        }
    }

    // Check max uses
    if max_uses != -1 && used_count >= max_uses {
        return Ok(HttpResponse::Forbidden().json(serde_json::json!({
            "error": "Token has reached maximum uses"
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
        100 * 1024 * 1024
    };

    let upload_id = Uuid::new_v4().to_string();
    let mut uploaded_files = Vec::new();
    let mut total_size: i64 = 0;
    let mut email = String::new();
    let mut validity = String::from("7days");

    std::fs::create_dir_all("./uploads").map_err(error::ErrorInternalServerError)?;

    // Process multipart
    while let Some(item) = payload.next().await {
        let mut field = item.map_err(error::ErrorBadRequest)?;
        let content_disposition = field.content_disposition();
        let field_name = content_disposition
            .as_ref()
            .and_then(|cd| cd.get_name())
            .unwrap_or("");

        match field_name {
            "files" => {
                let filename = content_disposition
                    .as_ref()
                    .and_then(|cd| cd.get_filename())
                    .ok_or_else(|| error::ErrorBadRequest("No filename"))?
                    .to_string();

                let sanitized = sanitize_filename_safe(&filename);
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
                total_size += file_size;

                if total_size > max_size {
                    return Ok(HttpResponse::BadRequest().json(serde_json::json!({
                        "error": format!("Total file size ({} bytes) exceeds maximum allowed size ({} bytes)", total_size, max_size)
                    })));
                }

                let file_path = format!("./uploads/{}_{}", upload_id, sanitized);
                let mut file = std::fs::File::create(&file_path)
                    .map_err(error::ErrorInternalServerError)?;
                file.write_all(&file_data)
                    .map_err(error::ErrorInternalServerError)?;

                uploaded_files.push(filename);
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
        validity = settings.max_validity.clone();
    }

    let expires_at = calculate_expiry_time(&validity);
    let files_json = serde_json::to_string(&uploaded_files)
        .map_err(error::ErrorInternalServerError)?;
    let download_url = format!("/download/{}", upload_id);

    let email_value = if email.is_empty() { None } else { Some(email) };

    // Save upload
    sqlx::query(
        r#"
        INSERT INTO uploads (user_id, upload_id, files, total_size, email, download_url, expires_at, is_available, is_reverse, reverse_token)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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
    .bind(true)
    .bind(token.as_str())
    .execute(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    // Update token usage
    sqlx::query("UPDATE reverse_share_tokens SET used_count = used_count + 1 WHERE id = $1")
        .bind(token_id)
        .execute(pool.as_ref())
        .await
        .map_err(|e| {
            log::error!("Failed to update token usage count: {}", e);
            error::ErrorInternalServerError("Failed to update token usage")
        })?;

    let files_count = uploaded_files.len();
    Ok(HttpResponse::Ok().json(UploadResponse {
        download_url,
        files: uploaded_files,
        count: files_count,
        expires_at,
    }))
}
