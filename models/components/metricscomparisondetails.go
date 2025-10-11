

package components

type Item struct {
	// The specific metric value calculated based on the applied filters.
	Value *int64 `json:"value,omitempty"`
	// value can be score that ranges from 0 to 100
	Type *string `json:"type,omitempty"`
	// value can be score that ranges from 0 to 100
	Name *string `json:"name,omitempty"`
	// The metric field represents the name of the Key Performance Indicator (KPI) being tracked or analyzed. It identifies a specific measurable aspect of the video playback experience, such as buffering time, video start failure rate, or playback quality.
	//
	Metric *string `json:"metric,omitempty"`
	// value can be avg, sum, count or 95th
	Measurement *string `json:"measurement,omitempty"`
}

func (i *Item) GetValue() *int64 {
	if i == nil {
		return nil
	}
	return i.Value
}

func (i *Item) GetType() *string {
	if i == nil {
		return nil
	}
	return i.Type
}

func (i *Item) GetName() *string {
	if i == nil {
		return nil
	}
	return i.Name
}

func (i *Item) GetMetric() *string {
	if i == nil {
		return nil
	}
	return i.Metric
}

func (i *Item) GetMeasurement() *string {
	if i == nil {
		return nil
	}
	return i.Measurement
}

type MetricsComparisonDetails struct {
	// The specific metric value calculated based on the applied filters.
	Value *int64 `json:"value,omitempty"`
	// value can be score that ranges from 0 to 100
	Type *string `json:"type,omitempty"`
	// value can be score that ranges from 0 to 100
	Name *string `json:"name,omitempty"`
	// The metric field represents the name of the Key Performance Indicator (KPI) being tracked or analyzed. It identifies a specific measurable aspect of the video playback experience, such as buffering time, video start failure rate, or playback quality.
	//
	Metric *string `json:"metric,omitempty"`
	Items  []Item  `json:"items,omitempty"`
}

func (m *MetricsComparisonDetails) GetValue() *int64 {
	if m == nil {
		return nil
	}
	return m.Value
}

func (m *MetricsComparisonDetails) GetType() *string {
	if m == nil {
		return nil
	}
	return m.Type
}

func (m *MetricsComparisonDetails) GetName() *string {
	if m == nil {
		return nil
	}
	return m.Name
}

func (m *MetricsComparisonDetails) GetMetric() *string {
	if m == nil {
		return nil
	}
	return m.Metric
}

func (m *MetricsComparisonDetails) GetItems() []Item {
	if m == nil {
		return nil
	}
	return m.Items
}
