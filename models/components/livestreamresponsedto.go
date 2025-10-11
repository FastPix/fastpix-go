

package components

// LiveStreamResponseDTO - Displays the result of the request.
type LiveStreamResponseDTO struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data *GetCreateLiveStreamResponseDTO `json:"data,omitempty"`
}

func (l *LiveStreamResponseDTO) GetSuccess() *bool {
	if l == nil {
		return nil
	}
	return l.Success
}

func (l *LiveStreamResponseDTO) GetData() *GetCreateLiveStreamResponseDTO {
	if l == nil {
		return nil
	}
	return l.Data
}
