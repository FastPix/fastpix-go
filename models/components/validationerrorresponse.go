

package components

// ValidationErrorResponseError - Displays details about the reasons behind the request's failure.
type ValidationErrorResponseError struct {
	// Displays the error code indicating the type of the error.
	Code *int64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error.
	Message *string `json:"message,omitempty"`
	// It is an collection of objects, where each object contains information about a specific field and a corresponding error message.
	Fields []FieldError `json:"fields,omitempty"`
}

func (v *ValidationErrorResponseError) GetCode() *int64 {
	if v == nil {
		return nil
	}
	return v.Code
}

func (v *ValidationErrorResponseError) GetMessage() *string {
	if v == nil {
		return nil
	}
	return v.Message
}

func (v *ValidationErrorResponseError) GetFields() []FieldError {
	if v == nil {
		return nil
	}
	return v.Fields
}
