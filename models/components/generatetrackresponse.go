package components

import (
	"encoding/json"
	"fmt"
)

// GenerateTrackResponseType - The type of track generated ("subtitle").
type GenerateTrackResponseType string

const (
	GenerateTrackResponseTypeSubtitle GenerateTrackResponseType = "subtitle"
)

func (e GenerateTrackResponseType) ToPointer() *GenerateTrackResponseType {
	return &e
}
func (e *GenerateTrackResponseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "subtitle":
		*e = GenerateTrackResponseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GenerateTrackResponseType: %v", v)
	}
}

// GenerateTrackResponseLanguageCode - The BCP 47 language code representing the language of the generated track.
type GenerateTrackResponseLanguageCode string

const (
	GenerateTrackResponseLanguageCodeArSa GenerateTrackResponseLanguageCode = "ar-SA"
	GenerateTrackResponseLanguageCodeBnBd GenerateTrackResponseLanguageCode = "bn-BD"
	GenerateTrackResponseLanguageCodeBnIn GenerateTrackResponseLanguageCode = "bn-IN"
	GenerateTrackResponseLanguageCodeCaEs GenerateTrackResponseLanguageCode = "ca-ES"
	GenerateTrackResponseLanguageCodeCsCz GenerateTrackResponseLanguageCode = "cs-CZ"
	GenerateTrackResponseLanguageCodeDaDk GenerateTrackResponseLanguageCode = "da-DK"
	GenerateTrackResponseLanguageCodeDeAt GenerateTrackResponseLanguageCode = "de-AT"
	GenerateTrackResponseLanguageCodeDeCh GenerateTrackResponseLanguageCode = "de-CH"
	GenerateTrackResponseLanguageCodeDeDe GenerateTrackResponseLanguageCode = "de-DE"
	GenerateTrackResponseLanguageCodeElGr GenerateTrackResponseLanguageCode = "el-GR"
	GenerateTrackResponseLanguageCodeEnAu GenerateTrackResponseLanguageCode = "en-AU"
	GenerateTrackResponseLanguageCodeEnCa GenerateTrackResponseLanguageCode = "en-CA"
	GenerateTrackResponseLanguageCodeEnGb GenerateTrackResponseLanguageCode = "en-GB"
	GenerateTrackResponseLanguageCodeEnIe GenerateTrackResponseLanguageCode = "en-IE"
	GenerateTrackResponseLanguageCodeEnIn GenerateTrackResponseLanguageCode = "en-IN"
	GenerateTrackResponseLanguageCodeEnNz GenerateTrackResponseLanguageCode = "en-NZ"
	GenerateTrackResponseLanguageCodeEnUs GenerateTrackResponseLanguageCode = "en-US"
	GenerateTrackResponseLanguageCodeEnZa GenerateTrackResponseLanguageCode = "en-ZA"
	GenerateTrackResponseLanguageCodeEsAr GenerateTrackResponseLanguageCode = "es-AR"
	GenerateTrackResponseLanguageCodeEsCl GenerateTrackResponseLanguageCode = "es-CL"
	GenerateTrackResponseLanguageCodeEsCo GenerateTrackResponseLanguageCode = "es-CO"
	GenerateTrackResponseLanguageCodeEsEs GenerateTrackResponseLanguageCode = "es-ES"
	GenerateTrackResponseLanguageCodeEsMx GenerateTrackResponseLanguageCode = "es-MX"
	GenerateTrackResponseLanguageCodeEsUs GenerateTrackResponseLanguageCode = "es-US"
	GenerateTrackResponseLanguageCodeFiFi GenerateTrackResponseLanguageCode = "fi-FI"
	GenerateTrackResponseLanguageCodeFrBe GenerateTrackResponseLanguageCode = "fr-BE"
	GenerateTrackResponseLanguageCodeFrCa GenerateTrackResponseLanguageCode = "fr-CA"
	GenerateTrackResponseLanguageCodeFrCh GenerateTrackResponseLanguageCode = "fr-CH"
	GenerateTrackResponseLanguageCodeFrFr GenerateTrackResponseLanguageCode = "fr-FR"
	GenerateTrackResponseLanguageCodeHeIl GenerateTrackResponseLanguageCode = "he-IL"
	GenerateTrackResponseLanguageCodeHiIn GenerateTrackResponseLanguageCode = "hi-IN"
	GenerateTrackResponseLanguageCodeHrHr GenerateTrackResponseLanguageCode = "hr-HR"
	GenerateTrackResponseLanguageCodeHuHu GenerateTrackResponseLanguageCode = "hu-HU"
	GenerateTrackResponseLanguageCodeIDID GenerateTrackResponseLanguageCode = "id-ID"
	GenerateTrackResponseLanguageCodeItCh GenerateTrackResponseLanguageCode = "it-CH"
	GenerateTrackResponseLanguageCodeItIt GenerateTrackResponseLanguageCode = "it-IT"
	GenerateTrackResponseLanguageCodeJaJp GenerateTrackResponseLanguageCode = "ja-JP"
	GenerateTrackResponseLanguageCodeKoKr GenerateTrackResponseLanguageCode = "ko-KR"
	GenerateTrackResponseLanguageCodeNlBe GenerateTrackResponseLanguageCode = "nl-BE"
	GenerateTrackResponseLanguageCodeNlNl GenerateTrackResponseLanguageCode = "nl-NL"
	GenerateTrackResponseLanguageCodeNoNo GenerateTrackResponseLanguageCode = "no-NO"
	GenerateTrackResponseLanguageCodePlPl GenerateTrackResponseLanguageCode = "pl-PL"
	GenerateTrackResponseLanguageCodePtBr GenerateTrackResponseLanguageCode = "pt-BR"
	GenerateTrackResponseLanguageCodePtPt GenerateTrackResponseLanguageCode = "pt-PT"
	GenerateTrackResponseLanguageCodeRoRo GenerateTrackResponseLanguageCode = "ro-RO"
	GenerateTrackResponseLanguageCodeRuRu GenerateTrackResponseLanguageCode = "ru-RU"
	GenerateTrackResponseLanguageCodeSkSk GenerateTrackResponseLanguageCode = "sk-SK"
	GenerateTrackResponseLanguageCodeSvSe GenerateTrackResponseLanguageCode = "sv-SE"
	GenerateTrackResponseLanguageCodeTaIn GenerateTrackResponseLanguageCode = "ta-IN"
	GenerateTrackResponseLanguageCodeTaLk GenerateTrackResponseLanguageCode = "ta-LK"
	GenerateTrackResponseLanguageCodeThTh GenerateTrackResponseLanguageCode = "th-TH"
	GenerateTrackResponseLanguageCodeTrTr GenerateTrackResponseLanguageCode = "tr-TR"
	GenerateTrackResponseLanguageCodeUkUa GenerateTrackResponseLanguageCode = "uk-UA"
	GenerateTrackResponseLanguageCodeBgBg GenerateTrackResponseLanguageCode = "bg-BG"
	GenerateTrackResponseLanguageCodeZhCn GenerateTrackResponseLanguageCode = "zh-CN"
	GenerateTrackResponseLanguageCodeZhHk GenerateTrackResponseLanguageCode = "zh-HK"
	GenerateTrackResponseLanguageCodeZhTw GenerateTrackResponseLanguageCode = "zh-TW"
)

func (e GenerateTrackResponseLanguageCode) ToPointer() *GenerateTrackResponseLanguageCode {
	return &e
}
func (e *GenerateTrackResponseLanguageCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ar-SA":
		fallthrough
	case "bn-BD":
		fallthrough
	case "bn-IN":
		fallthrough
	case "ca-ES":
		fallthrough
	case "cs-CZ":
		fallthrough
	case "da-DK":
		fallthrough
	case "de-AT":
		fallthrough
	case "de-CH":
		fallthrough
	case "de-DE":
		fallthrough
	case "el-GR":
		fallthrough
	case "en-AU":
		fallthrough
	case "en-CA":
		fallthrough
	case "en-GB":
		fallthrough
	case "en-IE":
		fallthrough
	case "en-IN":
		fallthrough
	case "en-NZ":
		fallthrough
	case "en-US":
		fallthrough
	case "en-ZA":
		fallthrough
	case "es-AR":
		fallthrough
	case "es-CL":
		fallthrough
	case "es-CO":
		fallthrough
	case "es-ES":
		fallthrough
	case "es-MX":
		fallthrough
	case "es-US":
		fallthrough
	case "fi-FI":
		fallthrough
	case "fr-BE":
		fallthrough
	case "fr-CA":
		fallthrough
	case "fr-CH":
		fallthrough
	case "fr-FR":
		fallthrough
	case "he-IL":
		fallthrough
	case "hi-IN":
		fallthrough
	case "hr-HR":
		fallthrough
	case "hu-HU":
		fallthrough
	case "id-ID":
		fallthrough
	case "it-CH":
		fallthrough
	case "it-IT":
		fallthrough
	case "ja-JP":
		fallthrough
	case "ko-KR":
		fallthrough
	case "nl-BE":
		fallthrough
	case "nl-NL":
		fallthrough
	case "no-NO":
		fallthrough
	case "pl-PL":
		fallthrough
	case "pt-BR":
		fallthrough
	case "pt-PT":
		fallthrough
	case "ro-RO":
		fallthrough
	case "ru-RU":
		fallthrough
	case "sk-SK":
		fallthrough
	case "sv-SE":
		fallthrough
	case "ta-IN":
		fallthrough
	case "ta-LK":
		fallthrough
	case "th-TH":
		fallthrough
	case "tr-TR":
		fallthrough
	case "uk-UA":
		fallthrough
	case "bg-BG":
		fallthrough
	case "zh-CN":
		fallthrough
	case "zh-HK":
		fallthrough
	case "zh-TW":
		*e = GenerateTrackResponseLanguageCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GenerateTrackResponseLanguageCode: %v", v)
	}
}

// GenerateTrackResponse - Represents the response for a successfully generated subtitle track.
type GenerateTrackResponse struct {
	// A unique identifier for the generated track.
	ID *string `json:"id,omitempty"`
	// The type of track generated ("subtitle").
	Type *GenerateTrackResponseType `json:"type,omitempty"`
	// The BCP 47 language code representing the language of the generated track.
	//
	LanguageCode *GenerateTrackResponseLanguageCode `json:"languageCode,omitempty"`
	// The full name of the language for the generated track.
	LanguageName *string `json:"languageName,omitempty"`
	// You can search for videos with specific key value pairs using metadata, when you tag a video in "key" : "value" pairs. Dynamic Metadata allows you to define a key that allows any value pair. You can have maximum of 255 characters and upto 10 entries are allowed.
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (g *GenerateTrackResponse) GetID() *string {
	if g == nil {
		return nil
	}
	return g.ID
}

func (g *GenerateTrackResponse) GetType() *GenerateTrackResponseType {
	if g == nil {
		return nil
	}
	return g.Type
}

func (g *GenerateTrackResponse) GetLanguageCode() *GenerateTrackResponseLanguageCode {
	if g == nil {
		return nil
	}
	return g.LanguageCode
}

func (g *GenerateTrackResponse) GetLanguageName() *string {
	if g == nil {
		return nil
	}
	return g.LanguageName
}

func (g *GenerateTrackResponse) GetMetadata() map[string]string {
	if g == nil {
		return nil
	}
	return g.Metadata
}
