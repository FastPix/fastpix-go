package components

type SimulcastRequest struct {
	// The RTMPS hostname, combined with the application name, is crucial for connecting to third-party live streaming services and transmitting the live stream.
	URL *string `json:"url,omitempty"`
	// A unique stream key is generated for streaming, allowing the user to start streaming on any third-party platform using this key.
	StreamKey *string `json:"streamKey,omitempty"`
	// You can search for videos with specific key value pairs using metadata, when you tag a video in "key":"value"s pairs. Dynamic Metadata allows you to define a key that allows any value pair. You can have maximum of 255 characters and upto 10 entries are allowed.
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (s *SimulcastRequest) GetURL() *string {
	if s == nil {
		return nil
	}
	return s.URL
}

func (s *SimulcastRequest) GetStreamKey() *string {
	if s == nil {
		return nil
	}
	return s.StreamKey
}

func (s *SimulcastRequest) GetMetadata() map[string]string {
	if s == nil {
		return nil
	}
	return s.Metadata
}
