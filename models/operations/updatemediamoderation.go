

package operations

import (
	"github.com/fastpix/fastpix-go/models/components"
)

type UpdateMediaModerationModeration struct {
	// Type of media content
	Type *components.MediaType `json:"type,omitempty"`
}

func (u *UpdateMediaModerationModeration) GetType() *components.MediaType {
	if u == nil {
		return nil
	}
	return u.Type
}

type UpdateMediaModerationRequestBody struct {
	Moderation *UpdateMediaModerationModeration `json:"moderation,omitempty"`
}

func (u *UpdateMediaModerationRequestBody) GetModeration() *UpdateMediaModerationModeration {
	if u == nil {
		return nil
	}
	return u.Moderation
}

type UpdateMediaModerationRequest struct {
	// The unique identifier assigned to the media when created. The value should be a valid UUID.
	//
	MediaID     string                           `pathParam:"style=simple,explode=false,name=mediaId"`
	RequestBody UpdateMediaModerationRequestBody `request:"mediaType=application/json"`
}

func (u *UpdateMediaModerationRequest) GetMediaID() string {
	if u == nil {
		return ""
	}
	return u.MediaID
}

func (u *UpdateMediaModerationRequest) GetRequestBody() UpdateMediaModerationRequestBody {
	if u == nil {
		return UpdateMediaModerationRequestBody{}
	}
	return u.RequestBody
}

// UpdateMediaModerationResponseBody - Media details updated successfully with the named entity extraction feature enabled or disabled
type UpdateMediaModerationResponseBody struct {
	// Indicates if the request was successful or not.
	Success *bool                          `json:"success,omitempty"`
	Data    *components.ModerationResponse `json:"data,omitempty"`
}

func (u *UpdateMediaModerationResponseBody) GetSuccess() *bool {
	if u == nil {
		return nil
	}
	return u.Success
}

func (u *UpdateMediaModerationResponseBody) GetData() *components.ModerationResponse {
	if u == nil {
		return nil
	}
	return u.Data
}

type UpdateMediaModerationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Media details updated successfully with the named entity extraction feature enabled or disabled
	Object *UpdateMediaModerationResponseBody
}

func (u *UpdateMediaModerationResponse) GetHTTPMeta() components.HTTPMetadata {
	if u == nil {
		return components.HTTPMetadata{}
	}
	return u.HTTPMeta
}

func (u *UpdateMediaModerationResponse) GetObject() *UpdateMediaModerationResponseBody {
	if u == nil {
		return nil
	}
	return u.Object
}
