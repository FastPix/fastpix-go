package tests

import (
	"context"
	"encoding/json"
	"testing"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
)

type SimulcastStreamTest struct {
	sdk *fastpixgo.FastPixSDK
}

func setupSimulcastStreamTest(t *testing.T) *SimulcastStreamTest {
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

	return &SimulcastStreamTest{
		sdk: sdk,
	}
}

func TestSimulcastStream(t *testing.T) {
	test := setupSimulcastStreamTest(t)
	ctx := context.Background()

	// First, get a list of livestreams to use as source
	resp, err := test.sdk.ManageLiveStream.GetAllStreams(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("Failed to get livestreams: %v", err)
	}
	if len(resp.GetStreamsResponse.Data) == 0 {
		t.Fatal("No livestreams available for testing")
	}

	// Use the first stream as source
	sourceStream := resp.GetStreamsResponse.Data[0]
	sourceStreamID := *sourceStream.StreamID

	t.Run("CreateSimulcastOfStream", func(t *testing.T) {
		// Create a simulcast stream using the working RTMP URL and stream key format
		createReq := &components.SimulcastRequest{
			URL:       stringPtr("rtmp://hyd01.contribute.live-video.net/app/"),
			StreamKey: stringPtr("live_1012464221_DuM8W004MoZYNxQEZ0czODgfHCFBhk"),
		}

		resp, err := test.sdk.SimulcastStream.CreateSimulcastOfStream(ctx, sourceStreamID, createReq)
		if err != nil {
			t.Fatalf("CreateSimulcastOfStream failed: %v", err)
		}
		if resp == nil || resp.SimulcastResponse == nil {
			t.Fatal("CreateSimulcastOfStream response is nil")
		}

		// Log the response
		if raw, err := json.MarshalIndent(resp.SimulcastResponse, "", "  "); err == nil {
			t.Logf("CreateSimulcastOfStream response: %s", string(raw))
		}

		// Store the simulcast stream ID for later tests
		if resp.SimulcastResponse.Data != nil && resp.SimulcastResponse.Data.SimulcastID != nil {
			simulcastStreamID := *resp.SimulcastResponse.Data.SimulcastID
			
			t.Run("GetSpecificSimulcastOfStream", func(t *testing.T) {
				resp, err := test.sdk.SimulcastStream.GetSpecificSimulcastOfStream(ctx, sourceStreamID, simulcastStreamID)
				if err != nil {
					t.Fatalf("GetSpecificSimulcastOfStream failed: %v", err)
				}
				if resp == nil || resp.SimulcastResponse == nil {
					t.Fatal("GetSpecificSimulcastOfStream response is nil")
				}

				if raw, err := json.MarshalIndent(resp.SimulcastResponse, "", "  "); err == nil {
					t.Logf("GetSpecificSimulcastOfStream response: %s", string(raw))
				}
			})

			t.Run("UpdateSpecificSimulcastOfStream", func(t *testing.T) {
				updateReq := &components.SimulcastUpdateRequest{
					IsEnabled: boolPtr(false),
				}

				resp, err := test.sdk.SimulcastStream.UpdateSpecificSimulcastOfStream(ctx, sourceStreamID, simulcastStreamID, updateReq)
				if err != nil {
					t.Fatalf("UpdateSpecificSimulcastOfStream failed: %v", err)
				}
				if resp == nil || resp.SimulcastUpdateResponse == nil {
					t.Fatal("UpdateSpecificSimulcastOfStream response is nil")
				}

				if raw, err := json.MarshalIndent(resp.SimulcastUpdateResponse, "", "  "); err == nil {
					t.Logf("UpdateSpecificSimulcastOfStream response: %s", string(raw))
				}
			})

			t.Run("DeleteSimulcastOfStream", func(t *testing.T) {
				resp, err := test.sdk.SimulcastStream.DeleteSimulcastOfStream(ctx, sourceStreamID, simulcastStreamID)
				if err != nil {
					t.Fatalf("DeleteSimulcastOfStream failed: %v", err)
				}
				if resp == nil || resp.SimulcastdeleteResponse == nil {
					t.Fatal("DeleteSimulcastOfStream response is nil")
				}

				if raw, err := json.MarshalIndent(resp.SimulcastdeleteResponse, "", "  "); err == nil {
					t.Logf("DeleteSimulcastOfStream response: %s", string(raw))
				}
			})
		}
	})
}

// Helper functions for creating pointers to primitive types
func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
} 