

package components

type SuccessResponseData struct {
}

type SuccessResponse struct {
	// Demonstrates whether the request is successful or not.
	Success bool `json:"success"`
	// Array of response data
	Data []SuccessResponseData `json:"data"`
}

func (s *SuccessResponse) GetSuccess() bool {
	if s == nil {
		return false
	}
	return s.Success
}

func (s *SuccessResponse) GetData() []SuccessResponseData {
	if s == nil {
		return []SuccessResponseData{}
	}
	return s.Data
}
