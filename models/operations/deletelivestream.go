

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type DeleteLiveStreamRequest struct {
	// Upon creating a new live stream, FastPix assigns a unique identifier to the stream.
	StreamID string `pathParam:"style=simple,explode=false,name=streamId"`
}

func (d *DeleteLiveStreamRequest) GetStreamID() string {
	if d == nil {
		return ""
	}
	return d.StreamID
}

type DeleteLiveStreamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Stream deleted successfully
	LiveStreamDeleteResponse *components.LiveStreamDeleteResponse
}

func (d *DeleteLiveStreamResponse) GetHTTPMeta() components.HTTPMetadata {
	if d == nil {
		return components.HTTPMetadata{}
	}
	return d.HTTPMeta
}

func (d *DeleteLiveStreamResponse) GetLiveStreamDeleteResponse() *components.LiveStreamDeleteResponse {
	if d == nil {
		return nil
	}
	return d.LiveStreamDeleteResponse
}
