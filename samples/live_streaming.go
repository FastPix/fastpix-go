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
	reconnectWindow := int64(60)
	createStreamRequest := components.CreateLiveStreamRequest{
		PlaybackSettings: components.PlaybackSettings{
			AccessPolicy: components.BasicAccessPolicyPublic.ToPointer(),
		},
		InputMediaSettings: components.InputMediaSettings{
			MaxResolution:   components.CreateLiveStreamRequestMaxResolutionOneThousandAndEightyp.ToPointer(),
			ReconnectWindow: &reconnectWindow,
			MediaPolicy:     components.BasicAccessPolicyPublic.ToPointer(),
		},
	}
	createResponse, err := client.StartLiveStream.Create(ctx, createStreamRequest)
	if err != nil {
		log.Printf("Error creating live stream: %v", err)
	} else {
		printCreateStreamResult(createResponse.LiveStreamResponseDTO)
	}

	// 2. List All Live Streams
	fmt.Println("\n=== Listing All Live Streams ===")
	limit := int64(10)
	offset := int64(0)
	orderBy := operations.OrderByDesc

	streamsResponse, err := client.ManageLiveStream.List(ctx, &limit, &offset, &orderBy)
	if err != nil {
		log.Printf("Error listing streams: %v", err)
		return
	}
	if streamsResponse.GetStreamsResponse == nil || len(streamsResponse.GetStreamsResponse.Data) == 0 {
		fmt.Println("No live streams found")
		return
	}
	printStreamList(streamsResponse.GetStreamsResponse.Data)

	streamID := *streamsResponse.GetStreamsResponse.Data[0].StreamID
	manageStream(ctx, client, streamID)
}

func printCreateStreamResult(dto *components.LiveStreamResponseDTO) {
	if dto == nil || dto.Data == nil {
		fmt.Println("Live stream created but no data returned")
		return
	}

	data := dto.Data
	if data.StreamID != nil {
		fmt.Printf("Live stream created successfully! Stream ID: %s\n", *data.StreamID)
	}
	fmt.Printf("Stream Key: %s\n", getStringValue(data.StreamKey))
	fmt.Printf("Status: %s\n", getStringValue(data.Status))
}

func printStreamList(streams []components.GetCreateLiveStreamResponseDTO) {
	fmt.Printf("Found %d live streams:\n", len(streams))
	for i, stream := range streams {
		fmt.Printf("  %d. ID: %s, Status: %s\n",
			i+1, getStringValue(stream.StreamID), getStringValue(stream.Status))
	}
}

func manageStream(ctx context.Context, client *fastpixgo.Fastpixgo, streamID string) {
	getStreamDetails(ctx, client, streamID)
	updateStream(ctx, client, streamID)
	enableStream(ctx, client, streamID)

	playbackID := managePlayback(ctx, client, streamID)
	if playbackID != "" {
		managePlaybackDetails(ctx, client, streamID, playbackID)
	}

	simulcastID := manageSimulcast(ctx, client, streamID)
	if simulcastID != "" {
		manageSimulcastDetails(ctx, client, streamID, simulcastID)
	}

	completeStream(ctx, client, streamID)
	deleteStream(ctx, client, streamID)
}

func getStreamDetails(ctx context.Context, client *fastpixgo.Fastpixgo, streamID string) {
	fmt.Printf("\n=== Getting Stream Details for ID: %s ===\n", streamID)

	streamResponse, err := client.LiveStreams.GetByID(ctx, streamID)
	if err != nil {
		log.Printf("Error getting stream details: %v", err)
		return
	}
	if streamResponse.LivestreamgetResponse == nil || streamResponse.LivestreamgetResponse.Data == nil {
		fmt.Println("No stream data returned")
		return
	}

	stream := streamResponse.LivestreamgetResponse.Data
	fmt.Printf("Stream ID: %s\n", getStringValue(stream.StreamID))
	fmt.Printf("Status: %s\n", getStringValue(stream.Status))
	if stream.CreatedAt != nil {
		fmt.Printf("Created: %s\n", stream.CreatedAt.Format(time.RFC3339))
	}
	fmt.Printf("Max Resolution: %s\n", getStringValue(stream.MaxResolution))
}

func updateStream(ctx context.Context, client *fastpixgo.Fastpixgo, streamID string) {
	fmt.Printf("\n=== Updating Live Stream: %s ===\n", streamID)

	reconnectWindow := int64(120)
	updateRequest := components.PatchLiveStreamRequest{
		ReconnectWindow: &reconnectWindow,
		Metadata: map[string]string{
			"name":     "Updated Live Stream",
			"category": "updated-gaming",
		},
	}

	_, err := client.ManageLiveStream.UpdateLiveStream(ctx, streamID, updateRequest)
	if err != nil {
		log.Printf("Error updating stream: %v", err)
	} else {
		fmt.Println("Live stream updated successfully!")
	}
}

func enableStream(ctx context.Context, client *fastpixgo.Fastpixgo, streamID string) {
	fmt.Printf("\n=== Enabling Live Stream: %s ===\n", streamID)

	_, err := client.ManageLiveStream.Enable(ctx, streamID)
	if err != nil {
		log.Printf("Error enabling stream: %v", err)
	} else {
		fmt.Println("Live stream enabled successfully!")
	}
}

func managePlayback(ctx context.Context, client *fastpixgo.Fastpixgo, streamID string) string {
	fmt.Printf("\n=== Creating Playback ID for Stream: %s ===\n", streamID)

	playbackRequest := components.PlaybackIDRequest{
		AccessPolicy: components.BasicAccessPolicyPublic.ToPointer(),
	}

	playbackResponse, err := client.LivePlayback.Create(ctx, streamID, playbackRequest)
	if err != nil {
		log.Printf("Error creating playback ID: %v", err)
		return ""
	}
	if playbackResponse.PlaybackIDSuccessResponse == nil ||
		playbackResponse.PlaybackIDSuccessResponse.Data == nil ||
		playbackResponse.PlaybackIDSuccessResponse.Data.ID == nil {
		return ""
	}

	playbackID := *playbackResponse.PlaybackIDSuccessResponse.Data.ID
	fmt.Printf("Playback ID created successfully! ID: %s\n", playbackID)
	return playbackID
}

func managePlaybackDetails(ctx context.Context, client *fastpixgo.Fastpixgo, streamID, playbackID string) {
	fmt.Printf("\n=== Getting Playback ID Details: %s ===\n", playbackID)

	playbackDetailsResponse, err := client.LivePlayback.GetPlaybackIDDetails(ctx, streamID, playbackID)
	if err != nil {
		log.Printf("Error getting playback ID details: %v", err)
	} else if playbackDetailsResponse.PlaybackIDSuccessResponse != nil &&
		playbackDetailsResponse.PlaybackIDSuccessResponse.Data != nil {
		fmt.Println("Playback ID details retrieved successfully!")
		fmt.Printf("Access Policy: %s\n", getStringValue(playbackDetailsResponse.PlaybackIDSuccessResponse.Data.AccessPolicy))
	}

	deletePlayback(ctx, client, streamID, playbackID)
}

func deletePlayback(ctx context.Context, client *fastpixgo.Fastpixgo, streamID, playbackID string) {
	fmt.Printf("\n=== Deleting Playback ID: %s ===\n", playbackID)

	_, err := client.LivePlayback.DeletePlaybackID(ctx, streamID, playbackID)
	if err != nil {
		log.Printf("Error deleting playback ID: %v", err)
	} else {
		fmt.Println("Playback ID deleted successfully!")
	}
}

func manageSimulcast(ctx context.Context, client *fastpixgo.Fastpixgo, streamID string) string {
	fmt.Printf("\n=== Creating Simulcast for Stream: %s ===\n", streamID)

	simulcastRequest := components.SimulcastRequest{
		URL:       fastpixgo.Pointer("rtmp://example.contribute.live-video.net/app/"),
		StreamKey: fastpixgo.Pointer("live_example_streamkey"),
	}

	simulcastResponse, err := client.SimulcastStreams.Create(ctx, streamID, simulcastRequest)
	if err != nil {
		log.Printf("Error creating simulcast: %v", err)
		return ""
	}
	if simulcastResponse.SimulcastResponse == nil ||
		simulcastResponse.SimulcastResponse.Data == nil ||
		simulcastResponse.SimulcastResponse.Data.SimulcastID == nil {
		return ""
	}

	simulcastID := *simulcastResponse.SimulcastResponse.Data.SimulcastID
	fmt.Printf("Simulcast created successfully! ID: %s\n", simulcastID)
	return simulcastID
}

func manageSimulcastDetails(ctx context.Context, client *fastpixgo.Fastpixgo, streamID, simulcastID string) {
	fmt.Printf("\n=== Getting Simulcast Details: %s ===\n", simulcastID)

	simulcastDetailsResponse, err := client.SimulcastStreams.GetSpecific(ctx, streamID, simulcastID)
	if err != nil {
		log.Printf("Error getting simulcast details: %v", err)
	} else if simulcastDetailsResponse.SimulcastResponse != nil &&
		simulcastDetailsResponse.SimulcastResponse.Data != nil {
		fmt.Println("Simulcast details retrieved successfully!")
		fmt.Printf("URL: %s\n", getStringValue(simulcastDetailsResponse.SimulcastResponse.Data.URL))
	}

	updateSimulcast(ctx, client, streamID, simulcastID)
	deleteSimulcast(ctx, client, streamID, simulcastID)
}

func updateSimulcast(ctx context.Context, client *fastpixgo.Fastpixgo, streamID, simulcastID string) {
	fmt.Printf("\n=== Updating Simulcast: %s ===\n", simulcastID)

	_, err := client.SimulcastStreams.Update(ctx, streamID, simulcastID, components.SimulcastUpdateRequest{
		IsEnabled: fastpixgo.Pointer(false),
	})
	if err != nil {
		log.Printf("Error updating simulcast: %v", err)
	} else {
		fmt.Println("Simulcast updated successfully!")
	}
}

func deleteSimulcast(ctx context.Context, client *fastpixgo.Fastpixgo, streamID, simulcastID string) {
	fmt.Printf("\n=== Deleting Simulcast: %s ===\n", simulcastID)

	_, err := client.SimulcastStream.Delete(ctx, streamID, simulcastID)
	if err != nil {
		log.Printf("Error deleting simulcast: %v", err)
	} else {
		fmt.Println("Simulcast deleted successfully!")
	}
}

func completeStream(ctx context.Context, client *fastpixgo.Fastpixgo, streamID string) {
	fmt.Printf("\n=== Completing Live Stream: %s ===\n", streamID)

	_, err := client.ManageLiveStream.Complete(ctx, streamID)
	if err != nil {
		log.Printf("Error completing stream: %v", err)
	} else {
		fmt.Println("Live stream completed successfully!")
	}
}

func deleteStream(ctx context.Context, client *fastpixgo.Fastpixgo, streamID string) {
	fmt.Printf("\n=== Deleting Live Stream: %s ===\n", streamID)

	_, err := client.LiveStreams.Delete(ctx, streamID)
	if err != nil {
		log.Printf("Error deleting stream: %v", err)
	} else {
		fmt.Println("Live stream deleted successfully!")
	}
}

// Helper to safely dereference a string pointer.
func getStringValue(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}
