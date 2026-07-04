// Based on: https://github.com/tokio-rs/axum/blob/b90b8e02d0f761ce36a13610acd2afa60984a5e2/examples/prometheus-metrics/src/main.rs#L79
use std::time::{Duration, Instant};

use axum::{
    extract::{MatchedPath, Request},
    middleware::Next,
    response::IntoResponse,
};
use metrics_exporter_prometheus::{Matcher, PrometheusBuilder, PrometheusHandle};

#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("prometheus build error: {0}")]
    Build(#[from] metrics_exporter_prometheus::BuildError),
}

pub async fn track_metrics(request: Request, next: Next) -> impl IntoResponse {
    let Some(matched_path) = request.extensions().get::<MatchedPath>() else {
        return next.run(request).await;
    };

    let path = matched_path.as_str().to_owned();
    let method = request.method().clone().to_string();

    let start_time = Instant::now();
    let response = next.run(request).await;
    let latency = start_time.elapsed().as_secs_f64();

    let status = response.status().as_u16().to_string();

    let labels = [("method", method), ("path", path), ("status", status)];
    metrics::counter!("http_requests_total", &labels).increment(1);
    metrics::histogram!("http_requests_duration_seconds", &labels).record(latency);

    response
}

pub fn prometheus_handler() -> Result<PrometheusHandle, Error> {
    const EXPONENTIAL_SECONDS: &[f64] = &[
        0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1.0, 2.5, 5.0, 10.0,
    ];

    let recorder_handle = PrometheusBuilder::new()
        .set_buckets_for_metric(
            Matcher::Full("http_requests_duration_seconds".to_string()),
            EXPONENTIAL_SECONDS,
        )?
        .install_recorder()?;

    let upkeep_handle = recorder_handle.clone();
    tokio::spawn(async move {
        loop {
            tokio::time::sleep(Duration::from_secs(5)).await;
            upkeep_handle.run_upkeep();
        }
    });

    Ok(recorder_handle)
}
