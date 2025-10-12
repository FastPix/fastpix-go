package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
)

func main() {
	// Initialize context
	ctx := context.Background()

	// Get credentials from environment variables
	username := os.Getenv("FASTPIX_USERNAME")
	password := os.Getenv("FASTPIX_PASSWORD")

	if username == "" || password == "" {
		log.Fatal("Please set FASTPIX_USERNAME and FASTPIX_PASSWORD environment variables")
	}

	// Initialize the FastPix SDK with authentication
	client := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(username),
			Password: fastpixgo.Pointer(password),
		}),
		fastpixgo.WithTimeout(30*time.Second),
	)

	fmt.Println("FastPix SDK initialized successfully!")
	fmt.Printf("SDK Version: %s\n", client.SDKVersion)

	// Example: List all media
	fmt.Println("\n=== Listing Media ===")
	mediaResponse, err := client.ManageVideos.ListMedia(ctx, nil, nil, nil)
	if err != nil {
		log.Printf("Error listing media: %v", err)
	} else {
		fmt.Printf("Found %d media items\n", len(mediaResponse.Object.Data))
	}

	// Example: List all live streams
	fmt.Println("\n=== Listing Live Streams ===")
	streamsResponse, err := client.ManageLiveStream.GetAllStreams(ctx, nil, nil, nil)
	if err != nil {
		log.Printf("Error listing streams: %v", err)
	} else {
		fmt.Printf("Found %d live streams\n", len(streamsResponse.GetStreamsResponse.Data))
	}

	// Example: List signing keys
	fmt.Println("\n=== Listing Signing Keys ===")
	keysResponse, err := client.SigningKeys.ListSigningKeys(ctx, nil, nil)
	if err != nil {
		log.Printf("Error listing signing keys: %v", err)
	} else {
		fmt.Printf("Found %d signing keys\n", len(keysResponse.GetAllSigningKeyResponse.Data))
	}
}
