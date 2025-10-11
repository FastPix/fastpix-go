

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type UpdateMediaChaptersRequestBody struct {
	// Enable or disable the chapters feature for the media. Set to `true` to enable chapters or `false` to disable.
	//
	Chapters bool `json:"chapters"`
}

func (u *UpdateMediaChaptersRequestBody) GetChapters() bool {
	if u == nil {
		return false
	}
	return u.Chapters
}

type UpdateMediaChaptersRequest struct {
	// The unique identifier assigned to the media when created. The value should be a valid UUID.
	//
	MediaID     string                         `pathParam:"style=simple,explode=false,name=mediaId"`
	RequestBody UpdateMediaChaptersRequestBody `request:"mediaType=application/json"`
}

func (u *UpdateMediaChaptersRequest) GetMediaID() string {
	if u == nil {
		return ""
	}
	return u.MediaID
}

func (u *UpdateMediaChaptersRequest) GetRequestBody() UpdateMediaChaptersRequestBody {
	if u == nil {
		return UpdateMediaChaptersRequestBody{}
	}
	return u.RequestBody
}

// UpdateMediaChaptersResponseBody - Media details updated successfully with the chapters feature enabled or disabled
type UpdateMediaChaptersResponseBody struct {
	// Indicates if the request was successful or not.
	Success *bool                        `json:"success,omitempty"`
	Data    *components.ChaptersResponse `json:"data,omitempty"`
}

func (u *UpdateMediaChaptersResponseBody) GetSuccess() *bool {
	if u == nil {
		return nil
	}
	return u.Success
}

func (u *UpdateMediaChaptersResponseBody) GetData() *components.ChaptersResponse {
	if u == nil {
		return nil
	}
	return u.Data
}

type UpdateMediaChaptersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Media details updated successfully with the chapters feature enabled or disabled
	Object *UpdateMediaChaptersResponseBody
}

func (u *UpdateMediaChaptersResponse) GetHTTPMeta() components.HTTPMetadata {
	if u == nil {
		return components.HTTPMetadata{}
	}
	return u.HTTPMeta
}

func (u *UpdateMediaChaptersResponse) GetObject() *UpdateMediaChaptersResponseBody {
	if u == nil {
		return nil
	}
	return u.Object
}
