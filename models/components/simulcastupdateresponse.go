

package components

// SimulcastUpdateResponseData - Displays the result of the request.
type SimulcastUpdateResponseData struct {
	// When you create the new simulcast, FastPix assign a universal unique identifier which can contain a maximum of 255 characters.
	SimulcastID *string `json:"simulcastId,omitempty"`
	// The RTMP hostname, combined with the application name, is crucial for connecting to third-party live streaming services and transmitting the live stream.
	URL *string `json:"url,omitempty"`
	// A unique stream key is generated for streaming, allowing the user to start streaming on any third-party platform using this key.
	StreamKey *string `json:"streamKey,omitempty"`
	// When the value is set to false, the simulcast will be disabled for the given stream
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// You can search for videos with specific key value pairs using metadata, when you tag a video in "key":"value"s pairs. Dynamic Metadata allows you to define a key that allows any value pair. You can have maximum of 255 characters and upto 10 entries are allowed.
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (s *SimulcastUpdateResponseData) GetSimulcastID() *string {
	if s == nil {
		return nil
	}
	return s.SimulcastID
}

func (s *SimulcastUpdateResponseData) GetURL() *string {
	if s == nil {
		return nil
	}
	return s.URL
}

func (s *SimulcastUpdateResponseData) GetStreamKey() *string {
	if s == nil {
		return nil
	}
	return s.StreamKey
}

func (s *SimulcastUpdateResponseData) GetIsEnabled() *bool {
	if s == nil {
		return nil
	}
	return s.IsEnabled
}

func (s *SimulcastUpdateResponseData) GetMetadata() map[string]string {
	if s == nil {
		return nil
	}
	return s.Metadata
}

// SimulcastUpdateResponse - Displays the result of the request.
type SimulcastUpdateResponse struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data *SimulcastUpdateResponseData `json:"data,omitempty"`
}

func (s *SimulcastUpdateResponse) GetSuccess() *bool {
	if s == nil {
		return nil
	}
	return s.Success
}

func (s *SimulcastUpdateResponse) GetData() *SimulcastUpdateResponseData {
	if s == nil {
		return nil
	}
	return s.Data
}
