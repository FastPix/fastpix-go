package operations

import (
	"github.com/FastPix/fastpix-go/internal/utils"
	"github.com/FastPix/fastpix-go/models/components"
)

type ListUploadsRequest struct {
	// Limit specifies the maximum number of items to display per page.
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	// Offset determines the starting point for data retrieval within a paginated list.
	Offset *int64 `default:"1" queryParam:"style=form,explode=true,name=offset"`
	// The values in the list can be arranged in two ways: DESC (Descending) or ASC (Ascending).
	OrderBy *components.SortOrder `default:"desc" queryParam:"style=form,explode=true,name=orderBy"`
}

func (l ListUploadsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListUploadsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListUploadsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListUploadsRequest) GetOffset() *int64 {
	if l == nil {
		return nil
	}
	return l.Offset
}

func (l *ListUploadsRequest) GetOrderBy() *components.SortOrder {
	if l == nil {
		return nil
	}
	return l.OrderBy
}

// ListUploadsResponseBody - List of video media
type ListUploadsResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data []components.DirectUpload `json:"data,omitempty"`
	// Pagination organizes content into pages for better readability and navigation.
	Pagination *components.Pagination `json:"pagination,omitempty"`
}

func (l *ListUploadsResponseBody) GetSuccess() *bool {
	if l == nil {
		return nil
	}
	return l.Success
}

func (l *ListUploadsResponseBody) GetData() []components.DirectUpload {
	if l == nil {
		return nil
	}
	return l.Data
}

func (l *ListUploadsResponseBody) GetPagination() *components.Pagination {
	if l == nil {
		return nil
	}
	return l.Pagination
}

type ListUploadsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List of video media
	Object *ListUploadsResponseBody
}

func (l *ListUploadsResponse) GetHTTPMeta() components.HTTPMetadata {
	if l == nil {
		return components.HTTPMetadata{}
	}
	return l.HTTPMeta
}

func (l *ListUploadsResponse) GetObject() *ListUploadsResponseBody {
	if l == nil {
		return nil
	}
	return l.Object
}
