

package operations

import (
	"github.com/fastpix/fastpix-go/internal/utils"
	"github.com/fastpix/fastpix-go/models/components"
)

type GetMediaClipsRequest struct {
	// The unique identifier of the source media.
	SourceMediaID string `pathParam:"style=simple,explode=false,name=sourceMediaId"`
	// Offset determines the starting point for data retrieval within a paginated list.
	Offset *int64 `default:"1" queryParam:"style=form,explode=true,name=offset"`
	// The number of media clips to retrieve per request.
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	// The values in the list can be arranged in two ways DESC (Descending) or ASC (Ascending).
	OrderBy *components.SortOrder `default:"desc" queryParam:"style=form,explode=true,name=orderBy"`
}

func (g GetMediaClipsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetMediaClipsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, []string{"sourceMediaId"}); err != nil {
		return err
	}
	return nil
}

func (g *GetMediaClipsRequest) GetSourceMediaID() string {
	if g == nil {
		return ""
	}
	return g.SourceMediaID
}

func (g *GetMediaClipsRequest) GetOffset() *int64 {
	if g == nil {
		return nil
	}
	return g.Offset
}

func (g *GetMediaClipsRequest) GetLimit() *int64 {
	if g == nil {
		return nil
	}
	return g.Limit
}

func (g *GetMediaClipsRequest) GetOrderBy() *components.SortOrder {
	if g == nil {
		return nil
	}
	return g.OrderBy
}

type GetMediaClipsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Get media clips
	MediaClipResponse *components.MediaClipResponse
}

func (g *GetMediaClipsResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GetMediaClipsResponse) GetMediaClipResponse() *components.MediaClipResponse {
	if g == nil {
		return nil
	}
	return g.MediaClipResponse
}
