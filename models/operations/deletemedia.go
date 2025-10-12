

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type DeleteMediaRequest struct {
	// When creating the media, FastPix assigns a universally unique identifier with a maximum length of 255 characters.
	MediaID string `pathParam:"style=simple,explode=false,name=mediaId"`
}

func (d *DeleteMediaRequest) GetMediaID() string {
	if d == nil {
		return ""
	}
	return d.MediaID
}

// DeleteMediaResponseBody - Delete a video media
type DeleteMediaResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
}

func (d *DeleteMediaResponseBody) GetSuccess() *bool {
	if d == nil {
		return nil
	}
	return d.Success
}

type DeleteMediaResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Delete a video media
	Object *DeleteMediaResponseBody
}

func (d *DeleteMediaResponse) GetHTTPMeta() components.HTTPMetadata {
	if d == nil {
		return components.HTTPMetadata{}
	}
	return d.HTTPMeta
}

func (d *DeleteMediaResponse) GetObject() *DeleteMediaResponseBody {
	if d == nil {
		return nil
	}
	return d.Object
}
