

package components

// UpdateTrackRequest - Contains details about the track being added to the media file.
type UpdateTrackRequest struct {
	// The direct URL of the track file. It should point to a valid audio or subtitle file.
	URL *string `json:"url,omitempty"`
	// The BCP 47 language code representing the track's language.
	LanguageCode *string `json:"languageCode,omitempty"`
	// The full name of the language corresponding to the `languageCode`.
	LanguageName *string `json:"languageName,omitempty"`
}

func (u *UpdateTrackRequest) GetURL() *string {
	if u == nil {
		return nil
	}
	return u.URL
}

func (u *UpdateTrackRequest) GetLanguageCode() *string {
	if u == nil {
		return nil
	}
	return u.LanguageCode
}

func (u *UpdateTrackRequest) GetLanguageName() *string {
	if u == nil {
		return nil
	}
	return u.LanguageName
}
