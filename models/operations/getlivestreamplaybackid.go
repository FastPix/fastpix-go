package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type GetLiveStreamPlaybackIDRequest struct {
	// Upon creating a new live stream, FastPix assigns a unique identifier to the stream.
	StreamID string `pathParam:"style=simple,explode=false,name=streamId"`
	// Upon creating a new playbackId, FastPix assigns a unique identifier to the playback.
	PlaybackID string `pathParam:"style=simple,explode=false,name=playbackId"`
}

func (g *GetLiveStreamPlaybackIDRequest) GetStreamID() string {
	if g == nil {
		return ""
	}
	return g.StreamID
}

func (g *GetLiveStreamPlaybackIDRequest) GetPlaybackID() string {
	if g == nil {
		return ""
	}
	return g.PlaybackID
}

type GetLiveStreamPlaybackIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Stream details retrieved successfully
	PlaybackIDSuccessResponse *components.PlaybackIDSuccessResponse
}

func (g *GetLiveStreamPlaybackIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GetLiveStreamPlaybackIDResponse) GetPlaybackIDSuccessResponse() *components.PlaybackIDSuccessResponse {
	if g == nil {
		return nil
	}
	return g.PlaybackIDSuccessResponse
}
