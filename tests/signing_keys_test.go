package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSigningKeys(t *testing.T) {
	config := GetTestConfig(t)
	ctx := context.Background()

	t.Run("Test Create Signing Key", func(t *testing.T) {
		// Test creating a new signing key
		response, err := config.Client.SigningKeys.CreateSigningKey(ctx)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.CreateResponse)
		
		// Store the signing key ID for cleanup
		if response.CreateResponse != nil && response.CreateResponse.Data != nil {
			t.Logf("Created signing key with ID: %s", *response.CreateResponse.Data.ID)
		}
	})

	t.Run("Test List Signing Keys", func(t *testing.T) {
		// Test listing all signing keys
		response, err := config.Client.SigningKeys.ListSigningKeys(ctx, nil, nil)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.GetAllSigningKeyResponse)
	})

	t.Run("Test List Signing Keys with Pagination", func(t *testing.T) {
		// Test listing signing keys with pagination
		limit := 10.0
		offset := 1.0

		response, err := config.Client.SigningKeys.ListSigningKeys(ctx, &limit, &offset)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.GetAllSigningKeyResponse)
	})

	t.Run("Test Get Signing Key by ID", func(t *testing.T) {
		// First, get a list of signing keys to find a valid ID
		listResponse, err := config.Client.SigningKeys.ListSigningKeys(ctx, nil, nil)
		require.NoError(t, err)

		if listResponse.GetAllSigningKeyResponse != nil && len(listResponse.GetAllSigningKeyResponse.Data) > 0 {
			signingKeyID := listResponse.GetAllSigningKeyResponse.Data[0].ID
			require.NotNil(t, signingKeyID)

			// Test getting specific signing key
			response, err := config.Client.SigningKeys.GetSigningKeyByID(ctx, *signingKeyID)
			require.NoError(t, err)
			assert.NotNil(t, response)
			assert.NotNil(t, response.GetPublicPemUsingSigningKeyIDResponseDTO)
		} else {
			t.Skip("No signing keys found to test GetSigningKeyByID")
		}
	})

	t.Run("Test Delete Signing Key", func(t *testing.T) {
		// First, get a list of signing keys to find a valid ID
		listResponse, err := config.Client.SigningKeys.ListSigningKeys(ctx, nil, nil)
		require.NoError(t, err)

		if listResponse.GetAllSigningKeyResponse != nil && len(listResponse.GetAllSigningKeyResponse.Data) > 0 {
			signingKeyID := listResponse.GetAllSigningKeyResponse.Data[0].ID
			require.NotNil(t, signingKeyID)

			// Test deleting signing key
			response, err := config.Client.SigningKeys.DeleteSigningKey(ctx, *signingKeyID)
			require.NoError(t, err)
			assert.NotNil(t, response)
		} else {
			t.Skip("No signing keys found to test DeleteSigningKey")
		}
	})
}
