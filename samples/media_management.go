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

	// 1. Create Media from URL
	fmt.Println("=== Creating Media from URL ===")
	createRequest := components.CreateMediaRequest{
		Inputs: []components.Input{
			{
				VideoInput: &components.VideoInput{
					Type: "url",
					URL:  "https://sample-videos.com/zip/10/mp4/SampleVideo_1280x720_1mb.mp4",
				},
				Type: components.InputTypeVideoInput,
			},
		},
		AccessPolicy: components.CreateMediaRequestAccessPolicyPublic,
		Metadata: map[string]string{
			"title":       "Sample Video",
			"description": "A sample video for testing",
			"category":    "demo",
		},
	}

    createResponse, err := client.InputVideo.CreateMedia(ctx, createRequest)
	if err != nil {
		log.Printf("Error creating media: %v", err)
	} else {
		printCreateMediaResult(createResponse)
	}

	fmt.Println("\n=== Listing All Media ===")
	limit := int64(10)
	offset := int64(0)
	orderBy := components.SortOrderDesc

	listResponse, err := client.ManageVideos.ListMedia(ctx, &limit, &offset, &orderBy)
	if err != nil {
		log.Printf("Error listing media: %v", err)
	} else {
		printMediaList(listResponse)
	}

	if listResponse.Object != nil && len(listResponse.Object.Data) > 0 {
		manageMedia(ctx, client, *listResponse.Object.Data[0].ID)
	}

	listUploads(ctx, client)
	listLiveClips(ctx, client)
}

func printCreateMediaResult(createResponse *operations.CreateMediaResponse) {
	if createResponse.CreateMediaSuccessResponse != nil && createResponse.CreateMediaSuccessResponse.Data.ID != nil {
		fmt.Printf("Media created successfully! ID: %s\n", *createResponse.CreateMediaSuccessResponse.Data.ID)
		return
	}
	fmt.Println("Media created but no ID returned")
}

func printMediaList(listResponse *operations.ListMediaResponse) {
	if listResponse.Object == nil {
		fmt.Println("No media data returned")
		return
	}
	fmt.Printf("Found %d media items:\n", len(listResponse.Object.Data))
	for i, media := range listResponse.Object.Data {
		fmt.Printf("  %d. ID: %s, Title: %s\n", i+1, getStringValue(media.ID), getStringValue(media.Title))
	}
}

func manageMedia(ctx context.Context, client *fastpixgo.FastPix, mediaID string) {
	getMediaDetails(ctx, client, mediaID)
	updateMedia(ctx, client, mediaID)
	addAudioTrack(ctx, client, mediaID)
	generateSubtitleTrack(ctx, client, mediaID)
	getMediaClips(ctx, client, mediaID)
	updateSourceAccess(ctx, client, mediaID)
	updateMp4Support(ctx, client, mediaID)
	retrieveMediaInputInfo(ctx, client, mediaID)
}

func getMediaDetails(ctx context.Context, client *fastpixgo.FastPix, mediaID string) {
	fmt.Printf("\n=== Getting Media Details for ID: %s ===\n", mediaID)

	mediaResponse, err := client.ManageVideos.GetMedia(ctx, mediaID)
	if err != nil {
		log.Printf("Error getting media: %v", err)
		return
	}

	media := mediaResponse.Object.Data
	fmt.Printf("Title: %s\n", getStringValue(media.Title))
	fmt.Printf("Duration: %d seconds\n", getInt64Value(media.Duration))
	fmt.Printf("Status: %s\n", getStringValue(media.Status))
	fmt.Printf("Created: %s\n", getStringValue(media.CreatedAt))
}

func updateMedia(ctx context.Context, client *fastpixgo.FastPix, mediaID string) {
	fmt.Printf("\n=== Updating Media: %s ===\n", mediaID)

	updateRequest := operations.UpdatedMediaRequestBody{
		Metadata: map[string]string{
			"title":       "Updated Sample Video",
			"description": "Updated description for testing",
			"category":    "updated-demo",
			"updated":     "true",
		},
	}

	updateResponse, err := client.ManageVideos.UpdatedMedia(ctx, mediaID, updateRequest)
	if err != nil {
		log.Printf("Error updating media: %v", err)
		return
	}

	fmt.Println("Media updated successfully!")
	fmt.Printf("Updated metadata: %+v\n", updateResponse.Object.Data.Metadata)
}

func addAudioTrack(ctx context.Context, client *fastpixgo.FastPix, mediaID string) {
	fmt.Printf("\n=== Adding Audio Track to Media: %s ===\n", mediaID)

	audioTrackRequest := operations.AddMediaTrackRequestBody{
		Type:     components.TrackTypeAudio,
		Language: "en",
		URL:      "https://example.com/audio-track.mp3",
	}

	audioResponse, err := client.ManageVideos.AddMediaTrack(ctx, mediaID, audioTrackRequest)
	if err != nil {
		log.Printf("Error adding audio track: %v", err)
		return
	}

	fmt.Printf("Audio track added successfully! Track ID: %s\n", *audioResponse.Object.Data.ID)
}

func generateSubtitleTrack(ctx context.Context, client *fastpixgo.FastPix, mediaID string) {
	fmt.Printf("\n=== Generating Subtitle Track for Media: %s ===\n", mediaID)

	subtitleRequest := operations.GenerateSubtitleTrackRequestBody{
		Language: "en",
	}

	subtitleResponse, err := client.ManageVideos.GenerateSubtitleTrack(ctx, mediaID, subtitleRequest)
	if err != nil {
		log.Printf("Error generating subtitle track: %v", err)
		return
	}

	fmt.Printf("Subtitle track generated successfully! Track ID: %s\n", *subtitleResponse.Object.Data.ID)
}

func getMediaClips(ctx context.Context, client *fastpixgo.FastPix, mediaID string) {
	fmt.Printf("\n=== Listing Clips for Media: %s ===\n", mediaID)

	clipsResponse, err := client.ManageVideos.GetMediaClips(ctx, mediaID)
	if err != nil {
		log.Printf("Error listing media clips: %v", err)
		return
	}

	fmt.Printf("Found %d clips for this media\n", len(clipsResponse.Object.Data))
}

func updateSourceAccess(ctx context.Context, client *fastpixgo.FastPix, mediaID string) {
	fmt.Printf("\n=== Updating Source Access for Media: %s ===\n", mediaID)

	sourceAccessRequest := operations.UpdatedSourceAccessRequestBody{
		SourceAccess: components.SourceAccessPublic,
	}

	_, err := client.ManageVideos.UpdatedSourceAccess(ctx, mediaID, sourceAccessRequest)
	if err != nil {
		log.Printf("Error updating source access: %v", err)
		return
	}

	fmt.Println("Source access updated successfully!")
}

func updateMp4Support(ctx context.Context, client *fastpixgo.FastPix, mediaID string) {
	fmt.Printf("\n=== Updating MP4 Support for Media: %s ===\n", mediaID)

	mp4SupportRequest := operations.UpdatedMp4SupportRequestBody{
		Mp4Support: true,
	}

	_, err := client.ManageVideos.UpdatedMp4Support(ctx, mediaID, mp4SupportRequest)
	if err != nil {
		log.Printf("Error updating MP4 support: %v", err)
		return
	}

	fmt.Println("MP4 support updated successfully!")
}

func retrieveMediaInputInfo(ctx context.Context, client *fastpixgo.FastPix, mediaID string) {
	fmt.Printf("\n=== Getting Input Info for Media: %s ===\n", mediaID)

	inputInfoResponse, err := client.ManageVideos.RetrieveMediaInputInfo(ctx, mediaID)
	if err != nil {
		log.Printf("Error getting media input info: %v", err)
		return
	}

	fmt.Printf("Input info retrieved successfully! Inputs: %+v\n", inputInfoResponse.Object.Data)
}

func listUploads(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n=== Listing Uploads ===")

	uploadsResponse, err := client.ManageVideos.ListUploads(ctx, nil, nil, nil)
	if err != nil {
		log.Printf("Error listing uploads: %v", err)
		return
	}

	fmt.Printf("Found %d uploads\n", len(uploadsResponse.Object.Data))
}

func listLiveClips(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n=== Listing Live Clips ===")

	liveClipsResponse, err := client.ManageVideos.ListLiveClips(ctx, nil, nil, nil)
	if err != nil {
		log.Printf("Error listing live clips: %v", err)
		return
	}

	fmt.Printf("Found %d live clips\n", len(liveClipsResponse.Object.Data))
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