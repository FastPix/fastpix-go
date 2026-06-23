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
	videoInput := components.PullVideoInput{
		Type: fastpixgo.Pointer("video"),
		URL:  fastpixgo.Pointer("https://sample-videos.com/zip/10/mp4/SampleVideo_1280x720_1mb.mp4"),
	}

	createRequest := components.CreateMediaRequest{
		Inputs:       []components.Input{components.CreateInputPullVideoInput(videoInput)},
		AccessPolicy: components.CreateMediaRequestAccessPolicyPublic.ToPointer(),
		Metadata: map[string]string{
			"title":       "Sample Video",
			"description": "A sample video for testing",
			"category":    "demo",
		},
	}

	createResponse, err := client.InputVideo.Create(ctx, createRequest)
	if err != nil {
		log.Printf("Error creating media: %v", err)
	} else {
		printCreateMediaResult(createResponse)
	}

	// 2. List all media
	fmt.Println("\n=== Listing All Media ===")
	limit := int64(10)
	offset := int64(0)
	orderBy := components.SortOrderDesc

	listResponse, err := client.ManageVideos.List(ctx, &limit, &offset, &orderBy)
	if err != nil {
		log.Printf("Error listing media: %v", err)
	} else {
		printMediaList(listResponse)
	}

	if listResponse != nil && listResponse.Object != nil && len(listResponse.Object.Data) > 0 {
		manageMedia(ctx, client, *listResponse.Object.Data[0].ID)
	}

	listUploads(ctx, client)
}

func printCreateMediaResult(createResponse *operations.CreateMediaResponse) {
	if createResponse.CreateMediaSuccessResponse != nil &&
		createResponse.CreateMediaSuccessResponse.Data != nil &&
		createResponse.CreateMediaSuccessResponse.Data.ID != nil {
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
		fmt.Printf("  %d. ID: %s, Title: %s\n", i+1, getStringValue(media.ID), optString(media.Title))
	}
}

func manageMedia(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	getMediaDetails(ctx, client, mediaID)
	updateMedia(ctx, client, mediaID)
	addAudioTrack(ctx, client, mediaID)
	generateSubtitleTrack(ctx, client, mediaID)
	getMediaClips(ctx, client, mediaID)
	updateSourceAccess(ctx, client, mediaID)
	updateMp4Support(ctx, client, mediaID)
	retrieveMediaInputInfo(ctx, client, mediaID)
	deleteMedia(ctx, client, mediaID)
}

func getMediaDetails(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	fmt.Printf("\n=== Getting Media Details for ID: %s ===\n", mediaID)

	mediaResponse, err := client.Videos.Get(ctx, mediaID)
	if err != nil {
		log.Printf("Error getting media: %v", err)
		return
	}

	if mediaResponse.Object == nil || mediaResponse.Object.Data == nil {
		fmt.Println("No media details returned")
		return
	}

	media := mediaResponse.Object.Data
	fmt.Printf("Title: %s\n", optString(media.Title))
	fmt.Printf("Duration: %s seconds\n", getStringValue(media.Duration))
	if media.Status != nil {
		fmt.Printf("Status: %s\n", string(*media.Status))
	}
	if media.CreatedAt != nil {
		fmt.Printf("Created: %s\n", media.CreatedAt.Format(time.RFC3339))
	}
}

func updateMedia(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	fmt.Printf("\n=== Updating Media: %s ===\n", mediaID)

	updateRequest := operations.UpdatedMediaRequestBody{
		Title: fastpixgo.Pointer("Updated Sample Video"),
		Metadata: map[string]string{
			"description": "Updated description for testing",
			"category":    "updated-demo",
			"updated":     "true",
		},
	}

	updateResponse, err := client.Videos.Update(ctx, mediaID, updateRequest)
	if err != nil {
		log.Printf("Error updating media: %v", err)
		return
	}

	fmt.Println("Media updated successfully!")
	if updateResponse.Object != nil && updateResponse.Object.Data != nil {
		fmt.Printf("Updated metadata: %+v\n", updateResponse.Object.Data.Metadata)
	}
}

func addAudioTrack(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	fmt.Printf("\n=== Adding Audio Track to Media: %s ===\n", mediaID)

	audioTrackRequest := operations.AddMediaTrackRequestBody{
		Tracks: components.AddTrackRequest{
			Type:         components.AddTrackRequestTypeAudio.ToPointer(),
			URL:          fastpixgo.Pointer("https://example.com/audio-track.mp3"),
			LanguageCode: fastpixgo.Pointer("en"),
			LanguageName: fastpixgo.Pointer("English"),
		},
	}

	audioResponse, err := client.Videos.AddMediaTrack(ctx, mediaID, audioTrackRequest)
	if err != nil {
		log.Printf("Error adding audio track: %v", err)
		return
	}

	if audioResponse.Object != nil && audioResponse.Object.Data != nil && audioResponse.Object.Data.ID != nil {
		fmt.Printf("Audio track added successfully! Track ID: %s\n", *audioResponse.Object.Data.ID)
	}
}

func generateSubtitleTrack(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	fmt.Printf("\n=== Generating Subtitle Track for Media: %s ===\n", mediaID)

	// A track ID is required to generate subtitles for an existing audio track.
	trackID := "audio-track-id"
	subtitleRequest := components.TrackSubtitlesGenerateRequest{
		LanguageName: fastpixgo.Pointer("English"),
	}

	subtitleResponse, err := client.ManageVideos.GenerateSubtitleTrack(ctx, mediaID, trackID, subtitleRequest)
	if err != nil {
		log.Printf("Error generating subtitle track: %v", err)
		return
	}

	if subtitleResponse.Object != nil && subtitleResponse.Object.Data != nil && subtitleResponse.Object.Data.ID != nil {
		fmt.Printf("Subtitle track generated successfully! Track ID: %s\n", *subtitleResponse.Object.Data.ID)
	}
}

func getMediaClips(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	fmt.Printf("\n=== Listing Clips for Media: %s ===\n", mediaID)

	clipsResponse, err := client.ManageVideos.GetMediaClips(ctx, mediaID, nil, nil, nil)
	if err != nil {
		log.Printf("Error listing media clips: %v", err)
		return
	}

	if clipsResponse.MediaClipResponse != nil {
		fmt.Printf("Found %d clips for this media\n", len(clipsResponse.MediaClipResponse.Data))
	}
}

func updateSourceAccess(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	fmt.Printf("\n=== Updating Source Access for Media: %s ===\n", mediaID)

	sourceAccessRequest := operations.UpdatedSourceAccessRequestBody{
		SourceAccess: true,
	}

	_, err := client.Videos.UpdateSourceAccess(ctx, mediaID, sourceAccessRequest)
	if err != nil {
		log.Printf("Error updating source access: %v", err)
		return
	}

	fmt.Println("Source access updated successfully!")
}

func updateMp4Support(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	fmt.Printf("\n=== Updating MP4 Support for Media: %s ===\n", mediaID)

	mp4SupportRequest := operations.UpdatedMp4SupportRequestBody{
		Mp4Support: operations.UpdatedMp4SupportMp4SupportCapped4k.ToPointer(),
	}

	_, err := client.Videos.UpdateMp4Support(ctx, mediaID, mp4SupportRequest)
	if err != nil {
		log.Printf("Error updating MP4 support: %v", err)
		return
	}

	fmt.Println("MP4 support updated successfully!")
}

func retrieveMediaInputInfo(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	fmt.Printf("\n=== Getting Input Info for Media: %s ===\n", mediaID)

	inputInfoResponse, err := client.ManageVideos.GetInputInfo(ctx, mediaID)
	if err != nil {
		log.Printf("Error getting media input info: %v", err)
		return
	}

	if inputInfoResponse.Object != nil {
		fmt.Printf("Input info retrieved successfully! Inputs: %+v\n", inputInfoResponse.Object.Data)
	}
}

func deleteMedia(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	fmt.Printf("\n=== Deleting Media: %s ===\n", mediaID)

	_, err := client.ManageVideos.Delete(ctx, mediaID)
	if err != nil {
		log.Printf("Error deleting media: %v", err)
		return
	}

	fmt.Println("Media deleted successfully!")
}

func listUploads(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n=== Listing Uploads ===")

	uploadsResponse, err := client.ManageVideos.ListUnusedUploadUrls(ctx, nil, nil, nil)
	if err != nil {
		log.Printf("Error listing uploads: %v", err)
		return
	}

	if uploadsResponse.Object != nil {
		fmt.Printf("Found %d uploads\n", len(uploadsResponse.Object.Data))
	}
}

// Helper functions to safely get values from pointers / optional-nullable fields.
func getStringValue(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func optString(v interface{ GetOrZero() (string, bool) }) string {
	val, _ := v.GetOrZero()
	return val
}
