package tests

import (
	"context"
	"encoding/json"
	"testing"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/models/operations"
)

type InputVideoTest struct {
	sdk *fastpixgo.FastPixSDK
}

func setupInputVideoTest(t *testing.T) *InputVideoTest {
	_, serverURL, username, password := LoadConfig()

	sdk := fastpixgo.New(
		fastpixgo.WithServerURL(serverURL),
		fastpixgo.WithSecurity(components.Security{
			Username: &username,
			Password: &password,
		}),
	)

	return &InputVideoTest{
		sdk: sdk,
	}
}

func TestCreateMediaFromURL(t *testing.T) {
	test := setupInputVideoTest(t)

	ctx := context.Background()

	// Create media from a public URL
	videoInput := components.VideoInput{
		Type: "video",
		URL:  "https://example.com/sample-video.mp4", // Replace with a valid public video URL
	}
	input := components.CreateInputVideoInput(videoInput)

	request := &components.CreateMediaRequest{
		Inputs:        []components.Input{input},
		AccessPolicy:  components.CreateMediaRequestAccessPolicyPublic,
		MaxResolution: components.CreateMediaRequestMaxResolution("1080p").ToPointer(),
	}

	resp, err := test.sdk.InputVideo.CreateMedia(ctx, request)
	if err != nil {
		t.Fatalf("CreateMedia failed: %v", err)
	}

	if resp == nil || resp.Object == nil || resp.Object.Data == nil {
		t.Fatal("CreateMedia response is nil")
	}

	// Verify the response contains an ID
	if resp.Object.Data.ID == nil {
		t.Fatal("Media ID is nil in response")
	}

	t.Logf("Successfully created media with ID: %s", *resp.Object.Data.ID)
}

func TestDirectUploadVideoMedia(t *testing.T) {
	test := setupInputVideoTest(t)

	ctx := context.Background()

	// Request direct upload URL
	request := &operations.DirectUploadVideoMediaRequest{
		CorsOrigin: "*", // Allow uploads from any origin
		PushMediaSettings: &operations.PushMediaSettings{
			AccessPolicy:  operations.DirectUploadVideoMediaAccessPolicyPublic,
			MaxResolution: operations.MaxResolution("1080p").ToPointer(),
		},
	}

	resp, err := test.sdk.InputVideo.DirectUploadVideoMedia(ctx, request)
	if err != nil {
		t.Fatalf("DirectUploadVideoMedia failed: %v", err)
	}

	if resp == nil || resp.Object == nil || resp.Object.Data == nil {
		t.Fatal("DirectUploadVideoMedia response is nil")
	}

	// Print full response details
	t.Logf("Full Response Details:")
	if resp.Object.Success != nil {
		t.Logf("Success: %v", *resp.Object.Success)
	}
	if resp.Object.Data.ID != nil {
		t.Logf("Upload ID: %s", *resp.Object.Data.ID)
	}
	if resp.Object.Data.MediaID != nil {
		t.Logf("Media ID: %s", *resp.Object.Data.MediaID)
	}
	if resp.Object.Data.Status != nil {
		t.Logf("Status: %s", *resp.Object.Data.Status)
	}
	if resp.Object.Data.URL != nil {
		t.Logf("Upload URL: %s", *resp.Object.Data.URL)
	}
	if resp.Object.Data.Timeout != nil {
		t.Logf("Timeout: %v", *resp.Object.Data.Timeout)
	}
	if resp.Object.Data.CorsOrigin != nil {
		t.Logf("CORS Origin: %s", *resp.Object.Data.CorsOrigin)
	}

	// Print the full raw JSON of resp.Object.Data for debugging
	if raw, err := json.MarshalIndent(resp.Object.Data, "", "  "); err == nil {
		t.Logf("Raw resp.Object.Data: %s", string(raw))
	} else {
		t.Logf("Failed to marshal resp.Object.Data: %v", err)
	}

	// Verify the response contains upload URL
	if resp.Object.Data.URL == nil {
		t.Fatal("Upload URL is nil in response")
	}

	t.Logf("Successfully got upload URL: %s", *resp.Object.Data.URL)
} 