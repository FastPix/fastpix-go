

package components

// ForbiddenError - Displays details about the reasons behind the request's failure.
type ForbiddenError struct {
	// Displays the error code indicating the type of the error.
	Code *float64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error.
	Message *string `json:"message,omitempty"`
	// A detailed explanation of the possible causes for the error.
	//
	Description *string `json:"description,omitempty"`
}

func (f *ForbiddenError) GetCode() *float64 {
	if f == nil {
		return nil
	}
	return f.Code
}

func (f *ForbiddenError) GetMessage() *string {
	if f == nil {
		return nil
	}
	return f.Message
}

func (f *ForbiddenError) GetDescription() *string {
	if f == nil {
		return nil
	}
	return f.Description
}
