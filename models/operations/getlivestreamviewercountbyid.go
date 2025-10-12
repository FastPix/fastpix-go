package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type GetLiveStreamViewerCountByIDRequest struct {
	// Upon creating a new live stream, FastPix assigns a unique identifier to the stream.
	StreamID string `pathParam:"style=simple,explode=false,name=streamId"`
}

func (g *GetLiveStreamViewerCountByIDRequest) GetStreamID() string {
	if g == nil {
		return ""
	}
	return g.StreamID
}

type GetLiveStreamViewerCountByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Stream viewer count retrieved successfully
	ViewsCountResponse *components.ViewsCountResponse
}

func (g *GetLiveStreamViewerCountByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GetLiveStreamViewerCountByIDResponse) GetViewsCountResponse() *components.ViewsCountResponse {
	if g == nil {
		return nil
	}
	return g.ViewsCountResponse
}
