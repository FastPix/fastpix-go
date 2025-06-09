package tests

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
	"time"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/models/operations"
)

type PlaybackTest struct {
	livestreamSDK *fastpixgo.FastPixSDK
	mediaSDK      *fastpixgo.FastPixSDK
	t             *testing.T
}

func setupPlaybackTest(t *testing.T) *PlaybackTest {
	livestreamServerURL, baseServerUrl, username, password := LoadConfig()
	t.Logf("Using livestream server URL: %s", livestreamServerURL)
	t.Logf("Using base server URL: %s", baseServerUrl)

	livestreamSDK := fastpixgo.New(
		fastpixgo.WithServerURL(livestreamServerURL),
		fastpixgo.WithSecurity(components.Security{
			Username: &username,
			Password: &password,
		}),
	)

	mediaSDK := fastpixgo.New(
		fastpixgo.WithServerURL(baseServerUrl),
		fastpixgo.WithSecurity(components.Security{
			Username: &username,
			Password: &password,
		}),
	)

	return &PlaybackTest{
		livestreamSDK: livestreamSDK,
		mediaSDK:      mediaSDK,
		t:             t,
	}
}

func (pt *PlaybackTest) getStreamID() string {
	resp, err := pt.livestreamSDK.ManageLiveStream.GetAllStreams(context.Background(), nil, nil, nil)
	if err != nil {
		pt.t.Fatalf("Failed to get livestreams: %v", err)
	}
	if len(resp.GetStreamsResponse.Data) == 0 {
		pt.t.Fatal("No livestreams available for testing")
	}

	streamID := *resp.GetStreamsResponse.Data[0].StreamID
	pt.t.Logf("Using stream ID: %s", streamID)
	return streamID
}

func (pt *PlaybackTest) createMediaAsset() string {
	ctx := context.Background()

	// Create a media asset from a URL
	createReq := &components.CreateMediaRequest{
		Inputs: []components.Input{
			{
				Type: components.InputTypeVideoInput,
				VideoInput: &components.VideoInput{
					Type: "video",
					URL: "https://storage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4",
				},
			},
		},
		AccessPolicy: components.CreateMediaRequestAccessPolicyPublic,
	}

	resp, err := pt.mediaSDK.InputVideo.CreateMedia(ctx, createReq)
	if err != nil {
		pt.t.Fatalf("Failed to create media asset: %v", err)
	}
	if resp == nil || resp.Object == nil || resp.Object.Data == nil || resp.Object.Data.ID == nil {
		pt.t.Fatal("Failed to create media asset: response is nil or missing ID")
	}

	mediaID := *resp.Object.Data.ID
	pt.t.Logf("Created media asset with ID: %s", mediaID)

	// Wait for the media asset to be processed
	for i := 0; i < 45; i++ {
		time.Sleep(4 * time.Second)
		getResp, err := pt.mediaSDK.ManageVideos.GetMedia(ctx, mediaID)
		if err == nil && getResp != nil && getResp.Object != nil && getResp.Object.Data != nil {
			if getResp.Object.Data.Status != nil {
				pt.t.Logf("Media asset status: %s", *getResp.Object.Data.Status)
				if strings.ToLower(*getResp.Object.Data.Status) == "ready" {
					pt.t.Logf("Media asset is ready")
					return mediaID
				}
			}
		}
	}

	pt.t.Fatal("Media asset failed to process within timeout")
	return ""
}

func TestLiveStreamPlayback(t *testing.T) {
	test := setupPlaybackTest(t)
	ctx := context.Background()

	streamID := test.getStreamID()

	t.Run("CreatePlaybackID", func(t *testing.T) {
		// Create a playback ID for the live stream
		createReq := &components.PlaybackIDRequest{
			AccessPolicy: components.PlaybackIDRequestAccessPolicyPublic.ToPointer(),
		}

		resp, err := test.livestreamSDK.Playback.CreatePlaybackIDOfStream(ctx, streamID, createReq)
		if err != nil {
			t.Fatalf("CreatePlaybackIDOfStream failed: %v", err)
		}
		if resp == nil || resp.PlaybackIDResponse == nil {
			t.Fatal("CreatePlaybackIDOfStream response is nil")
		}

		// Log the response
		if raw, err := json.MarshalIndent(resp.PlaybackIDResponse, "", "  "); err == nil {
			t.Logf("CreatePlaybackIDOfStream response: %s", string(raw))
		}

		// Store the playback ID for later tests
		if resp.PlaybackIDResponse.Data != nil && resp.PlaybackIDResponse.Data.ID != nil {
			playbackID := *resp.PlaybackIDResponse.Data.ID

			t.Run("GetPlaybackID", func(t *testing.T) {
				resp, err := test.livestreamSDK.Playback.GetLiveStreamPlaybackID(ctx, streamID, playbackID)
				if err != nil {
					t.Fatalf("GetLiveStreamPlaybackID failed: %v", err)
				}
				if resp == nil || resp.PlaybackIDResponse == nil {
					t.Fatal("GetLiveStreamPlaybackID response is nil")
				}

				if raw, err := json.MarshalIndent(resp.PlaybackIDResponse, "", "  "); err == nil {
					t.Logf("GetLiveStreamPlaybackID response: %s", string(raw))
				}
			})

			t.Run("DeletePlaybackID", func(t *testing.T) {
				resp, err := test.livestreamSDK.Playback.DeletePlaybackIDOfStream(ctx, streamID, playbackID)
				if err != nil {
					t.Fatalf("DeletePlaybackIDOfStream failed: %v", err)
				}
				if resp == nil || resp.LiveStreamDeleteResponse == nil {
					t.Fatal("DeletePlaybackIDOfStream response is nil")
				}

				if raw, err := json.MarshalIndent(resp.LiveStreamDeleteResponse, "", "  "); err == nil {
					t.Logf("DeletePlaybackIDOfStream response: %s", string(raw))
				}
			})
		}
	})
}

func TestMediaPlayback(t *testing.T) {
	test := setupPlaybackTest(t)
	ctx := context.Background()

	// Create a media asset first
	mediaID := test.createMediaAsset()
	if mediaID == "" {
		t.Fatal("Failed to create media asset")
	}

	t.Run("CreatePlaybackID", func(t *testing.T) {
		// Create a playback ID for the media
		createReq := &operations.CreateMediaPlaybackIDRequestBody{
			AccessPolicy: operations.CreateMediaPlaybackIDAccessPolicyPublic,
		}

		resp, err := test.mediaSDK.Playback.CreateMediaPlaybackID(ctx, mediaID, createReq)
		if err != nil {
			t.Fatalf("CreateMediaPlaybackID failed: %v", err)
		}
		if resp == nil || resp.Object == nil {
			t.Fatal("CreateMediaPlaybackID response is nil")
		}

		// Log the response
		if raw, err := json.MarshalIndent(resp.Object, "", "  "); err == nil {
			t.Logf("CreateMediaPlaybackID response: %s", string(raw))
		}

		// Store the playback ID for later tests
		if resp.Object.Data != nil && len(resp.Object.Data.PlaybackIds) > 0 {
			playbackID := *resp.Object.Data.PlaybackIds[0].ID

			t.Run("DeletePlaybackID", func(t *testing.T) {
				resp, err := test.mediaSDK.Playback.DeleteMediaPlaybackID(ctx, mediaID, playbackID)
				if err != nil {
					t.Fatalf("DeleteMediaPlaybackID failed: %v", err)
				}
				if resp == nil || resp.Object == nil {
					t.Fatal("DeleteMediaPlaybackID response is nil")
				}

				if raw, err := json.MarshalIndent(resp.Object, "", "  "); err == nil {
					t.Logf("DeleteMediaPlaybackID response: %s", string(raw))
				}
			})
		}
	})
} 