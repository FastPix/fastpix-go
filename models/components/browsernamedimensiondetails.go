

package components

type BrowserNameDimensiondetails struct {
	// The specific metric value calculated based on the applied filters.
	Value string `json:"value"`
	// The count of unique viewers who interacted with the content.
	UniqueCount *int64 `json:"uniqueCount,omitempty"`
	// The count of viewers.
	Count int64 `json:"count"`
}

func (b *BrowserNameDimensiondetails) GetValue() string {
	if b == nil {
		return ""
	}
	return b.Value
}

func (b *BrowserNameDimensiondetails) GetUniqueCount() *int64 {
	if b == nil {
		return nil
	}
	return b.UniqueCount
}

func (b *BrowserNameDimensiondetails) GetCount() int64 {
	if b == nil {
		return 0
	}
	return b.Count
}
