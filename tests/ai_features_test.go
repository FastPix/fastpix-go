package tests

import (
	"context"
	"testing"

	"github.com/fastpix/fastpix-go/models/components"
	"github.com/fastpix/fastpix-go/models/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAIFeatures(t *testing.T) {
	config := GetTestConfig(t)
	ctx := context.Background()

	t.Run("Test Update Media Summary", func(t *testing.T) {
		// First, get a list of media to find a valid ID
		listResponse, err := config.Client.ManageVideos.ListMedia(ctx, nil, nil, nil)
		require.NoError(t, err)

		if listResponse.Object != nil && len(listResponse.Object.Data) > 0 {
			mediaID := listResponse.Object.Data[0].ID
			require.NotNil(t, mediaID)

			// Test updating media summary
			request := operations.UpdateMediaSummaryRequestBody{
				Generate:      true,
				SummaryLength: int64Ptr(100),
			}

			response, err := config.Client.InVideoAIFeatures.UpdateMediaSummary(ctx, *mediaID, request)
			require.NoError(t, err)
			assert.NotNil(t, response)
		} else {
			t.Skip("No media found to test UpdateMediaSummary")
		}
	})

	t.Run("Test Update Media Chapters", func(t *testing.T) {
		// First, get a list of media to find a valid ID
		listResponse, err := config.Client.ManageVideos.ListMedia(ctx, nil, nil, nil)
		require.NoError(t, err)

		if listResponse.Object != nil && len(listResponse.Object.Data) > 0 {
			mediaID := listResponse.Object.Data[0].ID
			require.NotNil(t, mediaID)

			// Test updating media chapters
			request := operations.UpdateMediaChaptersRequestBody{
				Chapters: true,
			}

			response, err := config.Client.InVideoAIFeatures.UpdateMediaChapters(ctx, *mediaID, request)
			require.NoError(t, err)
			assert.NotNil(t, response)
		} else {
			t.Skip("No media found to test UpdateMediaChapters")
		}
	})

	t.Run("Test Update Media Named Entities", func(t *testing.T) {
		// First, get a list of media to find a valid ID
		listResponse, err := config.Client.ManageVideos.ListMedia(ctx, nil, nil, nil)
		require.NoError(t, err)

		if listResponse.Object != nil && len(listResponse.Object.Data) > 0 {
			mediaID := listResponse.Object.Data[0].ID
			require.NotNil(t, mediaID)

			// Test updating media named entities
			request := operations.UpdateMediaNamedEntitiesRequestBody{
				NamedEntities: true,
			}

			response, err := config.Client.InVideoAIFeatures.UpdateMediaNamedEntities(ctx, *mediaID, request)
			require.NoError(t, err)
			assert.NotNil(t, response)
		} else {
			t.Skip("No media found to test UpdateMediaNamedEntities")
		}
	})

	t.Run("Test Update Media Moderation", func(t *testing.T) {
		// First, get a list of media to find a valid ID
		listResponse, err := config.Client.ManageVideos.ListMedia(ctx, nil, nil, nil)
		require.NoError(t, err)

		if listResponse.Object != nil && len(listResponse.Object.Data) > 0 {
			mediaID := listResponse.Object.Data[0].ID
			require.NotNil(t, mediaID)

			// Test updating media moderation
			request := operations.UpdateMediaModerationRequestBody{
				Moderation: &operations.UpdateMediaModerationModeration{
					Type: components.MediaTypeVideo.ToPointer(),
				},
			}

			response, err := config.Client.InVideoAIFeatures.UpdateMediaModeration(ctx, *mediaID, request)
			require.NoError(t, err)
			assert.NotNil(t, response)
		} else {
			t.Skip("No media found to test UpdateMediaModeration")
		}
	})
}

