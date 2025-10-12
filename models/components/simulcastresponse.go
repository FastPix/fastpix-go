package components

// SimulcastResponseData - Displays the result of the request.
type SimulcastResponseData struct {
	// When you create the new simulcast, FastPix assign a universal unique identifier which can contain a maximum of 255 characters.
	SimulcastID *string `json:"simulcastId,omitempty"`
	// The RTMPS hostname, combined with the application name, is crucial for connecting to third-party live streaming services and transmitting the live stream.
	URL *string `json:"url,omitempty"`
	// A unique stream key is generated for streaming, allowing the user to start streaming on any third-party platform using this key.
	StreamKey *string `json:"streamKey,omitempty"`
	// When the value is true, the simulcast will be enabled for the given stream
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// You can search for videos with specific key value pairs using metadata, when you tag a video in "key":"value"s pairs. Dynamic Metadata allows you to define a key that allows any value pair. You can have maximum of 255 characters and upto 10 entries are allowed.
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (s *SimulcastResponseData) GetSimulcastID() *string {
	if s == nil {
		return nil
	}
	return s.SimulcastID
}

func (s *SimulcastResponseData) GetURL() *string {
	if s == nil {
		return nil
	}
	return s.URL
}

func (s *SimulcastResponseData) GetStreamKey() *string {
	if s == nil {
		return nil
	}
	return s.StreamKey
}

func (s *SimulcastResponseData) GetIsEnabled() *bool {
	if s == nil {
		return nil
	}
	return s.IsEnabled
}

func (s *SimulcastResponseData) GetMetadata() map[string]string {
	if s == nil {
		return nil
	}
	return s.Metadata
}

// SimulcastResponse - Displays the result of the request.
type SimulcastResponse struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data *SimulcastResponseData `json:"data,omitempty"`
}

func (s *SimulcastResponse) GetSuccess() *bool {
	if s == nil {
		return nil
	}
	return s.Success
}

func (s *SimulcastResponse) GetData() *SimulcastResponseData {
	if s == nil {
		return nil
	}
	return s.Data
}
