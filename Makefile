# Build targets
.PHONY: build build-windows build-linux build-macos build-plugin test clean help

# Default build
build: build-windows

# Windows build
build-windows:
	@echo "Building for Windows..."
	@set GOOS=windows&set GOARCH=amd64&go build -o bin/loglint.exe ./cmd/loglint

# Linux build
build-linux:
	@echo "Building for Linux..."
	@GOOS=linux GOARCH=amd64 go build -o bin/loglint ./cmd/loglint

# macOS build
build-macos:
	@echo "Building for macOS..."
	@GOOS=darwin GOARCH=amd64 go build -o bin/loglint ./cmd/loglint

# Plugin build (Linux only)
build-plugin:
	@echo "Building plugin for Linux..."
	@GOOS=linux CGO_ENABLED=1 go build -buildmode=plugin -o loglint.so ./plugin

# Tests
test:
	go test -v ./...

# Clean
clean:
	@echo "Cleaning up..."
	@go clean ./...
	@rm -f bin/loglint bin/loglint.exe loglint.so

# Help
help:
	@echo "Available targets:"
	@echo "  make build          - Build for Windows (default)"
	@echo "  make build-windows  - Build for Windows"
	@echo "  make build-linux    - Build for Linux"
	@echo "  make build-macos    - Build for macOS"
	@echo "  make build-plugin   - Build plugin for Linux"
	@echo "  make test           - Run tests"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make help           - Show this help"

.DEFAULT_GOAL := help
