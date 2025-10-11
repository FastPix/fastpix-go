

package components

type GetAllSigningKeyResponse struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data []GetAllSigningKeyResponseDTO `json:"data,omitempty"`
}

func (g *GetAllSigningKeyResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetAllSigningKeyResponse) GetData() []GetAllSigningKeyResponseDTO {
	if g == nil {
		return nil
	}
	return g.Data
}
