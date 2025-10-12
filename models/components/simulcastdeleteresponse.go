package components

type SimulcastdeleteResponse struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
}

func (s *SimulcastdeleteResponse) GetSuccess() *bool {
	if s == nil {
		return nil
	}
	return s.Success
}
