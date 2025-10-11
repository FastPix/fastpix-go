

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type GetPlaybackIDRequest struct {
	MediaID    string `pathParam:"style=simple,explode=false,name=mediaId"`
	PlaybackID string `pathParam:"style=simple,explode=false,name=playbackId"`
}

func (g *GetPlaybackIDRequest) GetMediaID() string {
	if g == nil {
		return ""
	}
	return g.MediaID
}

func (g *GetPlaybackIDRequest) GetPlaybackID() string {
	if g == nil {
		return ""
	}
	return g.PlaybackID
}

type GetPlaybackIDData struct {
	// The unique identifier for the playback ID.
	ID *string `json:"id,omitempty"`
	// Access policy for media content
	AccessPolicy *components.AccessPolicy `json:"accessPolicy,omitempty"`
}

func (g *GetPlaybackIDData) GetID() *string {
	if g == nil {
		return nil
	}
	return g.ID
}

func (g *GetPlaybackIDData) GetAccessPolicy() *components.AccessPolicy {
	if g == nil {
		return nil
	}
	return g.AccessPolicy
}

// GetPlaybackIDResponseBody - Successfully retrieved playback ID details
type GetPlaybackIDResponseBody struct {
	// Indicates if the request was successful or not.
	Success *bool              `json:"success,omitempty"`
	Data    *GetPlaybackIDData `json:"data,omitempty"`
}

func (g *GetPlaybackIDResponseBody) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetPlaybackIDResponseBody) GetData() *GetPlaybackIDData {
	if g == nil {
		return nil
	}
	return g.Data
}

type GetPlaybackIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successfully retrieved playback ID details
	Object *GetPlaybackIDResponseBody
}

func (g *GetPlaybackIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GetPlaybackIDResponse) GetObject() *GetPlaybackIDResponseBody {
	if g == nil {
		return nil
	}
	return g.Object
}
