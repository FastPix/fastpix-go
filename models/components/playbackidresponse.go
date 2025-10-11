

package components

// PlaybackIDResponse - A collection of Playback ID objects utilized for crafting HLS playback urls.
type PlaybackIDResponse struct {
	// Unique identifier for the playbackId
	ID *string `json:"id,omitempty"`
	// Determines if access to the streamed content is kept private or available to all.
	AccessPolicy *string `json:"accessPolicy,omitempty"`
}

func (p *PlaybackIDResponse) GetID() *string {
	if p == nil {
		return nil
	}
	return p.ID
}

func (p *PlaybackIDResponse) GetAccessPolicy() *string {
	if p == nil {
		return nil
	}
	return p.AccessPolicy
}
