

package components

// ViewNotFoundError - Returns the problem that has occured
type ViewNotFoundError struct {
	// An error code indicating the type of the error.
	Code *float64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error.
	Message *string `json:"message,omitempty"`
}

func (v *ViewNotFoundError) GetCode() *float64 {
	if v == nil {
		return nil
	}
	return v.Code
}

func (v *ViewNotFoundError) GetMessage() *string {
	if v == nil {
		return nil
	}
	return v.Message
}
