

package components

// PlaylistCreatedResponse - Displays the result of the request.
type PlaylistCreatedResponse struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data *PlaylistCreatedSchema `json:"data,omitempty"`
}

func (p *PlaylistCreatedResponse) GetSuccess() *bool {
	if p == nil {
		return nil
	}
	return p.Success
}

func (p *PlaylistCreatedResponse) GetData() *PlaylistCreatedSchema {
	if p == nil {
		return nil
	}
	return p.Data
}
