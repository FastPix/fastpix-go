

package operations

import (
	"github.com/fastpix/fastpix-go/models/components"
)

type UpdateAPlaylistRequest struct {
	// The unique id of the playlist you want to retrieve.
	PlaylistID            string                           `pathParam:"style=simple,explode=false,name=playlistId"`
	UpdatePlaylistRequest components.UpdatePlaylistRequest `request:"mediaType=application/json"`
}

func (u *UpdateAPlaylistRequest) GetPlaylistID() string {
	if u == nil {
		return ""
	}
	return u.PlaylistID
}

func (u *UpdateAPlaylistRequest) GetUpdatePlaylistRequest() components.UpdatePlaylistRequest {
	if u == nil {
		return components.UpdatePlaylistRequest{}
	}
	return u.UpdatePlaylistRequest
}

type UpdateAPlaylistResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Playlist updated successfully
	PlaylistCreatedResponse *components.PlaylistCreatedResponse
}

func (u *UpdateAPlaylistResponse) GetHTTPMeta() components.HTTPMetadata {
	if u == nil {
		return components.HTTPMetadata{}
	}
	return u.HTTPMeta
}

func (u *UpdateAPlaylistResponse) GetPlaylistCreatedResponse() *components.PlaylistCreatedResponse {
	if u == nil {
		return nil
	}
	return u.PlaylistCreatedResponse
}
