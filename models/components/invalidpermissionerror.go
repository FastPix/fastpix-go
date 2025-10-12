package components

// InvalidPermissionError - Displays details about the reasons behind the request's failure.
type InvalidPermissionError struct {
	// Displays the error code indicating the type of the error.
	Code *int64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error.
	Message *string `json:"message,omitempty"`
	// A detailed explanation of the possible causes for the error.
	//
	Description *string `json:"description,omitempty"`
}

func (i *InvalidPermissionError) GetCode() *int64 {
	if i == nil {
		return nil
	}
	return i.Code
}

func (i *InvalidPermissionError) GetMessage() *string {
	if i == nil {
		return nil
	}
	return i.Message
}

func (i *InvalidPermissionError) GetDescription() *string {
	if i == nil {
		return nil
	}
	return i.Description
}
