package components

type CreateMediaSuccessResponse struct {
	// Demonstrates whether the request is successful or not.
	Success bool                `json:"success"`
	Data    CreateMediaResponse `json:"data"`
}

func (c *CreateMediaSuccessResponse) GetSuccess() bool {
	if c == nil {
		return false
	}
	return c.Success
}

func (c *CreateMediaSuccessResponse) GetData() CreateMediaResponse {
	if c == nil {
		return CreateMediaResponse{}
	}
	return c.Data
}
