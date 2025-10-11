

package components

// NotFoundErrorPlaybackIDError - Displays details about the reasons behind the request's failure.
type NotFoundErrorPlaybackIDError struct {
	// Displays the error code indicating the type of the error.
	Code *float64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error.
	Message *string `json:"message,omitempty"`
	// A detailed explanation of the possible causes for the error.
	//
	Description *string `json:"description,omitempty"`
}

func (n *NotFoundErrorPlaybackIDError) GetCode() *float64 {
	if n == nil {
		return nil
	}
	return n.Code
}

func (n *NotFoundErrorPlaybackIDError) GetMessage() *string {
	if n == nil {
		return nil
	}
	return n.Message
}

func (n *NotFoundErrorPlaybackIDError) GetDescription() *string {
	if n == nil {
		return nil
	}
	return n.Description
}
