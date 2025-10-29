use sqlx::{Pool, Postgres};

pub async fn run_migrations(pool: &Pool<Postgres>) -> Result<(), sqlx::Error> {
    // Users table
    sqlx::query(
        r#"
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username VARCHAR(255) UNIQUE NOT NULL,
            email VARCHAR(255) UNIQUE NOT NULL,
            password_hash VARCHAR(255) NOT NULL,
            is_admin BOOLEAN DEFAULT FALSE,
            is_blocked BOOLEAN DEFAULT FALSE,
            avatar VARCHAR(500),
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
        )
        "#,
    )
    .execute(pool)
    .await?;

    // Settings table
    sqlx::query(
        r#"
        CREATE TABLE IF NOT EXISTS settings (
            id SERIAL PRIMARY KEY,
            theme VARCHAR(50) DEFAULT 'light',
            logo_path VARCHAR(500),
            background_path VARCHAR(500),
            navbar_title VARCHAR(255) DEFAULT 'PInGO Share',
            max_upload_size BIGINT DEFAULT 104857600,
            upload_box_transparency INTEGER DEFAULT 0,
            blur_intensity INTEGER DEFAULT 0,
            max_validity VARCHAR(20) DEFAULT '7days',
            allow_registration BOOLEAN DEFAULT TRUE,
            expiration_action VARCHAR(20) DEFAULT 'unavailable',
            website_color VARCHAR(7) DEFAULT '#3b82f6',
            gradient_color_1 VARCHAR(7) DEFAULT '#3b82f6',
            gradient_color_2 VARCHAR(7) DEFAULT '#8b5cf6',
            gradient_color_3 VARCHAR(7) DEFAULT '#ec4899',
            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
        )
        "#,
    )
    .execute(pool)
    .await?;

    // Uploads table
    sqlx::query(
        r#"
        CREATE TABLE IF NOT EXISTS uploads (
            id SERIAL PRIMARY KEY,
            user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
            upload_id VARCHAR(255) UNIQUE NOT NULL,
            files TEXT NOT NULL,
            total_size BIGINT NOT NULL,
            email VARCHAR(255),
            download_url VARCHAR(255) NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            expires_at TIMESTAMP WITH TIME ZONE NULL,
            is_available BOOLEAN DEFAULT TRUE,
            is_reverse BOOLEAN DEFAULT FALSE,
            reverse_token VARCHAR(255),
            is_deleted BOOLEAN DEFAULT FALSE,
            deleted_at TIMESTAMP WITH TIME ZONE NULL,
            deletion_reason VARCHAR(255) DEFAULT NULL
        )
        "#,
    )
    .execute(pool)
    .await?;

    // Reverse share tokens table
    sqlx::query(
        r#"
        CREATE TABLE IF NOT EXISTS reverse_share_tokens (
            id SERIAL PRIMARY KEY,
            user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
            token VARCHAR(255) UNIQUE NOT NULL,
            name VARCHAR(255) NOT NULL,
            used_count INTEGER DEFAULT 0,
            max_uses INTEGER DEFAULT -1,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            expires_at TIMESTAMP WITH TIME ZONE NULL
        )
        "#,
    )
    .execute(pool)
    .await?;

    // Deletion logs table
    sqlx::query(
        r#"
        CREATE TABLE IF NOT EXISTS deletion_logs (
            id SERIAL PRIMARY KEY,
            user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
            username VARCHAR(255) NOT NULL,
            upload_id VARCHAR(255) NOT NULL,
            files TEXT NOT NULL,
            total_size BIGINT NOT NULL,
            email VARCHAR(255),
            download_url VARCHAR(255) NOT NULL,
            uploaded_at TIMESTAMP WITH TIME ZONE NOT NULL,
            deleted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            expires_at TIMESTAMP WITH TIME ZONE NULL,
            is_reverse BOOLEAN DEFAULT FALSE,
            reverse_token VARCHAR(255),
            deletion_reason VARCHAR(500) DEFAULT 'User deleted'
        )
        "#,
    )
    .execute(pool)
    .await?;

    // Create indexes
    sqlx::query("CREATE INDEX IF NOT EXISTS idx_users_email ON users(email)")
        .execute(pool)
        .await?;
    sqlx::query("CREATE INDEX IF NOT EXISTS idx_uploads_user_id ON uploads(user_id)")
        .execute(pool)
        .await?;
    sqlx::query("CREATE INDEX IF NOT EXISTS idx_uploads_upload_id ON uploads(upload_id)")
        .execute(pool)
        .await?;
    sqlx::query("CREATE INDEX IF NOT EXISTS idx_uploads_expires_at ON uploads(expires_at)")
        .execute(pool)
        .await?;
    sqlx::query("CREATE INDEX IF NOT EXISTS idx_reverse_tokens_token ON reverse_share_tokens(token)")
        .execute(pool)
        .await?;
    sqlx::query("CREATE INDEX IF NOT EXISTS idx_reverse_tokens_user_id ON reverse_share_tokens(user_id)")
        .execute(pool)
        .await?;

    // Insert default settings
    sqlx::query(
        r#"
        INSERT INTO settings (
            theme, max_upload_size, upload_box_transparency, blur_intensity,
            max_validity, allow_registration, expiration_action,
            website_color, gradient_color_1, gradient_color_2, gradient_color_3
        )
        SELECT 'light', 104857600, 0, 0, '7days', TRUE, 'unavailable',
               '#3b82f6', '#3b82f6', '#8b5cf6', '#ec4899'
        WHERE NOT EXISTS (SELECT 1 FROM settings)
        "#,
    )
    .execute(pool)
    .await?;

    log::info!("Database tables created successfully");
    Ok(())
}
