package tests

import (
	"context"
	"os"
	"testing"
	"time"

	fastpix "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
)

// TestConfig holds configuration for tests
type TestConfig struct {
	Username string
	Password string
	BaseURL  string
	Client   *fastpix.Fastpixgo
}

// GetTestConfig returns test configuration
func GetTestConfig(t *testing.T) *TestConfig {
	username := os.Getenv("FASTPIX_USERNAME")
	password := os.Getenv("FASTPIX_PASSWORD")
	baseURL := os.Getenv("FASTPIX_BASE_URL")

	if username == "" {
		t.Skip("FASTPIX_USERNAME environment variable not set")
	}
	if password == "" {
		t.Skip("FASTPIX_PASSWORD environment variable not set")
	}

	if baseURL == "" {
		baseURL = "https://api.fastpix.io/v1"
	}

	// Create SDK client with basic auth
	client := fastpix.New(
		fastpix.WithSecuritySource(func(ctx context.Context) (components.Security, error) {
			return components.Security{
				Username: &username,
				Password: &password,
			}, nil
		}),
		fastpix.WithTimeout(30*time.Second),
	)

	return &TestConfig{
		Username: username,
		Password: password,
		BaseURL:  baseURL,
		Client:   client,
	}
}

// TestMediaID is a test media ID for testing purposes
const TestMediaID = "test-media-id"

// TestStreamID is a test stream ID for testing purposes
const TestStreamID = "test-stream-id"

// TestPlaylistID is a test playlist ID for testing purposes
const TestPlaylistID = "test-playlist-id"

// TestSigningKeyID is a test signing key ID for testing purposes
const TestSigningKeyID = "test-signing-key-id"
