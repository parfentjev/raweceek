mod countdown;
mod error;
mod handler;
mod metrics;
mod session;

use core::future;
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
    #[error("metrics setup error: {0}")]
    Metrics(#[from] metrics::Error),
}

#[tokio::main]
async fn main() {
    let (server, result) = tokio::select! {
        result = run_app_server() => ("app_server", result),
        result = run_metrics_server() => ("metrics_server", result),
    };

    println!("{server} exited, stopping");
    if let Err(e) = result {
        eprintln!("unhandled service error: {}", e);
        process::exit(1);
    };
}

async fn run_app_server() -> Result<(), Error> {
    let database_url = env::var("DATABASE_URL")?;
    let db = PgPool::connect(&database_url).await?;

    let state = AppState { db };
    let app = Router::new()
        .route("/", routing::get_service(handler::index()))
        .route("/api/status", routing::get(handler::status))
        .route("/api/next-session", routing::get(handler::next_session))
        .route("/api/v2/status", routing::get(handler::status_v2))
        .route(
            "/api/v2/next-session",
            routing::get(handler::next_session_v2),
        )
        .route_layer(axum::middleware::from_fn(metrics::track_metrics))
        .fallback_service(handler::fallback())
        .with_state(state);

    let listener = TcpListener::bind("0.0.0.0:8080").await?;
    axum::serve(listener, app).await?;

    Ok(())
}

async fn run_metrics_server() -> Result<(), Error> {
    let metrics_handler = metrics::prometheus_handler()?;
    let app = Router::new().route(
        "/metrics",
        routing::get(move || future::ready(metrics_handler.render())),
    );

    let listener = TcpListener::bind("0.0.0.0:8081").await?;
    axum::serve(listener, app).await?;

    Ok(())
}
