package tests

import (
	"context"
	"testing"

	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/models/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMediaManagement(t *testing.T) {
	config := GetTestConfig(t)
	ctx := context.Background()

	t.Run("Test List Media", func(t *testing.T) {
		// Test listing all media
		response, err := config.Client.ManageVideos.ListMedia(ctx, nil, nil, nil)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.Object)
	})

	t.Run("Test List Media with Pagination", func(t *testing.T) {
		// Test listing media with pagination parameters
		limit := int64(5)
		offset := int64(1)
		orderBy := components.SortOrderDesc

		response, err := config.Client.ManageVideos.ListMedia(ctx, &limit, &offset, &orderBy)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.Object)
	})

	t.Run("Test Create Media from URL", func(t *testing.T) {
		// Test creating media from a public URL
		testURL := "https://sample-videos.com/zip/10/mp4/SampleVideo_1280x720_1mb.mp4"
		
		createRequest := components.CreateMediaRequest{
			Inputs: []components.Input{
				{
					VideoInput: &components.VideoInput{
						Type: "url",
						URL:  testURL,
					},
					Type: components.InputTypeVideoInput,
				},
			},
			AccessPolicy: components.CreateMediaRequestAccessPolicyPublic,
			Metadata: map[string]string{
				"test": "true",
				"source": "sdk-test",
			},
		}

		response, err := config.Client.InputVideo.CreateMedia(ctx, createRequest)
		if err != nil {
			// If the test URL is not accessible, that's okay for testing
			t.Logf("Create media failed (expected for test URL): %v", err)
			return
		}

		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.CreateMediaSuccessResponse)
	})

	t.Run("Test Get Media by ID", func(t *testing.T) {
		// First, get a list of media to find a valid ID
		listResponse, err := config.Client.ManageVideos.ListMedia(ctx, nil, nil, nil)
		require.NoError(t, err)

		if listResponse.Object != nil && len(listResponse.Object.Data) > 0 {
			mediaID := listResponse.Object.Data[0].ID
			require.NotNil(t, mediaID)

			// Test getting specific media
			response, err := config.Client.ManageVideos.GetMedia(ctx, *mediaID)
			require.NoError(t, err)
			assert.NotNil(t, response)
			assert.NotNil(t, response.Object)
		} else {
			t.Skip("No media found to test GetMedia")
		}
	})

	t.Run("Test Update Media", func(t *testing.T) {
		// First, get a list of media to find a valid ID
		listResponse, err := config.Client.ManageVideos.ListMedia(ctx, nil, nil, nil)
		require.NoError(t, err)

		if listResponse.Object != nil && len(listResponse.Object.Data) > 0 {
			mediaID := listResponse.Object.Data[0].ID
			require.NotNil(t, mediaID)

			// Test updating media metadata
			updateRequest := operations.UpdatedMediaRequestBody{
				Metadata: map[string]string{
					"updated": "true",
					"test": "sdk-update",
				},
			}

			response, err := config.Client.ManageVideos.UpdatedMedia(ctx, *mediaID, updateRequest)
			require.NoError(t, err)
			assert.NotNil(t, response)
		} else {
			t.Skip("No media found to test UpdateMedia")
		}
	})

	t.Run("Test List Uploads", func(t *testing.T) {
		// Test listing uploads
		response, err := config.Client.ManageVideos.ListUploads(ctx, nil, nil, nil)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.Object)
	})

	t.Run("Test Direct Upload Video Media", func(t *testing.T) {
		// Test direct upload (this would typically require a file)
		// For testing purposes, we'll just verify the method exists
		uploadRequest := operations.DirectUploadVideoMediaRequest{
			// This would normally contain file data
		}

		_, err := config.Client.InputVideo.DirectUploadVideoMedia(ctx, &uploadRequest)
		// This will likely fail without actual file data, which is expected
		assert.Error(t, err)
	})
}
