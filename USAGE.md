<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.String(""),
			Password: fastpixgo.String(""),
		}),
	)

	res, err := s.StartLiveStream.CreateNewStream(ctx, &components.CreateLiveStreamRequest{
		PlaybackSettings: components.PlaybackSettings{},
		InputMediaSettings: components.InputMediaSettings{
			Metadata: &components.CreateLiveStreamRequestMetadata{},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.LiveStreamResponseDTO != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->