use anyhow::Result;
use axum::{
    Json,
    extract::State,
    http::StatusCode,
    response::{IntoResponse, Response},
};
use serde::Serialize;
use tower_http::services::ServeDir;

use crate::{
    AppState,
    session::{self, SessionDto},
};

pub struct HandlerError {
    error: anyhow::Error,
}

impl HandlerError {
    fn status_code(&self) -> StatusCode {
        match self.error.downcast_ref::<sqlx::Error>() {
            Some(sqlx::Error::RowNotFound) => StatusCode::NOT_FOUND,
            _ => StatusCode::INTERNAL_SERVER_ERROR,
        }
    }
}

pub type ResponseBody<T> = Result<T, HandlerError>;

impl<E: Into<anyhow::Error>> From<E> for HandlerError {
    fn from(value: E) -> Self {
        Self {
            error: value.into(),
        }
    }
}

impl IntoResponse for HandlerError {
    fn into_response(self) -> Response {
        eprintln!("handler error: {}", self.error);
        self.status_code().into_response()
    }
}

#[derive(Serialize)]
#[serde(rename_all = "camelCase")]
pub struct StatusDto {
    pub race_week: bool,
    pub next_session: SessionDto,
}

/// GET /api/status
pub async fn status(State(state): State<AppState>) -> ResponseBody<Json<StatusDto>> {
    let race_week = session::count_this_week(&state.db).await? > 0;
    let next_session = session::find_next(&state.db).await?;

    Ok(Json(StatusDto {
        race_week,
        next_session,
    }))
}

/// GET /api/next-session
pub async fn next_session(State(state): State<AppState>) -> ResponseBody<Json<SessionDto>> {
    let session = session::find_next(&state.db).await?;

    Ok(Json(session))
}

/// Fallback for any other request (static files)
pub fn fallback() -> ServeDir {
    ServeDir::new("public")
}
