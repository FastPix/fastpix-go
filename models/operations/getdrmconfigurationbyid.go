

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type GetDrmConfigurationByIDRequest struct {
	// The unique identifier of the DRM configuration.
	DrmConfigurationID string `pathParam:"style=simple,explode=false,name=drmConfigurationId"`
}

func (g *GetDrmConfigurationByIDRequest) GetDrmConfigurationID() string {
	if g == nil {
		return ""
	}
	return g.DrmConfigurationID
}

// GetDrmConfigurationByIDResponseBody - DRM configuration retrieved successfully
type GetDrmConfigurationByIDResponseBody struct {
	Success *bool                     `json:"success,omitempty"`
	Data    *components.DrmIDResponse `json:"data,omitempty"`
}

func (g *GetDrmConfigurationByIDResponseBody) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetDrmConfigurationByIDResponseBody) GetData() *components.DrmIDResponse {
	if g == nil {
		return nil
	}
	return g.Data
}

type GetDrmConfigurationByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// DRM configuration retrieved successfully
	Object *GetDrmConfigurationByIDResponseBody
}

func (g *GetDrmConfigurationByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GetDrmConfigurationByIDResponse) GetObject() *GetDrmConfigurationByIDResponseBody {
	if g == nil {
		return nil
	}
	return g.Object
}
