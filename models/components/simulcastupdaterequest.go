package components

type SimulcastUpdateRequest struct {
	// When the value is set to false, the simulcast will be disabled for the given stream.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// You can search for videos with specific key value pairs using metadata, when you tag a video in "key":"value"s pairs. Dynamic Metadata allows you to define a key that allows any value pair. You can have maximum of 255 characters and upto 10 entries are allowed.
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (s *SimulcastUpdateRequest) GetIsEnabled() *bool {
	if s == nil {
		return nil
	}
	return s.IsEnabled
}

func (s *SimulcastUpdateRequest) GetMetadata() map[string]string {
	if s == nil {
		return nil
	}
	return s.Metadata
}
