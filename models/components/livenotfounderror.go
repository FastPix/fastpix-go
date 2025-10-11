

package components

// LiveNotFoundErrorError - Displays details about the reasons behind the request's failure.
type LiveNotFoundErrorError struct {
	// Displays the error code indicating the type of the error.
	Code *float64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error.
	Message *string `json:"message,omitempty"`
	// A detailed explanation of the possible causes for the error.
	//
	Description *string `json:"description,omitempty"`
}

func (l *LiveNotFoundErrorError) GetCode() *float64 {
	if l == nil {
		return nil
	}
	return l.Code
}

func (l *LiveNotFoundErrorError) GetMessage() *string {
	if l == nil {
		return nil
	}
	return l.Message
}

func (l *LiveNotFoundErrorError) GetDescription() *string {
	if l == nil {
		return nil
	}
	return l.Description
}
