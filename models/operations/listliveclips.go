

package operations

import (
	"github.com/FastPix/fastpix-go/internal/utils"
	"github.com/FastPix/fastpix-go/models/components"
)

type ListLiveClipsRequest struct {
	// The stream Id is unique identifier assigned to the live stream.
	LivestreamID string `pathParam:"style=simple,explode=false,name=livestreamId"`
	// Limit specifies the maximum number of items to display per page.
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	// Offset determines the starting point for data retrieval within a paginated list.
	Offset *int64 `default:"1" queryParam:"style=form,explode=true,name=offset"`
	// The values in the list can be arranged in two ways: DESC (Descending) or ASC (Ascending).
	OrderBy *components.SortOrder `default:"desc" queryParam:"style=form,explode=true,name=orderBy"`
}

func (l ListLiveClipsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListLiveClipsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, []string{"livestreamId"}); err != nil {
		return err
	}
	return nil
}

func (l *ListLiveClipsRequest) GetLivestreamID() string {
	if l == nil {
		return ""
	}
	return l.LivestreamID
}

func (l *ListLiveClipsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListLiveClipsRequest) GetOffset() *int64 {
	if l == nil {
		return nil
	}
	return l.Offset
}

func (l *ListLiveClipsRequest) GetOrderBy() *components.SortOrder {
	if l == nil {
		return nil
	}
	return l.OrderBy
}

// ListLiveClipsResponseBody - List of video media
type ListLiveClipsResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data []components.Media `json:"data,omitempty"`
	// Pagination organizes content into pages for better readability and navigation.
	Pagination *components.Pagination `json:"pagination,omitempty"`
}

func (l *ListLiveClipsResponseBody) GetSuccess() *bool {
	if l == nil {
		return nil
	}
	return l.Success
}

func (l *ListLiveClipsResponseBody) GetData() []components.Media {
	if l == nil {
		return nil
	}
	return l.Data
}

func (l *ListLiveClipsResponseBody) GetPagination() *components.Pagination {
	if l == nil {
		return nil
	}
	return l.Pagination
}

type ListLiveClipsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List of video media
	Object *ListLiveClipsResponseBody
}

func (l *ListLiveClipsResponse) GetHTTPMeta() components.HTTPMetadata {
	if l == nil {
		return components.HTTPMetadata{}
	}
	return l.HTTPMeta
}

func (l *ListLiveClipsResponse) GetObject() *ListLiveClipsResponseBody {
	if l == nil {
		return nil
	}
	return l.Object
}
