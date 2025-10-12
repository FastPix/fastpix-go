package operations

import (
	"github.com/FastPix/fastpix-go/internal/utils"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/optionalnullable"
	"time"
)

type GetDataViewlistCurrentViewsGetTimeseriesViewsData struct {
	// The timestamp for the interval (ISO 8601 format).
	IntervalTime *time.Time `json:"intervalTime,omitempty"`
	// Reserved for future metric values (currently null).
	MetricValue optionalnullable.OptionalNullable[int64] `json:"metricValue,omitempty"`
	// Number of concurrent viewers at the given interval.
	NumberOfViews *int64 `json:"numberOfViews,omitempty"`
}

func (g GetDataViewlistCurrentViewsGetTimeseriesViewsData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetDataViewlistCurrentViewsGetTimeseriesViewsData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (g *GetDataViewlistCurrentViewsGetTimeseriesViewsData) GetIntervalTime() *time.Time {
	if g == nil {
		return nil
	}
	return g.IntervalTime
}

func (g *GetDataViewlistCurrentViewsGetTimeseriesViewsData) GetMetricValue() optionalnullable.OptionalNullable[int64] {
	if g == nil {
		return nil
	}
	return g.MetricValue
}

func (g *GetDataViewlistCurrentViewsGetTimeseriesViewsData) GetNumberOfViews() *int64 {
	if g == nil {
		return nil
	}
	return g.NumberOfViews
}

// GetDataViewlistCurrentViewsGetTimeseriesViewsResponseBody - Successfully retrieved concurrent viewers timeseries.
type GetDataViewlistCurrentViewsGetTimeseriesViewsResponseBody struct {
	// Indicates if the request was successful.
	Success *bool `json:"success,omitempty"`
	// Time series data for concurrent viewers.
	Data []GetDataViewlistCurrentViewsGetTimeseriesViewsData `json:"data,omitempty"`
	// Start and end epoch timestamps (milliseconds) for the timeseries window.
	Timespan []int64 `json:"timespan,omitempty"`
}

func (g *GetDataViewlistCurrentViewsGetTimeseriesViewsResponseBody) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetDataViewlistCurrentViewsGetTimeseriesViewsResponseBody) GetData() []GetDataViewlistCurrentViewsGetTimeseriesViewsData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetDataViewlistCurrentViewsGetTimeseriesViewsResponseBody) GetTimespan() []int64 {
	if g == nil {
		return nil
	}
	return g.Timespan
}

type GetDataViewlistCurrentViewsGetTimeseriesViewsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successfully retrieved concurrent viewers timeseries.
	Object *GetDataViewlistCurrentViewsGetTimeseriesViewsResponseBody
}

func (g *GetDataViewlistCurrentViewsGetTimeseriesViewsResponse) GetHTTPMeta() components.HTTPMetadata {
	if g == nil {
		return components.HTTPMetadata{}
	}
	return g.HTTPMeta
}

func (g *GetDataViewlistCurrentViewsGetTimeseriesViewsResponse) GetObject() *GetDataViewlistCurrentViewsGetTimeseriesViewsResponseBody {
	if g == nil {
		return nil
	}
	return g.Object
}
