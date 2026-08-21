-include .env
export

.PHONY: fmt lint build

fmt:
	cargo fmt

lint:
	cargo clippy --all
