

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type GetMediaRequest struct {
	// The Media Id is assigned a universal unique identifier, which can contain a maximum of 255 characters.
	MediaID string `pathParam:"style=simple,explode=false,name=mediaId"`
}

func (g *GetMediaRequest) GetMediaID() string {
	if g == nil {
		return ""
	}
	return g.MediaID
}

// GetMediaResponseBody - Get a video media by id
type GetMediaResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool             `json:"success,omitempty"`
	Data    *components.Media `json:"data,omitempty"`
}

func (g *GetMediaResponseBody) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetMediaResponseBody) GetData() *components.Media {
	if g == nil {
		return nil
	}
	return g.Data
}

type GetMediaResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Get a video media by id
	Object *GetMediaResponseBody
}

func (g *GetMediaResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GetMediaResponse) GetObject() *GetMediaResponseBody {
	if g == nil {
		return nil
	}
	return g.Object
}
