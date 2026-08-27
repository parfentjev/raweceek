use axum::{Json, extract::State};
use serde::Serialize;
use tower_http::services::{ServeDir, ServeFile};

use crate::{
    AppState,
    error::AppError,
    session::{self, SessionDto, SessionDtoV2},
};

#[derive(Serialize)]
#[serde(rename_all = "camelCase")]
pub struct StatusDto {
    pub race_week: bool,
    pub next_session: SessionDto,
}

#[derive(Serialize)]
#[serde(rename_all = "snake_case")]
pub struct StatusDtoV2 {
    upcoming_sessions: Vec<SessionDtoV2>,
    next_session_index: usize,
    race_week: bool,
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

/// GET /api/v2/status
pub async fn status_v2(State(state): State<AppState>) -> Result<Json<StatusDtoV2>, AppError> {
    let upcoming_sessions = session::find_next_v2(&state.db).await?;

    let (next_session_index, race_week) = upcoming_sessions
        .iter()
        .enumerate()
        .find(|(_, s)| !s.started)
        .map(|(i, s)| (i, s.this_week))
        .expect("upcoming_sessions is guaranteed to contain at least 1 session");

    Ok(Json(StatusDtoV2 {
        upcoming_sessions,
        next_session_index,
        race_week,
    }))
}

/// GET /api/next-session
pub async fn next_session(State(state): State<AppState>) -> Result<Json<SessionDto>, AppError> {
    let session = session::find_next(&state.db).await?;

    Ok(Json(session))
}

/// GET /api/v2/next-session
pub async fn next_session_v2(
    State(state): State<AppState>,
) -> Result<Json<SessionDtoV2>, AppError> {
    let session = session::find_next_session_v2(&state.db).await?;

    Ok(Json(session))
}

/// Handlers for static files
pub fn index() -> ServeFile {
    ServeFile::new("public/index.html")
}

pub fn fallback() -> ServeDir {
    ServeDir::new("public")
}
