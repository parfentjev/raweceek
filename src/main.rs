mod countdown;
mod handler;
mod session;

use std::{env, process};

use anyhow::Result;
use axum::{Router, routing};
use sqlx::PgPool;
use tokio::net::TcpListener;

#[derive(Clone)]
pub struct AppState {
    pub db: PgPool,
}

#[tokio::main]
async fn main() {
    if let Err(e) = run().await {
        eprintln!("service failed: {}", e);
        process::exit(1);
    };
}

async fn run() -> Result<()> {
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
