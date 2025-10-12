package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/models/operations"
)

func main() {
	ctx := context.Background()

	// Initialize SDK
	client := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")),
			Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")),
		}),
		fastpixgo.WithTimeout(30*time.Second),
	)

	// 1. Create New Live Stream
	fmt.Println("=== Creating New Live Stream ===")
	createStreamRequest := components.CreateLiveStreamRequest{
		PlaybackSettings: components.PlaybackSettings{
			// Configure playback settings
		},
		InputMediaSettings: components.InputMediaSettings{
			// Configure input media settings
		},
		Metadata: map[string]string{
			"name":        "My Live Stream",
			"description": "A test live stream",
			"category":    "gaming",
		},
	}

	createResponse, err := client.StartLiveStream.CreateNewStream(ctx, createStreamRequest)
	if err != nil {
		log.Printf("Error creating live stream: %v", err)
	} else {
		if createResponse.LiveStreamResponseDTO != nil && createResponse.LiveStreamResponseDTO.Data != nil {
			streamID := createResponse.LiveStreamResponseDTO.Data.StreamID
			if streamID != nil {
				fmt.Printf("Live stream created successfully! Stream ID: %s\n", *streamID)
			}
			fmt.Printf("RTMP URL: %s\n", getStringValue(createResponse.LiveStreamResponseDTO.Data.RtmpURL))
			fmt.Printf("Stream Key: %s\n", getStringValue(createResponse.LiveStreamResponseDTO.Data.StreamKey))
		} else {
			fmt.Println("Live stream created but no data returned")
		}
	}

	// 2. List All Live Streams
	fmt.Println("\n=== Listing All Live Streams ===")
	limit := int64(10)
	offset := int64(0)
	orderBy := components.SortOrderDesc

	streamsResponse, err := client.ManageLiveStream.GetAllStreams(ctx, &limit, &offset, &orderBy)
	if err != nil {
		log.Printf("Error listing streams: %v", err)
	} else {
		fmt.Printf("Found %d live streams:\n", len(streamsResponse.GetStreamsResponse.Data))
		for i, stream := range streamsResponse.GetStreamsResponse.Data {
			fmt.Printf("  %d. ID: %s, Name: %s, Status: %s\n",
				i+1, *stream.StreamID, getStringValue(stream.Name), getStringValue(stream.Status))
		}
	}

	// 3. Get Specific Stream Details
	if streamsResponse.GetStreamsResponse != nil && len(streamsResponse.GetStreamsResponse.Data) > 0 {
		streamID := streamsResponse.GetStreamsResponse.Data[0].StreamID
		fmt.Printf("\n=== Getting Stream Details for ID: %s ===\n", *streamID)

		streamResponse, err := client.ManageLiveStream.GetLiveStreamByID(ctx, *streamID)
		if err != nil {
			log.Printf("Error getting stream details: %v", err)
		} else {
			stream := streamResponse.LivestreamgetResponse.Data
			fmt.Printf("Name: %s\n", getStringValue(stream.Name))
			fmt.Printf("Status: %s\n", getStringValue(stream.Status))
			fmt.Printf("Created: %s\n", getStringValue(stream.CreatedAt))
			fmt.Printf("RTMP URL: %s\n", getStringValue(stream.RtmpURL))
		}

		// 4. Update Live Stream
		fmt.Printf("\n=== Updating Live Stream: %s ===\n", *streamID)
		updateRequest := components.PatchLiveStreamRequest{
			Metadata: map[string]string{
				"name":        "Updated Live Stream",
				"description": "Updated description for testing",
				"category":    "updated-gaming",
			},
		}

		updateResponse, err := client.ManageLiveStream.UpdateLiveStream(ctx, *streamID, updateRequest)
		if err != nil {
			log.Printf("Error updating stream: %v", err)
		} else {
			fmt.Println("Live stream updated successfully!")
		}

		// 5. Enable Live Stream
		fmt.Printf("\n=== Enabling Live Stream: %s ===\n", *streamID)
		enableResponse, err := client.ManageLiveStream.EnableLiveStream(ctx, *streamID)
		if err != nil {
			log.Printf("Error enabling stream: %v", err)
		} else {
			fmt.Println("Live stream enabled successfully!")
			fmt.Printf("Response: %+v\n", enableResponse)
		}

		// 6. Get Stream Viewer Count
		fmt.Printf("\n=== Getting Viewer Count for Stream: %s ===\n", *streamID)
		viewerResponse, err := client.ManageLiveStream.GetLiveStreamViewerCountByID(ctx, *streamID)
		if err != nil {
			log.Printf("Error getting viewer count: %v", err)
		} else {
			fmt.Printf("Current viewer count: %d\n", getInt64Value(viewerResponse.Object.Data.ViewerCount))
		}

		// 7. Create Playback ID for Live Stream
		fmt.Printf("\n=== Creating Playback ID for Stream: %s ===\n", *streamID)
		playbackRequest := operations.CreatePlaybackIDOfStreamRequestBody{
			AccessPolicy: components.AccessPolicyPublic,
		}

		playbackResponse, err := client.LivePlayback.CreatePlaybackIDOfStream(ctx, *streamID, &playbackRequest)
		if err != nil {
			log.Printf("Error creating playback ID: %v", err)
		} else {
			fmt.Printf("Playback ID created successfully! ID: %s\n", *playbackResponse.Object.Data.PlaybackID)
		}

		// 8. Get Live Stream Playback ID Details
		if playbackResponse.Object != nil && playbackResponse.Object.Data != nil {
			playbackID := playbackResponse.Object.Data.PlaybackID
			fmt.Printf("\n=== Getting Playback ID Details: %s ===\n", *playbackID)

			playbackDetailsResponse, err := client.LivePlayback.GetLiveStreamPlaybackID(ctx, *streamID, *playbackID)
			if err != nil {
				log.Printf("Error getting playback ID details: %v", err)
			} else {
				fmt.Printf("Playback ID details retrieved successfully!\n")
				fmt.Printf("Access Policy: %s\n", getStringValue(playbackDetailsResponse.Object.Data.AccessPolicy))
			}

			// 9. Delete Playback ID
			fmt.Printf("\n=== Deleting Playback ID: %s ===\n", *playbackID)
			deletePlaybackResponse, err := client.LivePlayback.DeletePlaybackIDOfStream(ctx, *streamID, *playbackID)
			if err != nil {
				log.Printf("Error deleting playback ID: %v", err)
			} else {
				fmt.Println("Playback ID deleted successfully!")
				fmt.Printf("Response: %+v\n", deletePlaybackResponse)
			}
		}

		// 10. Create Simulcast
		fmt.Printf("\n=== Creating Simulcast for Stream: %s ===\n", *streamID)
		simulcastRequest := components.CreateSimulcastRequest{
			Platform: components.SimulcastPlatformYoutube,
			URL:      "https://youtube.com/watch?v=example",
		}

		simulcastResponse, err := client.SimulcastStream.CreateSimulcastOfStream(ctx, *streamID, simulcastRequest)
		if err != nil {
			log.Printf("Error creating simulcast: %v", err)
		} else {
			fmt.Printf("Simulcast created successfully! ID: %s\n", *simulcastResponse.Object.Data.ID)
		}

		// 11. Get Simulcast Details
		if simulcastResponse.Object != nil && simulcastResponse.Object.Data != nil {
			simulcastID := simulcastResponse.Object.Data.ID
			fmt.Printf("\n=== Getting Simulcast Details: %s ===\n", *simulcastID)

			simulcastDetailsResponse, err := client.SimulcastStream.GetSpecificSimulcastOfStream(ctx, *streamID, *simulcastID)
			if err != nil {
				log.Printf("Error getting simulcast details: %v", err)
			} else {
				fmt.Printf("Simulcast details retrieved successfully!\n")
				fmt.Printf("Platform: %s\n", getStringValue(simulcastDetailsResponse.Object.Data.Platform))
			}

			// 12. Update Simulcast
			fmt.Printf("\n=== Updating Simulcast: %s ===\n", *simulcastID)
			updateSimulcastRequest := components.UpdateSimulcastRequest{
				URL: "https://youtube.com/watch?v=updated-example",
			}

			updateSimulcastResponse, err := client.SimulcastStream.UpdateSpecificSimulcastOfStream(ctx, *streamID, *simulcastID, updateSimulcastRequest)
			if err != nil {
				log.Printf("Error updating simulcast: %v", err)
			} else {
				fmt.Println("Simulcast updated successfully!")
			}

			// 13. Delete Simulcast
			fmt.Printf("\n=== Deleting Simulcast: %s ===\n", *simulcastID)
			deleteSimulcastResponse, err := client.SimulcastStream.DeleteSimulcastOfStream(ctx, *streamID, *simulcastID)
			if err != nil {
				log.Printf("Error deleting simulcast: %v", err)
			} else {
				fmt.Println("Simulcast deleted successfully!")
			}
		}

		// 14. Disable Live Stream
		fmt.Printf("\n=== Disabling Live Stream: %s ===\n", *streamID)
		disableResponse, err := client.ManageLiveStream.DisableLiveStream(ctx, *streamID)
		if err != nil {
			log.Printf("Error disabling stream: %v", err)
		} else {
			fmt.Println("Live stream disabled successfully!")
		}

		// 15. Complete Live Stream
		fmt.Printf("\n=== Completing Live Stream: %s ===\n", *streamID)
		completeResponse, err := client.ManageLiveStream.CompleteLiveStream(ctx, *streamID)
		if err != nil {
			log.Printf("Error completing stream: %v", err)
		} else {
			fmt.Println("Live stream completed successfully!")
		}

		// 16. Delete Live Stream
		fmt.Printf("\n=== Deleting Live Stream: %s ===\n", *streamID)
		deleteResponse, err := client.ManageLiveStream.DeleteLiveStream(ctx, *streamID)
		if err != nil {
			log.Printf("Error deleting stream: %v", err)
		} else {
			fmt.Println("Live stream deleted successfully!")
		}
	}
}

// Helper functions to safely get values from pointers
func getStringValue(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func getInt64Value(ptr *int64) int64 {
	if ptr == nil {
		return 0
	}
	return *ptr
}
