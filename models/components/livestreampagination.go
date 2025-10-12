package components

// LiveStreamPagination - Pagination organizes content into pages for better readability and navigation.
type LiveStreamPagination struct {
	// It gives the total number of media assets that are accessible overall.
	TotalRecords *int64 `json:"totalRecords,omitempty"`
	// Determines the current point for data retrieval within a paginated list.
	CurrentOffset *int64 `json:"currentOffset,omitempty"`
	// The offset count is expressed as total records by limit.
	OffsetCount *int64 `json:"offsetCount,omitempty"`
}

func (l *LiveStreamPagination) GetTotalRecords() *int64 {
	if l == nil {
		return nil
	}
	return l.TotalRecords
}

func (l *LiveStreamPagination) GetCurrentOffset() *int64 {
	if l == nil {
		return nil
	}
	return l.CurrentOffset
}

func (l *LiveStreamPagination) GetOffsetCount() *int64 {
	if l == nil {
		return nil
	}
	return l.OffsetCount
}
