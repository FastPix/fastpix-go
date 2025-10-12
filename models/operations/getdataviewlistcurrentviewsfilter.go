

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
	"github.com/FastPix/fastpix-go/models/components"
)

// GetDataViewlistCurrentViewsFilterDimension - The dimension to group and breakdown the concurrent viewers data by.
// This determines how the results will be categorized and aggregated.
// Choose from geographic, content, technical, or behavioral dimensions.
type GetDataViewlistCurrentViewsFilterDimension string

const (
	GetDataViewlistCurrentViewsFilterDimensionCountry               GetDataViewlistCurrentViewsFilterDimension = "country"
	GetDataViewlistCurrentViewsFilterDimensionRegion                GetDataViewlistCurrentViewsFilterDimension = "region"
	GetDataViewlistCurrentViewsFilterDimensionAsnID                 GetDataViewlistCurrentViewsFilterDimension = "asn_id"
	GetDataViewlistCurrentViewsFilterDimensionCdn                   GetDataViewlistCurrentViewsFilterDimension = "cdn"
	GetDataViewlistCurrentViewsFilterDimensionVideoTitle            GetDataViewlistCurrentViewsFilterDimension = "video_title"
	GetDataViewlistCurrentViewsFilterDimensionVideoSeries           GetDataViewlistCurrentViewsFilterDimension = "video_series"
	GetDataViewlistCurrentViewsFilterDimensionVideoID               GetDataViewlistCurrentViewsFilterDimension = "video_id"
	GetDataViewlistCurrentViewsFilterDimensionSubPropertyID         GetDataViewlistCurrentViewsFilterDimension = "sub_property_id"
	GetDataViewlistCurrentViewsFilterDimensionVideoSourceStreamType GetDataViewlistCurrentViewsFilterDimension = "video_source_stream_type"
	GetDataViewlistCurrentViewsFilterDimensionOsName                GetDataViewlistCurrentViewsFilterDimension = "os_name"
	GetDataViewlistCurrentViewsFilterDimensionPlayerName            GetDataViewlistCurrentViewsFilterDimension = "player_name"
	GetDataViewlistCurrentViewsFilterDimensionMediaID               GetDataViewlistCurrentViewsFilterDimension = "media_id"
	GetDataViewlistCurrentViewsFilterDimensionFpPlaybackID          GetDataViewlistCurrentViewsFilterDimension = "fp_playback_id"
	GetDataViewlistCurrentViewsFilterDimensionViewID                GetDataViewlistCurrentViewsFilterDimension = "view_id"
)

func (e GetDataViewlistCurrentViewsFilterDimension) ToPointer() *GetDataViewlistCurrentViewsFilterDimension {
	return &e
}
func (e *GetDataViewlistCurrentViewsFilterDimension) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "country":
		fallthrough
	case "region":
		fallthrough
	case "asn_id":
		fallthrough
	case "cdn":
		fallthrough
	case "video_title":
		fallthrough
	case "video_series":
		fallthrough
	case "video_id":
		fallthrough
	case "sub_property_id":
		fallthrough
	case "video_source_stream_type":
		fallthrough
	case "os_name":
		fallthrough
	case "player_name":
		fallthrough
	case "media_id":
		fallthrough
	case "fp_playback_id":
		fallthrough
	case "view_id":
		*e = GetDataViewlistCurrentViewsFilterDimension(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDataViewlistCurrentViewsFilterDimension: %v", v)
	}
}

type GetDataViewlistCurrentViewsFilterRequest struct {
	// The dimension to group and breakdown the concurrent viewers data by.
	// This determines how the results will be categorized and aggregated.
	// Choose from geographic, content, technical, or behavioral dimensions.
	//
	Dimension *GetDataViewlistCurrentViewsFilterDimension `queryParam:"style=form,explode=true,name=dimension"`
	// Maximum number of results to return. Controls the number of dimension values
	// that will be included in the response. Useful for pagination and performance.
	// Higher limits provide more detailed breakdowns but may impact response time.
	//
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
}

func (g GetDataViewlistCurrentViewsFilterRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetDataViewlistCurrentViewsFilterRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (g *GetDataViewlistCurrentViewsFilterRequest) GetDimension() *GetDataViewlistCurrentViewsFilterDimension {
	if g == nil {
		return nil
	}
	return g.Dimension
}

func (g *GetDataViewlistCurrentViewsFilterRequest) GetLimit() *int64 {
	if g == nil {
		return nil
	}
	return g.Limit
}

type GetDataViewlistCurrentViewsFilterData struct {
	// Number of concurrent viewers for this dimension value.
	ConcurrentViewers *int64 `json:"concurrent_viewers,omitempty"`
	// Name of the dimension (e.g., country, device type, etc.).
	DimensionName *string `json:"dimension_name,omitempty"`
	// Additional metric value for this dimension (if applicable).
	MetricValue *int64 `json:"metricValue,omitempty"`
}

func (g *GetDataViewlistCurrentViewsFilterData) GetConcurrentViewers() *int64 {
	if g == nil {
		return nil
	}
	return g.ConcurrentViewers
}

func (g *GetDataViewlistCurrentViewsFilterData) GetDimensionName() *string {
	if g == nil {
		return nil
	}
	return g.DimensionName
}

func (g *GetDataViewlistCurrentViewsFilterData) GetMetricValue() *int64 {
	if g == nil {
		return nil
	}
	return g.MetricValue
}

// GetDataViewlistCurrentViewsFilterResponseBody - Successfully retrieved concurrent viewers breakdown by the specified dimension.
type GetDataViewlistCurrentViewsFilterResponseBody struct {
	// Indicates if the request was successful.
	Success *bool `json:"success,omitempty"`
	// Breakdown of concurrent viewers by the specified dimension.
	Data []GetDataViewlistCurrentViewsFilterData `json:"data,omitempty"`
	// Start and end epoch timestamps (milliseconds) for the timespan window.
	Timespan []int64 `json:"timespan,omitempty"`
}

func (g *GetDataViewlistCurrentViewsFilterResponseBody) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetDataViewlistCurrentViewsFilterResponseBody) GetData() []GetDataViewlistCurrentViewsFilterData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetDataViewlistCurrentViewsFilterResponseBody) GetTimespan() []int64 {
	if g == nil {
		return nil
	}
	return g.Timespan
}

type GetDataViewlistCurrentViewsFilterResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successfully retrieved concurrent viewers breakdown by the specified dimension.
	Object *GetDataViewlistCurrentViewsFilterResponseBody
}

func (g *GetDataViewlistCurrentViewsFilterResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GetDataViewlistCurrentViewsFilterResponse) GetObject() *GetDataViewlistCurrentViewsFilterResponseBody {
	if g == nil {
		return nil
	}
	return g.Object
}
