.PHONY: build
build:
	docker buildx use container-builder
	docker buildx build --platform linux/arm64 -t rawe-ceek-frontend:latest --output type=docker,dest=image.tar .

.PHONY: push
push:
	rsync -r image.tar $(RAWE_CEEK_FRONTEND_PWD)
