use std::env;

#[derive(Clone)]
pub struct Config {
    pub database_url: String,
    pub jwt_secret: String,
    pub server_port: String,
    pub allowed_origins: String,
}

impl Config {
    pub fn from_env() -> Self {
        dotenv::dotenv().ok();

        let jwt_secret = env::var("JWT_SECRET")
            .expect("JWT_SECRET must be set");
        
        if jwt_secret.len() < 32 {
            panic!("JWT_SECRET must be at least 32 characters long");
        }

        let db_host = env::var("DB_HOST").unwrap_or_else(|_| "localhost".to_string());
        let db_port = env::var("DB_PORT").unwrap_or_else(|_| "5432".to_string());
        let db_user = env::var("DB_USER").unwrap_or_else(|_| "user".to_string());
        let db_password = env::var("DB_PASSWORD").unwrap_or_else(|_| "pass".to_string());
        let db_name = env::var("DB_NAME").unwrap_or_else(|_| "filesharing".to_string());

        let database_url = format!(
            "postgres://{}:{}@{}:{}/{}",
            db_user, db_password, db_host, db_port, db_name
        );

        Self {
            database_url,
            jwt_secret,
            server_port: env::var("SERVER_PORT").unwrap_or_else(|_| "8080".to_string()),
            allowed_origins: env::var("ALLOWED_ORIGINS").unwrap_or_else(|_| "*".to_string()),
        }
    }
}
