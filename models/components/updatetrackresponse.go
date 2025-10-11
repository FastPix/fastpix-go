

package components

import (
	"encoding/json"
	"fmt"
)

// UpdateTrackResponseType - Specifies the type of track (audio or subtitle).
type UpdateTrackResponseType string

const (
	UpdateTrackResponseTypeAudio    UpdateTrackResponseType = "audio"
	UpdateTrackResponseTypeSubtitle UpdateTrackResponseType = "subtitle"
)

func (e UpdateTrackResponseType) ToPointer() *UpdateTrackResponseType {
	return &e
}
func (e *UpdateTrackResponseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "audio":
		fallthrough
	case "subtitle":
		*e = UpdateTrackResponseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateTrackResponseType: %v", v)
	}
}

// UpdateTrackResponse - Contains details about the track that was added or updated.
type UpdateTrackResponse struct {
	// The unique identifier of the track.
	ID *string `json:"id,omitempty"`
	// Specifies the type of track (audio or subtitle).
	Type *UpdateTrackResponseType `json:"type,omitempty"`
	// The direct URL of the track file.
	URL *string `json:"url,omitempty"`
	// The BCP 47 language code representing the track's language.
	LanguageCode *string `json:"languageCode,omitempty"`
	// The full name of the language corresponding to the `languageCode`.
	LanguageName *string `json:"languageName,omitempty"`
}

func (u *UpdateTrackResponse) GetID() *string {
	if u == nil {
		return nil
	}
	return u.ID
}

func (u *UpdateTrackResponse) GetType() *UpdateTrackResponseType {
	if u == nil {
		return nil
	}
	return u.Type
}

func (u *UpdateTrackResponse) GetURL() *string {
	if u == nil {
		return nil
	}
	return u.URL
}

func (u *UpdateTrackResponse) GetLanguageCode() *string {
	if u == nil {
		return nil
	}
	return u.LanguageCode
}

func (u *UpdateTrackResponse) GetLanguageName() *string {
	if u == nil {
		return nil
	}
	return u.LanguageName
}
