.PHONY: build clean install test help run-examples

# Build the nanobanana binary
build:
	go build -o nanobanana

# Build with optimizations for production
build-release:
	CGO_ENABLED=0 go build -ldflags="-s -w" -o nanobanana

# Install to $GOPATH/bin
install:
	go install

# Clean build artifacts
clean:
	rm -f nanobanana
	rm -rf examples/*.png examples/*.jpg

# Run tests
test:
	go test -v ./...

# Run examples (requires API key)
run-examples:
	./examples.sh

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	go vet ./...

# Show help
help:
	@echo "Available targets:"
	@echo "  build          - Build the nanobanana binary"
	@echo "  build-release  - Build optimized binary for production"
	@echo "  install        - Install to \$$GOPATH/bin"
	@echo "  clean          - Clean build artifacts"
	@echo "  test           - Run tests"
	@echo "  run-examples   - Run example commands"
	@echo "  fmt            - Format code"
	@echo "  lint           - Run linter"
	@echo "  help           - Show this help message"
