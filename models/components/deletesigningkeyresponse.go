package components

type DeleteSigningKeyResponse struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
}

func (d *DeleteSigningKeyResponse) GetSuccess() *bool {
	if d == nil {
		return nil
	}
	return d.Success
}
