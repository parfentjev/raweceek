use axum::{http::StatusCode, response::IntoResponse};

use crate::session;

#[derive(thiserror::Error, Debug)]
pub enum AppError {
    #[error("session module error: {0}")]
    Session(#[from] session::Error),
}

impl IntoResponse for AppError {
    fn into_response(self) -> axum::response::Response {
        let (status_code, error) = match &self {
            Self::Session(session::Error::NotFound) => (StatusCode::NOT_FOUND, None),
            Self::Session(_) => (StatusCode::INTERNAL_SERVER_ERROR, Some(self)),
        };

        let response = status_code.into_response();
        if let Some(error) = error {
            // todo: add a logging middleware
            // https://github.com/tokio-rs/axum/blob/b90b8e02d0f761ce36a13610acd2afa60984a5e2/examples/error-handling/src/main.rs#L190
            eprintln!("handler error: {error}");
        }

        response
    }
}
