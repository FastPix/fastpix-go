

package operations

import (
	"github.com/fastpix/fastpix-go/models/components"
)

type EnableLiveStreamRequest struct {
	// Upon creating a new live stream, FastPix assigns a unique identifier to the stream.
	StreamID string `pathParam:"style=simple,explode=false,name=streamId"`
}

func (e *EnableLiveStreamRequest) GetStreamID() string {
	if e == nil {
		return ""
	}
	return e.StreamID
}

type EnableLiveStreamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Stream details updated successfully
	LiveStreamDeleteResponse *components.LiveStreamDeleteResponse
}

func (e *EnableLiveStreamResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EnableLiveStreamResponse) GetLiveStreamDeleteResponse() *components.LiveStreamDeleteResponse {
	if e == nil {
		return nil
	}
	return e.LiveStreamDeleteResponse
}
