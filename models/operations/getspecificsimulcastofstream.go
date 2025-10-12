package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type GetSpecificSimulcastOfStreamRequest struct {
	// Upon creating a new live stream, FastPix assigns a unique identifier to the stream.
	StreamID string `pathParam:"style=simple,explode=false,name=streamId"`
	// When you create the new simulcast, FastPix assign a universal unique identifier which can contain a maximum of 255 characters.
	SimulcastID string `pathParam:"style=simple,explode=false,name=simulcastId"`
}

func (g *GetSpecificSimulcastOfStreamRequest) GetStreamID() string {
	if g == nil {
		return ""
	}
	return g.StreamID
}

func (g *GetSpecificSimulcastOfStreamRequest) GetSimulcastID() string {
	if g == nil {
		return ""
	}
	return g.SimulcastID
}

type GetSpecificSimulcastOfStreamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Stream's simulcast details fetched successfully
	SimulcastResponse *components.SimulcastResponse
}

func (g *GetSpecificSimulcastOfStreamResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GetSpecificSimulcastOfStreamResponse) GetSimulcastResponse() *components.SimulcastResponse {
	if g == nil {
		return nil
	}
	return g.SimulcastResponse
}
