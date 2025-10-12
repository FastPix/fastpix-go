

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type GenerateSubtitleTrackRequest struct {
	// A universally unique identifier (UUID) assigned to the media by FastPix.
	MediaID string `pathParam:"style=simple,explode=false,name=mediaId"`
	// A universally unique identifier (UUID) assigned to the specific track for which subtitles should be generated.
	TrackID                       string                                   `pathParam:"style=simple,explode=false,name=trackId"`
	TrackSubtitlesGenerateRequest components.TrackSubtitlesGenerateRequest `request:"mediaType=application/json"`
}

func (g *GenerateSubtitleTrackRequest) GetMediaID() string {
	if g == nil {
		return ""
	}
	return g.MediaID
}

func (g *GenerateSubtitleTrackRequest) GetTrackID() string {
	if g == nil {
		return ""
	}
	return g.TrackID
}

func (g *GenerateSubtitleTrackRequest) GetTrackSubtitlesGenerateRequest() components.TrackSubtitlesGenerateRequest {
	if g == nil {
		return components.TrackSubtitlesGenerateRequest{}
	}
	return g.TrackSubtitlesGenerateRequest
}

// GenerateSubtitleTrackResponseBody - Media details updated successfully
type GenerateSubtitleTrackResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Represents the response for a successfully generated subtitle track.
	Data *components.GenerateTrackResponse `json:"data,omitempty"`
}

func (g *GenerateSubtitleTrackResponseBody) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GenerateSubtitleTrackResponseBody) GetData() *components.GenerateTrackResponse {
	if g == nil {
		return nil
	}
	return g.Data
}

type GenerateSubtitleTrackResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Media details updated successfully
	Object *GenerateSubtitleTrackResponseBody
}

func (g *GenerateSubtitleTrackResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GenerateSubtitleTrackResponse) GetObject() *GenerateSubtitleTrackResponseBody {
	if g == nil {
		return nil
	}
	return g.Object
}
