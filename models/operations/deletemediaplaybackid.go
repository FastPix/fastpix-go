package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type DeleteMediaPlaybackIDRequest struct {
	// Return the universal unique identifier for media which can contain a maximum of 255 characters.
	MediaID string `pathParam:"style=simple,explode=false,name=mediaId"`
	// Return the universal unique identifier for playbacks  which can contain a maximum of 255 characters.
	PlaybackID string `queryParam:"style=form,explode=true,name=playbackId"`
}

func (d *DeleteMediaPlaybackIDRequest) GetMediaID() string {
	if d == nil {
		return ""
	}
	return d.MediaID
}

func (d *DeleteMediaPlaybackIDRequest) GetPlaybackID() string {
	if d == nil {
		return ""
	}
	return d.PlaybackID
}

// DeleteMediaPlaybackIDResponseBody - Deleted a Playback Id successfully
type DeleteMediaPlaybackIDResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
}

func (d *DeleteMediaPlaybackIDResponseBody) GetSuccess() *bool {
	if d == nil {
		return nil
	}
	return d.Success
}

type DeleteMediaPlaybackIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Deleted a Playback Id successfully
	Object *DeleteMediaPlaybackIDResponseBody
}

func (d *DeleteMediaPlaybackIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if d == nil {
		return components.HTTPMetadata{}
	}
	return d.HTTPMeta
}

func (d *DeleteMediaPlaybackIDResponse) GetObject() *DeleteMediaPlaybackIDResponseBody {
	if d == nil {
		return nil
	}
	return d.Object
}
