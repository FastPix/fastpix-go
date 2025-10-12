

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type GetSigningKeyByIDRequest struct {
	// When creating the signing key, FastPix assigns a universally unique identifier with a maximum length of 255 characters.
	SigningKeyID string `pathParam:"style=simple,explode=false,name=signingKeyId"`
}

func (g *GetSigningKeyByIDRequest) GetSigningKeyID() string {
	if g == nil {
		return ""
	}
	return g.SigningKeyID
}

type GetSigningKeyByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// successfully fetched signing key
	GetPublicPemUsingSigningKeyIDResponseDTO *components.GetPublicPemUsingSigningKeyIDResponseDTO
}

func (g *GetSigningKeyByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GetSigningKeyByIDResponse) GetGetPublicPemUsingSigningKeyIDResponseDTO() *components.GetPublicPemUsingSigningKeyIDResponseDTO {
	if g == nil {
		return nil
	}
	return g.GetPublicPemUsingSigningKeyIDResponseDTO
}
