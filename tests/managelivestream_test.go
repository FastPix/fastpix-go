package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/models/operations"
)

type ManageLivestreamTest struct {
	sdk *fastpixgo.Fastpixgo
}

func setupManageLivestreamTest(t *testing.T) *ManageLivestreamTest {
	RequireCredentials(t)
	livestreamServerURL, _, username, password := LoadConfig()
	t.Logf("Using livestream server URL: %s", livestreamServerURL)

	sdk := fastpixgo.New(
		fastpixgo.WithServerURL(livestreamServerURL),
		fastpixgo.WithSecurity(components.Security{
			Username: &username,
			Password: &password,
		}),
	)
	return &ManageLivestreamTest{sdk: sdk}
}

func TestManageLivestream(t *testing.T) {
	test := setupManageLivestreamTest(t)
	ctx := context.Background()

	resp, err := test.sdk.ManageLiveStream.List(ctx, nil, nil, nil)
	if err != nil {
		printListStreamsError(resp, err)
		t.Fatalf("GetAllStreams failed: %v", err)
	}

	if resp == nil || resp.GetStreamsResponse == nil {
		t.Fatal("GetAllStreams response is nil")
	}

	logListStreamsResponse(t, resp)

	if len(resp.GetStreamsResponse.Data) == 0 {
		t.Fatal("No livestreams found in response")
	}

	t.Logf("Successfully listed livestreams")

	streamID := extractStreamID(t, resp)
	stream := resp.GetStreamsResponse.Data[0]

	t.Run("GetLiveStreamByID", func(t *testing.T) {
		testGetLiveStreamByID(t, ctx, test, streamID)
	})

	t.Run("UpdateLiveStream", func(t *testing.T) {
		testUpdateLiveStream(t, ctx, test, streamID, stream)
	})

	t.Run("DeleteLiveStream", func(t *testing.T) {
		testDeleteLiveStream(t, ctx, test, streamID)
	})
}

func printListStreamsError(resp *operations.GetAllStreamsResponse, err error) {
	fmt.Printf("GetAllStreams error: %T: %v\n", err, err)
	if resp == nil {
		fmt.Printf("Response is nil. No HTTP request to print.\n")
		return
	}
	if resp.HTTPMeta.Request == nil {
		fmt.Printf("HTTPMeta.Request is nil.\n")
		return
	}
	req := resp.HTTPMeta.Request
	fmt.Printf("Request Method: %s\n", req.Method)
	fmt.Printf("Request URL: %s\n", req.URL.String())
	fmt.Printf("Request Headers:\n")
	for k, v := range req.Header {
		fmt.Printf("  %s: %v\n", k, v)
	}
}

func logListStreamsResponse(t *testing.T, resp *operations.GetAllStreamsResponse) {
	t.Logf("Full Response Details:")
	if resp.GetStreamsResponse.Success != nil {
		t.Logf("Success: %v", *resp.GetStreamsResponse.Success)
	}
	if raw, err := json.MarshalIndent(resp.GetStreamsResponse.Data, "", "  "); err == nil {
		t.Logf("Raw resp.GetStreamsResponse.Data: %s", string(raw))
	} else {
		t.Logf("Failed to marshal resp.GetStreamsResponse.Data: %v", err)
	}
}

func extractStreamID(t *testing.T, resp *operations.GetAllStreamsResponse) string {
	stream := resp.GetStreamsResponse.Data[0]
	if stream.StreamID == nil {
		t.Fatal("First stream has no StreamID")
	}
	return *stream.StreamID
}

func testGetLiveStreamByID(t *testing.T, ctx context.Context, test *ManageLivestreamTest, streamID string) {
	resp, err := test.sdk.LiveStreams.GetByID(ctx, streamID)
	if err != nil {
		t.Fatalf("GetLiveStreamByID failed: %v", err)
	}
	if resp == nil || resp.LivestreamgetResponse == nil {
		t.Fatal("GetLiveStreamByID response is nil")
	}
	if raw, err := json.MarshalIndent(resp.LivestreamgetResponse, "", "  "); err == nil {
		t.Logf("GetLiveStreamByID response: %s", string(raw))
	}
}

func testUpdateLiveStream(t *testing.T, ctx context.Context, test *ManageLivestreamTest, streamID string, stream components.GetCreateLiveStreamResponseDTO) {
	patch := components.PatchLiveStreamRequest{
		ReconnectWindow: stream.ReconnectWindow,
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
}

func testDeleteLiveStream(t *testing.T, ctx context.Context, test *ManageLivestreamTest, streamID string) {
	resp, err := test.sdk.LiveStreams.Delete(ctx, streamID)
	if err != nil {
		t.Fatalf("DeleteLiveStream failed: %v", err)
	}
	if resp == nil || resp.LiveStreamDeleteResponse == nil {
		t.Fatal("DeleteLiveStream response is nil")
	}
	if raw, err := json.MarshalIndent(resp.LiveStreamDeleteResponse, "", "  "); err == nil {
		t.Logf("DeleteLiveStream response: %s", string(raw))
	}
}
