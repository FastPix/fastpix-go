package components

import (
	"encoding/json"
	"fmt"
)

// AddTrackRequestType - Specifies the type of track being added. It can be either `audio` or `subtitle`.
type AddTrackRequestType string

const (
	AddTrackRequestTypeAudio    AddTrackRequestType = "audio"
	AddTrackRequestTypeSubtitle AddTrackRequestType = "subtitle"
)

func (e AddTrackRequestType) ToPointer() *AddTrackRequestType {
	return &e
}
func (e *AddTrackRequestType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "audio":
		fallthrough
	case "subtitle":
		*e = AddTrackRequestType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddTrackRequestType: %v", v)
	}
}

// AddTrackRequest - Contains details about the track being added to the media file.
type AddTrackRequest struct {
	// The direct URL of the track file. It should point to a valid audio or subtitle file.
	URL *string `json:"url,omitempty"`
	// Specifies the type of track being added. It can be either `audio` or `subtitle`.
	Type *AddTrackRequestType `json:"type,omitempty"`
	// The BCP 47 language code representing the track's language.
	LanguageCode *string `json:"languageCode,omitempty"`
	// The full name of the language corresponding to the `languageCode`.
	LanguageName *string `json:"languageName,omitempty"`
}

func (a *AddTrackRequest) GetURL() *string {
	if a == nil {
		return nil
	}
	return a.URL
}

func (a *AddTrackRequest) GetType() *AddTrackRequestType {
	if a == nil {
		return nil
	}
	return a.Type
}

func (a *AddTrackRequest) GetLanguageCode() *string {
	if a == nil {
		return nil
	}
	return a.LanguageCode
}

func (a *AddTrackRequest) GetLanguageName() *string {
	if a == nil {
		return nil
	}
	return a.LanguageName
}
