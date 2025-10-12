package components

type ModerationResponse struct {
	MediaID             *string `json:"mediaId,omitempty"`
	IsModerationEnabled *bool   `json:"isModerationEnabled,omitempty"`
}

func (m *ModerationResponse) GetMediaID() *string {
	if m == nil {
		return nil
	}
	return m.MediaID
}

func (m *ModerationResponse) GetIsModerationEnabled() *bool {
	if m == nil {
		return nil
	}
	return m.IsModerationEnabled
}
