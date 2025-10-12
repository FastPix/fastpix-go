package operations

import (
	"github.com/FastPix/fastpix-go/internal/utils"
	"github.com/FastPix/fastpix-go/models/components"
)

type GetDrmConfigurationRequest struct {
	// Offset determines the starting point for data retrieval within a paginated list.
	Offset *int64 `default:"1" queryParam:"style=form,explode=true,name=offset"`
	// Limit specifies the maximum number of items to display per page.
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
}

func (g GetDrmConfigurationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetDrmConfigurationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (g *GetDrmConfigurationRequest) GetOffset() *int64 {
	if g == nil {
		return nil
	}
	return g.Offset
}

func (g *GetDrmConfigurationRequest) GetLimit() *int64 {
	if g == nil {
		return nil
	}
	return g.Limit
}

// GetDrmConfigurationResponseBody - DRM configuration(s) retrieved successfully
type GetDrmConfigurationResponseBody struct {
	Success *bool                      `json:"success,omitempty"`
	Data    []components.DrmIDResponse `json:"data,omitempty"`
	// Pagination organizes content into pages for better readability and navigation.
	Pagination *components.Pagination `json:"pagination,omitempty"`
}

func (g *GetDrmConfigurationResponseBody) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetDrmConfigurationResponseBody) GetData() []components.DrmIDResponse {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetDrmConfigurationResponseBody) GetPagination() *components.Pagination {
	if g == nil {
		return nil
	}
	return g.Pagination
}

type GetDrmConfigurationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// DRM configuration(s) retrieved successfully
	Object *GetDrmConfigurationResponseBody
}

func (g *GetDrmConfigurationResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GetDrmConfigurationResponse) GetObject() *GetDrmConfigurationResponseBody {
	if g == nil {
		return nil
	}
	return g.Object
}
