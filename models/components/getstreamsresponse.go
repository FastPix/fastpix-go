package components

// GetStreamsResponse - Displays the result of the request.
type GetStreamsResponse struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data []GetCreateLiveStreamResponseDTO `json:"data,omitempty"`
	// Pagination organizes content into pages for better readability and navigation.
	Pagination *LiveStreamPagination `json:"pagination,omitempty"`
}

func (g *GetStreamsResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetStreamsResponse) GetData() []GetCreateLiveStreamResponseDTO {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetStreamsResponse) GetPagination() *LiveStreamPagination {
	if g == nil {
		return nil
	}
	return g.Pagination
}
