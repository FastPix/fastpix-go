package components

// MediaIdsRequest - The list of mediaId(s) you want to perform the operation on.rds by limit.
type MediaIdsRequest struct {
	MediaIds []string `json:"mediaIds"`
}

func (m *MediaIdsRequest) GetMediaIds() []string {
	if m == nil {
		return []string{}
	}
	return m.MediaIds
}
