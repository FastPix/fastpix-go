package components

// DateRange - Date range with start and end dates.
type DateRange struct {
	StartDate *string `json:"startDate,omitempty"`
	EndDate   *string `json:"endDate,omitempty"`
}

func (d *DateRange) GetStartDate() *string {
	if d == nil {
		return nil
	}
	return d.StartDate
}

func (d *DateRange) GetEndDate() *string {
	if d == nil {
		return nil
	}
	return d.EndDate
}
