

package components

type UpdatePlaylistRequest struct {
	// New name to the playlist.
	Name string `json:"name"`
	// Updated description to the playlist.
	Description string `json:"description"`
}

func (u *UpdatePlaylistRequest) GetName() string {
	if u == nil {
		return ""
	}
	return u.Name
}

func (u *UpdatePlaylistRequest) GetDescription() string {
	if u == nil {
		return ""
	}
	return u.Description
}
