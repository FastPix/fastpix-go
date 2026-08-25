# FastPix Go SDK Examples Makefile

.PHONY: help basic media live playback ai analytics security errors create-upload verify-webhook all clean

# Default target
help:
	@echo "FastPix Go SDK Examples"
	@echo "======================="
	@echo ""
	@echo "Available targets:"
	@echo "  basic      - Run basic usage example"
	@echo "  media      - Run media management example"
	@echo "  live       - Run live streaming example"
	@echo "  playback   - Run playback and playlist example"
	@echo "  ai         - Run AI features example"
	@echo "  analytics  - Run analytics and metrics example"
	@echo "  security   - Run security and signing keys example"
	@echo "  errors     - Run error handling example"
	@echo "  create-upload  - Create a signed direct-upload URL"
	@echo "  verify-webhook - Verify a FastPix webhook signature (offline)"
	@echo "  all        - Run all examples"
	@echo "  clean      - Clean up generated files"
	@echo ""
	@echo "Prerequisites:"
	@echo "  - Set FASTPIX_USERNAME and FASTPIX_PASSWORD environment variables"
	@echo "  - Go 1.22 or later installed"
	@echo ""

# Check if environment variables are set
check-env:
	@if [ -z "$$FASTPIX_USERNAME" ]; then \
		echo "Error: FASTPIX_USERNAME environment variable not set"; \
		exit 1; \
	fi
	@if [ -z "$$FASTPIX_PASSWORD" ]; then \
		echo "Error: FASTPIX_PASSWORD environment variable not set"; \
		exit 1; \
	fi

# Basic usage example
basic: check-env
	@echo "Running basic usage example..."
	@go run ./basic-usage

# Media management example
media: check-env
	@echo "Running media management example..."
	@go run ./media-management

# Live streaming example
live: check-env
	@echo "Running live streaming example..."
	@go run ./live-streaming

# Playback and playlist example
playback: check-env
	@echo "Running playback and playlist example..."
	@go run ./playback-playlist

# AI features example
ai: check-env
	@echo "Running AI features example..."
	@go run ./ai-features

# Analytics and metrics example
analytics: check-env
	@echo "Running analytics and metrics example..."
	@go run ./analytics-metrics

# Security and signing keys example
security: check-env
	@echo "Running security and signing keys example..."
	@go run ./security-signing-keys

# Error handling example
errors: check-env
	@echo "Running error handling example..."
	@go run ./error-handling

# Create a signed direct-upload URL (its own package -> run by folder path)
create-upload: check-env
	@echo "Creating a signed upload URL..."
	@go run ./create-upload

# Verify a FastPix webhook signature (offline; no credentials needed)
verify-webhook:
	@echo "Verifying a webhook signature..."
	@go run ./verify-webhook

# Run all examples
all: check-env
	@echo "Running all examples..."
	@echo "======================"
	@echo ""
	@echo "1. Basic Usage"
	@echo "=============="
	@$(MAKE) basic
	@echo ""
	@echo "2. Media Management"
	@echo "==================="
	@$(MAKE) media
	@echo ""
	@echo "3. Live Streaming"
	@echo "================="
	@$(MAKE) live
	@echo ""
	@echo "4. Playback & Playlist"
	@echo "======================"
	@$(MAKE) playback
	@echo ""
	@echo "5. AI Features"
	@echo "=============="
	@$(MAKE) ai
	@echo ""
	@echo "6. Analytics & Metrics"
	@echo "======================"
	@$(MAKE) analytics
	@echo ""
	@echo "7. Security & Signing Keys"
	@echo "=========================="
	@$(MAKE) security
	@echo ""
	@echo "8. Error Handling"
	@echo "================="
	@$(MAKE) errors
	@echo ""
	@echo "9. Create Upload"
	@echo "================"
	@$(MAKE) create-upload
	@echo ""
	@echo "10. Verify Webhook"
	@echo "=================="
	@$(MAKE) verify-webhook
	@echo ""
	@echo "All examples completed!"

# Clean up generated files
clean:
	@echo "Cleaning up..."
	@rm -f *.exe
	@rm -f *.test
	@echo "Cleanup complete!"

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@echo "Dependencies installed!"

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Code formatted!"

# Run tests (if any)
test:
	@echo "Running tests..."
	@go test ./...
	@echo "Tests completed!"

# Compile-check all examples
build:
	@echo "Building all examples..."
	@go build ./...
	@echo "All examples built!"

# Run specific example by name
run-%: check-env
	@echo "Running $* example..."
	@go run ./$*
