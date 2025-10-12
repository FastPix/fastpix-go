

package operations

import (
	"github.com/fastpix/fastpix-go/models/components"
)

type UpdateSpecificSimulcastOfStreamRequest struct {
	// Upon creating a new live stream, FastPix assigns a unique identifier to the stream.
	StreamID string `pathParam:"style=simple,explode=false,name=streamId"`
	// When you create the new simulcast, FastPix assign a universal unique identifier which can contain a maximum of 255 characters.
	SimulcastID            string                            `pathParam:"style=simple,explode=false,name=simulcastId"`
	SimulcastUpdateRequest components.SimulcastUpdateRequest `request:"mediaType=application/json"`
}

func (u *UpdateSpecificSimulcastOfStreamRequest) GetStreamID() string {
	if u == nil {
		return ""
	}
	return u.StreamID
}

func (u *UpdateSpecificSimulcastOfStreamRequest) GetSimulcastID() string {
	if u == nil {
		return ""
	}
	return u.SimulcastID
}

func (u *UpdateSpecificSimulcastOfStreamRequest) GetSimulcastUpdateRequest() components.SimulcastUpdateRequest {
	if u == nil {
		return components.SimulcastUpdateRequest{}
	}
	return u.SimulcastUpdateRequest
}

type UpdateSpecificSimulcastOfStreamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Stream's simulcast details fetched successfully
	SimulcastUpdateResponse *components.SimulcastUpdateResponse
}

func (u *UpdateSpecificSimulcastOfStreamResponse) GetHTTPMeta() components.HTTPMetadata {
	if u == nil {
		return components.HTTPMetadata{}
	}
	return u.HTTPMeta
}

func (u *UpdateSpecificSimulcastOfStreamResponse) GetSimulcastUpdateResponse() *components.SimulcastUpdateResponse {
	if u == nil {
		return nil
	}
	return u.SimulcastUpdateResponse
}
