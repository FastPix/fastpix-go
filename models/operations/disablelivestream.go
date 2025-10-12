

package operations

import (
	"github.com/fastpix/fastpix-go/models/components"
)

type DisableLiveStreamRequest struct {
	// Upon creating a new live stream, FastPix assigns a unique identifier to the stream.
	StreamID string `pathParam:"style=simple,explode=false,name=streamId"`
}

func (d *DisableLiveStreamRequest) GetStreamID() string {
	if d == nil {
		return ""
	}
	return d.StreamID
}

type DisableLiveStreamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Stream details updated successfully
	LiveStreamDeleteResponse *components.LiveStreamDeleteResponse
}

func (d *DisableLiveStreamResponse) GetHTTPMeta() components.HTTPMetadata {
	if d == nil {
		return components.HTTPMetadata{}
	}
	return d.HTTPMeta
}

func (d *DisableLiveStreamResponse) GetLiveStreamDeleteResponse() *components.LiveStreamDeleteResponse {
	if d == nil {
		return nil
	}
	return d.LiveStreamDeleteResponse
}
