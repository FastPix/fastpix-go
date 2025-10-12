package components

import (
	"github.com/FastPix/fastpix-go/internal/utils"
)

// SubtitleInput - Generates subtitle files for audio/video files.
type SubtitleInput struct {
	// Defines the type of input.
	//
	Type string `json:"type"`
	// The direct URL of the subtitle file.
	URL string `json:"url"`
	// Name of the language in which the subtitles will be generated.
	LanguageName string `json:"languageName"`
	// Language code for content localization
	LanguageCode LanguageCode `json:"languageCode"`
}

func (s SubtitleInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SubtitleInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, []string{"type", "url", "languageName", "languageCode"}); err != nil {
		return err
	}
	return nil
}

func (s *SubtitleInput) GetType() string {
	if s == nil {
		return ""
	}
	return s.Type
}

func (s *SubtitleInput) GetURL() string {
	if s == nil {
		return ""
	}
	return s.URL
}

func (s *SubtitleInput) GetLanguageName() string {
	if s == nil {
		return ""
	}
	return s.LanguageName
}

func (s *SubtitleInput) GetLanguageCode() LanguageCode {
	if s == nil {
		return LanguageCode("")
	}
	return s.LanguageCode
}
