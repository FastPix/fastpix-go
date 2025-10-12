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

	// Get media for AI operations
	mediaResponse, err := client.ManageVideos.ListMedia(ctx, nil, nil, nil)
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

	// 1. Generate Video Summary
	fmt.Println("\n=== Generating Video Summary ===")
	summaryRequest := operations.UpdateMediaSummaryRequestBody{
		Generate:      true,
		SummaryLength: fastpixgo.Int64(200), // Summary length in words
	}

	summaryResponse, err := client.InVideoAIFeatures.UpdateMediaSummary(ctx, *mediaID, summaryRequest)
	if err != nil {
		log.Printf("Error generating video summary: %v", err)
	} else {
		fmt.Println("Video summary generation initiated successfully!")
		fmt.Printf("Response: %+v\n", summaryResponse)
	}

	// 2. Generate Video Chapters
	fmt.Println("\n=== Generating Video Chapters ===")
	chaptersRequest := operations.UpdateMediaChaptersRequestBody{
		Chapters: true,
	}

	chaptersResponse, err := client.InVideoAIFeatures.UpdateMediaChapters(ctx, *mediaID, chaptersRequest)
	if err != nil {
		log.Printf("Error generating video chapters: %v", err)
	} else {
		fmt.Println("Video chapters generation initiated successfully!")
		fmt.Printf("Response: %+v\n", chaptersResponse)
	}

	// 3. Generate Named Entities
	fmt.Println("\n=== Generating Named Entities ===")
	entitiesRequest := operations.UpdateMediaNamedEntitiesRequestBody{
		NamedEntities: true,
	}

	entitiesResponse, err := client.InVideoAIFeatures.UpdateMediaNamedEntities(ctx, *mediaID, entitiesRequest)
	if err != nil {
		log.Printf("Error generating named entities: %v", err)
	} else {
		fmt.Println("Named entities generation initiated successfully!")
		fmt.Printf("Response: %+v\n", entitiesResponse)
	}

	// 4. Enable Video Moderation
	fmt.Println("\n=== Enabling Video Moderation ===")
	moderationRequest := operations.UpdateMediaModerationRequestBody{
		Moderation: &operations.UpdateMediaModerationModeration{
			Type: components.MediaTypeVideo.ToPointer(),
		},
	}

	moderationResponse, err := client.InVideoAIFeatures.UpdateMediaModeration(ctx, *mediaID, moderationRequest)
	if err != nil {
		log.Printf("Error enabling video moderation: %v", err)
	} else {
		fmt.Println("Video moderation enabled successfully!")
		fmt.Printf("Response: %+v\n", moderationResponse)
	}

	// 5. Advanced AI Features with Custom Parameters
	fmt.Println("\n=== Advanced AI Features ===")

	// Generate detailed summary with specific parameters
	detailedSummaryRequest := operations.UpdateMediaSummaryRequestBody{
		Generate:      true,
		SummaryLength: fastpixgo.Int64(500), // Longer summary
	}

	detailedSummaryResponse, err := client.InVideoAIFeatures.UpdateMediaSummary(ctx, *mediaID, detailedSummaryRequest)
	if err != nil {
		log.Printf("Error generating detailed summary: %v", err)
	} else {
		fmt.Println("Detailed summary generation initiated successfully!")
	}

	// Generate chapters with specific settings
	detailedChaptersRequest := operations.UpdateMediaChaptersRequestBody{
		Chapters: true,
	}

	detailedChaptersResponse, err := client.InVideoAIFeatures.UpdateMediaChapters(ctx, *mediaID, detailedChaptersRequest)
	if err != nil {
		log.Printf("Error generating detailed chapters: %v", err)
	} else {
		fmt.Println("Detailed chapters generation initiated successfully!")
	}

	// 6. Batch AI Processing
	fmt.Println("\n=== Batch AI Processing ===")

	// Process multiple media items with AI features
	if len(mediaResponse.Object.Data) > 1 {
		for i, media := range mediaResponse.Object.Data[1:3] { // Process up to 2 more media items
			fmt.Printf("\nProcessing media %d: %s\n", i+2, *media.ID)

			// Generate summary for each media
			batchSummaryRequest := operations.UpdateMediaSummaryRequestBody{
				Generate:      true,
				SummaryLength: fastpixgo.Int64(100),
			}

			_, err := client.InVideoAIFeatures.UpdateMediaSummary(ctx, *media.ID, batchSummaryRequest)
			if err != nil {
				log.Printf("Error generating summary for media %s: %v", *media.ID, err)
			} else {
				fmt.Printf("Summary generation initiated for media %s\n", *media.ID)
			}

			// Generate chapters for each media
			batchChaptersRequest := operations.UpdateMediaChaptersRequestBody{
				Chapters: true,
			}

			_, err = client.InVideoAIFeatures.UpdateMediaChapters(ctx, *media.ID, batchChaptersRequest)
			if err != nil {
				log.Printf("Error generating chapters for media %s: %v", *media.ID, err)
			} else {
				fmt.Printf("Chapters generation initiated for media %s\n", *media.ID)
			}
		}
	}

	// 7. AI Features Status Check
	fmt.Println("\n=== Checking AI Features Status ===")

	// Get media details to check AI features status
	mediaDetailsResponse, err := client.ManageVideos.GetMedia(ctx, *mediaID)
	if err != nil {
		log.Printf("Error getting media details: %v", err)
	} else {
		media := mediaDetailsResponse.Object.Data
		fmt.Printf("Media Title: %s\n", getStringValue(media.Title))
		fmt.Printf("Media Duration: %d seconds\n", getInt64Value(media.Duration))
		fmt.Printf("Media Status: %s\n", getStringValue(media.Status))

		// Check if AI features are available in metadata
		if media.Metadata != nil {
			fmt.Println("Media Metadata:")
			for key, value := range media.Metadata {
				fmt.Printf("  %s: %s\n", key, value)
			}
		}
	}

	// 8. Error Handling for AI Features
	fmt.Println("\n=== AI Features Error Handling ===")

	// Try to process a non-existent media ID to demonstrate error handling
	fakeMediaID := "non-existent-media-id"

	fakeSummaryRequest := operations.UpdateMediaSummaryRequestBody{
		Generate:      true,
		SummaryLength: fastpixgo.Int64(100),
	}

	_, err = client.InVideoAIFeatures.UpdateMediaSummary(ctx, fakeMediaID, fakeSummaryRequest)
	if err != nil {
		fmt.Printf("Expected error for non-existent media: %v\n", err)
	} else {
		fmt.Println("Unexpected success for non-existent media")
	}

	fmt.Println("\n=== AI Features Demo Complete ===")
	fmt.Println("Note: AI features are processed asynchronously.")
	fmt.Println("Check the media details after some time to see the generated content.")
}
