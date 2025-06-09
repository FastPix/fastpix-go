package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
)

type ManageLivestreamTest struct {
	sdk *fastpixgo.FastPixSDK
}

func setupManageLivestreamTest(t *testing.T) *ManageLivestreamTest {
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

	return &ManageLivestreamTest{
		sdk: sdk,
	}
}

func TestManageLivestream(t *testing.T) {
	test := setupManageLivestreamTest(t)

	ctx := context.Background()

	// List livestreams
	resp, err := test.sdk.ManageLiveStream.GetAllStreams(ctx, nil, nil, nil)
	if err != nil {
		fmt.Printf("GetAllStreams error: %T: %v\n", err, err)
		if resp == nil {
			fmt.Printf("Response is nil. No HTTP request to print.\n")
		} else if resp.HTTPMeta.Request != nil {
			req := resp.HTTPMeta.Request
			fmt.Printf("Request Method: %s\n", req.Method)
			fmt.Printf("Request URL: %s\n", req.URL.String())
			fmt.Printf("Request Headers:\n")
			for k, v := range req.Header {
				fmt.Printf("  %s: %v\n", k, v)
			}
		} else {
			fmt.Printf("HTTPMeta.Request is nil.\n")
		}
		t.Fatalf("GetAllStreams failed: %v", err)
	}

	if resp == nil || resp.GetStreamsResponse == nil {
		t.Fatal("GetAllStreams response is nil")
	}

	// Print full response details
	t.Logf("Full Response Details:")
	if resp.GetStreamsResponse.Success != nil {
		t.Logf("Success: %v", *resp.GetStreamsResponse.Success)
	}

	// Print the full raw JSON of resp.GetStreamsResponse.Data for debugging
	if raw, err := json.MarshalIndent(resp.GetStreamsResponse.Data, "", "  "); err == nil {
		t.Logf("Raw resp.GetStreamsResponse.Data: %s", string(raw))
	} else {
		t.Logf("Failed to marshal resp.GetStreamsResponse.Data: %v", err)
	}

	// Verify the response contains livestreams
	if len(resp.GetStreamsResponse.Data) == 0 {
		t.Fatal("No livestreams found in response")
	}

	t.Logf("Successfully listed livestreams")

	// Use the first stream for further tests
	stream := resp.GetStreamsResponse.Data[0]
	streamID := ""
	if stream.StreamID != nil {
		streamID = *stream.StreamID
	} else {
		t.Fatal("First stream has no StreamID")
	}

	t.Run("GetLiveStreamByID", func(t *testing.T) {
		resp, err := test.sdk.ManageLiveStream.GetLiveStreamByID(ctx, streamID)
		if err != nil {
			t.Fatalf("GetLiveStreamByID failed: %v", err)
		}
		if resp == nil || resp.LivestreamgetResponse == nil {
			t.Fatal("GetLiveStreamByID response is nil")
		}
		if raw, err := json.MarshalIndent(resp.LivestreamgetResponse, "", "  "); err == nil {
			t.Logf("GetLiveStreamByID response: %s", string(raw))
		}
	})

	t.Run("UpdateLiveStream", func(t *testing.T) {
		patch := &components.PatchLiveStreamRequest{
			ReconnectWindow: stream.ReconnectWindow, // no-op update, just for test
		}
		resp, err := test.sdk.ManageLiveStream.UpdateLiveStream(ctx, streamID, patch)
		if err != nil {
			t.Fatalf("UpdateLiveStream failed: %v", err)
		}
		if resp == nil || resp.PatchResponseDTO == nil {
			t.Fatal("UpdateLiveStream response is nil")
		}
		if raw, err := json.MarshalIndent(resp.PatchResponseDTO, "", "  "); err == nil {
			t.Logf("UpdateLiveStream response: %s", string(raw))
		}
	})

	t.Run("DeleteLiveStream", func(t *testing.T) {
		resp, err := test.sdk.ManageLiveStream.DeleteLiveStream(ctx, streamID)
		if err != nil {
			t.Fatalf("DeleteLiveStream failed: %v", err)
		}
		if resp == nil || resp.LiveStreamDeleteResponse == nil {
			t.Fatal("DeleteLiveStream response is nil")
		}
		if raw, err := json.MarshalIndent(resp.LiveStreamDeleteResponse, "", "  "); err == nil {
			t.Logf("DeleteLiveStream response: %s", string(raw))
		}
	})
} 