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

const responseFormat = "Response: %+v\n"

func main() {
	ctx := context.Background()

	client := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")),
			Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")),
		}),
		fastpixgo.WithTimeout(30*time.Second),
	)

	mediaResponse, err := client.ManageVideos.List(ctx, nil, nil, nil)
	if err != nil {
		log.Printf("Error listing media: %v", err)
		return
	}

	if mediaResponse.Object == nil || len(mediaResponse.Object.Data) == 0 {
		fmt.Println("No media available for AI operations")
		return
	}

	mediaID := mediaResponse.Object.Data[0].ID
	fmt.Printf("Working with media: %s\n", *mediaID)

	runAIFeatures(ctx, client, *mediaID)
	runAdvancedAIFeatures(ctx, client, *mediaID)
	runBatchAIProcessing(ctx, client, mediaResponse)
	checkAIFeaturesStatus(ctx, client, *mediaID)
	runAIErrorHandling(ctx, client)

	fmt.Println("\n=== AI Features Demo Complete ===")
	fmt.Println("Note: AI features are processed asynchronously.")
	fmt.Println("Check the media details after some time to see the generated content.")
}

func runAIFeatures(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	generateSummary(ctx, client, mediaID, 200)
	generateChapters(ctx, client, mediaID)
	generateNamedEntities(ctx, client, mediaID)
	enableModeration(ctx, client, mediaID)
}

func generateSummary(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string, length int64) {
	fmt.Println("\n=== Generating Video Summary ===")
	req := operations.UpdateMediaSummaryRequestBody{
		Generate:      true,
		SummaryLength: fastpixgo.Int64(length),
	}
	resp, err := client.InVideoAI.GenerateSummary(ctx, mediaID, req)
	if err != nil {
		log.Printf("Error generating video summary: %v", err)
		return
	}
	fmt.Println("Video summary generation initiated successfully!")
	fmt.Printf(responseFormat, resp)
}

func generateChapters(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	fmt.Println("\n=== Generating Video Chapters ===")
	req := operations.UpdateMediaChaptersRequestBody{
		Chapters: fastpixgo.Bool(true),
	}
	resp, err := client.InVideoAIFeatures.UpdateChapters(ctx, mediaID, req)
	if err != nil {
		log.Printf("Error generating video chapters: %v", err)
		return
	}
	fmt.Println("Video chapters generation initiated successfully!")
	fmt.Printf(responseFormat, resp)
}

func generateNamedEntities(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	fmt.Println("\n=== Generating Named Entities ===")
	req := operations.UpdateMediaNamedEntitiesRequestBody{
		NamedEntities: true,
	}
	resp, err := client.InVideoAIFeatures.UpdateMediaNamedEntities(ctx, mediaID, req)
	if err != nil {
		log.Printf("Error generating named entities: %v", err)
		return
	}
	fmt.Println("Named entities generation initiated successfully!")
	fmt.Printf(responseFormat, resp)
}

func enableModeration(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	fmt.Println("\n=== Enabling Video Moderation ===")
	req := operations.UpdateMediaModerationRequestBody{
		Moderation: &operations.UpdateMediaModerationModeration{
			Type: components.MediaTypeVideo.ToPointer(),
		},
	}
	resp, err := client.InVideoAIFeatures.UpdateModeration(ctx, mediaID, req)
	if err != nil {
		log.Printf("Error enabling video moderation: %v", err)
		return
	}
	fmt.Println("Video moderation enabled successfully!")
	fmt.Printf(responseFormat, resp)
}

func runAdvancedAIFeatures(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	fmt.Println("\n=== Advanced AI Features ===")

	_, err := client.InVideoAI.GenerateSummary(ctx, mediaID, operations.UpdateMediaSummaryRequestBody{
		Generate:      true,
		SummaryLength: fastpixgo.Int64(250),
	})
	if err != nil {
		log.Printf("Error generating detailed summary: %v", err)
	} else {
		fmt.Println("Detailed summary generation initiated successfully!")
	}

	_, err = client.InVideoAIFeatures.UpdateChapters(ctx, mediaID, operations.UpdateMediaChaptersRequestBody{
		Chapters: fastpixgo.Bool(true),
	})
	if err != nil {
		log.Printf("Error generating detailed chapters: %v", err)
	} else {
		fmt.Println("Detailed chapters generation initiated successfully!")
	}
}

func runBatchAIProcessing(ctx context.Context, client *fastpixgo.Fastpixgo, mediaResponse *operations.ListMediaResponse) {
	fmt.Println("\n=== Batch AI Processing ===")

	if len(mediaResponse.Object.Data) <= 1 {
		return
	}

	for i, media := range mediaResponse.Object.Data[1:3] {
		fmt.Printf("\nProcessing media %d: %s\n", i+2, *media.ID)
		processBatchMediaItem(ctx, client, *media.ID)
	}
}

func processBatchMediaItem(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	_, err := client.InVideoAI.GenerateSummary(ctx, mediaID, operations.UpdateMediaSummaryRequestBody{
		Generate:      true,
		SummaryLength: fastpixgo.Int64(100),
	})
	if err != nil {
		log.Printf("Error generating summary for media %s: %v", mediaID, err)
	} else {
		fmt.Printf("Summary generation initiated for media %s\n", mediaID)
	}

	_, err = client.InVideoAIFeatures.UpdateChapters(ctx, mediaID, operations.UpdateMediaChaptersRequestBody{
		Chapters: fastpixgo.Bool(true),
	})
	if err != nil {
		log.Printf("Error generating chapters for media %s: %v", mediaID, err)
	} else {
		fmt.Printf("Chapters generation initiated for media %s\n", mediaID)
	}
}

func checkAIFeaturesStatus(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	fmt.Println("\n=== Checking AI Features Status ===")

	summaryResponse, err := client.ManageVideos.GetMediaSummary(ctx, mediaID)
	if err != nil {
		log.Printf("Error getting media summary: %v", err)
		return
	}

	fmt.Printf("Media ID: %s\n", mediaID)
	fmt.Printf(responseFormat, summaryResponse)
}

func runAIErrorHandling(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n=== AI Features Error Handling ===")

	_, err := client.InVideoAI.GenerateSummary(ctx, "non-existent-media-id", operations.UpdateMediaSummaryRequestBody{
		Generate:      true,
		SummaryLength: fastpixgo.Int64(100),
	})
	if err != nil {
		fmt.Printf("Expected error for non-existent media: %v\n", err)
	} else {
		fmt.Println("Unexpected success for non-existent media")
	}
}
