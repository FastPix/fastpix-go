

package components

// TrackDuplicateRequestError - Displays details about the reasons behind the request's failure.
type TrackDuplicateRequestError struct {
	// Displays the error code indicating the type of the error.
	Code *int64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error.
	Message *string `json:"message,omitempty"`
	// A detailed explanation of the possible causes for the error.
	//
	Description *string `json:"description,omitempty"`
}

func (t *TrackDuplicateRequestError) GetCode() *int64 {
	if t == nil {
		return nil
	}
	return t.Code
}

func (t *TrackDuplicateRequestError) GetMessage() *string {
	if t == nil {
		return nil
	}
	return t.Message
}

func (t *TrackDuplicateRequestError) GetDescription() *string {
	if t == nil {
		return nil
	}
	return t.Description
}
