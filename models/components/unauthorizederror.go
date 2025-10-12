package components

// UnauthorizedError - Displays details about the reasons behind the request's failure.
type UnauthorizedError struct {
	// Displays the error code indicating the type of the error.
	Code *int64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error.
	Message *string `json:"message,omitempty"`
	// A detailed explanation of the possible causes for the error.
	//
	Description *string `json:"description,omitempty"`
}

func (u *UnauthorizedError) GetCode() *int64 {
	if u == nil {
		return nil
	}
	return u.Code
}

func (u *UnauthorizedError) GetMessage() *string {
	if u == nil {
		return nil
	}
	return u.Message
}

func (u *UnauthorizedError) GetDescription() *string {
	if u == nil {
		return nil
	}
	return u.Description
}
