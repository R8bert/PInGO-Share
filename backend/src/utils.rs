use actix_web::{error::ErrorUnauthorized, Error, HttpMessage, HttpRequest};
use chrono::{DateTime, Utc};
use sanitize_filename::sanitize;

use crate::auth::{extract_token_from_header, validate_jwt};
use crate::config::Config;

pub fn sanitize_filename_safe(filename: &str) -> String {
    let sanitized = sanitize(filename);
    sanitized.replace("..", "").replace("/", "").replace("\\", "")
}

pub fn is_allowed_file_type(filename: &str) -> bool {
    let blocked_extensions = [
        ".exe", ".bat", ".cmd", ".com", ".scr", ".pif", ".msi", ".vbs", ".js", ".jar",
        ".sh", ".bash", ".ps1", ".psm1",
    ];

    let lowercase = filename.to_lowercase();
    !blocked_extensions.iter().any(|ext| lowercase.ends_with(ext))
}

pub fn calculate_expiry_time(validity: &str) -> Option<DateTime<Utc>> {
    let now = Utc::now();
    
    match validity {
        "1h" => Some(now + chrono::Duration::hours(1)),
        "1d" | "7days" => Some(now + chrono::Duration::days(7)),
        "7d" => Some(now + chrono::Duration::days(7)),
        "1month" | "30d" => Some(now + chrono::Duration::days(30)),
        "6months" => Some(now + chrono::Duration::days(180)),
        "1year" => Some(now + chrono::Duration::days(365)),
        "never" => None,
        _ => Some(now + chrono::Duration::days(7)), // Default to 7 days
    }
}

pub fn is_validity_allowed(requested: &str, max: &str) -> bool {
    let validity_order = [
        ("1h", 0),
        ("1d", 1),
        ("7d", 1),
        ("7days", 1),
        ("30d", 2),
        ("1month", 2),
        ("6months", 3),
        ("1year", 4),
        ("never", 5),
    ];

    let requested_level = validity_order
        .iter()
        .find(|(v, _)| v == &requested)
        .map(|(_, l)| l);
    let max_level = validity_order
        .iter()
        .find(|(v, _)| v == &max)
        .map(|(_, l)| l);

    match (requested_level, max_level) {
        (Some(r), Some(m)) => r <= m,
        _ => false,
    }
}

pub fn parse_size(size_str: &str) -> Result<i64, String> {
    let mut num_str = String::new();
    let mut unit_str = String::new();
    let mut parsing_num = true;

    for c in size_str.chars() {
        if c.is_numeric() && parsing_num {
            num_str.push(c);
        } else {
            parsing_num = false;
            unit_str.push(c);
        }
    }

    let num: i64 = num_str.parse().map_err(|_| "Invalid number")?;
    let unit = unit_str.to_uppercase();

    let multiplier = match unit.as_str() {
        "KB" => 1024,
        "MB" => 1024 * 1024,
        "GB" => 1024 * 1024 * 1024,
        "" => 1,
        _ => return Err(format!("Unsupported unit: {}", unit)),
    };

    Ok(num * multiplier)
}

pub fn extract_user_id_from_request(req: &HttpRequest, config: &Config) -> Result<i32, Error> {
    let token = extract_token_from_header(
        req.headers()
            .get(actix_web::http::header::AUTHORIZATION)
            .and_then(|h| h.to_str().ok()),
    )
    .or_else(|| {
        req.cookie("auth_token")
            .map(|c| c.value().to_string())
    })
    .ok_or_else(|| ErrorUnauthorized("No authorization token provided"))?;

    let claims = validate_jwt(&token, &config.jwt_secret)
        .map_err(|_| ErrorUnauthorized("Invalid token"))?;

    Ok(claims.user_id)
}

pub async fn check_is_admin(
    user_id: i32,
    pool: &sqlx::PgPool,
) -> Result<bool, sqlx::Error> {
    let result: (bool,) = sqlx::query_as("SELECT is_admin FROM users WHERE id = $1")
        .bind(user_id)
        .fetch_one(pool)
        .await?;
    
    Ok(result.0)
}

pub async fn check_is_blocked(
    user_id: i32,
    pool: &sqlx::PgPool,
) -> Result<bool, sqlx::Error> {
    let result: (Option<bool>,) = sqlx::query_as("SELECT is_blocked FROM users WHERE id = $1")
        .bind(user_id)
        .fetch_one(pool)
        .await?;
    
    Ok(result.0.unwrap_or(false))
}
