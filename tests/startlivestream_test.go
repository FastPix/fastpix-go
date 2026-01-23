package tests

import (
	"context"
	"testing"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
)

type StartLiveStreamTest struct {
	sdk *fastpixgo.FastPixSDK
}

func setupStartLiveStreamTest(t *testing.T) *StartLiveStreamTest {
	livestreamServerURL, _, username, password := LoadConfig()
	
	// Log the livestream server URL
	t.Logf("Using livestream server URL: %s", livestreamServerURL)

	sdk := fastpixgo.New(
		fastpixgo.WithServerURL(livestreamServerURL),
		fastpixgo.WithSecurity(components.Security{
			Username: &username,
			Password: &password,
		}),
	)

	return &StartLiveStreamTest{
		sdk: sdk,
	}
}

func TestCreateNewStream(t *testing.T) {
	test := setupStartLiveStreamTest(t)

	ctx := context.Background()
	reconnectWindow := int64(60)
	request := &components.CreateLiveStreamRequest{
		PlaybackSettings: components.PlaybackSettings{
			// Add playback settings if needed
		},
		InputMediaSettings: components.InputMediaSettings{
			MaxResolution:   components.CreateLiveStreamRequestMaxResolutionOneThousandAndEightyp.ToPointer(),
			ReconnectWindow: &reconnectWindow,
			MediaPolicy:     components.MediaPolicyPublic.ToPointer(),
		},
	}

	resp, err := test.sdk.StartLiveStream.CreateNewStream(ctx, request)
	if err != nil {
		t.Fatalf("CreateNewStream failed: %v", err)
	}

	if resp == nil || resp.LiveStreamResponseDTO == nil {
		t.Fatal("CreateNewStream response is nil")
	}

	t.Logf("Successfully created a new stream")
} 