

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type UpdatedSourceAccessRequestBody struct {
	// The sourceAccess parameter determines whether the original media file is accessible. Set to true to enable access or false to restrict it.
	SourceAccess *bool `json:"sourceAccess,omitempty"`
}

func (u *UpdatedSourceAccessRequestBody) GetSourceAccess() *bool {
	if u == nil {
		return nil
	}
	return u.SourceAccess
}

type UpdatedSourceAccessRequest struct {
	// When creating the media, FastPix assigns a universally unique identifier with a maximum length of 255 characters.
	//
	MediaID     string                         `pathParam:"style=simple,explode=false,name=mediaId"`
	RequestBody UpdatedSourceAccessRequestBody `request:"mediaType=application/json"`
}

func (u *UpdatedSourceAccessRequest) GetMediaID() string {
	if u == nil {
		return ""
	}
	return u.MediaID
}

func (u *UpdatedSourceAccessRequest) GetRequestBody() UpdatedSourceAccessRequestBody {
	if u == nil {
		return UpdatedSourceAccessRequestBody{}
	}
	return u.RequestBody
}

// UpdatedSourceAccessResponseBody - Media details updated successfully
type UpdatedSourceAccessResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool             `json:"success,omitempty"`
	Data    *components.Media `json:"data,omitempty"`
}

func (u *UpdatedSourceAccessResponseBody) GetSuccess() *bool {
	if u == nil {
		return nil
	}
	return u.Success
}

func (u *UpdatedSourceAccessResponseBody) GetData() *components.Media {
	if u == nil {
		return nil
	}
	return u.Data
}

type UpdatedSourceAccessResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Media details updated successfully
	Object *UpdatedSourceAccessResponseBody
}

func (u *UpdatedSourceAccessResponse) GetHTTPMeta() components.HTTPMetadata {
	if u == nil {
		return components.HTTPMetadata{}
	}
	return u.HTTPMeta
}

func (u *UpdatedSourceAccessResponse) GetObject() *UpdatedSourceAccessResponseBody {
	if u == nil {
		return nil
	}
	return u.Object
}
