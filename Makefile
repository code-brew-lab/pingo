PROJECT_NAME ?= pingo

GOOS ?= $(shell uname | tr '[:upper:]' '[:lower:]')
GOARCH ?= $(shell arch)

CGO_ENABLED ?= 0

SRC_DIR = ./cmd
BIN_DIR = ./bin

PROD_OS_ARCH = "linux/amd64 darwin/amd64 windows/amd64 linux/arm64 darwin/arm64"

.PHONY: format install tidy clean-cache build clean build-prod help

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  format        Format the code using gofmt"
	@echo "  install       Install dependencies using go mod"
	@echo "  tidy          Tidy dependencies using go mod"
	@echo "  clean-cache   Clean build and package cache"
	@echo "  build         Build the project for development"
	@echo "  clean         Clean the build artifacts"
	@echo "  build-prod    Build the project for production for multiple OS/ARCH"
	@echo "  help          Display this help message"

format:
	@echo "Formatting..."
	@gofmt -s -w .
	@echo "Done!"

install:
	@echo "Installing dependencies..."
	@go mod download && go mod verify
	@echo "Done!"

tidy:
	@echo "Tidying dependencies..."
	@go mod tidy
	@echo "Done!"

clean-cache:
	@echo "Cleaning build and package cache..."
	@go clean -modcache
	@go clean -cache
	@echo "Done!"

build:
	@echo "Building for development..."
	@echo "GOOS         = $(GOOS)"
	@echo "GOARCH       = $(GOARCH)"
	@echo "CGO_ENABLED  = $(CGO_ENABLED)"
	@go build -o $(BIN_DIR)/$(PROJECT_NAME)-$(GOOS)-$(GOARCH) -v -x $(SRC_DIR)/$(PROJECT_NAME)
	@echo "Built!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf $(BIN_DIR)/*
	@echo "Cleaned!"