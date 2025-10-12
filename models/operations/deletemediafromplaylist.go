

package operations

import (
	"github.com/fastpix/fastpix-go/models/components"
)

type DeleteMediaFromPlaylistRequest struct {
	// The unique id of the playlist you want to perform the operation on.
	PlaylistID      string                      `pathParam:"style=simple,explode=false,name=playlistId"`
	MediaIdsRequest *components.MediaIdsRequest `request:"mediaType=application/json"`
}

func (d *DeleteMediaFromPlaylistRequest) GetPlaylistID() string {
	if d == nil {
		return ""
	}
	return d.PlaylistID
}

func (d *DeleteMediaFromPlaylistRequest) GetMediaIdsRequest() *components.MediaIdsRequest {
	if d == nil {
		return nil
	}
	return d.MediaIdsRequest
}

type DeleteMediaFromPlaylistResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Deleted media from playlist successfully
	PlaylistByIDResponse *components.PlaylistByIDResponse
}

func (d *DeleteMediaFromPlaylistResponse) GetHTTPMeta() components.HTTPMetadata {
	if d == nil {
		return components.HTTPMetadata{}
	}
	return d.HTTPMeta
}

func (d *DeleteMediaFromPlaylistResponse) GetPlaylistByIDResponse() *components.PlaylistByIDResponse {
	if d == nil {
		return nil
	}
	return d.PlaylistByIDResponse
}
