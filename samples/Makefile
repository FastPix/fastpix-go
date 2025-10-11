# FastPix Go SDK Examples Makefile

.PHONY: help basic media live playback ai analytics security errors all clean

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
	@go run basic_usage.go

# Media management example
media: check-env
	@echo "Running media management example..."
	@go run media_management.go

# Live streaming example
live: check-env
	@echo "Running live streaming example..."
	@go run live_streaming.go

# Playback and playlist example
playback: check-env
	@echo "Running playback and playlist example..."
	@go run playback_playlist.go

# AI features example
ai: check-env
	@echo "Running AI features example..."
	@go run ai_features.go

# Analytics and metrics example
analytics: check-env
	@echo "Running analytics and metrics example..."
	@go run analytics_metrics.go

# Security and signing keys example
security: check-env
	@echo "Running security and signing keys example..."
	@go run security_signing_keys.go

# Error handling example
errors: check-env
	@echo "Running error handling example..."
	@go run error_handling.go

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

# Build all examples
build:
	@echo "Building all examples..."
	@go build -o basic_usage basic_usage.go
	@go build -o media_management media_management.go
	@go build -o live_streaming live_streaming.go
	@go build -o playback_playlist playback_playlist.go
	@go build -o ai_features ai_features.go
	@go build -o analytics_metrics analytics_metrics.go
	@go build -o security_signing_keys security_signing_keys.go
	@go build -o error_handling error_handling.go
	@echo "All examples built!"

# Run specific example by name
run-%: check-env
	@echo "Running $* example..."
	@go run $*.go
