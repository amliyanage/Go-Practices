APP_NAME=go-jwt-tasks
APP_NAME_DEV=$(APP_NAME)-dev
DOCKER_DIR=docker
PORT=8080

.PHONY: help build run dev dev-stop compose-up compose-down logs clean

help:
	@echo "Available commands:"
	@echo "  make build          - Build Docker image"
	@echo "  make run            - Run Docker container"
	@echo "  make dev            - Run with Air hot-reload (development mode)"
	@echo "  make dev-stop       - Stop development container"
	@echo "  make up             - Docker compose up"
	@echo "  make down           - Docker compose down"
	@echo "  make logs           - View docker compose logs (follow mode)"
	@echo "  make clean          - Remove Docker container and image"

build:
	docker build -t $(APP_NAME) -f $(DOCKER_DIR)/Dockerfile .

run:
	docker run -p $(PORT):8080 --name $(APP_NAME) $(APP_NAME)

dev:
	docker build -t $(APP_NAME_DEV) -f $(DOCKER_DIR)/Dockerfile.dev .
	docker run -p $(PORT):8080 --name $(APP_NAME_DEV) -v $(PWD):/app $(APP_NAME_DEV)

dev-stop:
	docker rm -f $(APP_NAME_DEV) || true
	docker rmi -f $(APP_NAME_DEV) || true

up:
	docker compose -f $(DOCKER_DIR)/docker-compose.yml up --build -d

down:
	docker compose -f $(DOCKER_DIR)/docker-compose.yml down

logs:
	docker compose -f $(DOCKER_DIR)/docker-compose.yml logs -f

clean:
	docker rm -f $(APP_NAME) || true
	docker rmi -f $(APP_NAME) || true
