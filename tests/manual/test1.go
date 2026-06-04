package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
)

func main() {
	ctx := context.Background()

	s := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer("1b92c0d6-5548-4642-b13e-4bb7d77dbaf4"), // ← replace with real credentials
			Password: fastpixgo.Pointer("ff32012b-ec02-40ca-b0d4-711d81537e73"),  // ← replace with real credentials
		}),
	)

	// your-video-id   : FastPix Video ID returned from upload/create API
	// your-language-id: Track ID for subtitle generation
	res, err := s.ManageVideos.GenerateSubtitleTrack(
		ctx,
		"your-video-id",    // ← replace with real media ID
		"your-language-id", // ← replace with real track ID
		components.TrackSubtitlesGenerateRequest{
			LanguageName: fastpixgo.Pointer("Italian"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	if res.Object != nil {
		if res.HTTPMeta.Response != nil && res.HTTPMeta.Response.Body != nil {
			rawBody, err := io.ReadAll(res.HTTPMeta.Response.Body)
			if err == nil && len(rawBody) > 0 {
				var buf bytes.Buffer
				if err := json.Indent(&buf, rawBody, "", "  "); err == nil {
					fmt.Println(buf.String())
				} else {
					fmt.Println(string(rawBody))
				}
			} else {
				responseJSON, err := json.MarshalIndent(res.Object, "", "  ")
				if err != nil {
					log.Printf("Error marshaling response: %v", err)
					fmt.Printf("Response: %+v\n", res.Object)
				} else {
					fmt.Println(string(responseJSON))
				}
			}
		} else {
			responseJSON, err := json.MarshalIndent(res.Object, "", "  ")
			if err != nil {
				log.Printf("Error marshaling response: %v", err)
				fmt.Printf("Response: %+v\n", res.Object)
			} else {
				fmt.Println(string(responseJSON))
			}
		}
	} else if res.DefaultError != nil {
		fmt.Printf("Error: %+v\n", res.DefaultError)
	}
}