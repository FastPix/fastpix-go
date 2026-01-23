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
			Username: fastpixgo.Pointer("your access-token"),
			Password: fastpixgo.Pointer("your-secret-key"),
		}),
	)

	res, err := s.InputVideo.Create(ctx, components.CreateMediaRequest{
		Inputs: []components.Input{
			components.CreateInputPullVideoInput(
				components.PullVideoInput{},
			),
		},
		Metadata: map[string]string{
			"your-metadata-key": "your-metadata-value",
		},
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