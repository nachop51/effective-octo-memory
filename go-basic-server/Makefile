# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=server
BINARY_UNIX=$(BINARY_NAME)_unix

# Test parameters
TEST_TIMEOUT=30s
TEST_VERBOSE=-v
COVERAGE_FILE=coverage.out

.PHONY: help build clean test test-unit test-integration test-coverage deps run migrate lint format

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Build the application
build: ## Build the application
	$(GOBUILD) -o $(BINARY_NAME) -v ./...

# Clean build artifacts
clean: ## Clean build artifacts
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(COVERAGE_FILE)

# Install dependencies
deps: ## Download and install dependencies
	$(GOMOD) download
	$(GOMOD) tidy

# Run the application
run: ## Run the application
	$(GOCMD) run main.go

# Run database migrations
migrate: ## Run database migrations
	$(GOCMD) run main.go -migrate

# Run all tests
test: ## Run all tests
	$(GOTEST) $(TEST_VERBOSE) -timeout $(TEST_TIMEOUT) ./tests/...

# Run only unit tests
test-unit: ## Run unit tests only
	$(GOTEST) $(TEST_VERBOSE) -timeout $(TEST_TIMEOUT) ./tests/unit/...

# Run only integration tests
test-integration: ## Run integration tests only
	$(GOTEST) $(TEST_VERBOSE) -timeout $(TEST_TIMEOUT) ./tests/integration/...

# Run tests with coverage
test-coverage: ## Run tests with coverage report
	$(GOTEST) $(TEST_VERBOSE) -timeout $(TEST_TIMEOUT) -coverprofile=$(COVERAGE_FILE) ./tests/...
	$(GOCMD) tool cover -html=$(COVERAGE_FILE) -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run tests with coverage (console output)
test-coverage-console: ## Run tests with coverage in console
	$(GOTEST) $(TEST_VERBOSE) -timeout $(TEST_TIMEOUT) -coverprofile=$(COVERAGE_FILE) ./tests/...
	$(GOCMD) tool cover -func=$(COVERAGE_FILE)

# Watch for changes and run tests
test-watch: ## Watch for changes and run tests (requires entr)
	find . -name "*.go" | entr -c make test

# Lint the code
lint: ## Run golangci-lint
	golangci-lint run

# Format the code
format: ## Format Go code
	$(GOCMD) fmt ./...
	$(GOCMD) mod tidy

# Generate mocks (if using mockery)
generate-mocks: ## Generate mocks for interfaces
	@echo "Generating mocks..."
	@if command -v mockery > /dev/null; then \
		mockery --all --output ./tests/mocks --case underscore; \
	else \
		echo "mockery not installed. Install with: go install github.com/vektra/mockery/v2@latest"; \
	fi

# Setup development environment
setup-dev: deps ## Setup development environment
	@echo "Setting up development environment..."
	@if ! command -v golangci-lint > /dev/null; then \
		echo "Installing golangci-lint..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.54.2; \
	fi
	@echo "Development environment setup complete!"

# Run tests in CI environment
test-ci: ## Run tests in CI environment
	$(GOTEST) -timeout $(TEST_TIMEOUT) -race -coverprofile=$(COVERAGE_FILE) ./tests/...

# Build for Linux
build-linux: ## Build for Linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

# Run with air for hot reload (development)
dev: ## Run with hot reload using air
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "air not installed. Install with: go install github.com/cosmtrek/air@latest"; \
		echo "Or run with: make run"; \
	fi

# Security scan
security: ## Run security scan with gosec
	@if command -v gosec > /dev/null; then \
		gosec ./...; \
	else \
		echo "gosec not installed. Install with: go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest"; \
	fi

# Run all quality checks
quality: format lint test security ## Run all quality checks (format, lint, test, security)

# Docker build
docker-build: ## Build Docker image
	docker build -t $(BINARY_NAME) .

# Docker run
docker-run: ## Run Docker container
	docker run -p 1234:1234 $(BINARY_NAME)
