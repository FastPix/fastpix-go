

package apierrors

import (
	"encoding/json"
	"github.com/FastPix/fastpix-go/models/components"
)

type MediaClipNotFoundError struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays details about the reasons behind the request's failure.
	Error_   *components.MediaClipNotFoundError `json:"error,omitempty"`
	HTTPMeta components.HTTPMetadata            `json:"-"`
}

var _ error = &MediaClipNotFoundError{}

func (e *MediaClipNotFoundError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
