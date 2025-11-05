use actix_multipart::Multipart;
use actix_web::{cookie::Cookie, error, web, Error, HttpRequest, HttpResponse};
use futures_util::StreamExt;
use sha2::{Digest, Sha256};
use sqlx::PgPool;
use std::io::Write;

use crate::auth::{generate_jwt, hash_password, verify_password};
use crate::config::Config;
use crate::models::{AuthResponse, LoginRequest, RegisterRequest, Settings, User, UserInfo};
use crate::utils::extract_user_id_from_request;

pub async fn register(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: web::Json<RegisterRequest>,
) -> Result<HttpResponse, Error> {
    // Check if registration is allowed
    let settings: Option<Settings> = sqlx::query_as(
        "SELECT * FROM settings ORDER BY id LIMIT 1"
    )
    .fetch_optional(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    if let Some(s) = settings {
        if !s.allow_registration {
            return Ok(HttpResponse::Forbidden().json(serde_json::json!({
                "error": "User registration is currently disabled"
            })));
        }
    }

    // Check if user exists
    let existing: Option<(i32,)> = sqlx::query_as(
        "SELECT id FROM users WHERE email = $1 OR username = $2"
    )
    .bind(&req.email)
    .bind(&req.username)
    .fetch_optional(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    if existing.is_some() {
        return Ok(HttpResponse::Conflict().json(serde_json::json!({
            "error": "User already exists"
        })));
    }

    // Check if this is the first user
    let user_count: (i64,) = sqlx::query_as("SELECT COUNT(*) FROM users")
        .fetch_one(pool.as_ref())
        .await
        .map_err(error::ErrorInternalServerError)?;

    let is_first_user = user_count.0 == 0;

    // Hash password
    let password_hash = hash_password(&req.password)
        .map_err(|_| error::ErrorInternalServerError("Failed to hash password"))?;

    // Create user
    let user_id: (i32,) = sqlx::query_as(
        "INSERT INTO users (username, email, password_hash, is_admin) VALUES ($1, $2, $3, $4) RETURNING id"
    )
    .bind(&req.username)
    .bind(&req.email)
    .bind(&password_hash)
    .bind(is_first_user)
    .fetch_one(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    // Generate JWT
    let token = generate_jwt(user_id.0, &config.jwt_secret)
        .map_err(|_| error::ErrorInternalServerError("Failed to generate token"))?;

    let message = if is_first_user {
        "First user created successfully with admin privileges"
    } else {
        "User created successfully"
    };

    let cookie = Cookie::build("auth_token", &token)
        .path("/")
        .max_age(actix_web::cookie::time::Duration::days(1))
        .http_only(true)
        .finish();

    Ok(HttpResponse::Created()
        .cookie(cookie)
        .json(AuthResponse {
            message: message.to_string(),
            token,
            user: UserInfo {
                id: user_id.0,
                username: req.username.clone(),
                email: req.email.clone(),
                is_admin: is_first_user,
                avatar: None,
            },
        }))
}

pub async fn login(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: web::Json<LoginRequest>,
) -> Result<HttpResponse, Error> {
    // Get user by username or email
    let user: Option<User> = sqlx::query_as(
        "SELECT id, username, email, password_hash, is_admin, is_blocked, avatar, created_at FROM users WHERE email = $1 OR username = $1"
    )
    .bind(&req.username_or_email)
    .fetch_optional(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    let user = user.ok_or_else(|| {
        error::ErrorUnauthorized("Invalid credentials")
    })?;

    // Verify password
    if !verify_password(&req.password, &user.password_hash)
        .map_err(|_| error::ErrorInternalServerError("Password verification failed"))?
    {
        return Err(error::ErrorUnauthorized("Invalid credentials"));
    }

    // Generate JWT
    let token = generate_jwt(user.id, &config.jwt_secret)
        .map_err(|_| error::ErrorInternalServerError("Failed to generate token"))?;

    let cookie = Cookie::build("auth_token", &token)
        .path("/")
        .max_age(actix_web::cookie::time::Duration::days(1))
        .http_only(true)
        .finish();

    Ok(HttpResponse::Ok()
        .cookie(cookie)
        .json(AuthResponse {
            message: "Login successful".to_string(),
            token,
            user: UserInfo {
                id: user.id,
                username: user.username,
                email: user.email,
                is_admin: user.is_admin,
                avatar: user.avatar,
            },
        }))
}

pub async fn logout() -> Result<HttpResponse, Error> {
    let cookie = Cookie::build("auth_token", "")
        .path("/")
        .max_age(actix_web::cookie::time::Duration::seconds(-1))
        .http_only(true)
        .finish();

    Ok(HttpResponse::Ok()
        .cookie(cookie)
        .json(serde_json::json!({
            "message": "Logged out successfully"
        })))
}

pub async fn me(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: HttpRequest,
) -> Result<HttpResponse, Error> {
    let user_id = extract_user_id_from_request(&req, &config)?;

    let user: Option<User> = sqlx::query_as(
        "SELECT id, username, email, password_hash, is_admin, is_blocked, avatar, created_at FROM users WHERE id = $1"
    )
    .bind(user_id)
    .fetch_optional(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    let user = user.ok_or_else(|| error::ErrorNotFound("User not found"))?;

    Ok(HttpResponse::Ok().json(serde_json::json!({
        "user": {
            "id": user.id,
            "username": user.username,
            "email": user.email,
            "is_admin": user.is_admin,
            "avatar": user.avatar,
            "created_at": user.created_at,
        }
    })))
}

pub async fn upload_avatar(
    pool: web::Data<PgPool>,
    config: web::Data<Config>,
    req: HttpRequest,
    mut payload: Multipart,
) -> Result<HttpResponse, Error> {
    let user_id = extract_user_id_from_request(&req, &config)?;

    // Get user info
    let user: Option<User> = sqlx::query_as(
        "SELECT id, username, email, password_hash, is_admin, is_blocked, avatar, created_at FROM users WHERE id = $1"
    )
    .bind(user_id)
    .fetch_optional(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    let user = user.ok_or_else(|| error::ErrorNotFound("User not found"))?;

    // Create avatars directory
    std::fs::create_dir_all("./avatars").map_err(error::ErrorInternalServerError)?;

    let mut file_data = Vec::new();
    let mut content_type = String::new();

    // Process multipart
    while let Some(item) = payload.next().await {
        let mut field = item.map_err(error::ErrorBadRequest)?;
        content_type = field
            .content_type()
            .map(|ct| ct.to_string())
            .unwrap_or_default();

        while let Some(chunk) = field.next().await {
            let data = chunk.map_err(error::ErrorBadRequest)?;
            file_data.extend_from_slice(&data);
        }
    }

    // Validate content type
    let allowed_types = ["image/jpeg", "image/jpg", "image/png", "image/gif"];
    if !allowed_types.contains(&content_type.as_str()) {
        return Ok(HttpResponse::BadRequest().json(serde_json::json!({
            "error": "Invalid file type. Only PNG, JPG, JPEG, and GIF are allowed"
        })));
    }

    // Generate checksum
    let mut hasher = Sha256::new();
    hasher.update(&file_data);
    let hash = hasher.finalize();
    let hash_string = hex::encode(hash);

    // Determine extension
    let ext = match content_type.as_str() {
        "image/jpeg" | "image/jpg" => ".jpg",
        "image/png" => ".png",
        "image/gif" => ".gif",
        _ => ".jpg",
    };

    let filename = format!("{}${}{}", user.username, hash_string, ext);
    let avatar_path = format!("avatars/{}", filename);
    let full_path = format!("./{}", avatar_path);

    // Delete old avatar if exists
    if let Some(old_avatar) = &user.avatar {
        let old_path = format!(".{}", old_avatar);
        if std::path::Path::new(&old_path).exists() {
            std::fs::remove_file(&old_path).ok();
            log::info!("Deleted old avatar: {}", old_path);
        }
    }

    // Save new file
    let mut file = std::fs::File::create(&full_path)
        .map_err(error::ErrorInternalServerError)?;
    file.write_all(&file_data)
        .map_err(error::ErrorInternalServerError)?;

    // Update database
    sqlx::query("UPDATE users SET avatar = $1 WHERE id = $2")
        .bind(&avatar_path)
        .bind(user_id)
        .execute(pool.as_ref())
        .await
        .map_err(error::ErrorInternalServerError)?;

    Ok(HttpResponse::Ok().json(serde_json::json!({
        "message": "Avatar uploaded successfully",
        "avatar": avatar_path,
    })))
}
