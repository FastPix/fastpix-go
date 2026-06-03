//go:build ignore

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
		printCreateStreamResult(createResponse)
	}

	fmt.Println("\n=== Listing All Live Streams ===")
	limit := int64(10)
	offset := int64(0)
	orderBy := components.SortOrderDesc

	streamsResponse, err := client.ManageLiveStream.GetAllStreams(ctx, &limit, &offset, &orderBy)
	if err != nil {
		log.Printf("Error listing streams: %v", err)
	} else {
		printStreamList(streamsResponse)
	}

	if streamsResponse.GetStreamsResponse == nil || len(streamsResponse.GetStreamsResponse.Data) == 0 {
		return
	}

	streamID := *streamsResponse.GetStreamsResponse.Data[0].StreamID
	manageStream(ctx, client, streamID)

func printCreateStreamResult(createResponse *operations.CreateNewStreamResponse) {
	if createResponse.LiveStreamResponseDTO == nil || createResponse.LiveStreamResponseDTO.Data == nil {
		fmt.Println("Live stream created but no data returned")
		return
	}

	data := createResponse.LiveStreamResponseDTO.Data
	if data.StreamID != nil {
		fmt.Printf("Live stream created successfully! Stream ID: %s\n", *data.StreamID)
	}
	fmt.Printf("RTMP URL: %s\n", getStringValue(data.RtmpURL))
	fmt.Printf("Stream Key: %s\n", getStringValue(data.StreamKey))
}

func printStreamList(streamsResponse *operations.GetAllStreamsResponse) {
	fmt.Printf("Found %d live streams:\n", len(streamsResponse.GetStreamsResponse.Data))
	for i, stream := range streamsResponse.GetStreamsResponse.Data {
		fmt.Printf("  %d. ID: %s, Name: %s, Status: %s\n",
			i+1, *stream.StreamID, getStringValue(stream.Name), getStringValue(stream.Status))
	}
}

func manageStream(ctx context.Context, client *fastpixgo.FastPix, streamID string) {
	getStreamDetails(ctx, client, streamID)
	updateStream(ctx, client, streamID)
	enableStream(ctx, client, streamID)
	getViewerCount(ctx, client, streamID)

	playbackID := managePlayback(ctx, client, streamID)
	if playbackID != "" {
		managePlaybackDetails(ctx, client, streamID, playbackID)
	}

	simulcastID := manageSimulcast(ctx, client, streamID)
	if simulcastID != "" {
		manageSimulcastDetails(ctx, client, streamID, simulcastID)
	}

	disableStream(ctx, client, streamID)
	completeStream(ctx, client, streamID)
	deleteStream(ctx, client, streamID)
}

func getStreamDetails(ctx context.Context, client *fastpixgo.FastPix, streamID string) {
	fmt.Printf("\n=== Getting Stream Details for ID: %s ===\n", streamID)

	streamResponse, err := client.ManageLiveStream.GetLiveStreamByID(ctx, streamID)
	if err != nil {
		log.Printf("Error getting stream details: %v", err)
		return
	}

	stream := streamResponse.LivestreamgetResponse.Data
	fmt.Printf("Name: %s\n", getStringValue(stream.Name))
	fmt.Printf("Status: %s\n", getStringValue(stream.Status))
	fmt.Printf("Created: %s\n", getStringValue(stream.CreatedAt))
	fmt.Printf("RTMP URL: %s\n", getStringValue(stream.RtmpURL))
}

func updateStream(ctx context.Context, client *fastpixgo.FastPix, streamID string) {
	fmt.Printf("\n=== Updating Live Stream: %s ===\n", streamID)

	updateRequest := components.PatchLiveStreamRequest{
		Metadata: map[string]string{
			"name":        "Updated Live Stream",
			"description": "Updated description for testing",
			"category":    "updated-gaming",
		},
	}

	_, err := client.ManageLiveStream.UpdateLiveStream(ctx, streamID, updateRequest)
	if err != nil {
		log.Printf("Error updating stream: %v", err)
	} else {
		fmt.Println("Live stream updated successfully!")
	}
}

func enableStream(ctx context.Context, client *fastpixgo.FastPix, streamID string) {
	fmt.Printf("\n=== Enabling Live Stream: %s ===\n", streamID)

	enableResponse, err := client.ManageLiveStream.EnableLiveStream(ctx, streamID)
	if err != nil {
		log.Printf("Error enabling stream: %v", err)
	} else {
		fmt.Println("Live stream enabled successfully!")
		fmt.Printf("Response: %+v\n", enableResponse)
	}
}

func getViewerCount(ctx context.Context, client *fastpixgo.FastPix, streamID string) {
	fmt.Printf("\n=== Getting Viewer Count for Stream: %s ===\n", streamID)

	viewerResponse, err := client.ManageLiveStream.GetLiveStreamViewerCountByID(ctx, streamID)
	if err != nil {
		log.Printf("Error getting viewer count: %v", err)
	} else {
		fmt.Printf("Current viewer count: %d\n", getInt64Value(viewerResponse.Object.Data.ViewerCount))
	}
}

func managePlayback(ctx context.Context, client *fastpixgo.FastPix, streamID string) string {
	fmt.Printf("\n=== Creating Playback ID for Stream: %s ===\n", streamID)

	playbackRequest := operations.CreatePlaybackIDOfStreamRequestBody{
		AccessPolicy: components.AccessPolicyPublic,
	}

	playbackResponse, err := client.LivePlayback.CreatePlaybackIDOfStream(ctx, streamID, &playbackRequest)
	if err != nil {
		log.Printf("Error creating playback ID: %v", err)
		return ""
	}

	fmt.Printf("Playback ID created successfully! ID: %s\n", *playbackResponse.Object.Data.PlaybackID)
	return *playbackResponse.Object.Data.PlaybackID
}

func managePlaybackDetails(ctx context.Context, client *fastpixgo.FastPix, streamID, playbackID string) {
	fmt.Printf("\n=== Getting Playback ID Details: %s ===\n", playbackID)

	playbackDetailsResponse, err := client.LivePlayback.GetLiveStreamPlaybackID(ctx, streamID, playbackID)
	if err != nil {
		log.Printf("Error getting playback ID details: %v", err)
	} else {
		fmt.Printf("Playback ID details retrieved successfully!\n")
		fmt.Printf("Access Policy: %s\n", getStringValue(playbackDetailsResponse.Object.Data.AccessPolicy))
	}

	deletePlayback(ctx, client, streamID, playbackID)
}

func deletePlayback(ctx context.Context, client *fastpixgo.FastPix, streamID, playbackID string) {
	fmt.Printf("\n=== Deleting Playback ID: %s ===\n", playbackID)

	deletePlaybackResponse, err := client.LivePlayback.DeletePlaybackIDOfStream(ctx, streamID, playbackID)
	if err != nil {
		log.Printf("Error deleting playback ID: %v", err)
	} else {
		fmt.Println("Playback ID deleted successfully!")
		fmt.Printf("Response: %+v\n", deletePlaybackResponse)
	}
}

func manageSimulcast(ctx context.Context, client *fastpixgo.FastPix, streamID string) string {
	fmt.Printf("\n=== Creating Simulcast for Stream: %s ===\n", streamID)

	simulcastRequest := components.CreateSimulcastRequest{
		Platform: components.SimulcastPlatformYoutube,
		URL:      "https://youtube.com/watch?v=example",
	}

	simulcastResponse, err := client.SimulcastStream.CreateSimulcastOfStream(ctx, streamID, simulcastRequest)
	if err != nil {
		log.Printf("Error creating simulcast: %v", err)
		return ""
	}

	fmt.Printf("Simulcast created successfully! ID: %s\n", *simulcastResponse.Object.Data.ID)
	return *simulcastResponse.Object.Data.ID
}

func manageSimulcastDetails(ctx context.Context, client *fastpixgo.FastPix, streamID, simulcastID string) {
	fmt.Printf("\n=== Getting Simulcast Details: %s ===\n", simulcastID)

	simulcastDetailsResponse, err := client.SimulcastStream.GetSpecificSimulcastOfStream(ctx, streamID, simulcastID)
	if err != nil {
		log.Printf("Error getting simulcast details: %v", err)
	} else {
		fmt.Printf("Simulcast details retrieved successfully!\n")
		fmt.Printf("Platform: %s\n", getStringValue(simulcastDetailsResponse.Object.Data.Platform))
	}

	updateSimulcast(ctx, client, streamID, simulcastID)
	deleteSimulcast(ctx, client, streamID, simulcastID)
}

func updateSimulcast(ctx context.Context, client *fastpixgo.FastPix, streamID, simulcastID string) {
	fmt.Printf("\n=== Updating Simulcast: %s ===\n", simulcastID)

	_, err := client.SimulcastStream.UpdateSpecificSimulcastOfStream(ctx, streamID, simulcastID, components.UpdateSimulcastRequest{
		URL: "https://youtube.com/watch?v=updated-example",
	})
	if err != nil {
		log.Printf("Error updating simulcast: %v", err)
	} else {
		fmt.Println("Simulcast updated successfully!")
	}
}

func deleteSimulcast(ctx context.Context, client *fastpixgo.FastPix, streamID, simulcastID string) {
	fmt.Printf("\n=== Deleting Simulcast: %s ===\n", simulcastID)

	_, err := client.SimulcastStream.DeleteSimulcastOfStream(ctx, streamID, simulcastID)
	if err != nil {
		log.Printf("Error deleting simulcast: %v", err)
	} else {
		fmt.Println("Simulcast deleted successfully!")
	}
}

func disableStream(ctx context.Context, client *fastpixgo.FastPix, streamID string) {
	fmt.Printf("\n=== Disabling Live Stream: %s ===\n", streamID)

	_, err := client.ManageLiveStream.DisableLiveStream(ctx, streamID)
	if err != nil {
		log.Printf("Error disabling stream: %v", err)
	} else {
		fmt.Println("Live stream disabled successfully!")
	}
}

func completeStream(ctx context.Context, client *fastpixgo.FastPix, streamID string) {
	fmt.Printf("\n=== Completing Live Stream: %s ===\n", streamID)

	_, err := client.ManageLiveStream.CompleteLiveStream(ctx, streamID)
	if err != nil {
		log.Printf("Error completing stream: %v", err)
	} else {
		fmt.Println("Live stream completed successfully!")
	}
}

func deleteStream(ctx context.Context, client *fastpixgo.FastPix, streamID string) {
	fmt.Printf("\n=== Deleting Live Stream: %s ===\n", streamID)

	_, err := client.ManageLiveStream.DeleteLiveStream(ctx, streamID)
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