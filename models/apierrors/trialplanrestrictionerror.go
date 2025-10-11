

package apierrors

import (
	"encoding/json"
	"github.com/FastPix/fastpix-go/models/components"
)

type TrialPlanRestrictionError struct {
	// Indicates whether the request was successful or not.
	Success *bool `json:"success,omitempty"`
	// Contains details explaining why the request failed.
	Error_   *components.TrialPlanRestrictionErrorError `json:"error,omitempty"`
	HTTPMeta components.HTTPMetadata                    `json:"-"`
}

var _ error = &TrialPlanRestrictionError{}

func (e *TrialPlanRestrictionError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
