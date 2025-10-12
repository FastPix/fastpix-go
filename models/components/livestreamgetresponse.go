package components

// LivestreamgetResponse - Displays the result of the request.
type LivestreamgetResponse struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data *GetCreateLiveStreamResponseDTO `json:"data,omitempty"`
}

func (l *LivestreamgetResponse) GetSuccess() *bool {
	if l == nil {
		return nil
	}
	return l.Success
}

func (l *LivestreamgetResponse) GetData() *GetCreateLiveStreamResponseDTO {
	if l == nil {
		return nil
	}
	return l.Data
}
