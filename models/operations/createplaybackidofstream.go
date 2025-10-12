

package operations

import (
	"github.com/fastpix/fastpix-go/models/components"
)

type CreatePlaybackIDOfStreamRequest struct {
	// Upon creating a new live stream, FastPix assigns a unique identifier to the stream.
	StreamID          string                       `pathParam:"style=simple,explode=false,name=streamId"`
	PlaybackIDRequest components.PlaybackIDRequest `request:"mediaType=application/json"`
}

func (c *CreatePlaybackIDOfStreamRequest) GetStreamID() string {
	if c == nil {
		return ""
	}
	return c.StreamID
}

func (c *CreatePlaybackIDOfStreamRequest) GetPlaybackIDRequest() components.PlaybackIDRequest {
	if c == nil {
		return components.PlaybackIDRequest{}
	}
	return c.PlaybackIDRequest
}

type CreatePlaybackIDOfStreamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// New PlaybackId created successfully
	PlaybackIDSuccessResponse *components.PlaybackIDSuccessResponse
}

func (c *CreatePlaybackIDOfStreamResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CreatePlaybackIDOfStreamResponse) GetPlaybackIDSuccessResponse() *components.PlaybackIDSuccessResponse {
	if c == nil {
		return nil
	}
	return c.PlaybackIDSuccessResponse
}
