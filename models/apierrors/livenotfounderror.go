

package apierrors

import (
	"encoding/json"
	"github.com/FastPix/fastpix-go/models/components"
)

type LiveNotFoundError struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays details about the reasons behind the request's failure.
	Error_   *components.LiveNotFoundErrorError `json:"error,omitempty"`
	HTTPMeta components.HTTPMetadata            `json:"-"`
}

var _ error = &LiveNotFoundError{}

func (e *LiveNotFoundError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
