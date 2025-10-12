package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type CreateSigningKeyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// created a signing key successfully
	CreateResponse *components.CreateResponse
}

func (c *CreateSigningKeyResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CreateSigningKeyResponse) GetCreateResponse() *components.CreateResponse {
	if c == nil {
		return nil
	}
	return c.CreateResponse
}
