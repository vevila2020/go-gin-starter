.PHONY: help deps tidy fmt vet test build run clean docker

APP_NAME := go-gin-starter
VERSION ?= 0.1.0
BUILD_TIME := $(shell date +%Y-%m-%dT%H:%M:%S%z)

help:
	@echo "Available targets:"
	@echo "  deps     Download dependencies"
	@echo "  tidy     Tidy dependencies"
	@echo "  fmt      Format code"
	@echo "  vet      Run go vet"
	@echo "  test     Run tests"
	@echo "  build    Build the application"
	@echo "  run      Run the application"
	@echo "  clean    Clean build artifacts"
	@echo "  docker   Build docker image"

deps:
	go mod download

tidy:
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test -race -coverprofile=coverage.out ./...

build:
	go build -ldflags "-s -w -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" -o bin/$(APP_NAME) ./cmd/server

run:
	go run ./cmd/server -config config/config.yaml

clean:
	rm -rf bin/ logs/ coverage.out coverage.html

docker:
	docker build -t $(APP_NAME):$(VERSION) .