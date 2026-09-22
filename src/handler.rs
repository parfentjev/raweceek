use axum::{Json, extract::State};
use serde::Serialize;
use tower_http::services::{ServeDir, ServeFile};

use crate::{
    AppState,
    error::AppError,
    session::{self, SessionDtoV2},
};

#[derive(Serialize)]
#[serde(rename_all = "snake_case")]
pub struct StatusDtoV2 {
    upcoming_sessions: Vec<SessionDtoV2>,
    race_week: bool,
}

/// GET /api/v2/status
pub async fn status_v2(State(state): State<AppState>) -> Result<Json<StatusDtoV2>, AppError> {
    let upcoming_sessions = session::find_upcoming(&state.db).await?;
    let race_week = upcoming_sessions.first().is_some_and(|s| s.this_week);

    Ok(Json(StatusDtoV2 {
        upcoming_sessions,
        race_week,
    }))
}

/// Handlers for static files
pub fn index() -> ServeFile {
    ServeFile::new("public/index.html")
}

pub fn fallback() -> ServeDir {
    ServeDir::new("public")
}
