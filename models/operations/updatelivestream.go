package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type UpdateLiveStreamRequest struct {
	// Upon creating a new live stream, FastPix assigns a unique identifier to the stream.
	StreamID               string                            `pathParam:"style=simple,explode=false,name=streamId"`
	PatchLiveStreamRequest components.PatchLiveStreamRequest `request:"mediaType=application/json"`
}

func (u *UpdateLiveStreamRequest) GetStreamID() string {
	if u == nil {
		return ""
	}
	return u.StreamID
}

func (u *UpdateLiveStreamRequest) GetPatchLiveStreamRequest() components.PatchLiveStreamRequest {
	if u == nil {
		return components.PatchLiveStreamRequest{}
	}
	return u.PatchLiveStreamRequest
}

type UpdateLiveStreamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Stream details updated successfully
	PatchResponseDTO *components.PatchResponseDTO
}

func (u *UpdateLiveStreamResponse) GetHTTPMeta() components.HTTPMetadata {
	if u == nil {
		return components.HTTPMetadata{}
	}
	return u.HTTPMeta
}

func (u *UpdateLiveStreamResponse) GetPatchResponseDTO() *components.PatchResponseDTO {
	if u == nil {
		return nil
	}
	return u.PatchResponseDTO
}
