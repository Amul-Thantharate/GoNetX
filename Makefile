# Go executable name
BINARY_NAME := gonetx

# Go source files
GO_FILES := $(wildcard *.go)

# Build target
build:
	@echo "Building..."
	go build -o $(BINARY_NAME) $(GO_FILES)
	@echo "Build complete."

# Run target
run: build
	@echo "Running..."
	./$(BINARY_NAME)
	@echo "Run complete."

# Clean target
clean:
	@echo "Cleaning..."
	go clean
	rm -f $(BINARY_NAME)
	@echo "Clean complete."

# Default target
all: build run

# Install target
install: build
	@echo "Installing..."
	go install
	@echo "Installation complete.  Binary installed to $(go env GOPATH)/bin"

# Test target
test:
	@echo "Testing..."
	go test ./...
	@echo "Testing complete."

# Show help (using make itself)
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  all     Build and run the application (default)"
	@echo "  build   Build the application"
	@echo "  run     Run the application (requires build)"
	@echo "  clean   Remove build artifacts"
	@echo "  test    Run unit tests"
	@echo "  install Install the application to $(go env GOPATH)/bin"
	@echo "  help    Show this help message"
	@echo ""
	@echo "Example: make build"
	@echo "Example: make run"
	@echo "Example: make clean"

# Phony targets (not actual files)
.PHONY: all build run clean test install help