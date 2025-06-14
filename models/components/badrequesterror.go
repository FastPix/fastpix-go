// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

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

func (o *BadRequestError) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *BadRequestError) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *BadRequestError) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
