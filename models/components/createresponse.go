package components

type CreateResponse struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data *CreateSigningKeyResponseDTO `json:"data,omitempty"`
}

func (c *CreateResponse) GetSuccess() *bool {
	if c == nil {
		return nil
	}
	return c.Success
}

func (c *CreateResponse) GetData() *CreateSigningKeyResponseDTO {
	if c == nil {
		return nil
	}
	return c.Data
}
