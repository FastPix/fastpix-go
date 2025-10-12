package components

import (
	"encoding/json"
	"fmt"
)

// AddTrackResponseType - Specifies the type of track (audio or subtitle).
type AddTrackResponseType string

const (
	AddTrackResponseTypeAudio    AddTrackResponseType = "audio"
	AddTrackResponseTypeSubtitle AddTrackResponseType = "subtitle"
)

func (e AddTrackResponseType) ToPointer() *AddTrackResponseType {
	return &e
}
func (e *AddTrackResponseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "audio":
		fallthrough
	case "subtitle":
		*e = AddTrackResponseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddTrackResponseType: %v", v)
	}
}

// AddTrackResponse - Contains details about the track that was added or updated.
type AddTrackResponse struct {
	// The unique identifier of the track.
	ID *string `json:"id,omitempty"`
	// Specifies the type of track (audio or subtitle).
	Type *AddTrackResponseType `json:"type,omitempty"`
	// The direct URL of the track file.
	URL *string `json:"url,omitempty"`
	// The BCP 47 language code representing the track's language.
	LanguageCode *string `json:"languageCode,omitempty"`
	// The full name of the language corresponding to the `languageCode`.
	LanguageName *string `json:"languageName,omitempty"`
}

func (a *AddTrackResponse) GetID() *string {
	if a == nil {
		return nil
	}
	return a.ID
}

func (a *AddTrackResponse) GetType() *AddTrackResponseType {
	if a == nil {
		return nil
	}
	return a.Type
}

func (a *AddTrackResponse) GetURL() *string {
	if a == nil {
		return nil
	}
	return a.URL
}

func (a *AddTrackResponse) GetLanguageCode() *string {
	if a == nil {
		return nil
	}
	return a.LanguageCode
}

func (a *AddTrackResponse) GetLanguageName() *string {
	if a == nil {
		return nil
	}
	return a.LanguageName
}
