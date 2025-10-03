.PHONY: build test clean install lint fmt help

# Variables
BINARY_NAME=pdf-converter
BUILD_DIR=bin
MAIN_PATH=./cmd/pdf-converter
VERSION?=1.0.0
LDFLAGS=-ldflags "-X main.version=${VERSION}"

# CGO flags for Tesseract (adjust paths based on your OS)
# macOS Homebrew
export CGO_CPPFLAGS=-I/opt/homebrew/include
export CGO_LDFLAGS=-L/opt/homebrew/lib

## help: Display this help message
help:
	@echo "Available targets:"
	@echo "  build            - Build the binary"
	@echo "  test             - Run unit tests"
	@echo "  test-race        - Run tests with race detector"
	@echo "  test-integration - Run integration tests"
	@echo "  clean            - Remove build artifacts"
	@echo "  install          - Install binary to GOPATH/bin"
	@echo "  lint             - Run linters"
	@echo "  fmt              - Format code"
	@echo "  run              - Run with example (requires PDF_FILE env var)"
	@echo "  coverage         - Run tests with coverage"
	@echo "  install-hooks    - Install git pre-commit hooks"
	@echo "  check            - Run all checks (fmt, lint, test)"

## build: Build the binary
build:
	@echo "Building ${BINARY_NAME}..."
	@mkdir -p ${BUILD_DIR}
	go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY_NAME} ${MAIN_PATH}
	@echo "✅ Build complete: ${BUILD_DIR}/${BINARY_NAME}"

## test: Run unit tests
test:
	@echo "Running unit tests..."
	go test -v ./...

## test-race: Run tests with race detector (requires CGO)
test-race:
	@echo "Running tests with race detector..."
	go test -v -race ./...

## test-integration: Run integration tests (requires test PDFs)
test-integration:
	@echo "Running integration tests..."
	go test -v -tags=integration ./...

## coverage: Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -html=coverage.txt -o coverage.html
	@echo "✅ Coverage report: coverage.html"

## clean: Remove build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf ${BUILD_DIR}
	@rm -f ${BINARY_NAME}
	@rm -f coverage.txt coverage.html
	@echo "✅ Clean complete"

## install: Install binary to GOPATH/bin
install: build
	@echo "Installing ${BINARY_NAME}..."
	go install ${LDFLAGS} ${MAIN_PATH}
	@echo "✅ Installed to ${GOPATH}/bin/${BINARY_NAME}"

## lint: Run linters
lint:
	@echo "Running linters..."
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./...

## fmt: Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...
	@echo "✅ Format complete"

## run: Run the application (set PDF_FILE env var)
run: build
	@if [ -z "${PDF_FILE}" ]; then \
		echo "Error: PDF_FILE environment variable not set"; \
		echo "Usage: PDF_FILE=path/to/file.pdf make run"; \
		exit 1; \
	fi
	@${BUILD_DIR}/${BINARY_NAME} ${PDF_FILE}

## deps: Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy
	@echo "✅ Dependencies updated"

## check: Run all checks (fmt, lint, test)
check: fmt lint test
	@echo "✅ All checks passed"

## install-hooks: Install git pre-commit hooks
install-hooks:
	@echo "Installing git hooks..."
	@bash scripts/install-hooks.sh

