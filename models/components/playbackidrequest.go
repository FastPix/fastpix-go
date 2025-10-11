

package components

type PlaybackIDRequest struct {
	// Basic access policy for media content
	AccessPolicy *BasicAccessPolicy `json:"accessPolicy,omitempty"`
}

func (p *PlaybackIDRequest) GetAccessPolicy() *BasicAccessPolicy {
	if p == nil {
		return nil
	}
	return p.AccessPolicy
}
