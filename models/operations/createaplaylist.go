package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type CreateAPlaylistResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Playlist created successfully
	PlaylistCreatedResponse *components.PlaylistCreatedResponse
}

func (c *CreateAPlaylistResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CreateAPlaylistResponse) GetPlaylistCreatedResponse() *components.PlaylistCreatedResponse {
	if c == nil {
		return nil
	}
	return c.PlaylistCreatedResponse
}
