

package components

// SigningKeyNotFoundErrorError - Displays details about the reasons behind the request's failure.
type SigningKeyNotFoundErrorError struct {
	// An error code indicating the type of the error.
	Code *int64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error.
	Message *string `json:"message,omitempty"`
}

func (s *SigningKeyNotFoundErrorError) GetCode() *int64 {
	if s == nil {
		return nil
	}
	return s.Code
}

func (s *SigningKeyNotFoundErrorError) GetMessage() *string {
	if s == nil {
		return nil
	}
	return s.Message
}
