package tests

import (
	"os"
	"testing"
)

// RequireCredentials skips the calling test when FastPix credentials are not
// configured in the environment. This keeps the live integration suite from
// failing in environments (such as CI) where credentials are intentionally
// absent, while still running fully when they are provided.
func RequireCredentials(t *testing.T) {
	t.Helper()
	if os.Getenv("FASTPIX_USERNAME") == "" || os.Getenv("FASTPIX_PASSWORD") == "" {
		t.Skip("skipping integration test: set FASTPIX_USERNAME and FASTPIX_PASSWORD to run")
	}
}

// LoadConfig returns the FastPix SDK configuration for integration tests.
//
// Credentials are read from the environment so they are never committed to
// source control. Set the following before running the integration tests:
//
//	export FASTPIX_USERNAME="<your-token-id>"
//	export FASTPIX_PASSWORD="<your-secret-key>"
//
// The server URLs may optionally be overridden via FASTPIX_SERVER_URL and
// FASTPIX_LIVE_SERVER_URL; otherwise the production defaults are used.
func LoadConfig() (liveserverUrl, serverURL, username, password string) {
	liveserverUrl = getEnv("FASTPIX_LIVE_SERVER_URL", "https://api.fastpix.com/live")
	serverURL = getEnv("FASTPIX_SERVER_URL", "https://api.fastpix.com")
	username = os.Getenv("FASTPIX_USERNAME")
	password = os.Getenv("FASTPIX_PASSWORD")
	return
}

// getEnv returns the value of the environment variable named by key, or
// fallback when the variable is unset or empty.
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
