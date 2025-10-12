

package apierrors

import (
	"encoding/json"
	"github.com/fastpix/fastpix-go/models/components"
)

type InvalidPlaylistIDResponseError struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays details about the reasons behind the request's failure.
	Error_   *components.InvalidPlaylistIDResponseError `json:"error,omitempty"`
	HTTPMeta components.HTTPMetadata                    `json:"-"`
}

var _ error = &InvalidPlaylistIDResponseError{}

func (e *InvalidPlaylistIDResponseError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
