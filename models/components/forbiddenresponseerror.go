package components

// ForbiddenResponseError - Displays details about the reasons behind the request's failure.
type ForbiddenResponseError struct {
	// Forbidden response
	Code *int64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error.
	Message *string `json:"message,omitempty"`
}

func (f *ForbiddenResponseError) GetCode() *int64 {
	if f == nil {
		return nil
	}
	return f.Code
}

func (f *ForbiddenResponseError) GetMessage() *string {
	if f == nil {
		return nil
	}
	return f.Message
}
