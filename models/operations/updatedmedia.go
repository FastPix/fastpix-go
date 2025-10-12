package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type UpdatedMediaRequestBody struct {
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (u *UpdatedMediaRequestBody) GetMetadata() map[string]string {
	if u == nil {
		return nil
	}
	return u.Metadata
}

type UpdatedMediaRequest struct {
	// When creating the media, FastPix assigns a universally unique identifier with a maximum length of 255 characters.
	MediaID     string                  `pathParam:"style=simple,explode=false,name=mediaId"`
	RequestBody UpdatedMediaRequestBody `request:"mediaType=application/json"`
}

func (u *UpdatedMediaRequest) GetMediaID() string {
	if u == nil {
		return ""
	}
	return u.MediaID
}

func (u *UpdatedMediaRequest) GetRequestBody() UpdatedMediaRequestBody {
	if u == nil {
		return UpdatedMediaRequestBody{}
	}
	return u.RequestBody
}

// UpdatedMediaResponseBody - Media details updated successfully
type UpdatedMediaResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool             `json:"success,omitempty"`
	Data    *components.Media `json:"data,omitempty"`
}

func (u *UpdatedMediaResponseBody) GetSuccess() *bool {
	if u == nil {
		return nil
	}
	return u.Success
}

func (u *UpdatedMediaResponseBody) GetData() *components.Media {
	if u == nil {
		return nil
	}
	return u.Data
}

type UpdatedMediaResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Media details updated successfully
	Object *UpdatedMediaResponseBody
}

func (u *UpdatedMediaResponse) GetHTTPMeta() components.HTTPMetadata {
	if u == nil {
		return components.HTTPMetadata{}
	}
	return u.HTTPMeta
}

func (u *UpdatedMediaResponse) GetObject() *UpdatedMediaResponseBody {
	if u == nil {
		return nil
	}
	return u.Object
}
