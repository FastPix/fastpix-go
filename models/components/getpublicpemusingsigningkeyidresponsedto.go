

package components

// GetPublicPemUsingSigningKeyIDResponseDTOData - Displays the result of the request.
type GetPublicPemUsingSigningKeyIDResponseDTOData struct {
	// FastPix generates a unique identifier for each workspace.
	WorkspaceID  *string `json:"workspaceId,omitempty"`
	SigningKeyID *string `json:"signingKeyId,omitempty"`
	// A public key is a byte encoded key used to create a signed JSON Web Token (JWT) for authentication.
	PublicKey *string `json:"publicKey,omitempty"`
}

func (g *GetPublicPemUsingSigningKeyIDResponseDTOData) GetWorkspaceID() *string {
	if g == nil {
		return nil
	}
	return g.WorkspaceID
}

func (g *GetPublicPemUsingSigningKeyIDResponseDTOData) GetSigningKeyID() *string {
	if g == nil {
		return nil
	}
	return g.SigningKeyID
}

func (g *GetPublicPemUsingSigningKeyIDResponseDTOData) GetPublicKey() *string {
	if g == nil {
		return nil
	}
	return g.PublicKey
}

// GetPublicPemUsingSigningKeyIDResponseDTO - Displays the result of the request.
type GetPublicPemUsingSigningKeyIDResponseDTO struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data *GetPublicPemUsingSigningKeyIDResponseDTOData `json:"data,omitempty"`
}

func (g *GetPublicPemUsingSigningKeyIDResponseDTO) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetPublicPemUsingSigningKeyIDResponseDTO) GetData() *GetPublicPemUsingSigningKeyIDResponseDTOData {
	if g == nil {
		return nil
	}
	return g.Data
}
