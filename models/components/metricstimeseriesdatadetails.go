

package components

import (
	"errors"
	"fmt"
	"github.com/fastpix/fastpix-go/internal/utils"
	"github.com/fastpix/fastpix-go/optionalnullable"
	"time"
)

type MetricValueType string

const (
	MetricValueTypeInteger MetricValueType = "integer"
	MetricValueTypeNumber  MetricValueType = "number"
)

// MetricValue - The value of the specified metric at the given interval.
type MetricValue struct {
	Integer *int64   `queryParam:"inline,name=metricValue"`
	Number  *float64 `queryParam:"inline,name=metricValue"`

	Type MetricValueType
}

func CreateMetricValueInteger(integer int64) MetricValue {
	typ := MetricValueTypeInteger

	return MetricValue{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateMetricValueNumber(number float64) MetricValue {
	typ := MetricValueTypeNumber

	return MetricValue{
		Number: &number,
		Type:   typ,
	}
}

func (u *MetricValue) UnmarshalJSON(data []byte) error {

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, nil); err == nil {
		u.Integer = &integer
		u.Type = MetricValueTypeInteger
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, nil); err == nil {
		u.Number = &number
		u.Type = MetricValueTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for MetricValue", string(data))
}

func (u MetricValue) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type MetricValue: all fields are null")
}

// MetricsTimeseriesDataDetails - The metric's value at specific time intervals.
type MetricsTimeseriesDataDetails struct {
	// The timestamp for the data point indicating when the metric value was recorded.
	IntervalTime *time.Time `json:"intervalTime,omitempty"`
	// The value of the specified metric at the given interval.
	MetricValue optionalnullable.OptionalNullable[MetricValue] `json:"metricValue,omitempty"`
	// The total number of views recorded during that interval.
	NumberOfViews optionalnullable.OptionalNullable[int64] `json:"numberOfViews,omitempty"`
}

func (m MetricsTimeseriesDataDetails) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MetricsTimeseriesDataDetails) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (m *MetricsTimeseriesDataDetails) GetIntervalTime() *time.Time {
	if m == nil {
		return nil
	}
	return m.IntervalTime
}

func (m *MetricsTimeseriesDataDetails) GetMetricValue() optionalnullable.OptionalNullable[MetricValue] {
	if m == nil {
		return nil
	}
	return m.MetricValue
}

func (m *MetricsTimeseriesDataDetails) GetNumberOfViews() optionalnullable.OptionalNullable[int64] {
	if m == nil {
		return nil
	}
	return m.NumberOfViews
}
