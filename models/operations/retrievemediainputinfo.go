

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type RetrieveMediaInputInfoRequest struct {
	// Pass the list of the input objects used to create the media, along with applied settings.
	MediaID string `pathParam:"style=simple,explode=false,name=mediaId"`
}

func (r *RetrieveMediaInputInfoRequest) GetMediaID() string {
	if r == nil {
		return ""
	}
	return r.MediaID
}

// RetrieveMediaInputInfoData - Displays the result of the request.
type RetrieveMediaInputInfoData struct {
}

// RetrieveMediaInputInfoResponseBody - Get video media input information
type RetrieveMediaInputInfoResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data *RetrieveMediaInputInfoData `json:"data,omitempty"`
}

func (r *RetrieveMediaInputInfoResponseBody) GetSuccess() *bool {
	if r == nil {
		return nil
	}
	return r.Success
}

func (r *RetrieveMediaInputInfoResponseBody) GetData() *RetrieveMediaInputInfoData {
	if r == nil {
		return nil
	}
	return r.Data
}

type RetrieveMediaInputInfoResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Get video media input information
	Object *RetrieveMediaInputInfoResponseBody
}

func (r *RetrieveMediaInputInfoResponse) GetHTTPMeta() components.HTTPMetadata {
	if r == nil {
		return components.HTTPMetadata{}
	}
	return r.HTTPMeta
}

func (r *RetrieveMediaInputInfoResponse) GetObject() *RetrieveMediaInputInfoResponseBody {
	if r == nil {
		return nil
	}
	return r.Object
}
