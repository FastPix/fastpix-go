

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type CreateNewStreamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Stream created successfully
	LiveStreamResponseDTO *components.LiveStreamResponseDTO
}

func (c *CreateNewStreamResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CreateNewStreamResponse) GetLiveStreamResponseDTO() *components.LiveStreamResponseDTO {
	if c == nil {
		return nil
	}
	return c.LiveStreamResponseDTO
}
