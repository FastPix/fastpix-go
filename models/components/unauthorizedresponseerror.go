

package components

// UnAuthorizedResponseError - Displays details about the reasons behind the request's failure.
type UnAuthorizedResponseError struct {
	// UnAuthorized response
	Code *int64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error.
	Message *string `json:"message,omitempty"`
}

func (u *UnAuthorizedResponseError) GetCode() *int64 {
	if u == nil {
		return nil
	}
	return u.Code
}

func (u *UnAuthorizedResponseError) GetMessage() *string {
	if u == nil {
		return nil
	}
	return u.Message
}
