

package components

// MetricsTimeseriesMetaDataDetails - Retrieves breakdown values for a specified metric and timespan
type MetricsTimeseriesMetaDataDetails struct {
	// the unit for aggregating the timeseries data.
	Granularity *string `json:"granularity,omitempty"`
	// defines the field or dimension on which the aggregation is to be   applied.
	Aggregation *string `json:"aggregation,omitempty"`
}

func (m *MetricsTimeseriesMetaDataDetails) GetGranularity() *string {
	if m == nil {
		return nil
	}
	return m.Granularity
}

func (m *MetricsTimeseriesMetaDataDetails) GetAggregation() *string {
	if m == nil {
		return nil
	}
	return m.Aggregation
}
