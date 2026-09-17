// Create a signed direct-upload URL, then PUT your file to it.
//
// Run:
//
//	export FASTPIX_USERNAME="your-access-token"
//	export FASTPIX_PASSWORD="your-secret-key"
//	go run ./examples/create-upload
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/models/operations"
)

func main() {
	ctx := context.Background()

	username := os.Getenv("FASTPIX_USERNAME")
	password := os.Getenv("FASTPIX_PASSWORD")
	if username == "" || password == "" {
		log.Fatal("Please set FASTPIX_USERNAME and FASTPIX_PASSWORD environment variables")
	}

	client := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(username),
			Password: fastpixgo.Pointer(password),
		}),
	)

	// Create a direct upload. corsOrigin "*" allows a browser upload from any origin.
	resp, err := client.InputVideo.DirectUploadMedia(ctx, &operations.DirectUploadVideoMediaRequest{
		CorsOrigin: fastpixgo.Pointer("*"),
		PushMediaSettings: &operations.PushMediaSettings{
			AccessPolicy: operations.DirectUploadVideoMediaAccessPolicyPublic.ToPointer(),
			Metadata:     map[string]string{"key1": "value1"},
		},
	})
	if err != nil {
		log.Fatalf("create upload failed: %v", err)
	}
	if resp.Object == nil {
		log.Fatal("empty response from FastPix")
	}

	upload := resp.Object.Data
	fmt.Printf("uploadId: %s\n", deref(upload.UploadID))
	// Upload your file with an HTTP PUT to this URL, e.g. `curl -X PUT --upload-file video.mp4 "<url>"`.
	fmt.Printf("PUT your file to: %s\n", deref(upload.URL))

	pretty, _ := json.MarshalIndent(resp.Object, "", "  ")
	fmt.Println(string(pretty))
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
