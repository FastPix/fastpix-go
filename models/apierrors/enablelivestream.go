

package apierrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
	"github.com/FastPix/fastpix-go/models/components"
)

type BadRequestType string

const (
	BadRequestTypeTrialPlanRestrictionError BadRequestType = "TrialPlanRestrictionError"
	BadRequestTypeStreamAlreadyEnabledError BadRequestType = "StreamAlreadyEnabledError"
)

// BadRequest - Bad Request – Stream is either already enabled or cannot be enabled on trial plan.
type BadRequest struct {
	TrialPlanRestrictionError *TrialPlanRestrictionError `queryParam:"inline,name=ResponseBody"`
	StreamAlreadyEnabledError *StreamAlreadyEnabledError `queryParam:"inline,name=ResponseBody"`

	Type BadRequestType

	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &BadRequest{}

func CreateBadRequestTrialPlanRestrictionError(trialPlanRestrictionError TrialPlanRestrictionError) BadRequest {
	typ := BadRequestTypeTrialPlanRestrictionError

	return BadRequest{
		TrialPlanRestrictionError: &trialPlanRestrictionError,
		Type:                      typ,
	}
}

func CreateBadRequestStreamAlreadyEnabledError(streamAlreadyEnabledError StreamAlreadyEnabledError) BadRequest {
	typ := BadRequestTypeStreamAlreadyEnabledError

	return BadRequest{
		StreamAlreadyEnabledError: &streamAlreadyEnabledError,
		Type:                      typ,
	}
}

func (u *BadRequest) UnmarshalJSON(data []byte) error {

	var trialPlanRestrictionError TrialPlanRestrictionError = TrialPlanRestrictionError{}
	if err := utils.UnmarshalJSON(data, &trialPlanRestrictionError, "", true, nil); err == nil {
		u.TrialPlanRestrictionError = &trialPlanRestrictionError
		u.Type = BadRequestTypeTrialPlanRestrictionError
		return nil
	}

	var streamAlreadyEnabledError StreamAlreadyEnabledError = StreamAlreadyEnabledError{}
	if err := utils.UnmarshalJSON(data, &streamAlreadyEnabledError, "", true, nil); err == nil {
		u.StreamAlreadyEnabledError = &streamAlreadyEnabledError
		u.Type = BadRequestTypeStreamAlreadyEnabledError
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for BadRequest", string(data))
}

func (u BadRequest) MarshalJSON() ([]byte, error) {
	if u.TrialPlanRestrictionError != nil {
		return utils.MarshalJSON(u.TrialPlanRestrictionError, "", true)
	}

	if u.StreamAlreadyEnabledError != nil {
		return utils.MarshalJSON(u.StreamAlreadyEnabledError, "", true)
	}

	return nil, errors.New("could not marshal union type BadRequest: all fields are null")
}

func (u BadRequest) Error() string {
	switch u.Type {
	case BadRequestTypeTrialPlanRestrictionError:
		data, _ := json.Marshal(u.TrialPlanRestrictionError)
		return string(data)
	case BadRequestTypeStreamAlreadyEnabledError:
		data, _ := json.Marshal(u.StreamAlreadyEnabledError)
		return string(data)
	default:
		return "unknown error"
	}
}
