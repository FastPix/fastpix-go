

package operations

import (
	"github.com/fastpix/fastpix-go/models/components"
)

type GetPlaylistByIDRequest struct {
	// The unique id of the playlist you want to retrieve.
	PlaylistID string `pathParam:"style=simple,explode=false,name=playlistId"`
}

func (g *GetPlaylistByIDRequest) GetPlaylistID() string {
	if g == nil {
		return ""
	}
	return g.PlaylistID
}

type GetPlaylistByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successfully retrieved all playlists
	PlaylistByIDResponse *components.PlaylistByIDResponse
}

func (g *GetPlaylistByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GetPlaylistByIDResponse) GetPlaylistByIDResponse() *components.PlaylistByIDResponse {
	if g == nil {
		return nil
	}
	return g.PlaylistByIDResponse
}
