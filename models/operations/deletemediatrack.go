

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type DeleteMediaTrackRequest struct {
	// When creating the media, FastPix assigns a universally unique identifier with a maximum length of 255 characters.
	MediaID string `pathParam:"style=simple,explode=false,name=mediaId"`
	// When creating the media, FastPix assigns a universally unique identifier with a maximum length of 255 characters.
	TrackID string `pathParam:"style=simple,explode=false,name=trackId"`
}

func (d *DeleteMediaTrackRequest) GetMediaID() string {
	if d == nil {
		return ""
	}
	return d.MediaID
}

func (d *DeleteMediaTrackRequest) GetTrackID() string {
	if d == nil {
		return ""
	}
	return d.TrackID
}

// DeleteMediaTrackResponseBody - Delete a video media
type DeleteMediaTrackResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
}

func (d *DeleteMediaTrackResponseBody) GetSuccess() *bool {
	if d == nil {
		return nil
	}
	return d.Success
}

type DeleteMediaTrackResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Delete a video media
	Object *DeleteMediaTrackResponseBody
}

func (d *DeleteMediaTrackResponse) GetHTTPMeta() components.HTTPMetadata {
	if d == nil {
		return components.HTTPMetadata{}
	}
	return d.HTTPMeta
}

func (d *DeleteMediaTrackResponse) GetObject() *DeleteMediaTrackResponseBody {
	if d == nil {
		return nil
	}
	return d.Object
}
