

package components

type ChaptersResponse struct {
	MediaID             *string `json:"mediaId,omitempty"`
	IsGeneratedChapters *bool   `json:"isGeneratedChapters,omitempty"`
}

func (c *ChaptersResponse) GetMediaID() *string {
	if c == nil {
		return nil
	}
	return c.MediaID
}

func (c *ChaptersResponse) GetIsGeneratedChapters() *bool {
	if c == nil {
		return nil
	}
	return c.IsGeneratedChapters
}
