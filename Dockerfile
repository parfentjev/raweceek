FROM rust:1.96.0-slim-trixie AS builder
WORKDIR /usr/src/raweceek
COPY . .
RUN cargo install --path .

FROM debian:trixie-slim
COPY --from=builder /usr/local/cargo/bin/raweceek /usr/local/bin/app
COPY ./public /public
CMD ["app"]
