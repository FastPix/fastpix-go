package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDRMConfigurations(t *testing.T) {
	config := GetTestConfig(t)
	ctx := context.Background()

	t.Run("Test Get DRM Configuration", func(t *testing.T) {
		// Test listing DRM configurations
		response, err := config.Client.DRMConfigurations.GetDrmConfiguration(ctx, nil, nil)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.Object)
	})

	t.Run("Test Get DRM Configuration with Pagination", func(t *testing.T) {
		// Test listing DRM configurations with pagination
		limit := int64(10)
		offset := int64(1)

		response, err := config.Client.DRMConfigurations.GetDrmConfiguration(ctx, &limit, &offset)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.Object)
	})

	t.Run("Test Get DRM Configuration by ID", func(t *testing.T) {
		// First, get a list of DRM configurations to find a valid ID
		listResponse, err := config.Client.DRMConfigurations.GetDrmConfiguration(ctx, nil, nil)
		require.NoError(t, err)

		if listResponse.Object != nil && len(listResponse.Object.Data) > 0 {
			drmConfigID := listResponse.Object.Data[0].ID
			require.NotNil(t, drmConfigID)

			// Test getting specific DRM configuration
			response, err := config.Client.DRMConfigurations.GetDrmConfigurationByID(ctx, *drmConfigID)
			require.NoError(t, err)
			assert.NotNil(t, response)
			assert.NotNil(t, response.Object)
		} else {
			t.Skip("No DRM configurations found to test GetDrmConfigurationByID")
		}
	})
}
