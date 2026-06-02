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
	livestreamSDK *fastpixgo.Fastpixgo
	mediaSDK      *fastpixgo.Fastpixgo
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
	resp, err := pt.livestreamSDK.ManageLiveStream.List(context.Background(), nil, nil, nil)
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

	createReq := components.CreateMediaRequest{
		Inputs: []components.Input{
			components.CreateInputPullVideoInput(components.PullVideoInput{
				Type: fastpixgo.Pointer("video"),
				URL:  fastpixgo.Pointer("https://storage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4"),
			}),
		},
		AccessPolicy: components.CreateMediaRequestAccessPolicyPublic.ToPointer(),
	}

	resp, err := pt.mediaSDK.InputVideo.Create(ctx, createReq)
	if err != nil {
		pt.t.Fatalf("Failed to create media asset: %v", err)
	}
	if resp == nil || resp.CreateMediaSuccessResponse == nil ||
		resp.CreateMediaSuccessResponse.Data == nil ||
		resp.CreateMediaSuccessResponse.Data.ID == nil {
		pt.t.Fatal("Failed to create media asset: response is nil or missing ID")
	}

	mediaID := *resp.CreateMediaSuccessResponse.Data.ID
	pt.t.Logf("Created media asset with ID: %s", mediaID)

	return pt.waitForMediaReady(ctx, mediaID)
}

func (pt *PlaybackTest) waitForMediaReady(ctx context.Context, mediaID string) string {
	for i := 0; i < 45; i++ {
		time.Sleep(4 * time.Second)
		if pt.isMediaReady(ctx, mediaID) {
			pt.t.Logf("Media asset is ready")
			return mediaID
		}
	}
	pt.t.Fatal("Media asset failed to process within timeout")
	return ""
}

func (pt *PlaybackTest) isMediaReady(ctx context.Context, mediaID string) bool {
	getResp, err := pt.mediaSDK.Videos.Get(ctx, mediaID)
	if err != nil || getResp == nil || getResp.Object == nil ||
		getResp.Object.Data == nil || getResp.Object.Data.Status == nil {
		return false
	}
	status := strings.ToLower(string(*getResp.Object.Data.Status))
	pt.t.Logf("Media asset status: %s", string(*getResp.Object.Data.Status))
	return status == "ready"
}

func TestLiveStreamPlayback(t *testing.T) {
	test := setupPlaybackTest(t)
	ctx := context.Background()

	streamID := test.getStreamID()

	t.Run("CreatePlaybackID", func(t *testing.T) {
		testCreateLiveStreamPlaybackID(t, ctx, test, streamID)
	})
}

func testCreateLiveStreamPlaybackID(t *testing.T, ctx context.Context, test *PlaybackTest, streamID string) {
	createReq := &components.PlaybackIDRequest{
		AccessPolicy: components.BasicAccessPolicyPublic.ToPointer(),
	}

	resp, err := test.livestreamSDK.LivePlayback.Create(ctx, streamID, *createReq)
	if err != nil {
		t.Fatalf("CreatePlaybackIDOfStream failed: %v", err)
	}
	if resp == nil || resp.PlaybackIDSuccessResponse == nil {
		t.Fatal("CreatePlaybackIDOfStream response is nil")
	}

	if raw, err := json.MarshalIndent(resp.PlaybackIDSuccessResponse, "", "  "); err == nil {
		t.Logf("CreatePlaybackIDOfStream response: %s", string(raw))
	}

	if resp.PlaybackIDSuccessResponse.Data == nil || resp.PlaybackIDSuccessResponse.Data.ID == nil {
		return
	}

	playbackID := *resp.PlaybackIDSuccessResponse.Data.ID

	t.Run("GetPlaybackID", func(t *testing.T) {
		testGetLiveStreamPlaybackID(t, ctx, test, streamID, playbackID)
	})

	t.Run("DeletePlaybackID", func(t *testing.T) {
		testDeleteLiveStreamPlaybackID(t, ctx, test, streamID, playbackID)
	})
}

func testGetLiveStreamPlaybackID(t *testing.T, ctx context.Context, test *PlaybackTest, streamID, playbackID string) {
	resp, err := test.livestreamSDK.LivePlayback.GetPlaybackIDDetails(ctx, streamID, playbackID)
	if err != nil {
		t.Fatalf("GetLiveStreamPlaybackID failed: %v", err)
	}
	if resp == nil || resp.PlaybackIDSuccessResponse == nil {
		t.Fatal("GetLiveStreamPlaybackID response is nil")
	}
	if raw, err := json.MarshalIndent(resp.PlaybackIDSuccessResponse, "", "  "); err == nil {
		t.Logf("GetLiveStreamPlaybackID response: %s", string(raw))
	}
}

func testDeleteLiveStreamPlaybackID(t *testing.T, ctx context.Context, test *PlaybackTest, streamID, playbackID string) {
	resp, err := test.livestreamSDK.LivePlayback.DeletePlaybackID(ctx, streamID, playbackID)
	if err != nil {
		t.Fatalf("DeletePlaybackIDOfStream failed: %v", err)
	}
	if resp == nil || resp.LiveStreamDeleteResponse == nil {
		t.Fatal("DeletePlaybackIDOfStream response is nil")
	}
	if raw, err := json.MarshalIndent(resp.LiveStreamDeleteResponse, "", "  "); err == nil {
		t.Logf("DeletePlaybackIDOfStream response: %s", string(raw))
	}
}

func TestMediaPlayback(t *testing.T) {
	test := setupPlaybackTest(t)
	ctx := context.Background()

	mediaID := test.createMediaAsset()
	if mediaID == "" {
		t.Fatal("Failed to create media asset")
	}

	t.Run("CreatePlaybackID", func(t *testing.T) {
		testCreateMediaPlaybackID(t, ctx, test, mediaID)
	})
}

func testCreateMediaPlaybackID(t *testing.T, ctx context.Context, test *PlaybackTest, mediaID string) {
	createReq := &operations.CreateMediaPlaybackIDRequestBody{
		AccessPolicy: components.AccessPolicyPublic,
	}

	resp, err := test.mediaSDK.Playback.Create(ctx, mediaID, createReq)
	if err != nil {
		t.Fatalf("CreateMediaPlaybackID failed: %v", err)
	}
	if resp == nil || resp.Object == nil {
		t.Fatal("CreateMediaPlaybackID response is nil")
	}

	if raw, err := json.MarshalIndent(resp.Object, "", "  "); err == nil {
		t.Logf("CreateMediaPlaybackID response: %s", string(raw))
	}

	if resp.Object.Data == nil || resp.Object.Data.ID == nil {
		return
	}

	playbackID := *resp.Object.Data.ID

	t.Run("DeletePlaybackID", func(t *testing.T) {
		testDeleteMediaPlaybackID(t, ctx, test, mediaID, playbackID)
	})
}

func testDeleteMediaPlaybackID(t *testing.T, ctx context.Context, test *PlaybackTest, mediaID, playbackID string) {
	resp, err := test.mediaSDK.Playback.Delete(ctx, mediaID, playbackID)
	if err != nil {
		t.Fatalf("DeleteMediaPlaybackID failed: %v", err)
	}
	if resp == nil || resp.Object == nil {
		t.Fatal("DeleteMediaPlaybackID response is nil")
	}
	if raw, err := json.MarshalIndent(resp.Object, "", "  "); err == nil {
		t.Logf("DeleteMediaPlaybackID response: %s", string(raw))
	}
}