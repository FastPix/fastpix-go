// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/FastPix/fastpix-go/models/components"
)

type NotFoundError struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays details about the reasons behind the request's failure.
	Error_   *components.NotFoundErrorError `json:"error,omitempty"`
	HTTPMeta components.HTTPMetadata        `json:"-"`
}

var _ error = &NotFoundError{}

func (e *NotFoundError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
