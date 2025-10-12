

package components

import (
	"github.com/fastpix/fastpix-go/internal/utils"
)

// TrialPlanRestrictionErrorError - Contains details explaining why the request failed.
type TrialPlanRestrictionErrorError struct {
	// HTTP status code indicating the nature of the error.
	Code *float64 `json:"code,omitempty"`
	// A short message summarizing the error.
	Message *string `json:"message,omitempty"`
	// A detailed explanation of the error indicating that the operation is restricted under the trial plan. This typically occurs when certain features are gated for paid users only.
	//
	Description *string `json:"description,omitempty"`
}

func (t TrialPlanRestrictionErrorError) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TrialPlanRestrictionErrorError) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (t *TrialPlanRestrictionErrorError) GetCode() *float64 {
	if t == nil {
		return nil
	}
	return t.Code
}

func (t *TrialPlanRestrictionErrorError) GetMessage() *string {
	if t == nil {
		return nil
	}
	return t.Message
}

func (t *TrialPlanRestrictionErrorError) GetDescription() *string {
	if t == nil {
		return nil
	}
	return t.Description
}
