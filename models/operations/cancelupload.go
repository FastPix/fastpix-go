package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type CancelUploadRequest struct {
	// When uploading the media, FastPix assigns a universally unique identifier with a maximum length of 255 characters.
	UploadID string `pathParam:"style=simple,explode=false,name=uploadId"`
}

func (c *CancelUploadRequest) GetUploadID() string {
	if c == nil {
		return ""
	}
	return c.UploadID
}

// CancelUploadResponseBody - Upload cancelled successfully
type CancelUploadResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Response returned when an upload is cancelled.
	Data *components.MediaCancelResponse `json:"data,omitempty"`
}

func (c *CancelUploadResponseBody) GetSuccess() *bool {
	if c == nil {
		return nil
	}
	return c.Success
}

func (c *CancelUploadResponseBody) GetData() *components.MediaCancelResponse {
	if c == nil {
		return nil
	}
	return c.Data
}

type CancelUploadResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Upload cancelled successfully
	Object *CancelUploadResponseBody
}

func (c *CancelUploadResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CancelUploadResponse) GetObject() *CancelUploadResponseBody {
	if c == nil {
		return nil
	}
	return c.Object
}
