.PHONY: build run clean test lint fmt help

APP_NAME := go_basic_framework_template
VERSION := 1.0.0
BUILD_TIME := $(shell date +%Y-%m-%d\ %H:%M:%S)
GO_VERSION := $(shell go version | awk '{print $$3}')

help:
	@echo "Available targets:"
	@echo "  build       Build the application"
	@echo "  run         Run the application"
	@echo "  clean       Clean build artifacts"
	@echo "  test        Run tests"
	@echo "  lint        Run linter"
	@echo "  fmt         Format code"
	@echo "  deps        Download dependencies"
	@echo "  tidy        Tidy dependencies"
	@echo "  swagger     Generate swagger docs"
	@echo "  docker      Build docker image"
	@echo "  migrate-up  Run database migrations up"
	@echo "  migrate-down Run database migrations down"

deps:
	go mod download

tidy:
	go mod tidy

build:
	@echo "Building $(APP_NAME)..."
	go build -ldflags "-s -w -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" -o bin/$(APP_NAME) ./cmd/server

run:
	@echo "Running $(APP_NAME)..."
	go run ./cmd/server

clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -rf logs/

test:
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...

test-coverage:
	@echo "Running tests with coverage..."
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint:
	@echo "Running linter..."
	golangci-lint run ./...

fmt:
	@echo "Formatting code..."
	go fmt ./...

vet:
	@echo "Running go vet..."
	go vet ./...

swagger:
	@echo "Generating swagger docs..."
	swag init -g cmd/server/main.go -o docs

docker:
	@echo "Building docker image..."
	docker build -t $(APP_NAME):$(VERSION) .

docker-run:
	@echo "Running docker container..."
	docker run -p 8080:8080 $(APP_NAME):$(VERSION)

migrate-up:
	@echo "Running migrations up..."
	go run ./cmd/migrate -direction up

migrate-down:
	@echo "Running migrations down..."
	go run ./cmd/migrate -direction down

dev:
	@echo "Running in development mode..."
	go run ./cmd/server -config config/config.yaml

prod:
	@echo "Running in production mode..."
	go run ./cmd/server -config config/config.prod.yaml

all: deps tidy fmt vet test build
	@echo "All done!"
