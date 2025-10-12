

package components

import (
	"errors"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
	"github.com/FastPix/fastpix-go/optionalnullable"
)

type MetricsBreakdownDetailsValueType string

const (
	MetricsBreakdownDetailsValueTypeInteger MetricsBreakdownDetailsValueType = "integer"
	MetricsBreakdownDetailsValueTypeNumber  MetricsBreakdownDetailsValueType = "number"
)

// MetricsBreakdownDetailsValue - The specific metric value calculated based on the applied filters.
type MetricsBreakdownDetailsValue struct {
	Integer *int64   `queryParam:"inline,name=value"`
	Number  *float64 `queryParam:"inline,name=value"`

	Type MetricsBreakdownDetailsValueType
}

func CreateMetricsBreakdownDetailsValueInteger(integer int64) MetricsBreakdownDetailsValue {
	typ := MetricsBreakdownDetailsValueTypeInteger

	return MetricsBreakdownDetailsValue{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateMetricsBreakdownDetailsValueNumber(number float64) MetricsBreakdownDetailsValue {
	typ := MetricsBreakdownDetailsValueTypeNumber

	return MetricsBreakdownDetailsValue{
		Number: &number,
		Type:   typ,
	}
}

func (u *MetricsBreakdownDetailsValue) UnmarshalJSON(data []byte) error {

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, nil); err == nil {
		u.Integer = &integer
		u.Type = MetricsBreakdownDetailsValueTypeInteger
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, nil); err == nil {
		u.Number = &number
		u.Type = MetricsBreakdownDetailsValueTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for MetricsBreakdownDetailsValue", string(data))
}

func (u MetricsBreakdownDetailsValue) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type MetricsBreakdownDetailsValue: all fields are null")
}

type MetricsBreakdownDetails struct {
	// Total count of view sessions for a paricular video content.
	Views optionalnullable.OptionalNullable[int64] `json:"views,omitempty"`
	// The specific metric value calculated based on the applied filters.
	Value *MetricsBreakdownDetailsValue `json:"value"`
	// Total time watched across all views, represented in milliseconds.
	TotalWatchTime optionalnullable.OptionalNullable[int64] `json:"totalWatchTime,omitempty"`
	// Total time spent playing the video, represented in milliseconds.
	TotalPlayingTime optionalnullable.OptionalNullable[int64] `json:"totalPlayingTime,omitempty"`
	// the value of dimension or filter value on which the aggregation is to be applied.
	Field optionalnullable.OptionalNullable[string] `json:"field,omitempty"`
}

func (m *MetricsBreakdownDetails) GetViews() optionalnullable.OptionalNullable[int64] {
	if m == nil {
		return nil
	}
	return m.Views
}

func (m *MetricsBreakdownDetails) GetValue() *MetricsBreakdownDetailsValue {
	if m == nil {
		return nil
	}
	return m.Value
}

func (m *MetricsBreakdownDetails) GetTotalWatchTime() optionalnullable.OptionalNullable[int64] {
	if m == nil {
		return nil
	}
	return m.TotalWatchTime
}

func (m *MetricsBreakdownDetails) GetTotalPlayingTime() optionalnullable.OptionalNullable[int64] {
	if m == nil {
		return nil
	}
	return m.TotalPlayingTime
}

func (m *MetricsBreakdownDetails) GetField() optionalnullable.OptionalNullable[string] {
	if m == nil {
		return nil
	}
	return m.Field
}
