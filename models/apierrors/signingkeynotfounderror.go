

package apierrors

import (
	"encoding/json"
	"github.com/FastPix/fastpix-go/models/components"
)

type SigningKeyNotFoundError struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays details about the reasons behind the request's failure.
	Error_   *components.SigningKeyNotFoundErrorError `json:"error,omitempty"`
	HTTPMeta components.HTTPMetadata                  `json:"-"`
}

var _ error = &SigningKeyNotFoundError{}

func (e *SigningKeyNotFoundError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
