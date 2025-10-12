package components

type SummaryResponse struct {
	MediaID            *string `json:"mediaId,omitempty"`
	IsGeneratedSummary *bool   `json:"isGeneratedSummary,omitempty"`
}

func (s *SummaryResponse) GetMediaID() *string {
	if s == nil {
		return nil
	}
	return s.MediaID
}

func (s *SummaryResponse) GetIsGeneratedSummary() *bool {
	if s == nil {
		return nil
	}
	return s.IsGeneratedSummary
}
