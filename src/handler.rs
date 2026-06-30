use axum::{Json, extract::State};
use serde::Serialize;
use tower_http::services::ServeDir;

use crate::{AppState, error::AppError, session, session::SessionDto};

#[derive(Serialize)]
#[serde(rename_all = "camelCase")]
pub struct StatusDto {
    pub race_week: bool,
    pub next_session: SessionDto,
}

/// GET /api/status
pub async fn status(State(state): State<AppState>) -> Result<Json<StatusDto>, AppError> {
    let race_week = session::count_this_week(&state.db).await? > 0;
    let next_session = session::find_next(&state.db).await?;

    Ok(Json(StatusDto {
        race_week,
        next_session,
    }))
}

/// GET /api/next-session
pub async fn next_session(State(state): State<AppState>) -> Result<Json<SessionDto>, AppError> {
    let session = session::find_next(&state.db).await?;

    Ok(Json(session))
}

/// Fallback for any other request (static files)
pub fn fallback() -> ServeDir {
    ServeDir::new("public")
}
