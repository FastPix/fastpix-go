package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type UpdateMediaNamedEntitiesRequestBody struct {
	// Enable or disable named entity extraction. Set to `true` to enable or `false` to disable.
	//
	NamedEntities bool `json:"namedEntities"`
}

func (u *UpdateMediaNamedEntitiesRequestBody) GetNamedEntities() bool {
	if u == nil {
		return false
	}
	return u.NamedEntities
}

type UpdateMediaNamedEntitiesRequest struct {
	// The unique identifier assigned to the media when created. The value should be a valid UUID.
	//
	MediaID     string                              `pathParam:"style=simple,explode=false,name=mediaId"`
	RequestBody UpdateMediaNamedEntitiesRequestBody `request:"mediaType=application/json"`
}

func (u *UpdateMediaNamedEntitiesRequest) GetMediaID() string {
	if u == nil {
		return ""
	}
	return u.MediaID
}

func (u *UpdateMediaNamedEntitiesRequest) GetRequestBody() UpdateMediaNamedEntitiesRequestBody {
	if u == nil {
		return UpdateMediaNamedEntitiesRequestBody{}
	}
	return u.RequestBody
}

// UpdateMediaNamedEntitiesResponseBody - Media details updated successfully with the named entity extraction feature enabled or disabled
type UpdateMediaNamedEntitiesResponseBody struct {
	// Indicates if the request was successful or not.
	Success *bool                             `json:"success,omitempty"`
	Data    *components.NamedEntitiesResponse `json:"data,omitempty"`
}

func (u *UpdateMediaNamedEntitiesResponseBody) GetSuccess() *bool {
	if u == nil {
		return nil
	}
	return u.Success
}

func (u *UpdateMediaNamedEntitiesResponseBody) GetData() *components.NamedEntitiesResponse {
	if u == nil {
		return nil
	}
	return u.Data
}

type UpdateMediaNamedEntitiesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Media details updated successfully with the named entity extraction feature enabled or disabled
	Object *UpdateMediaNamedEntitiesResponseBody
}

func (u *UpdateMediaNamedEntitiesResponse) GetHTTPMeta() components.HTTPMetadata {
	if u == nil {
		return components.HTTPMetadata{}
	}
	return u.HTTPMeta
}

func (u *UpdateMediaNamedEntitiesResponse) GetObject() *UpdateMediaNamedEntitiesResponseBody {
	if u == nil {
		return nil
	}
	return u.Object
}
