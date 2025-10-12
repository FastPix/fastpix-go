package components

type FieldError struct {
	// Displays the specific field associated with the error.
	Field string `json:"field"`
	// Error message for the field
	Message string `json:"message"`
}

func (f *FieldError) GetField() string {
	if f == nil {
		return ""
	}
	return f.Field
}

func (f *FieldError) GetMessage() string {
	if f == nil {
		return ""
	}
	return f.Message
}
