

package apierrors

import (
	"encoding/json"
	"github.com/fastpix/fastpix-go/models/components"
)

type DuplicateMp4SupportError struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays details about the reasons behind the request's failure.
	Error_   *components.DuplicateMp4SupportError `json:"error,omitempty"`
	HTTPMeta components.HTTPMetadata              `json:"-"`
}

var _ error = &DuplicateMp4SupportError{}

func (e *DuplicateMp4SupportError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
