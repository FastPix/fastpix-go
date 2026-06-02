package tests

import (
	"context"
	"encoding/json"
	"testing"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
)

type SimulcastStreamTest struct {
	sdk *fastpixgo.Fastpixgo
}

func setupSimulcastStreamTest(t *testing.T) *SimulcastStreamTest {
	livestreamServerURL, _, username, password := LoadConfig()
	t.Logf("Using livestream server URL: %s", livestreamServerURL)

	sdk := fastpixgo.New(
		fastpixgo.WithServerURL(livestreamServerURL),
		fastpixgo.WithSecurity(components.Security{
			Username: &username,
			Password: &password,
		}),
	)
	return &SimulcastStreamTest{sdk: sdk}
}

func TestSimulcastStream(t *testing.T) {
	test := setupSimulcastStreamTest(t)
	ctx := context.Background()

	resp, err := test.sdk.ManageLiveStream.List(ctx, nil, nil, nil)
	if err != nil {
		t.Fatalf("Failed to get livestreams: %v", err)
	}
	if len(resp.GetStreamsResponse.Data) == 0 {
		t.Fatal("No livestreams available for testing")
	}

	sourceStreamID := *resp.GetStreamsResponse.Data[0].StreamID

	t.Run("CreateSimulcastOfStream", func(t *testing.T) {
		testCreateSimulcastOfStream(t, ctx, test, sourceStreamID)
	})
}

func testCreateSimulcastOfStream(t *testing.T, ctx context.Context, test *SimulcastStreamTest, sourceStreamID string) {
	createReq := &components.SimulcastRequest{
		URL:       stringPtr("rtmp://hyd01.contribute.live-video.net/app/"),
		StreamKey: stringPtr("live_1012464221_DuM8W004MoZYNxQEZ0czODgfHCFBhk"),
	}

	resp, err := test.sdk.SimulcastStreams.Create(ctx, sourceStreamID, *createReq)
	if err != nil {
		t.Fatalf("CreateSimulcastOfStream failed: %v", err)
	}
	if resp == nil || resp.SimulcastResponse == nil {
		t.Fatal("CreateSimulcastOfStream response is nil")
	}

	if raw, err := json.MarshalIndent(resp.SimulcastResponse, "", "  "); err == nil {
		t.Logf("CreateSimulcastOfStream response: %s", string(raw))
	}

	if resp.SimulcastResponse.Data == nil || resp.SimulcastResponse.Data.SimulcastID == nil {
		return
	}

	simulcastStreamID := *resp.SimulcastResponse.Data.SimulcastID

	t.Run("GetSpecificSimulcastOfStream", func(t *testing.T) {
		testGetSpecificSimulcastOfStream(t, ctx, test, sourceStreamID, simulcastStreamID)
	})

	t.Run("UpdateSpecificSimulcastOfStream", func(t *testing.T) {
		testUpdateSpecificSimulcastOfStream(t, ctx, test, sourceStreamID, simulcastStreamID)
	})

	t.Run("DeleteSimulcastOfStream", func(t *testing.T) {
		testDeleteSimulcastOfStream(t, ctx, test, sourceStreamID, simulcastStreamID)
	})
}

func testGetSpecificSimulcastOfStream(t *testing.T, ctx context.Context, test *SimulcastStreamTest, sourceStreamID, simulcastStreamID string) {
	resp, err := test.sdk.SimulcastStreams.GetSpecific(ctx, sourceStreamID, simulcastStreamID)
	if err != nil {
		t.Fatalf("GetSpecificSimulcastOfStream failed: %v", err)
	}
	if resp == nil || resp.SimulcastResponse == nil {
		t.Fatal("GetSpecificSimulcastOfStream response is nil")
	}
	if raw, err := json.MarshalIndent(resp.SimulcastResponse, "", "  "); err == nil {
		t.Logf("GetSpecificSimulcastOfStream response: %s", string(raw))
	}
}

func testUpdateSpecificSimulcastOfStream(t *testing.T, ctx context.Context, test *SimulcastStreamTest, sourceStreamID, simulcastStreamID string) {
	updateReq := components.SimulcastUpdateRequest{
		IsEnabled: boolPtr(false),
	}
	resp, err := test.sdk.SimulcastStreams.Update(ctx, sourceStreamID, simulcastStreamID, updateReq)
	if err != nil {
		t.Fatalf("UpdateSpecificSimulcastOfStream failed: %v", err)
	}
	if resp == nil || resp.SimulcastUpdateResponse == nil {
		t.Fatal("UpdateSpecificSimulcastOfStream response is nil")
	}
	if raw, err := json.MarshalIndent(resp.SimulcastUpdateResponse, "", "  "); err == nil {
		t.Logf("UpdateSpecificSimulcastOfStream response: %s", string(raw))
	}
}

func testDeleteSimulcastOfStream(t *testing.T, ctx context.Context, test *SimulcastStreamTest, sourceStreamID, simulcastStreamID string) {
	resp, err := test.sdk.SimulcastStream.Delete(ctx, sourceStreamID, simulcastStreamID)
	if err != nil {
		t.Fatalf("DeleteSimulcastOfStream failed: %v", err)
	}
	if resp == nil || resp.SimulcastdeleteResponse == nil {
		t.Fatal("DeleteSimulcastOfStream response is nil")
	}
	if raw, err := json.MarshalIndent(resp.SimulcastdeleteResponse, "", "  "); err == nil {
		t.Logf("DeleteSimulcastOfStream response: %s", string(raw))
	}
}

func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}