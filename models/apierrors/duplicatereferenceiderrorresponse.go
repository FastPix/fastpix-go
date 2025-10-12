

package apierrors

import (
	"encoding/json"
	"github.com/fastpix/fastpix-go/models/components"
)

// DuplicateReferenceIDErrorResponse - Displays the result of the request.
type DuplicateReferenceIDErrorResponse struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays details about the reasons behind the request's failure.
	Error_   *components.DuplicateReferenceIDErrorResponseError `json:"error,omitempty"`
	HTTPMeta components.HTTPMetadata                            `json:"-"`
}

var _ error = &DuplicateReferenceIDErrorResponse{}

func (e *DuplicateReferenceIDErrorResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
