// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/FastPix/fastpix-go/models/components"
)

type SimulcastUnavailableError struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Returns the problem that has occured.
	//
	Error_   *components.SimulcastUnavailableError `json:"error,omitempty"`
	HTTPMeta components.HTTPMetadata               `json:"-"`
}

var _ error = &SimulcastUnavailableError{}

func (e *SimulcastUnavailableError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
