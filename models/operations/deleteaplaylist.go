

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type DeleteAPlaylistRequest struct {
	// The unique id of the playlist you want to delete.
	PlaylistID string `pathParam:"style=simple,explode=false,name=playlistId"`
}

func (d *DeleteAPlaylistRequest) GetPlaylistID() string {
	if d == nil {
		return ""
	}
	return d.PlaylistID
}

type DeleteAPlaylistResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Playlist deleted successfully
	SuccessResponse *components.SuccessResponse
}

func (d *DeleteAPlaylistResponse) GetHTTPMeta() components.HTTPMetadata {
	if d == nil {
		return components.HTTPMetadata{}
	}
	return d.HTTPMeta
}

func (d *DeleteAPlaylistResponse) GetSuccessResponse() *components.SuccessResponse {
	if d == nil {
		return nil
	}
	return d.SuccessResponse
}
