package components

// PlaybackSettings - Displays the result of the playback settings.
type PlaybackSettings struct {
	// Basic access policy for media content
	AccessPolicy *BasicAccessPolicy `json:"accessPolicy,omitempty"`
}

func (p *PlaybackSettings) GetAccessPolicy() *BasicAccessPolicy {
	if p == nil {
		return nil
	}
	return p.AccessPolicy
}
