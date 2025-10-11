

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type GetLiveStreamByIDRequest struct {
	// Upon creating a new live stream, FastPix assigns a unique identifier to the stream.
	StreamID string `pathParam:"style=simple,explode=false,name=streamId"`
}

func (g *GetLiveStreamByIDRequest) GetStreamID() string {
	if g == nil {
		return ""
	}
	return g.StreamID
}

type GetLiveStreamByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Stream details retrieved successfully
	LivestreamgetResponse *components.LivestreamgetResponse
}

func (g *GetLiveStreamByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GetLiveStreamByIDResponse) GetLivestreamgetResponse() *components.LivestreamgetResponse {
	if g == nil {
		return nil
	}
	return g.LivestreamgetResponse
}
