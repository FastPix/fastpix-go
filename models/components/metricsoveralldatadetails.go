

package components

import (
	"errors"
	"fmt"
	"github.com/fastpix/fastpix-go/internal/utils"
	"github.com/fastpix/fastpix-go/optionalnullable"
)

type MetricsOverallDataDetailsValueType string

const (
	MetricsOverallDataDetailsValueTypeInteger MetricsOverallDataDetailsValueType = "integer"
	MetricsOverallDataDetailsValueTypeNumber  MetricsOverallDataDetailsValueType = "number"
)

// MetricsOverallDataDetailsValue - metric value calculated based on the applied filters.
type MetricsOverallDataDetailsValue struct {
	Integer *int64   `queryParam:"inline,name=value"`
	Number  *float64 `queryParam:"inline,name=value"`

	Type MetricsOverallDataDetailsValueType
}

func CreateMetricsOverallDataDetailsValueInteger(integer int64) MetricsOverallDataDetailsValue {
	typ := MetricsOverallDataDetailsValueTypeInteger

	return MetricsOverallDataDetailsValue{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateMetricsOverallDataDetailsValueNumber(number float64) MetricsOverallDataDetailsValue {
	typ := MetricsOverallDataDetailsValueTypeNumber

	return MetricsOverallDataDetailsValue{
		Number: &number,
		Type:   typ,
	}
}

func (u *MetricsOverallDataDetailsValue) UnmarshalJSON(data []byte) error {

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, nil); err == nil {
		u.Integer = &integer
		u.Type = MetricsOverallDataDetailsValueTypeInteger
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, nil); err == nil {
		u.Number = &number
		u.Type = MetricsOverallDataDetailsValueTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for MetricsOverallDataDetailsValue", string(data))
}

func (u MetricsOverallDataDetailsValue) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type MetricsOverallDataDetailsValue: all fields are null")
}

type GlobalValueType string

const (
	GlobalValueTypeInteger GlobalValueType = "integer"
	GlobalValueTypeNumber  GlobalValueType = "number"
)

// GlobalValue - A global metric value that reflects the overall performance of the specified metric across the entire dataset for the given timespan.
type GlobalValue struct {
	Integer *int64   `queryParam:"inline,name=globalValue"`
	Number  *float64 `queryParam:"inline,name=globalValue"`

	Type GlobalValueType
}

func CreateGlobalValueInteger(integer int64) GlobalValue {
	typ := GlobalValueTypeInteger

	return GlobalValue{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateGlobalValueNumber(number float64) GlobalValue {
	typ := GlobalValueTypeNumber

	return GlobalValue{
		Number: &number,
		Type:   typ,
	}
}

func (u *GlobalValue) UnmarshalJSON(data []byte) error {

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, nil); err == nil {
		u.Integer = &integer
		u.Type = GlobalValueTypeInteger
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, nil); err == nil {
		u.Number = &number
		u.Type = GlobalValueTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GlobalValue", string(data))
}

func (u GlobalValue) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type GlobalValue: all fields are null")
}

// MetricsOverallDataDetails - Retrieves overall values for a specified metric
type MetricsOverallDataDetails struct {
	// metric value calculated based on the applied filters.
	Value optionalnullable.OptionalNullable[MetricsOverallDataDetailsValue] `json:"value,omitempty"`
	// Total time watched across all views, represented in milliseconds.
	TotalWatchTime optionalnullable.OptionalNullable[int64] `json:"totalWatchTime,omitempty"`
	// The count of unique viewers who interacted with the content.
	UniqueViews optionalnullable.OptionalNullable[int64] `json:"uniqueViews,omitempty"`
	// The total number of views recorded.
	TotalViews optionalnullable.OptionalNullable[int64] `json:"totalViews,omitempty"`
	// Total time spent playing the video, represented in milliseconds.
	TotalPlayTime optionalnullable.OptionalNullable[int64] `json:"totalPlayTime,omitempty"`
	// A global metric value that reflects the overall performance of the specified metric across the entire dataset for the given timespan.
	GlobalValue optionalnullable.OptionalNullable[GlobalValue] `json:"globalValue,omitempty"`
}

func (m *MetricsOverallDataDetails) GetValue() optionalnullable.OptionalNullable[MetricsOverallDataDetailsValue] {
	if m == nil {
		return nil
	}
	return m.Value
}

func (m *MetricsOverallDataDetails) GetTotalWatchTime() optionalnullable.OptionalNullable[int64] {
	if m == nil {
		return nil
	}
	return m.TotalWatchTime
}

func (m *MetricsOverallDataDetails) GetUniqueViews() optionalnullable.OptionalNullable[int64] {
	if m == nil {
		return nil
	}
	return m.UniqueViews
}

func (m *MetricsOverallDataDetails) GetTotalViews() optionalnullable.OptionalNullable[int64] {
	if m == nil {
		return nil
	}
	return m.TotalViews
}

func (m *MetricsOverallDataDetails) GetTotalPlayTime() optionalnullable.OptionalNullable[int64] {
	if m == nil {
		return nil
	}
	return m.TotalPlayTime
}

func (m *MetricsOverallDataDetails) GetGlobalValue() optionalnullable.OptionalNullable[GlobalValue] {
	if m == nil {
		return nil
	}
	return m.GlobalValue
}
