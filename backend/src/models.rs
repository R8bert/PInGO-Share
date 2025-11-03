use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use sqlx::FromRow;

#[derive(Debug, Serialize, Deserialize, FromRow)]
pub struct User {
    pub id: i32,
    pub username: String,
    pub email: String,
    #[serde(skip_serializing)]
    pub password_hash: String,
    pub is_admin: bool,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub is_blocked: Option<bool>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub avatar: Option<String>,
    pub created_at: DateTime<Utc>,
}

#[derive(Debug, Serialize, Deserialize, FromRow)]
pub struct Settings {
    pub id: i32,
    pub theme: String,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub logo_path: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub background_path: Option<String>,
    pub navbar_title: String,
    pub max_upload_size: i64,
    pub upload_box_transparency: i32,
    pub blur_intensity: i32,
    pub max_validity: String,
    pub allow_registration: bool,
    pub expiration_action: String,
    pub website_color: String,
    pub gradient_color_1: String,
    pub gradient_color_2: String,
    pub gradient_color_3: String,
}

impl Default for Settings {
    fn default() -> Self {
        Self {
            id: 0,
            theme: "light".to_string(),
            logo_path: None,
            background_path: None,
            navbar_title: "RootDrop".to_string(),
            max_upload_size: 104857600, // 100MB
            upload_box_transparency: 0,
            blur_intensity: 0,
            max_validity: "7days".to_string(),
            allow_registration: true,
            expiration_action: "unavailable".to_string(),
            website_color: "#3b82f6".to_string(),
            gradient_color_1: "#3b82f6".to_string(),
            gradient_color_2: "#8b5cf6".to_string(),
            gradient_color_3: "#ec4899".to_string(),
        }
    }
}

#[derive(Debug, Serialize, Deserialize, FromRow)]
pub struct Upload {
    pub id: i32,
    pub user_id: i32,
    pub upload_id: String,
    pub files: String,
    pub total_size: i64,
    pub email: Option<String>,
    pub download_url: String,
    pub created_at: DateTime<Utc>,
    pub expires_at: Option<DateTime<Utc>>,
    pub is_available: bool,
    pub is_reverse: bool,
    pub reverse_token: Option<String>,
    pub is_deleted: bool,
    pub deleted_at: Option<DateTime<Utc>>,
    pub deletion_reason: Option<String>,
}

#[derive(Debug, Serialize, Deserialize, FromRow)]
pub struct ReverseShareToken {
    pub id: i32,
    pub user_id: i32,
    pub token: String,
    pub name: String,
    pub used_count: i32,
    pub max_uses: i32,
    pub created_at: DateTime<Utc>,
    pub expires_at: Option<DateTime<Utc>>,
}

#[derive(Debug, Serialize, Deserialize, FromRow)]
pub struct DeletionLog {
    pub id: i32,
    pub user_id: Option<i32>,
    pub username: String,
    pub upload_id: String,
    pub files: String,
    pub total_size: i64,
    pub email: Option<String>,
    pub download_url: String,
    pub uploaded_at: DateTime<Utc>,
    pub deleted_at: DateTime<Utc>,
    pub expires_at: Option<DateTime<Utc>>,
    pub is_reverse: bool,
    pub reverse_token: Option<String>,
    pub deletion_reason: String,
}

// Request/Response DTOs
#[derive(Debug, Deserialize)]
pub struct RegisterRequest {
    pub username: String,
    pub email: String,
    pub password: String,
}

#[derive(Debug, Deserialize)]
pub struct LoginRequest {
    pub email: String,
    pub password: String,
}

#[derive(Debug, Serialize)]
pub struct AuthResponse {
    pub message: String,
    pub token: String,
    pub user: UserInfo,
}

#[derive(Debug, Serialize)]
pub struct UserInfo {
    pub id: i32,
    pub username: String,
    pub email: String,
    pub is_admin: bool,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub avatar: Option<String>,
}

#[derive(Debug, Serialize)]
pub struct UploadResponse {
    pub download_url: String,
    pub files: Vec<String>,
    pub count: usize,
    pub expires_at: Option<DateTime<Utc>>,
}

#[derive(Debug, Deserialize)]
pub struct AvailabilityRequest {
    pub is_available: bool,
}

#[derive(Debug, Deserialize)]
pub struct ExpirationRequest {
    pub validity: String,
}

#[derive(Debug, Deserialize)]
pub struct CreateTokenRequest {
    pub name: String,
    #[serde(default)]
    pub max_uses: i32,
    #[serde(default)]
    pub expires_in: Option<String>,
}

#[derive(Debug, Deserialize)]
pub struct BlockUserRequest {
    pub blocked: bool,
}

#[derive(Debug, Deserialize)]
pub struct QuickSettingRequest {
    pub setting: String,
    pub value: serde_json::Value,
}

#[derive(Debug, Serialize)]
pub struct AdminStats {
    pub total_users: i64,
    pub total_uploads: i64,
    pub storage_used: i64,
}

#[derive(Debug, Serialize, FromRow)]
pub struct AdminUser {
    pub id: i32,
    pub username: String,
    pub email: String,
    pub avatar: Option<String>,
    pub is_admin: bool,
    pub is_blocked: bool,
    pub upload_count: i64,
    pub storage_used: i64,
    pub created_at: DateTime<Utc>,
    pub last_activity: Option<DateTime<Utc>>,
}

#[derive(Debug, Serialize)]
pub struct FileInfo {
    pub name: String,
    pub size: u64,
    pub url: String,
}

#[derive(Debug, Serialize)]
pub struct UploaderInfo {
    pub username: String,
    pub avatar: String,
    pub email: Option<String>,
    #[serde(rename = "expirationDate")]
    pub expiration_date: Option<DateTime<Utc>>,
}

#[derive(Debug, Serialize)]
pub struct FilesMetadataResponse {
    pub files: Vec<FileInfo>,
    pub uploader: UploaderInfo,
}
