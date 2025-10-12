package components

type GetAllPlaylistsResponse struct {
	Success *bool          `json:"success,omitempty"`
	Data    []PlaylistItem `json:"data,omitempty"`
	// Pagination organizes content into pages for better readability and navigation.
	Pagination *Pagination `json:"pagination,omitempty"`
}

func (g *GetAllPlaylistsResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetAllPlaylistsResponse) GetData() []PlaylistItem {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetAllPlaylistsResponse) GetPagination() *Pagination {
	if g == nil {
		return nil
	}
	return g.Pagination
}
