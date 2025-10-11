

package components

// MediaClipNotFoundError - Displays details about the reasons behind the request's failure.
type MediaClipNotFoundError struct {
	// Displays the error code indicating the type of the error.
	Code *float64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error.
	Message *string `json:"message,omitempty"`
	// A detailed explanation of the possible causes for the error.
	//
	Description *string `json:"description,omitempty"`
}

func (m *MediaClipNotFoundError) GetCode() *float64 {
	if m == nil {
		return nil
	}
	return m.Code
}

func (m *MediaClipNotFoundError) GetMessage() *string {
	if m == nil {
		return nil
	}
	return m.Message
}

func (m *MediaClipNotFoundError) GetDescription() *string {
	if m == nil {
		return nil
	}
	return m.Description
}
