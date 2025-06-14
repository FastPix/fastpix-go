// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

// CreateMediaResponseBody - Media is created successfully
type CreateMediaResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool                           `json:"success,omitempty"`
	Data    *components.CreateMediaResponse `json:"data,omitempty"`
}

func (o *CreateMediaResponseBody) GetSuccess() *bool {
	if o == nil {
		return nil
	}
	return o.Success
}

func (o *CreateMediaResponseBody) GetData() *components.CreateMediaResponse {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateMediaResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Media is created successfully
	Object *CreateMediaResponseBody
}

func (o *CreateMediaResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateMediaResponse) GetObject() *CreateMediaResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
