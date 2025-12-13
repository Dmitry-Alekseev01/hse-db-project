ENV_FILE=.env
EXAMPLE_ENV_FILE=.env.example

.PHONY: env up

env:
	@if [ ! -f $(ENV_FILE) ]; then \
		cp $(EXAMPLE_ENV_FILE) $(ENV_FILE); \
		echo "$(ENV_FILE) создан из $(EXAMPLE_ENV_FILE)"; \
	else \
		echo "$(ENV_FILE) уже существует"; \
	fi

up: env
	docker compose up -d --build
	@echo "Application: http://localhost:8080"
	@echo "Swagger: http://localhost:8080/swagger/index.html"
