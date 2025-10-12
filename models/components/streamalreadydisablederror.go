package components

// StreamAlreadyDisabledErrorError - Contains details explaining why the request failed.
type StreamAlreadyDisabledErrorError struct {
	// HTTP status code indicating the nature of the error.
	Code *float64 `json:"code,omitempty"`
	// A short message summarizing the error.
	Message *string `json:"message,omitempty"`
	// A detailed explanation indicating that the stream is already in a disabled state and cannot be disabled again.
	//
	Description *string `json:"description,omitempty"`
}

func (s *StreamAlreadyDisabledErrorError) GetCode() *float64 {
	if s == nil {
		return nil
	}
	return s.Code
}

func (s *StreamAlreadyDisabledErrorError) GetMessage() *string {
	if s == nil {
		return nil
	}
	return s.Message
}

func (s *StreamAlreadyDisabledErrorError) GetDescription() *string {
	if s == nil {
		return nil
	}
	return s.Description
}
