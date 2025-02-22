.PHONY: build
build:
	docker build -t rawe-ceek-frontend:latest .
	docker save rawe-ceek-frontend:latest > rawe-ceek-frontend.tar

.PHONY: push
push:
	rsync -r rawe-ceek-frontend.tar $(RAWE_CEEK_FRONTEND_PWD)
