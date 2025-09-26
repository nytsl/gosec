# Makefile for awesomeProject

# Variables
APP_NAME = awesomeProject
VERSION = 1.0.0
BUILD_DIR = build
BINARY_NAME = $(APP_NAME).exe

# Default target
.PHONY: all
all: clean build

# Build the application
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	@go mod tidy
	@go build -ldflags "-X awesomeProject/pkg/config.Version=$(VERSION)" -o $(BINARY_NAME) .
	@echo "Build completed: $(BINARY_NAME)"

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning..."
	@if exist $(BINARY_NAME) del $(BINARY_NAME)
	@if exist $(BUILD_DIR) rmdir /s /q $(BUILD_DIR)
	@echo "Clean completed"

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

# Run the application
.PHONY: run
run: build
	@echo "Running $(APP_NAME)..."
	@.\$(BINARY_NAME)

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Lint code
.PHONY: lint
lint:
	@echo "Linting code..."
	@go vet ./...

# Build for multiple platforms
.PHONY: build-all
build-all: clean
	@echo "Building for multiple platforms..."
	@mkdir $(BUILD_DIR)
	@set GOOS=windows&& set GOARCH=amd64&& go build -ldflags "-X awesomeProject/pkg/config.Version=$(VERSION)" -o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe .
	@set GOOS=linux&& set GOARCH=amd64&& go build -ldflags "-X awesomeProject/pkg/config.Version=$(VERSION)" -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64 .
	@set GOOS=darwin&& set GOARCH=amd64&& go build -ldflags "-X awesomeProject/pkg/config.Version=$(VERSION)" -o $(BUILD_DIR)/$(APP_NAME)-darwin-amd64 .
	@echo "Cross-platform build completed"

# Show help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all       - Clean and build the application"
	@echo "  build     - Build the application"
	@echo "  clean     - Clean build artifacts"
	@echo "  test      - Run tests"
	@echo "  run       - Build and run the application"
	@echo "  deps      - Install dependencies"
	@echo "  fmt       - Format code"
	@echo "  lint      - Lint code"
	@echo "  build-all - Build for multiple platforms"
	@echo "  help      - Show this help message"
