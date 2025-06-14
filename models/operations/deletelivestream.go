// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type DeleteLiveStreamRequest struct {
	// Upon creating a new live stream, FastPix assigns a unique identifier to the stream.
	StreamID string `pathParam:"style=simple,explode=false,name=streamId"`
}

func (o *DeleteLiveStreamRequest) GetStreamID() string {
	if o == nil {
		return ""
	}
	return o.StreamID
}

type DeleteLiveStreamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Stream deleted successfully
	LiveStreamDeleteResponse *components.LiveStreamDeleteResponse
}

func (o *DeleteLiveStreamResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteLiveStreamResponse) GetLiveStreamDeleteResponse() *components.LiveStreamDeleteResponse {
	if o == nil {
		return nil
	}
	return o.LiveStreamDeleteResponse
}
