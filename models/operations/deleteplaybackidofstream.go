package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type DeletePlaybackIDOfStreamRequest struct {
	// Upon creating a new live stream, FastPix assigns a unique identifier to the stream.
	StreamID string `pathParam:"style=simple,explode=false,name=streamId"`
	// Unique identifier for the playbackId
	PlaybackID string `queryParam:"style=form,explode=true,name=playbackId"`
}

func (d *DeletePlaybackIDOfStreamRequest) GetStreamID() string {
	if d == nil {
		return ""
	}
	return d.StreamID
}

func (d *DeletePlaybackIDOfStreamRequest) GetPlaybackID() string {
	if d == nil {
		return ""
	}
	return d.PlaybackID
}

type DeletePlaybackIDOfStreamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Stream's playbackId deleted successfully
	LiveStreamDeleteResponse *components.LiveStreamDeleteResponse
}

func (d *DeletePlaybackIDOfStreamResponse) GetHTTPMeta() components.HTTPMetadata {
	if d == nil {
		return components.HTTPMetadata{}
	}
	return d.HTTPMeta
}

func (d *DeletePlaybackIDOfStreamResponse) GetLiveStreamDeleteResponse() *components.LiveStreamDeleteResponse {
	if d == nil {
		return nil
	}
	return d.LiveStreamDeleteResponse
}
