

package components

// TrackSubtitlesGenerateRequest - Contains details for generating subtitle tracks for a media file.
type TrackSubtitlesGenerateRequest struct {
	// The full name of the language in which subtitles will be generated.
	LanguageName string `json:"languageName"`
	// You can search for videos with specific key value pairs using metadata, when you tag a video in "key" : "value" pairs. Dynamic Metadata allows you to define a key that allows any value pair. You can have maximum of 255 characters and upto 10 entries are allowed.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Language code for content localization
	LanguageCode LanguageCode `json:"languageCode"`
}

func (t *TrackSubtitlesGenerateRequest) GetLanguageName() string {
	if t == nil {
		return ""
	}
	return t.LanguageName
}

func (t *TrackSubtitlesGenerateRequest) GetMetadata() map[string]string {
	if t == nil {
		return nil
	}
	return t.Metadata
}

func (t *TrackSubtitlesGenerateRequest) GetLanguageCode() LanguageCode {
	if t == nil {
		return LanguageCode("")
	}
	return t.LanguageCode
}
