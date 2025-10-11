

package components

type LiveStreamDeleteResponse struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
}

func (l *LiveStreamDeleteResponse) GetSuccess() *bool {
	if l == nil {
		return nil
	}
	return l.Success
}
