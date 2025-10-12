package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type ChangeMediaOrderInPlaylistRequest struct {
	// The unique id of the playlist you want to perform the operation on.
	PlaylistID      string                     `pathParam:"style=simple,explode=false,name=playlistId"`
	MediaIdsRequest components.MediaIdsRequest `request:"mediaType=application/json"`
}

func (c *ChangeMediaOrderInPlaylistRequest) GetPlaylistID() string {
	if c == nil {
		return ""
	}
	return c.PlaylistID
}

func (c *ChangeMediaOrderInPlaylistRequest) GetMediaIdsRequest() components.MediaIdsRequest {
	if c == nil {
		return components.MediaIdsRequest{}
	}
	return c.MediaIdsRequest
}

type ChangeMediaOrderInPlaylistResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Added media to playlist successfully
	PlaylistByIDResponse *components.PlaylistByIDResponse
}

func (c *ChangeMediaOrderInPlaylistResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *ChangeMediaOrderInPlaylistResponse) GetPlaylistByIDResponse() *components.PlaylistByIDResponse {
	if c == nil {
		return nil
	}
	return c.PlaylistByIDResponse
}
