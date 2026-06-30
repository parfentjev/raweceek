mod countdown;
mod error;
mod handler;
mod session;

use std::{env, process};

use axum::{
    Router,
    routing::{self},
};
use sqlx::PgPool;
use tokio::net::TcpListener;

#[derive(Clone)]
pub struct AppState {
    pub db: PgPool,
}

#[derive(thiserror::Error, Debug)]
enum Error {
    #[error("env var read error: {0}")]
    EnvVar(#[from] std::env::VarError),
    #[error("database connection error: {0}")]
    Postgres(#[from] sqlx::Error),
    #[error("network io error: {0}")]
    IO(#[from] std::io::Error),
}

#[tokio::main]
async fn main() {
    if let Err(e) = run().await {
        eprintln!("unhandled service error: {}", e);
        process::exit(1);
    };
}

async fn run() -> Result<(), Error> {
    let database_url = env::var("DATABASE_URL")?;
    let db = PgPool::connect(&database_url).await?;

    let state = AppState { db };
    let app = Router::new()
        .route("/api/status", routing::get(handler::status))
        .route("/api/next-session", routing::get(handler::next_session))
        .fallback_service(handler::fallback())
        .with_state(state);

    let listener = TcpListener::bind("0.0.0.0:8080").await?;
    axum::serve(listener, app).await?;

    Ok(())
}
