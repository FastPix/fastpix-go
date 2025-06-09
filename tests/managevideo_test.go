package tests

import (
	"context"
	"testing"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/models/operations"
)

type ManageVideosTest struct {
	sdk *fastpixgo.FastPixSDK
}

func setupTest(t *testing.T) *ManageVideosTest {
	_, serverURL, username, password := LoadConfig()

	sdk := fastpixgo.New(
		fastpixgo.WithServerURL(serverURL),
		fastpixgo.WithSecurity(components.Security{
			Username: &username,
			Password: &password,
		}),
	)

	return &ManageVideosTest{
		sdk: sdk,
	}
}

func TestComprehensiveMediaOperations(t *testing.T) {
	test := setupTest(t)

	ctx := context.Background()
	limit := int64(20)
	offset := int64(1)
	orderBy := operations.ListMediaOrderByDesc

	// Step 1: List media to get a media ID
	listResp, err := test.sdk.ManageVideos.ListMedia(ctx, &limit, &offset, &orderBy)
	if err != nil {
		t.Fatalf("ListMedia failed: %v", err)
	}

	if listResp == nil || listResp.Object == nil || listResp.Object.Data == nil || len(listResp.Object.Data) == 0 {
		t.Fatal("No media found to test with")
	}

	mediaID := *listResp.Object.Data[0].ID // Dereference the media ID

	// Step 2: Get media details
	getResp, err := test.sdk.ManageVideos.GetMedia(ctx, mediaID)
	if err != nil {
		t.Fatalf("GetMedia failed: %v", err)
	}

	if getResp == nil || getResp.Object == nil {
		t.Fatal("GetMedia response is nil")
	}

	// Step 3: Update media metadata
	updateBody := &operations.UpdatedMediaRequestBody{
		Metadata: &operations.UpdatedMediaMetadata{},
	}

	updateResp, err := test.sdk.ManageVideos.UpdatedMedia(ctx, mediaID, updateBody)
	if err != nil {
		t.Fatalf("UpdatedMedia failed: %v", err)
	}

	if updateResp == nil || updateResp.Object == nil {
		t.Fatal("UpdatedMedia response is nil")
	}

	// Step 4: Delete media
	deleteResp, err := test.sdk.ManageVideos.DeleteMedia(ctx, mediaID)
	if err != nil {
		t.Fatalf("DeleteMedia failed: %v", err)
	}

	if deleteResp == nil || deleteResp.Object == nil {
		t.Fatal("DeleteMedia response is nil")
	}
} 