

package apierrors

import (
	"encoding/json"
	"github.com/FastPix/fastpix-go/models/components"
)

type ViewNotFoundError struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Returns the problem that has occured
	Error_   *components.ViewNotFoundError `json:"error,omitempty"`
	HTTPMeta components.HTTPMetadata       `json:"-"`
}

var _ error = &ViewNotFoundError{}

func (e *ViewNotFoundError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
