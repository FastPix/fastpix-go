package components

// SrtPlaybackResponse - This object contains the livestream playback response details for SRT Protocol
type SrtPlaybackResponse struct {
	// A unique identifier for the SRT playback stream. This ID is used to distinguish between different playback streams
	SrtPlaybackStreamID *string `json:"srtPlaybackStreamId,omitempty"`
	// A playback secret used for securing the SRT playback stream. This ensures that only authorized users can access the playback
	SrtPlaybackSecret *string `json:"srtPlaybackSecret,omitempty"`
}

func (s *SrtPlaybackResponse) GetSrtPlaybackStreamID() *string {
	if s == nil {
		return nil
	}
	return s.SrtPlaybackStreamID
}

func (s *SrtPlaybackResponse) GetSrtPlaybackSecret() *string {
	if s == nil {
		return nil
	}
	return s.SrtPlaybackSecret
}
