

package components

// InvalidPlaylistIDResponseError - Displays details about the reasons behind the request's failure.
type InvalidPlaylistIDResponseError struct {
	// Displays the error code indicating the type of the error.
	Code *int64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error.
	Message *string `json:"message,omitempty"`
	// It is an collection of objects, where each object contains information about a specific field and a corresponding error message.
	Fields []FieldError `json:"fields,omitempty"`
}

func (i *InvalidPlaylistIDResponseError) GetCode() *int64 {
	if i == nil {
		return nil
	}
	return i.Code
}

func (i *InvalidPlaylistIDResponseError) GetMessage() *string {
	if i == nil {
		return nil
	}
	return i.Message
}

func (i *InvalidPlaylistIDResponseError) GetFields() []FieldError {
	if i == nil {
		return nil
	}
	return i.Fields
}
