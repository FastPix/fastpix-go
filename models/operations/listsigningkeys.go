

package operations

import (
	"github.com/fastpix/fastpix-go/internal/utils"
	"github.com/fastpix/fastpix-go/models/components"
)

type ListSigningKeysRequest struct {
	// Limit specifies the maximum number of items to display per page.
	Limit *float64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	// It is used for pagination, indicating the starting point for fetching data.
	Offset *float64 `default:"1" queryParam:"style=form,explode=true,name=offset"`
}

func (l ListSigningKeysRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListSigningKeysRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListSigningKeysRequest) GetLimit() *float64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListSigningKeysRequest) GetOffset() *float64 {
	if l == nil {
		return nil
	}
	return l.Offset
}

type ListSigningKeysResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// successfully fetched all signing keys
	GetAllSigningKeyResponse *components.GetAllSigningKeyResponse
}

func (l *ListSigningKeysResponse) GetHTTPMeta() components.HTTPMetadata {
	if l == nil {
		return components.HTTPMetadata{}
	}
	return l.HTTPMeta
}

func (l *ListSigningKeysResponse) GetGetAllSigningKeyResponse() *components.GetAllSigningKeyResponse {
	if l == nil {
		return nil
	}
	return l.GetAllSigningKeyResponse
}
