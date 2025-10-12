package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type DeleteSigningKeyRequest struct {
	// When creating the signing key, FastPix assigns a universally unique identifier with a maximum length of 255 characters.
	SigningKeyID string `pathParam:"style=simple,explode=false,name=signingKeyId"`
}

func (d *DeleteSigningKeyRequest) GetSigningKeyID() string {
	if d == nil {
		return ""
	}
	return d.SigningKeyID
}

type DeleteSigningKeyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// successfully fetched all signing keys
	DeleteSigningKeyResponse *components.DeleteSigningKeyResponse
}

func (d *DeleteSigningKeyResponse) GetHTTPMeta() components.HTTPMetadata {
	if d == nil {
		return components.HTTPMetadata{}
	}
	return d.HTTPMeta
}

func (d *DeleteSigningKeyResponse) GetDeleteSigningKeyResponse() *components.DeleteSigningKeyResponse {
	if d == nil {
		return nil
	}
	return d.DeleteSigningKeyResponse
}
