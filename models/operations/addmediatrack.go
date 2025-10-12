

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type AddMediaTrackRequestBody struct {
	// Contains details about the track being added to the media file.
	Tracks *components.AddTrackRequest `json:"tracks,omitempty"`
}

func (a *AddMediaTrackRequestBody) GetTracks() *components.AddTrackRequest {
	if a == nil {
		return nil
	}
	return a.Tracks
}

type AddMediaTrackRequest struct {
	// When creating the media, FastPix assigns a universally unique identifier with a maximum length of 255 characters.
	MediaID     string                   `pathParam:"style=simple,explode=false,name=mediaId"`
	RequestBody AddMediaTrackRequestBody `request:"mediaType=application/json"`
}

func (a *AddMediaTrackRequest) GetMediaID() string {
	if a == nil {
		return ""
	}
	return a.MediaID
}

func (a *AddMediaTrackRequest) GetRequestBody() AddMediaTrackRequestBody {
	if a == nil {
		return AddMediaTrackRequestBody{}
	}
	return a.RequestBody
}

// AddMediaTrackResponseBody - Media details updated successfully
type AddMediaTrackResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Contains details about the track that was added or updated.
	Data *components.AddTrackResponse `json:"data,omitempty"`
}

func (a *AddMediaTrackResponseBody) GetSuccess() *bool {
	if a == nil {
		return nil
	}
	return a.Success
}

func (a *AddMediaTrackResponseBody) GetData() *components.AddTrackResponse {
	if a == nil {
		return nil
	}
	return a.Data
}

type AddMediaTrackResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Media details updated successfully
	Object *AddMediaTrackResponseBody
}

func (a *AddMediaTrackResponse) GetHTTPMeta() components.HTTPMetadata {
	if a == nil {
		return components.HTTPMetadata{}
	}
	return a.HTTPMeta
}

func (a *AddMediaTrackResponse) GetObject() *AddMediaTrackResponseBody {
	if a == nil {
		return nil
	}
	return a.Object
}
