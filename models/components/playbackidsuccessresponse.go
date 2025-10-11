

package components

type PlaybackIDSuccessResponseData struct {
	// Unique identifier for the playbackId
	ID *string `json:"id,omitempty"`
	// Determines if access to the streamed content is kept private or available to all.
	AccessPolicy *string `json:"accessPolicy,omitempty"`
}

func (p *PlaybackIDSuccessResponseData) GetID() *string {
	if p == nil {
		return nil
	}
	return p.ID
}

func (p *PlaybackIDSuccessResponseData) GetAccessPolicy() *string {
	if p == nil {
		return nil
	}
	return p.AccessPolicy
}

// PlaybackIDSuccessResponse - Displays the result of the request.
type PlaybackIDSuccessResponse struct {
	// It demonstrates whether the request is successful or not.
	Success *bool                          `json:"success,omitempty"`
	Data    *PlaybackIDSuccessResponseData `json:"data,omitempty"`
}

func (p *PlaybackIDSuccessResponse) GetSuccess() *bool {
	if p == nil {
		return nil
	}
	return p.Success
}

func (p *PlaybackIDSuccessResponse) GetData() *PlaybackIDSuccessResponseData {
	if p == nil {
		return nil
	}
	return p.Data
}
