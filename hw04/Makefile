APP=friends
SHELL := /bin/bash
DOCKER_COMPOSE_FILE=docker-compose.yaml
GO_SERVICE=friends1 friends2

.PHONY: help
help: Makefile ## Show this help
	@echo
	@echo "Choose a command run in "$(APP)":"
	@echo
	@fgrep -h "##" $(MAKEFILE_LIST) | sed -e 's/\(\:.*\#\#\)/\:\ /' | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build: ## Build app
	go build -o bin/server cmd/main.go

test: ## Test app
	go test -failfast -count=1 -v ./... -coverpkg=./... -coverprofile=coverpkg.out

migrate: ## Migrate DB
	psql postgres://postgres:postgres@localhost:5432/friends < ./migrations/001_friends.sql

start: ## Start all deployed services
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

stop: ## Stop all deployed services
	docker-compose -f $(DOCKER_COMPOSE_FILE) stop

redeploy: ## Redeploy go services	
	docker-compose -f $(DOCKER_COMPOSE_FILE) stop $(GO_SERVICE); \
	docker-compose -f $(DOCKER_COMPOSE_FILE) rm -f $(GO_SERVICE); \
	docker-compose -f $(DOCKER_COMPOSE_FILE) up --build -d $(GO_SERVICE); \
	docker-compose -f $(DOCKER_COMPOSE_FILE) restart proxy; \