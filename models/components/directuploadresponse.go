

package components

type DirectUploadResponse struct {
	PlaybackIds []PlaybackID `json:"playbackIds,omitempty"`
	// You can search for videos with specific key value pairs using metadata, when you tag a video in "key" : "value" pairs. Dynamic Metadata allows you to define a key that allows any value pair. You can have maximum of 255 characters and upto 10 entries are allowed.
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (d *DirectUploadResponse) GetPlaybackIds() []PlaybackID {
	if d == nil {
		return nil
	}
	return d.PlaybackIds
}

func (d *DirectUploadResponse) GetMetadata() map[string]string {
	if d == nil {
		return nil
	}
	return d.Metadata
}
