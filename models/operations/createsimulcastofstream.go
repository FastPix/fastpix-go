

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type CreateSimulcastOfStreamRequest struct {
	// Upon creating a new live stream, FastPix assigns a unique identifier to the stream.
	StreamID         string                      `pathParam:"style=simple,explode=false,name=streamId"`
	SimulcastRequest components.SimulcastRequest `request:"mediaType=application/json"`
}

func (c *CreateSimulcastOfStreamRequest) GetStreamID() string {
	if c == nil {
		return ""
	}
	return c.StreamID
}

func (c *CreateSimulcastOfStreamRequest) GetSimulcastRequest() components.SimulcastRequest {
	if c == nil {
		return components.SimulcastRequest{}
	}
	return c.SimulcastRequest
}

type CreateSimulcastOfStreamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// New Simulcast created successfully
	SimulcastResponse *components.SimulcastResponse
}

func (c *CreateSimulcastOfStreamResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CreateSimulcastOfStreamResponse) GetSimulcastResponse() *components.SimulcastResponse {
	if c == nil {
		return nil
	}
	return c.SimulcastResponse
}
