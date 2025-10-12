package components

// MetricsOverallMetaDataDetails - Metadata that has to be paased for metric calculations.
type MetricsOverallMetaDataDetails struct {
	// defines the field or dimension on which the aggregation is to be   applied.
	Aggregation *string `json:"aggregation,omitempty"`
}

func (m *MetricsOverallMetaDataDetails) GetAggregation() *string {
	if m == nil {
		return nil
	}
	return m.Aggregation
}
