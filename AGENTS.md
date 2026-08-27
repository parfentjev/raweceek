# AGENTS.md

## Working Agreement

- Use ASD-STE100 Simplified Technical English
- Do not modify files unless the user explicitly asks for a change.
- Treat questions as requests for concise answers, not implementation work.
- Do not provide code examples unless the user asks for them.
- Ask before adding dependencies, changing architecture, or expanding requested scope.
- Do not add tests or run build, test, lint, format, or application commands unless explicitly requested.
- Preserve unrelated worktree changes. Never revert changes made by the user.

## Project Overview

Rawe Ceek is a small Rust service that returns upcoming racing sessions and serves a static web application. It also exposes Prometheus metrics.

- Rust edition: 2024
- Build and dependencies: Cargo
- Entry point: `src/main.rs`
- Static web files: `public/`
- API contract: `spec/contract.yaml`
- Database design: `spec/schema.sql`

## Configuration

The service requires `DATABASE_URL`. The application listens on port 8080, and metrics listen on port 8081.

## Module Boundaries

Axum runs on Tokio and uses SQLx with PostgreSQL.

- `main`: application state, routing, configuration, and server startup.
- `handler`: HTTP handlers, response DTOs, and static file services.
- `session`: session queries and mapping database rows to API DTOs.
- `countdown`: countdown calculation and display formatting.
- `metrics`: request metrics and Prometheus setup.
- `error`: conversion of application errors to HTTP responses.

## Design Guidance

- Prefer smallest concrete design that solves current requirement.
- Keep the public API surface and module visibility minimal.
- Keep API behavior consistent with `spec/contract.yaml`.
- Keep database queries consistent with `spec/schema.sql`.
- Do not block the Tokio runtime with synchronous I/O.

## Rust Style

- Follow `rustfmt`, Clippy, and existing source conventions.
- Use `Result` and `thiserror`; preserve error sources and useful context.
- Avoid `unwrap` and `expect` in service code.
- Avoid speculative traits, helpers, abstractions, and compatibility code.
