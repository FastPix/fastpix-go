

package operations

import (
	"github.com/fastpix/fastpix-go/models/components"
)

type CompleteLiveStreamRequest struct {
	// Upon creating a new live stream, FastPix assigns a unique identifier to the stream.
	StreamID string `pathParam:"style=simple,explode=false,name=streamId"`
}

func (c *CompleteLiveStreamRequest) GetStreamID() string {
	if c == nil {
		return ""
	}
	return c.StreamID
}

type CompleteLiveStreamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Stream details updated successfully
	LiveStreamDeleteResponse *components.LiveStreamDeleteResponse
}

func (c *CompleteLiveStreamResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CompleteLiveStreamResponse) GetLiveStreamDeleteResponse() *components.LiveStreamDeleteResponse {
	if c == nil {
		return nil
	}
	return c.LiveStreamDeleteResponse
}
