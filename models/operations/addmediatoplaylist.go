

package operations

import (
	"github.com/fastpix/fastpix-go/models/components"
)

type AddMediaToPlaylistRequest struct {
	// The unique id of the playlist you want to perform the operation on.
	PlaylistID      string                     `pathParam:"style=simple,explode=false,name=playlistId"`
	MediaIdsRequest components.MediaIdsRequest `request:"mediaType=application/json"`
}

func (a *AddMediaToPlaylistRequest) GetPlaylistID() string {
	if a == nil {
		return ""
	}
	return a.PlaylistID
}

func (a *AddMediaToPlaylistRequest) GetMediaIdsRequest() components.MediaIdsRequest {
	if a == nil {
		return components.MediaIdsRequest{}
	}
	return a.MediaIdsRequest
}

type AddMediaToPlaylistResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Added media to playlist successfully
	PlaylistByIDResponse *components.PlaylistByIDResponse
}

func (a *AddMediaToPlaylistResponse) GetHTTPMeta() components.HTTPMetadata {
	if a == nil {
		return components.HTTPMetadata{}
	}
	return a.HTTPMeta
}

func (a *AddMediaToPlaylistResponse) GetPlaylistByIDResponse() *components.PlaylistByIDResponse {
	if a == nil {
		return nil
	}
	return a.PlaylistByIDResponse
}
