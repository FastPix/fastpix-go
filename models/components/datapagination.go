package components

// DataPagination - Pagination organizes content into pages for better readability and navigation.
type DataPagination struct {
	// The total number of records retrieved within the timeframe.
	//
	TotalRecords *int64 `json:"totalRecords,omitempty"`
	// The current offset value.
	//
	// Default: 1
	//
	CurrentOffset *int64 `json:"currentOffset,omitempty"`
	// The total number of offsets based on limit.
	//
	OffsetCount *int64 `json:"offsetCount,omitempty"`
}

func (d *DataPagination) GetTotalRecords() *int64 {
	if d == nil {
		return nil
	}
	return d.TotalRecords
}

func (d *DataPagination) GetCurrentOffset() *int64 {
	if d == nil {
		return nil
	}
	return d.CurrentOffset
}

func (d *DataPagination) GetOffsetCount() *int64 {
	if d == nil {
		return nil
	}
	return d.OffsetCount
}
