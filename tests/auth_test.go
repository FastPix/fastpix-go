package tests

import (
	"context"
	"os"
	"testing"
	"time"

	fastpix "github.com/fastpix/fastpix-go"
	"github.com/fastpix/fastpix-go/models/components"
	"github.com/stretchr/testify/assert"
)

func TestAuthentication(t *testing.T) {
	config := GetTestConfig(t)
	ctx := context.Background()

	t.Run("Test Basic Authentication", func(t *testing.T) {
		// Test that we can create a client with basic auth
		client := fastpix.New(
		fastpix.WithSecuritySource(func(ctx context.Context) (components.Security, error) {
			return components.Security{
				Username: &config.Username,
				Password: &config.Password,
			}, nil
		}),
			fastpix.WithTimeout(10*time.Second),
		)

		assert.NotNil(t, client)
		assert.NotNil(t, client.ManageVideos)
		assert.NotNil(t, client.ManageLiveStream)
		assert.NotNil(t, client.StartLiveStream)
	})

	t.Run("Test API Key Authentication", func(t *testing.T) {
		// Test API key authentication if available
		apiKey := os.Getenv("FASTPIX_API_KEY")
		if apiKey == "" {
			t.Skip("FASTPIX_API_KEY not set, skipping API key test")
		}

		client := fastpix.New(
		fastpix.WithSecuritySource(func(ctx context.Context) (components.Security, error) {
			return components.Security{
				Username: &apiKey,
				Password: &apiKey,
			}, nil
		}),
			fastpix.WithTimeout(10*time.Second),
		)

		assert.NotNil(t, client)
	})

	t.Run("Test Invalid Credentials", func(t *testing.T) {
		// Test with invalid credentials
		client := fastpix.New(
		fastpix.WithSecuritySource(func(ctx context.Context) (components.Security, error) {
			invalid := "invalid"
			return components.Security{
				Username: &invalid,
				Password: &invalid,
			}, nil
		}),
			fastpix.WithTimeout(5*time.Second),
		)

		// Try to make a request that should fail
		_, err := client.ManageVideos.ListMedia(ctx, nil, nil, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "401") // Should get unauthorized error
	})
}
