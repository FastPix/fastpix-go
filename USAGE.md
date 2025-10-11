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
			Username: fastpixgo.Pointer("your-access-token"),
			Password: fastpixgo.Pointer("your-secret-key"),
		}),
	)

	res, err := s.InputVideo.CreateMedia(ctx, components.CreateMediaRequest{
		Inputs: []components.Input{
			components.CreateInputVideoInput(
				components.VideoInput{
					Type: "video",
					URL:  "https://static.fastpix.io/sample.mp4",
				},
			),
		},
		Metadata: map[string]string{
			"key1": "value1",
		},
		AccessPolicy: components.CreateMediaRequestAccessPolicyPublic,
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.CreateMediaSuccessResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->