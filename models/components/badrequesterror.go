

package components

// BadRequestError - Displays details about the reasons behind the request's failure.
type BadRequestError struct {
	// Displays the error code indicating the type of the error.
	Code *float64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error.
	Message *string `json:"message,omitempty"`
	// A detailed explanation of the possible causes for the error.
	//
	Description *string `json:"description,omitempty"`
}

func (b *BadRequestError) GetCode() *float64 {
	if b == nil {
		return nil
	}
	return b.Code
}

func (b *BadRequestError) GetMessage() *string {
	if b == nil {
		return nil
	}
	return b.Message
}

func (b *BadRequestError) GetDescription() *string {
	if b == nil {
		return nil
	}
	return b.Description
}
