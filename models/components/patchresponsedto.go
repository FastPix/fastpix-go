package components

// PatchResponseDTO - Displays the result of the request.
type PatchResponseDTO struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data *PatchResponseData `json:"data,omitempty"`
}

func (p *PatchResponseDTO) GetSuccess() *bool {
	if p == nil {
		return nil
	}
	return p.Success
}

func (p *PatchResponseDTO) GetData() *PatchResponseData {
	if p == nil {
		return nil
	}
	return p.Data
}
