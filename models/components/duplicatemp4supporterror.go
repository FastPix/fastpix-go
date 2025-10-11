

package components

// DuplicateMp4SupportError - Displays details about the reasons behind the request's failure.
type DuplicateMp4SupportError struct {
	// Displays the error code indicating the type of the error.
	Code *float64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error.
	Message *string `json:"message,omitempty"`
	// A detailed explanation of the possible causes for the error.
	//
	Description *string `json:"description,omitempty"`
}

func (d *DuplicateMp4SupportError) GetCode() *float64 {
	if d == nil {
		return nil
	}
	return d.Code
}

func (d *DuplicateMp4SupportError) GetMessage() *string {
	if d == nil {
		return nil
	}
	return d.Message
}

func (d *DuplicateMp4SupportError) GetDescription() *string {
	if d == nil {
		return nil
	}
	return d.Description
}
