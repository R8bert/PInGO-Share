use actix_web::{error, web, Error, HttpResponse};
use sqlx::PgPool;

use crate::models::Settings;

pub async fn get_settings(pool: web::Data<PgPool>) -> Result<HttpResponse, Error> {
    let settings: Option<Settings> = sqlx::query_as(
        r#"
        SELECT id, theme, logo_path, background_path, navbar_title, max_upload_size,
               upload_box_transparency, blur_intensity, max_validity, allow_registration,
               expiration_action, website_color, gradient_color_1, gradient_color_2, gradient_color_3
        FROM settings ORDER BY id LIMIT 1
        "#
    )
    .fetch_optional(pool.as_ref())
    .await
    .map_err(error::ErrorInternalServerError)?;

    let settings = settings.unwrap_or_default();

    Ok(HttpResponse::Ok().json(settings))
}
