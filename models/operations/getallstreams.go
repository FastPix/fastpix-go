

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/fastpix/fastpix-go/internal/utils"
	"github.com/fastpix/fastpix-go/models/components"
)

// OrderBy - The list of value can be order in two ways DESC (Descending) or ASC (Ascending). In case not specified, by default it will be DESC.
type OrderBy string

const (
	OrderByAsc  OrderBy = "asc"
	OrderByDesc OrderBy = "desc"
)

func (e OrderBy) ToPointer() *OrderBy {
	return &e
}
func (e *OrderBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = OrderBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OrderBy: %v", v)
	}
}

type GetAllStreamsRequest struct {
	// Limit specifies the maximum number of items to display per page.
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	// Offset determines the starting point for data retrieval within a paginated list.
	Offset *int64 `default:"1" queryParam:"style=form,explode=true,name=offset"`
	// The list of value can be order in two ways DESC (Descending) or ASC (Ascending). In case not specified, by default it will be DESC.
	OrderBy *OrderBy `default:"desc" queryParam:"style=form,explode=true,name=orderBy"`
}

func (g GetAllStreamsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetAllStreamsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (g *GetAllStreamsRequest) GetLimit() *int64 {
	if g == nil {
		return nil
	}
	return g.Limit
}

func (g *GetAllStreamsRequest) GetOffset() *int64 {
	if g == nil {
		return nil
	}
	return g.Offset
}

func (g *GetAllStreamsRequest) GetOrderBy() *OrderBy {
	if g == nil {
		return nil
	}
	return g.OrderBy
}

type GetAllStreamsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// All streams retrieved sucessfully
	GetStreamsResponse *components.GetStreamsResponse
}

func (g *GetAllStreamsResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GetAllStreamsResponse) GetGetStreamsResponse() *components.GetStreamsResponse {
	if g == nil {
		return nil
	}
	return g.GetStreamsResponse
}
