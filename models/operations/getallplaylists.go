

package operations

import (
	"github.com/fastpix/fastpix-go/internal/utils"
	"github.com/fastpix/fastpix-go/models/components"
)

type GetAllPlaylistsRequest struct {
	// The number of playlists to return (default is 10, max is 50).
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	// The page number to retrieve, starting from 1. Used for paginating the playlist results.
	Offset *int64 `default:"1" queryParam:"style=form,explode=true,name=offset"`
}

func (g GetAllPlaylistsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetAllPlaylistsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (g *GetAllPlaylistsRequest) GetLimit() *int64 {
	if g == nil {
		return nil
	}
	return g.Limit
}

func (g *GetAllPlaylistsRequest) GetOffset() *int64 {
	if g == nil {
		return nil
	}
	return g.Offset
}

type GetAllPlaylistsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successfully retrieved all playlists
	GetAllPlaylistsResponse *components.GetAllPlaylistsResponse
}

func (g *GetAllPlaylistsResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GetAllPlaylistsResponse) GetGetAllPlaylistsResponse() *components.GetAllPlaylistsResponse {
	if g == nil {
		return nil
	}
	return g.GetAllPlaylistsResponse
}
