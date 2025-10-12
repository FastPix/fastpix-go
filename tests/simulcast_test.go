package tests

import (
	"context"
	"testing"

	"github.com/fastpix/fastpix-go/models/components"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimulcastManagement(t *testing.T) {
	config := GetTestConfig(t)
	ctx := context.Background()

	t.Run("Test Create Simulcast of Stream", func(t *testing.T) {
		// First, get a list of streams to find a valid ID
		listResponse, err := config.Client.ManageLiveStream.GetAllStreams(ctx, nil, nil, nil)
		require.NoError(t, err)

		if listResponse.GetStreamsResponse != nil && len(listResponse.GetStreamsResponse.Data) > 0 {
			streamID := listResponse.GetStreamsResponse.Data[0].StreamID
			require.NotNil(t, streamID)

			// Test creating simulcast
			simulcastRequest := components.SimulcastRequest{
				URL:      stringPtr("rtmps://test.example.com/live"),
				StreamKey: stringPtr("test-stream-key"),
				Metadata: map[string]string{
					"platform": "youtube",
					"test":     "true",
				},
			}

			response, err := config.Client.SimulcastStream.CreateSimulcastOfStream(ctx, *streamID, simulcastRequest)
			require.NoError(t, err)
			assert.NotNil(t, response)
			assert.NotNil(t, response.SimulcastResponse)
			
			// Store the simulcast ID for cleanup
			if response.SimulcastResponse != nil && response.SimulcastResponse.Data != nil {
				t.Logf("Created simulcast with ID: %s", *response.SimulcastResponse.Data.SimulcastID)
			}
		} else {
			t.Skip("No streams found to test CreateSimulcastOfStream")
		}
	})

	t.Run("Test Get Specific Simulcast of Stream", func(t *testing.T) {
		// First, get a list of streams to find a valid ID
		streamListResponse, err := config.Client.ManageLiveStream.GetAllStreams(ctx, nil, nil, nil)
		require.NoError(t, err)

		if streamListResponse.GetStreamsResponse != nil && len(streamListResponse.GetStreamsResponse.Data) > 0 {
			streamID := streamListResponse.GetStreamsResponse.Data[0].StreamID
			require.NotNil(t, streamID)

			// Create a simulcast first
			simulcastRequest := components.SimulcastRequest{
				URL:      stringPtr("rtmps://test.example.com/live"),
				StreamKey: stringPtr("test-stream-key"),
				Metadata: map[string]string{
					"platform": "youtube",
					"test":     "true",
				},
			}

			createResponse, err := config.Client.SimulcastStream.CreateSimulcastOfStream(ctx, *streamID, simulcastRequest)
			if err == nil && createResponse.SimulcastResponse != nil && createResponse.SimulcastResponse.Data != nil {
				simulcastID := createResponse.SimulcastResponse.Data.SimulcastID
				require.NotNil(t, simulcastID)

				// Test getting specific simulcast
				response, err := config.Client.SimulcastStream.GetSpecificSimulcastOfStream(ctx, *streamID, *simulcastID)
				require.NoError(t, err)
				assert.NotNil(t, response)
			} else {
				t.Skip("Could not create simulcast for testing")
			}
		} else {
			t.Skip("No streams found to test GetSpecificSimulcastOfStream")
		}
	})

	t.Run("Test Update Specific Simulcast of Stream", func(t *testing.T) {
		// First, get a list of streams to find a valid ID
		streamListResponse, err := config.Client.ManageLiveStream.GetAllStreams(ctx, nil, nil, nil)
		require.NoError(t, err)

		if streamListResponse.GetStreamsResponse != nil && len(streamListResponse.GetStreamsResponse.Data) > 0 {
			streamID := streamListResponse.GetStreamsResponse.Data[0].StreamID
			require.NotNil(t, streamID)

			// Create a simulcast first
			simulcastRequest := components.SimulcastRequest{
				URL:      stringPtr("rtmps://test.example.com/live"),
				StreamKey: stringPtr("test-stream-key"),
				Metadata: map[string]string{
					"platform": "youtube",
					"test":     "true",
				},
			}

			createResponse, err := config.Client.SimulcastStream.CreateSimulcastOfStream(ctx, *streamID, simulcastRequest)
			if err == nil && createResponse.SimulcastResponse != nil && createResponse.SimulcastResponse.Data != nil {
				simulcastID := createResponse.SimulcastResponse.Data.SimulcastID
				require.NotNil(t, simulcastID)

				// Test updating simulcast
				updateRequest := components.SimulcastUpdateRequest{
					IsEnabled: boolPtr(true),
					Metadata: map[string]string{
						"platform": "youtube",
						"test":     "updated",
					},
				}

				response, err := config.Client.SimulcastStream.UpdateSpecificSimulcastOfStream(ctx, *streamID, *simulcastID, updateRequest)
				require.NoError(t, err)
				assert.NotNil(t, response)
			} else {
				t.Skip("Could not create simulcast for testing")
			}
		} else {
			t.Skip("No streams found to test UpdateSpecificSimulcastOfStream")
		}
	})

	t.Run("Test Delete Simulcast of Stream", func(t *testing.T) {
		// First, get a list of streams to find a valid ID
		streamListResponse, err := config.Client.ManageLiveStream.GetAllStreams(ctx, nil, nil, nil)
		require.NoError(t, err)

		if streamListResponse.GetStreamsResponse != nil && len(streamListResponse.GetStreamsResponse.Data) > 0 {
			streamID := streamListResponse.GetStreamsResponse.Data[0].StreamID
			require.NotNil(t, streamID)

			// Create a simulcast first
			simulcastRequest := components.SimulcastRequest{
				URL:      stringPtr("rtmps://test.example.com/live"),
				StreamKey: stringPtr("test-stream-key"),
				Metadata: map[string]string{
					"platform": "youtube",
					"test":     "true",
				},
			}

			createResponse, err := config.Client.SimulcastStream.CreateSimulcastOfStream(ctx, *streamID, simulcastRequest)
			if err == nil && createResponse.SimulcastResponse != nil && createResponse.SimulcastResponse.Data != nil {
				simulcastID := createResponse.SimulcastResponse.Data.SimulcastID
				require.NotNil(t, simulcastID)

				// Test deleting simulcast
				response, err := config.Client.SimulcastStream.DeleteSimulcastOfStream(ctx, *streamID, *simulcastID)
				require.NoError(t, err)
				assert.NotNil(t, response)
			} else {
				t.Skip("Could not create simulcast for testing")
			}
		} else {
			t.Skip("No streams found to test DeleteSimulcastOfStream")
		}
	})
}

// Helper function to create bool pointers
func boolPtr(v bool) *bool {
	return &v
}
