package components

// Pagination organizes content into pages for better readability and navigation.
type Pagination struct {
	// It gives the total number of media assets that are accessible overall.
	TotalRecords *int64 `json:"totalRecords,omitempty"`
	// Offset determines the current point for data retrieval within a paginated list.
	CurrentOffset *int64 `json:"currentOffset,omitempty"`
	// The offset count is expressed as total records by limit
	OffsetCount *int64 `json:"offsetCount,omitempty"`
}

func (p *Pagination) GetTotalRecords() *int64 {
	if p == nil {
		return nil
	}
	return p.TotalRecords
}

func (p *Pagination) GetCurrentOffset() *int64 {
	if p == nil {
		return nil
	}
	return p.CurrentOffset
}

func (p *Pagination) GetOffsetCount() *int64 {
	if p == nil {
		return nil
	}
	return p.OffsetCount
}
