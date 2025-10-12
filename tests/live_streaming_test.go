package tests

import (
	"context"
	"testing"

	"github.com/fastpix/fastpix-go/models/components"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLiveStreaming(t *testing.T) {
	config := GetTestConfig(t)
	ctx := context.Background()

	t.Run("Test Create New Stream", func(t *testing.T) {
		// Test creating a new live stream
		createRequest := components.CreateLiveStreamRequest{
			PlaybackSettings: components.PlaybackSettings{
				// Add any required playback settings here
			},
			InputMediaSettings: components.InputMediaSettings{
				// Add any required input media settings here
			},
		}

		response, err := config.Client.StartLiveStream.CreateNewStream(ctx, createRequest)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.LiveStreamResponseDTO)
		
		// Store the stream ID for cleanup
		if response.LiveStreamResponseDTO != nil && response.LiveStreamResponseDTO.Data != nil && response.LiveStreamResponseDTO.Data.StreamID != nil {
			t.Logf("Created stream with ID: %s", *response.LiveStreamResponseDTO.Data.StreamID)
		}
	})

	t.Run("Test Get All Streams", func(t *testing.T) {
		// Test listing all live streams
		response, err := config.Client.ManageLiveStream.GetAllStreams(ctx, nil, nil, nil)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.GetStreamsResponse)
	})

	t.Run("Test Get Stream by ID", func(t *testing.T) {
		// First, get a list of streams to find a valid ID
		listResponse, err := config.Client.ManageLiveStream.GetAllStreams(ctx, nil, nil, nil)
		require.NoError(t, err)

		if listResponse.GetStreamsResponse != nil && len(listResponse.GetStreamsResponse.Data) > 0 {
			streamID := listResponse.GetStreamsResponse.Data[0].StreamID
			require.NotNil(t, streamID)

			// Test getting specific stream
		response, err := config.Client.ManageLiveStream.GetLiveStreamByID(ctx, *streamID)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.LivestreamgetResponse)
		} else {
			t.Skip("No streams found to test GetLiveStreamByID")
		}
	})

	t.Run("Test Update Live Stream", func(t *testing.T) {
		// First, get a list of streams to find a valid ID
		listResponse, err := config.Client.ManageLiveStream.GetAllStreams(ctx, nil, nil, nil)
		require.NoError(t, err)

		if listResponse.GetStreamsResponse != nil && len(listResponse.GetStreamsResponse.Data) > 0 {
			streamID := listResponse.GetStreamsResponse.Data[0].StreamID
			require.NotNil(t, streamID)

			// Test updating stream
			updateRequest := components.PatchLiveStreamRequest{
				Metadata: map[string]string{
					"name":        "Updated Test Stream",
					"description": "Updated by SDK test",
				},
			}

			response, err := config.Client.ManageLiveStream.UpdateLiveStream(ctx, *streamID, updateRequest)
			require.NoError(t, err)
			assert.NotNil(t, response)
		} else {
			t.Skip("No streams found to test UpdateLiveStream")
		}
	})

	t.Run("Test Enable Live Stream", func(t *testing.T) {
		// First, get a list of streams to find a valid ID
		listResponse, err := config.Client.ManageLiveStream.GetAllStreams(ctx, nil, nil, nil)
		require.NoError(t, err)

		if listResponse.GetStreamsResponse != nil && len(listResponse.GetStreamsResponse.Data) > 0 {
			streamID := listResponse.GetStreamsResponse.Data[0].StreamID
			require.NotNil(t, streamID)

			// Test enabling stream
			response, err := config.Client.ManageLiveStream.EnableLiveStream(ctx, *streamID)
			require.NoError(t, err)
			assert.NotNil(t, response)
		} else {
			t.Skip("No streams found to test EnableLiveStream")
		}
	})

	t.Run("Test Disable Live Stream", func(t *testing.T) {
		// First, get a list of streams to find a valid ID
		listResponse, err := config.Client.ManageLiveStream.GetAllStreams(ctx, nil, nil, nil)
		require.NoError(t, err)

		if listResponse.GetStreamsResponse != nil && len(listResponse.GetStreamsResponse.Data) > 0 {
			streamID := listResponse.GetStreamsResponse.Data[0].StreamID
			require.NotNil(t, streamID)

			// Test disabling stream
			response, err := config.Client.ManageLiveStream.DisableLiveStream(ctx, *streamID)
			require.NoError(t, err)
			assert.NotNil(t, response)
		} else {
			t.Skip("No streams found to test DisableLiveStream")
		}
	})

	t.Run("Test Get Stream Viewer Count", func(t *testing.T) {
		// First, get a list of streams to find a valid ID
		listResponse, err := config.Client.ManageLiveStream.GetAllStreams(ctx, nil, nil, nil)
		require.NoError(t, err)

		if listResponse.GetStreamsResponse != nil && len(listResponse.GetStreamsResponse.Data) > 0 {
			streamID := listResponse.GetStreamsResponse.Data[0].StreamID
			require.NotNil(t, streamID)

			// Test getting viewer count
			response, err := config.Client.ManageLiveStream.GetLiveStreamViewerCountByID(ctx, *streamID)
			require.NoError(t, err)
			assert.NotNil(t, response)
		} else {
			t.Skip("No streams found to test GetLiveStreamViewerCountByID")
		}
	})

	t.Run("Test Complete Live Stream", func(t *testing.T) {
		// First, get a list of streams to find a valid ID
		listResponse, err := config.Client.ManageLiveStream.GetAllStreams(ctx, nil, nil, nil)
		require.NoError(t, err)

		if listResponse.GetStreamsResponse != nil && len(listResponse.GetStreamsResponse.Data) > 0 {
			streamID := listResponse.GetStreamsResponse.Data[0].StreamID
			require.NotNil(t, streamID)

			// Test completing stream
			response, err := config.Client.ManageLiveStream.CompleteLiveStream(ctx, *streamID)
			require.NoError(t, err)
			assert.NotNil(t, response)
		} else {
			t.Skip("No streams found to test CompleteLiveStream")
		}
	})
}
