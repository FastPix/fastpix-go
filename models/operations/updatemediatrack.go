

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type UpdateMediaTrackRequest struct {
	// When creating the media, FastPix assigns a universally unique identifier with a maximum length of 255 characters.
	TrackID string `pathParam:"style=simple,explode=false,name=trackId"`
	// When creating the media, FastPix assigns a universally unique identifier with a maximum length of 255 characters.
	MediaID            string                        `pathParam:"style=simple,explode=false,name=mediaId"`
	UpdateTrackRequest components.UpdateTrackRequest `request:"mediaType=application/json"`
}

func (u *UpdateMediaTrackRequest) GetTrackID() string {
	if u == nil {
		return ""
	}
	return u.TrackID
}

func (u *UpdateMediaTrackRequest) GetMediaID() string {
	if u == nil {
		return ""
	}
	return u.MediaID
}

func (u *UpdateMediaTrackRequest) GetUpdateTrackRequest() components.UpdateTrackRequest {
	if u == nil {
		return components.UpdateTrackRequest{}
	}
	return u.UpdateTrackRequest
}

// UpdateMediaTrackResponseBody - Media details updated successfully
type UpdateMediaTrackResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Contains details about the track that was added or updated.
	Data *components.UpdateTrackResponse `json:"data,omitempty"`
}

func (u *UpdateMediaTrackResponseBody) GetSuccess() *bool {
	if u == nil {
		return nil
	}
	return u.Success
}

func (u *UpdateMediaTrackResponseBody) GetData() *components.UpdateTrackResponse {
	if u == nil {
		return nil
	}
	return u.Data
}

type UpdateMediaTrackResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Media details updated successfully
	Object *UpdateMediaTrackResponseBody
}

func (u *UpdateMediaTrackResponse) GetHTTPMeta() components.HTTPMetadata {
	if u == nil {
		return components.HTTPMetadata{}
	}
	return u.HTTPMeta
}

func (u *UpdateMediaTrackResponse) GetObject() *UpdateMediaTrackResponseBody {
	if u == nil {
		return nil
	}
	return u.Object
}
