package apierrors

import (
	"encoding/json"
	"github.com/FastPix/fastpix-go/models/components"
)

type NotFoundError struct {
	Success  *bool                          `json:"success,omitempty"`
	Error_   *components.NotFoundErrorError `json:"error,omitempty"`
	HTTPMeta components.HTTPMetadata        `json:"-"`
}

var _ error = &NotFoundError{}

func (e *NotFoundError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
