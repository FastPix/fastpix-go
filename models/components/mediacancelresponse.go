

package components

// MediaCancelResponse - Response returned when an upload is cancelled.
type MediaCancelResponse struct {
	// The unique identifier of the cancelled upload.
	UploadID *string `json:"uploadId,omitempty"`
	// Indicates if the upload was a trial.
	Trial *bool `json:"trial,omitempty"`
	// The status of the upload after cancellation.
	Status *string `json:"status,omitempty"`
	// The upload URL (if available) after cancellation.
	URL *string `json:"url,omitempty"`
	// The timeout value for the upload.
	Timeout *int64 `json:"timeout,omitempty"`
	// CORS origin allowed for the upload.
	CorsOrigin *string `json:"corsOrigin,omitempty"`
	// The maximum resolution allowed for the upload.
	MaxResolution *string `json:"maxResolution,omitempty"`
	// The access policy for the upload.
	AccessPolicy *string `json:"accessPolicy,omitempty"`
	// You can search for videos with specific key value pairs using metadata, when you tag a video in "key" : "value" pairs. Dynamic Metadata allows you to define a key that allows any value pair. You can have maximum of 255 characters and upto 10 entries are allowed.
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (m *MediaCancelResponse) GetUploadID() *string {
	if m == nil {
		return nil
	}
	return m.UploadID
}

func (m *MediaCancelResponse) GetTrial() *bool {
	if m == nil {
		return nil
	}
	return m.Trial
}

func (m *MediaCancelResponse) GetStatus() *string {
	if m == nil {
		return nil
	}
	return m.Status
}

func (m *MediaCancelResponse) GetURL() *string {
	if m == nil {
		return nil
	}
	return m.URL
}

func (m *MediaCancelResponse) GetTimeout() *int64 {
	if m == nil {
		return nil
	}
	return m.Timeout
}

func (m *MediaCancelResponse) GetCorsOrigin() *string {
	if m == nil {
		return nil
	}
	return m.CorsOrigin
}

func (m *MediaCancelResponse) GetMaxResolution() *string {
	if m == nil {
		return nil
	}
	return m.MaxResolution
}

func (m *MediaCancelResponse) GetAccessPolicy() *string {
	if m == nil {
		return nil
	}
	return m.AccessPolicy
}

func (m *MediaCancelResponse) GetMetadata() map[string]string {
	if m == nil {
		return nil
	}
	return m.Metadata
}
