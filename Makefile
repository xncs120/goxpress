# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

DB_STRING = postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable

# Docker cmd
docker-db-up:
	@docker-compose up db -d --build

docker-db-down:
	@docker-compose down db

ifeq ($(APP_ENV), development)
docker-app-up:
	@TRACK_DIR=. docker-compose up app -d --build
else
docker-app-up:
	@docker-compose up app -d --build
endif

docker-app-down:
	@docker-compose down app -v

# Goose cmd
goose-create:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir=./goose/migrations create $(filter-out $@,$(MAKECMDGOALS)) sql

goose-status:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir=./goose/migrations status

goose-up:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir=./goose/migrations up

goose-down:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir=./goose/migrations down

goose-reset:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir=./goose/migrations reset

goose-v:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose version

# Go cmd
run:
	@go run ./cmd/api/main.go

build:
	@go build -o ./tmp/api/main ./cmd/api/main.go

test:
	@go test -v ./...