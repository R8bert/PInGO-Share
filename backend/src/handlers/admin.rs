use actix_multipart::Multipart;
use actix_web::{error, web, Error, HttpRequest, HttpResponse};
use chrono::Utc;
use futures_util::StreamExt;
use sha2::{Digest, Sha256};
use sqlx::PgPool;
use std::io::Write;

use crate::config::Config;
use crate::models::{AdminStats, AdminUser, BlockUserRequest, QuickSettingRequest, Settings};
use crate::utils::{check_is_admin, extract_user_id_from_request, parse_size, sanitize_filename_safe};

pub async fn update_settings(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: HttpRequest,
    mut payload: Multipart,
) -> Result<HttpResponse, Error> {
    let user_id = extract_user_id_from_request(&req, &config)?;

    if !check_is_admin(user_id, &pool).await.unwrap_or(false) {
        return Ok(HttpResponse::Forbidden().json(serde_json::json!({
            "error": "Admin access required"
        })));
    }

    // Get username
    let username: (String,) = sqlx::query_as("SELECT username FROM users WHERE id = $1")
        .bind(user_id)
        .fetch_one(pool.as_ref())
        .await
        .map_err(error::ErrorInternalServerError)?;

    // Load existing settings
    let mut settings: Settings = sqlx::query_as("SELECT * FROM settings ORDER BY id LIMIT 1")
        .fetch_optional(pool.as_ref())
        .await
        .map_err(error::ErrorInternalServerError)?
        .unwrap_or_default();

    let mut form_data = std::collections::HashMap::new();
    let mut logo_data: Option<(Vec<u8>, String)> = None;
    let mut background_data: Option<(Vec<u8>, String)> = None;

    // Process multipart
    while let Some(item) = payload.next().await {
        let mut field = item.map_err(error::ErrorBadRequest)?;
        let content_disposition = field.content_disposition();
        let field_name = content_disposition.get_name().unwrap_or("").to_string();

        if field_name == "logo" {
            let mut data = Vec::new();
            while let Some(chunk) = field.next().await {
                let chunk_data = chunk.map_err(error::ErrorBadRequest)?;
                data.extend_from_slice(&chunk_data);
            }
            if !data.is_empty() {
                logo_data = Some((data, field_name));
            }
        } else if field_name == "backgroundImage" {
            let mut data = Vec::new();
            while let Some(chunk) = field.next().await {
                let chunk_data = chunk.map_err(error::ErrorBadRequest)?;
                data.extend_from_slice(&chunk_data);
            }
            if !data.is_empty() {
                background_data = Some((data, field_name));
            }
        } else {
            let mut value = String::new();
            while let Some(chunk) = field.next().await {
                let chunk_data = chunk.map_err(error::ErrorBadRequest)?;
                value.push_str(&String::from_utf8_lossy(&chunk_data));
            }
            form_data.insert(field_name, value);
        }
    }

    // Update settings from form data
    if let Some(theme) = form_data.get("theme") {
        if !theme.is_empty() {
            settings.theme = theme.clone();
        }
    }

    if let Some(navbar_title) = form_data.get("navbarTitle") {
        if !navbar_title.is_empty() {
            settings.navbar_title = navbar_title.clone();
        }
    }

    if let Some(max_validity) = form_data.get("maxValidity") {
        let valid_options = ["7days", "1month", "6months", "1year", "never"];
        if valid_options.contains(&max_validity.as_str()) {
            settings.max_validity = max_validity.clone();
        }
    }

    if let Some(max_size_str) = form_data.get("maxUploadSize") {
        if let Ok(size) = parse_size(max_size_str) {
            settings.max_upload_size = size;
        }
    }

    if let Some(transparency) = form_data.get("uploadBoxTransparency") {
        if let Ok(val) = transparency.parse::<i32>() {
            if (0..=100).contains(&val) {
                settings.upload_box_transparency = val;
            }
        }
    }

    if let Some(blur) = form_data.get("blurIntensity") {
        if let Ok(val) = blur.parse::<i32>() {
            if (0..=24).contains(&val) {
                settings.blur_intensity = val;
            }
        }
    }

    if let Some(allow_reg) = form_data.get("allowRegistration") {
        if let Ok(val) = allow_reg.parse::<bool>() {
            settings.allow_registration = val;
        }
    }

    if let Some(exp_action) = form_data.get("expirationAction") {
        if exp_action == "delete" || exp_action == "unavailable" {
            settings.expiration_action = exp_action.clone();
        }
    }

    // Handle color fields
    if let Some(color) = form_data.get("websiteColor") {
        if color.len() == 7 && color.starts_with('#') {
            settings.website_color = color.clone();
        }
    }

    if let Some(color) = form_data.get("gradientColor1") {
        if color.len() == 7 && color.starts_with('#') {
            settings.gradient_color_1 = color.clone();
        }
    }

    if let Some(color) = form_data.get("gradientColor2") {
        if color.len() == 7 && color.starts_with('#') {
            settings.gradient_color_2 = color.clone();
        }
    }

    if let Some(color) = form_data.get("gradientColor3") {
        if color.len() == 7 && color.starts_with('#') {
            settings.gradient_color_3 = color.clone();
        }
    }

    // Handle logo upload
    if let Some((data, _)) = logo_data {
        std::fs::create_dir_all("./logos").map_err(error::ErrorInternalServerError)?;

        // Delete old logo
        if let Some(old_logo) = &settings.logo_path {
            let old_path = format!(".{}", old_logo);
            std::fs::remove_file(&old_path).ok();
        }

        let mut hasher = Sha256::new();
        hasher.update(&data);
        let hash = hasher.finalize();
        let hash_string = hex::encode(hash);

        let ext = ".png";
        let filename = format!("{}${}{}", username.0, hash_string, ext);
        let logo_path = format!("logos/{}", filename);
        let full_path = format!("./{}", logo_path);

        let mut file = std::fs::File::create(&full_path)
            .map_err(error::ErrorInternalServerError)?;
        file.write_all(&data)
            .map_err(error::ErrorInternalServerError)?;

        settings.logo_path = Some(format!("/{}", logo_path));
    }

    // Handle background upload
    if let Some((data, _)) = background_data {
        std::fs::create_dir_all("./backgrounds").map_err(error::ErrorInternalServerError)?;

        // Delete old background
        if let Some(old_bg) = &settings.background_path {
            let old_path = format!(".{}", old_bg);
            std::fs::remove_file(&old_path).ok();
        }

        let mut hasher = Sha256::new();
        hasher.update(&data);
        let hash = hasher.finalize();
        let hash_string = hex::encode(hash);

        let ext = ".jpg";
        let filename = format!("{}{}", hash_string, ext);
        let bg_path = format!("backgrounds/{}", filename);
        let full_path = format!("./{}", bg_path);

        let mut file = std::fs::File::create(&full_path)
            .map_err(error::ErrorInternalServerError)?;
        file.write_all(&data)
            .map_err(error::ErrorInternalServerError)?;

        settings.background_path = Some(format!("/{}", bg_path));
    }

    // Save to database
    if settings.id == 0 {
        let id: (i32,) = sqlx::query_as(
            r#"
            INSERT INTO settings (theme, logo_path, background_path, navbar_title, max_upload_size, upload_box_transparency, blur_intensity, max_validity, allow_registration, expiration_action, website_color, gradient_color_1, gradient_color_2, gradient_color_3)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id
            "#
        )
        .bind(&settings.theme)
        .bind(&settings.logo_path)
        .bind(&settings.background_path)
        .bind(&settings.navbar_title)
        .bind(settings.max_upload_size)
        .bind(settings.upload_box_transparency)
        .bind(settings.blur_intensity)
        .bind(&settings.max_validity)
        .bind(settings.allow_registration)
        .bind(&settings.expiration_action)
        .bind(&settings.website_color)
        .bind(&settings.gradient_color_1)
        .bind(&settings.gradient_color_2)
        .bind(&settings.gradient_color_3)
        .fetch_one(pool.as_ref())
        .await
        .map_err(error::ErrorInternalServerError)?;
        
        settings.id = id.0;
    } else {
        sqlx::query(
            r#"
            UPDATE settings SET theme = $1, logo_path = $2, background_path = $3, navbar_title = $4, max_upload_size = $5,
            upload_box_transparency = $6, blur_intensity = $7, max_validity = $8, allow_registration = $9, expiration_action = $10,
            website_color = $11, gradient_color_1 = $12, gradient_color_2 = $13, gradient_color_3 = $14, updated_at = CURRENT_TIMESTAMP
            WHERE id = $15
            "#
        )
        .bind(&settings.theme)
        .bind(&settings.logo_path)
        .bind(&settings.background_path)
        .bind(&settings.navbar_title)
        .bind(settings.max_upload_size)
        .bind(settings.upload_box_transparency)
        .bind(settings.blur_intensity)
        .bind(&settings.max_validity)
        .bind(settings.allow_registration)
        .bind(&settings.expiration_action)
        .bind(&settings.website_color)
        .bind(&settings.gradient_color_1)
        .bind(&settings.gradient_color_2)
        .bind(&settings.gradient_color_3)
        .bind(settings.id)
        .execute(pool.as_ref())
        .await
        .map_err(error::ErrorInternalServerError)?;
    }

    Ok(HttpResponse::Ok().json(settings))
}

pub async fn get_stats(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: HttpRequest,
) -> Result<HttpResponse, Error> {
    let user_id = extract_user_id_from_request(&req, &config)?;

    if !check_is_admin(user_id, &pool).await.unwrap_or(false) {
        return Ok(HttpResponse::Forbidden().json(serde_json::json!({
            "error": "Admin access required"
        })));
    }

    let total_users: (i64,) = sqlx::query_as("SELECT COUNT(*) FROM users")
        .fetch_one(pool.as_ref())
        .await
        .map_err(error::ErrorInternalServerError)?;

    let total_uploads: (i64,) = sqlx::query_as("SELECT COUNT(*) FROM uploads")
        .fetch_one(pool.as_ref())
        .await
        .map_err(error::ErrorInternalServerError)?;

    let storage_used: (Option<i64>,) = sqlx::query_as(
        "SELECT COALESCE(SUM(total_size), 0) as total_size FROM uploads"
    )
    .fetch_one(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    Ok(HttpResponse::Ok().json(AdminStats {
        total_users: total_users.0,
        total_uploads: total_uploads.0,
        storage_used: storage_used.0.unwrap_or(0),
    }))
}

pub async fn get_users(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: HttpRequest,
) -> Result<HttpResponse, Error> {
    let user_id = extract_user_id_from_request(&req, &config)?;

    if !check_is_admin(user_id, &pool).await.unwrap_or(false) {
        return Ok(HttpResponse::Forbidden().json(serde_json::json!({
            "error": "Admin access required"
        })));
    }

    let users: Vec<AdminUser> = sqlx::query_as(
        r#"
        SELECT
            u.id, u.username, u.email, u.avatar, u.is_admin,
            COALESCE(u.is_blocked, false) as is_blocked,
            u.created_at,
            COUNT(CASE WHEN up.id IS NOT NULL THEN 1 END) as upload_count,
            COALESCE(SUM(up.total_size), 0) as storage_used,
            MAX(up.created_at) as last_activity
        FROM users u
        LEFT JOIN uploads up ON u.id = up.user_id
        GROUP BY u.id, u.username, u.email, u.avatar, u.is_admin, u.is_blocked, u.created_at
        ORDER BY u.created_at DESC
        "#
    )
    .fetch_all(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    Ok(HttpResponse::Ok().json(users))
}

pub async fn block_user(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: HttpRequest,
    target_user_id: web::Path<i32>,
    body: web::Json<BlockUserRequest>,
) -> Result<HttpResponse, Error> {
    let admin_id = extract_user_id_from_request(&req, &config)?;

    if !check_is_admin(admin_id, &pool).await.unwrap_or(false) {
        return Ok(HttpResponse::Forbidden().json(serde_json::json!({
            "error": "Admin access required"
        })));
    }

    // Check if target is admin
    let is_target_admin: (bool,) = sqlx::query_as(
        "SELECT COALESCE(is_admin, false) FROM users WHERE id = $1"
    )
    .bind(*target_user_id)
    .fetch_one(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    if is_target_admin.0 && body.blocked {
        return Ok(HttpResponse::Forbidden().json(serde_json::json!({
            "error": "Cannot block admin users"
        })));
    }

    sqlx::query("UPDATE users SET is_blocked = $1 WHERE id = $2")
        .bind(body.blocked)
        .bind(*target_user_id)
        .execute(pool.as_ref())
        .await
        .map_err(error::ErrorInternalServerError)?;

    let action = if body.blocked { "blocked" } else { "unblocked" };

    Ok(HttpResponse::Ok().json(serde_json::json!({
        "message": format!("User {} successfully", action)
    })))
}

pub async fn quick_settings(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: HttpRequest,
    body: web::Json<QuickSettingRequest>,
) -> Result<HttpResponse, Error> {
    let user_id = extract_user_id_from_request(&req, &config)?;

    if !check_is_admin(user_id, &pool).await.unwrap_or(false) {
        return Ok(HttpResponse::Forbidden().json(serde_json::json!({
            "error": "Admin access required"
        })));
    }

    match body.setting.as_str() {
        "allowRegistration" => {
            let value = body.value.as_bool().ok_or_else(|| {
                error::ErrorBadRequest("Invalid value for allowRegistration")
            })?;

            sqlx::query("UPDATE settings SET allow_registration = $1 WHERE id = (SELECT MIN(id) FROM settings)")
                .bind(value)
                .execute(pool.as_ref())
                .await
                .map_err(error::ErrorInternalServerError)?;

            Ok(HttpResponse::Ok().json(serde_json::json!({
                "message": "Setting updated successfully"
            })))
        }
        _ => Ok(HttpResponse::BadRequest().json(serde_json::json!({
            "error": "Unknown setting"
        }))),
    }
}

pub async fn promote_first_admin(pool: web::Data<PgPool>) -> Result<HttpResponse, Error> {
    sqlx::query("UPDATE users SET is_admin = TRUE WHERE id = (SELECT MIN(id) FROM users)")
        .execute(pool.as_ref())
        .await
        .map_err(error::ErrorInternalServerError)?;

    Ok(HttpResponse::Ok().json(serde_json::json!({
        "message": "First user promoted to admin"
    })))
}

pub async fn cleanup_expired_uploads(pool: PgPool) {
    let mut interval = tokio::time::interval(tokio::time::Duration::from_secs(3600)); // 1 hour

    loop {
        interval.tick().await;

        log::info!("Running expired uploads cleanup...");

        // Get settings
        let settings: Option<Settings> = sqlx::query_as("SELECT * FROM settings ORDER BY id LIMIT 1")
            .fetch_optional(&pool)
            .await
            .ok()
            .flatten();

        let expiration_action = settings
            .map(|s| s.expiration_action)
            .unwrap_or_else(|| "unavailable".to_string());

        // Find expired uploads
        let expired_uploads: Vec<(String, String)> = sqlx::query_as(
            "SELECT upload_id, files FROM uploads WHERE expires_at IS NOT NULL AND expires_at < NOW() AND is_deleted = FALSE"
        )
        .fetch_all(&pool)
        .await
        .unwrap_or_default();

        log::info!("Found {} expired uploads, action: {}", expired_uploads.len(), expiration_action);

        for (upload_id, files_json) in expired_uploads {
            if expiration_action == "delete" {
                // Parse files and delete
                if let Ok(files) = serde_json::from_str::<Vec<String>>(&files_json) {
                    for filename in files {
                        let sanitized = sanitize_filename_safe(&filename);
                        let file_path = format!("./uploads/{}_{}", upload_id, sanitized);
                        if let Err(e) = std::fs::remove_file(&file_path) {
                            log::warn!("Failed to delete file {}: {}", file_path, e);
                        } else {
                            log::info!("Deleted expired file: {}", file_path);
                        }
                    }
                }

                // Mark as deleted
                if let Err(e) = sqlx::query(
                    "UPDATE uploads SET is_deleted = TRUE, deleted_at = NOW(), deletion_reason = 'Expired' WHERE upload_id = $1"
                )
                .bind(&upload_id)
                .execute(&pool)
                .await
                {
                    log::error!("Failed to mark upload as deleted: {}", e);
                } else {
                    log::info!("Marked upload as deleted: {}", upload_id);
                }
            } else if expiration_action == "unavailable" {
                // Just mark as unavailable
                if let Err(e) = sqlx::query("UPDATE uploads SET is_available = FALSE WHERE upload_id = $1")
                    .bind(&upload_id)
                    .execute(&pool)
                    .await
                {
                    log::error!("Failed to mark upload as unavailable: {}", e);
                } else {
                    log::info!("Marked upload as unavailable: {}", upload_id);
                }
            }
        }
    }
}
