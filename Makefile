.PHONY: fmt lint build

fmt:
	cargo fmt

lint:
	cargo clippy --all

build:
	podman build -t raweceek:latest .
	rm raweceek.tar
	podman save -o raweceek.tar raweceek:latest

