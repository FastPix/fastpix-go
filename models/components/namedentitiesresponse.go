

package components

type NamedEntitiesResponse struct {
	MediaID                  *string `json:"mediaId,omitempty"`
	IsGeneratedNamedEntities *bool   `json:"isGeneratedNamedEntities,omitempty"`
}

func (n *NamedEntitiesResponse) GetMediaID() *string {
	if n == nil {
		return nil
	}
	return n.MediaID
}

func (n *NamedEntitiesResponse) GetIsGeneratedNamedEntities() *bool {
	if n == nil {
		return nil
	}
	return n.IsGeneratedNamedEntities
}
