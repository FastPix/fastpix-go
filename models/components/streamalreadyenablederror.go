

package components

import (
	"github.com/FastPix/fastpix-go/internal/utils"
)

// StreamAlreadyEnabledErrorError - Contains details explaining why the request failed.
type StreamAlreadyEnabledErrorError struct {
	// HTTP status code indicating the nature of the error.
	Code *float64 `json:"code,omitempty"`
	// A short message summarizing the error.
	Message *string `json:"message,omitempty"`
	// A detailed explanation indicating that the stream is already active, idle, or preparing, and therefore cannot be enabled again.
	//
	Description *string `json:"description,omitempty"`
}

func (s StreamAlreadyEnabledErrorError) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *StreamAlreadyEnabledErrorError) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *StreamAlreadyEnabledErrorError) GetCode() *float64 {
	if s == nil {
		return nil
	}
	return s.Code
}

func (s *StreamAlreadyEnabledErrorError) GetMessage() *string {
	if s == nil {
		return nil
	}
	return s.Message
}

func (s *StreamAlreadyEnabledErrorError) GetDescription() *string {
	if s == nil {
		return nil
	}
	return s.Description
}
