-include .env
export

.PHONY: fmt lint build run

fmt:
	cargo fmt

lint:
	cargo clippy --all

run:
	cargo run
