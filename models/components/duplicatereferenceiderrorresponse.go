

package components

// DuplicateReferenceIDErrorResponseError - Displays details about the reasons behind the request's failure.
type DuplicateReferenceIDErrorResponseError struct {
	// Displays the error code indicating the type of the error.
	Code *int64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error.
	Message *string `json:"message,omitempty"`
	// A detailed explanation of the possible causes for the error.
	//
	Description *string `json:"description,omitempty"`
}

func (d *DuplicateReferenceIDErrorResponseError) GetCode() *int64 {
	if d == nil {
		return nil
	}
	return d.Code
}

func (d *DuplicateReferenceIDErrorResponseError) GetMessage() *string {
	if d == nil {
		return nil
	}
	return d.Message
}

func (d *DuplicateReferenceIDErrorResponseError) GetDescription() *string {
	if d == nil {
		return nil
	}
	return d.Description
}
