

package operations

import (
	"github.com/fastpix/fastpix-go/models/components"
)

type GetVideoViewDetailsRequest struct {
	// Pass View id
	ViewID string `pathParam:"style=simple,explode=false,name=viewId"`
}

func (g *GetVideoViewDetailsRequest) GetViewID() string {
	if g == nil {
		return ""
	}
	return g.ViewID
}

// GetVideoViewDetailsResponseBody - Get a video view by id
type GetVideoViewDetailsResponseBody struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data *components.Views `json:"data,omitempty"`
}

func (g *GetVideoViewDetailsResponseBody) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetVideoViewDetailsResponseBody) GetData() *components.Views {
	if g == nil {
		return nil
	}
	return g.Data
}

type GetVideoViewDetailsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Get a video view by id
	Object *GetVideoViewDetailsResponseBody
}

func (g *GetVideoViewDetailsResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GetVideoViewDetailsResponse) GetObject() *GetVideoViewDetailsResponseBody {
	if g == nil {
		return nil
	}
	return g.Object
}
