package components

// SingingKeysPagination - Pagination organizes content into pages for better readability and navigation.
type SingingKeysPagination struct {
	// It gives the total number of media assets that are accessible overall.
	TotalRecords *int64 `json:"totalRecords,omitempty"`
	// Determines the current point for data retrieval within a paginated list.
	CurrentOffset *int64 `json:"currentOffset,omitempty"`
	// The offset count is expressed as total records by limit.
	OffsetCount *int64 `json:"offsetCount,omitempty"`
}

func (s *SingingKeysPagination) GetTotalRecords() *int64 {
	if s == nil {
		return nil
	}
	return s.TotalRecords
}

func (s *SingingKeysPagination) GetCurrentOffset() *int64 {
	if s == nil {
		return nil
	}
	return s.CurrentOffset
}

func (s *SingingKeysPagination) GetOffsetCount() *int64 {
	if s == nil {
		return nil
	}
	return s.OffsetCount
}
