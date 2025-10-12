package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type DeleteSimulcastOfStreamRequest struct {
	// Upon creating a new live stream, FastPix assigns a unique identifier to the stream.
	StreamID string `pathParam:"style=simple,explode=false,name=streamId"`
	// When you create the new simulcast, FastPix assign a universal unique identifier which can contain a maximum of 255 characters.
	SimulcastID string `pathParam:"style=simple,explode=false,name=simulcastId"`
}

func (d *DeleteSimulcastOfStreamRequest) GetStreamID() string {
	if d == nil {
		return ""
	}
	return d.StreamID
}

func (d *DeleteSimulcastOfStreamRequest) GetSimulcastID() string {
	if d == nil {
		return ""
	}
	return d.SimulcastID
}

type DeleteSimulcastOfStreamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Stream's simulcast deleted successfully
	SimulcastdeleteResponse *components.SimulcastdeleteResponse
}

func (d *DeleteSimulcastOfStreamResponse) GetHTTPMeta() components.HTTPMetadata {
	if d == nil {
		return components.HTTPMetadata{}
	}
	return d.HTTPMeta
}

func (d *DeleteSimulcastOfStreamResponse) GetSimulcastdeleteResponse() *components.SimulcastdeleteResponse {
	if d == nil {
		return nil
	}
	return d.SimulcastdeleteResponse
}
