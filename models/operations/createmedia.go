

package operations

import (
	"github.com/FastPix/fastpix-go/models/components"
)

type CreateMediaResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Media is created successfully
	CreateMediaSuccessResponse *components.CreateMediaSuccessResponse
}

func (c *CreateMediaResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CreateMediaResponse) GetCreateMediaSuccessResponse() *components.CreateMediaSuccessResponse {
	if c == nil {
		return nil
	}
	return c.CreateMediaSuccessResponse
}
