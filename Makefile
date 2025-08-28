.PHONY: setup-build
setup-build:
	docker buildx create --name container-builder

.PHONY: build
build:
	docker buildx use container-builder
	docker buildx build --platform linux/arm64 -t raweceek:latest --output type=docker,dest=image.tar .

.PHONY: push
push:
	rsync image.tar ${VPS_USER}@${VPS_HOST}:${VPS_PUSH_PATH}
