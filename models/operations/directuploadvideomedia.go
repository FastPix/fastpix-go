

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
	"github.com/FastPix/fastpix-go/models/components"
)

type InputType string

const (
	InputTypeVideoInput     InputType = "VideoInput"
	InputTypeWatermarkInput InputType = "WatermarkInput"
	InputTypeAudioInput     InputType = "AudioInput"
	InputTypeSubtitleInput  InputType = "SubtitleInput"
)

type Input struct {
	VideoInput     *components.VideoInput     `queryParam:"inline,name=input"`
	WatermarkInput *components.WatermarkInput `queryParam:"inline,name=input"`
	AudioInput     *components.AudioInput     `queryParam:"inline,name=input"`
	SubtitleInput  *components.SubtitleInput  `queryParam:"inline,name=input"`

	Type InputType
}

func CreateInputVideoInput(videoInput components.VideoInput) Input {
	typ := InputTypeVideoInput

	return Input{
		VideoInput: &videoInput,
		Type:       typ,
	}
}

func CreateInputWatermarkInput(watermarkInput components.WatermarkInput) Input {
	typ := InputTypeWatermarkInput

	return Input{
		WatermarkInput: &watermarkInput,
		Type:           typ,
	}
}

func CreateInputAudioInput(audioInput components.AudioInput) Input {
	typ := InputTypeAudioInput

	return Input{
		AudioInput: &audioInput,
		Type:       typ,
	}
}

func CreateInputSubtitleInput(subtitleInput components.SubtitleInput) Input {
	typ := InputTypeSubtitleInput

	return Input{
		SubtitleInput: &subtitleInput,
		Type:          typ,
	}
}

func (u *Input) UnmarshalJSON(data []byte) error {

	var subtitleInput components.SubtitleInput = components.SubtitleInput{}
	if err := utils.UnmarshalJSON(data, &subtitleInput, "", true, nil); err == nil {
		u.SubtitleInput = &subtitleInput
		u.Type = InputTypeSubtitleInput
		return nil
	}

	var videoInput components.VideoInput = components.VideoInput{}
	if err := utils.UnmarshalJSON(data, &videoInput, "", true, nil); err == nil {
		u.VideoInput = &videoInput
		u.Type = InputTypeVideoInput
		return nil
	}

	var watermarkInput components.WatermarkInput = components.WatermarkInput{}
	if err := utils.UnmarshalJSON(data, &watermarkInput, "", true, nil); err == nil {
		u.WatermarkInput = &watermarkInput
		u.Type = InputTypeWatermarkInput
		return nil
	}

	var audioInput components.AudioInput = components.AudioInput{}
	if err := utils.UnmarshalJSON(data, &audioInput, "", true, nil); err == nil {
		u.AudioInput = &audioInput
		u.Type = InputTypeAudioInput
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Input", string(data))
}

func (u Input) MarshalJSON() ([]byte, error) {
	if u.VideoInput != nil {
		return utils.MarshalJSON(u.VideoInput, "", true)
	}

	if u.WatermarkInput != nil {
		return utils.MarshalJSON(u.WatermarkInput, "", true)
	}

	if u.AudioInput != nil {
		return utils.MarshalJSON(u.AudioInput, "", true)
	}

	if u.SubtitleInput != nil {
		return utils.MarshalJSON(u.SubtitleInput, "", true)
	}

	return nil, errors.New("could not marshal union type Input: all fields are null")
}

// LanguageCode - Language codes (BCP 47 compliant) used for text files.
type LanguageCode string

const (
	LanguageCodeEn LanguageCode = "en"
	LanguageCodeIt LanguageCode = "it"
	LanguageCodePl LanguageCode = "pl"
	LanguageCodeEs LanguageCode = "es"
	LanguageCodeFr LanguageCode = "fr"
	LanguageCodeRu LanguageCode = "ru"
	LanguageCodeNl LanguageCode = "nl"
)

func (e LanguageCode) ToPointer() *LanguageCode {
	return &e
}
func (e *LanguageCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "en":
		fallthrough
	case "it":
		fallthrough
	case "pl":
		fallthrough
	case "es":
		fallthrough
	case "fr":
		fallthrough
	case "ru":
		fallthrough
	case "nl":
		*e = LanguageCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LanguageCode: %v", v)
	}
}

// Subtitles - Generates subtitle files for audio/video files.
type Subtitles struct {
	// Name of the language for the subtitles.
	LanguageName *string `json:"languageName,omitempty"`
	// Tag a video in "key" : "value" pairs for searchable metadata. Maximum 10 entries, 255 characters each.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Language codes (BCP 47 compliant) used for text files.
	//
	LanguageCode *LanguageCode `json:"languageCode,omitempty"`
}

func (s *Subtitles) GetLanguageName() *string {
	if s == nil {
		return nil
	}
	return s.LanguageName
}

func (s *Subtitles) GetMetadata() map[string]string {
	if s == nil {
		return nil
	}
	return s.Metadata
}

func (s *Subtitles) GetLanguageCode() *LanguageCode {
	if s == nil {
		return nil
	}
	return s.LanguageCode
}

// MaxResolution - Determines the highest quality resolution available.
type MaxResolution string

const (
	MaxResolutionTwoThousandOneHundredAndSixtyp  MaxResolution = "2160p"
	MaxResolutionOneThousandFourHundredAndFortyp MaxResolution = "1440p"
	MaxResolutionOneThousandAndEightyp           MaxResolution = "1080p"
	MaxResolutionSevenHundredAndTwentyp          MaxResolution = "720p"
	MaxResolutionFourHundredAndEightyp           MaxResolution = "480p"
	MaxResolutionThreeHundredAndSixtyp           MaxResolution = "360p"
)

func (e MaxResolution) ToPointer() *MaxResolution {
	return &e
}
func (e *MaxResolution) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "2160p":
		fallthrough
	case "1440p":
		fallthrough
	case "1080p":
		fallthrough
	case "720p":
		fallthrough
	case "480p":
		fallthrough
	case "360p":
		*e = MaxResolution(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MaxResolution: %v", v)
	}
}

// DirectUploadVideoMediaMp4Support - Generates MP4 video up to 4K ("capped_4k"), m4a audio only ("audioOnly"), or both for offline viewing.
type DirectUploadVideoMediaMp4Support string

const (
	DirectUploadVideoMediaMp4SupportCapped4k          DirectUploadVideoMediaMp4Support = "capped_4k"
	DirectUploadVideoMediaMp4SupportAudioOnly         DirectUploadVideoMediaMp4Support = "audioOnly"
	DirectUploadVideoMediaMp4SupportAudioOnlyCapped4k DirectUploadVideoMediaMp4Support = "audioOnly,capped_4k"
)

func (e DirectUploadVideoMediaMp4Support) ToPointer() *DirectUploadVideoMediaMp4Support {
	return &e
}
func (e *DirectUploadVideoMediaMp4Support) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "capped_4k":
		fallthrough
	case "audioOnly":
		fallthrough
	case "audioOnly,capped_4k":
		*e = DirectUploadVideoMediaMp4Support(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DirectUploadVideoMediaMp4Support: %v", v)
	}
}

type Summary struct {
	// Enable or disable the summary feature for the media. Set to true to enable summary or false to disable.
	//
	Generate *bool `json:"generate,omitempty"`
	// Specifies the desired word count for the generated summary.
	// - The value must be between **30** and **250** words.
	//
	SummaryLength *int64 `default:"100" json:"summaryLength"`
}

func (s Summary) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Summary) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *Summary) GetGenerate() *bool {
	if s == nil {
		return nil
	}
	return s.Generate
}

func (s *Summary) GetSummaryLength() *int64 {
	if s == nil {
		return nil
	}
	return s.SummaryLength
}

type DirectUploadVideoMediaModeration struct {
	// Type of media content
	Type *components.MediaType `json:"type,omitempty"`
}

func (d *DirectUploadVideoMediaModeration) GetType() *components.MediaType {
	if d == nil {
		return nil
	}
	return d.Type
}

type DirectUploadVideoMediaAccessRestrictions struct {
	// Restrictions based on the originating domain of a request
	Domains *components.DomainRestrictions `json:"domains,omitempty"`
	// Restrictions based on the user agent
	UserAgents *components.UserAgentRestrictions `json:"userAgents,omitempty"`
}

func (d *DirectUploadVideoMediaAccessRestrictions) GetDomains() *components.DomainRestrictions {
	if d == nil {
		return nil
	}
	return d.Domains
}

func (d *DirectUploadVideoMediaAccessRestrictions) GetUserAgents() *components.UserAgentRestrictions {
	if d == nil {
		return nil
	}
	return d.UserAgents
}

// PushMediaSettings - Configuration settings for media upload.
type PushMediaSettings struct {
	// Basic access policy for media content
	AccessPolicy components.BasicAccessPolicy `json:"accessPolicy"`
	// Start time indicates where encoding should begin within the video file, in seconds.
	StartTime *float64 `json:"startTime,omitempty"`
	// End time indicates where encoding should end within the video file, in seconds.
	EndTime *float64 `json:"endTime,omitempty"`
	Inputs  []Input  `json:"inputs,omitempty"`
	// Tag a video in "key" : "value" pairs for searchable metadata. Maximum 10 entries, 255 characters each.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Generates subtitle files for audio/video files.
	//
	Subtitles *Subtitles `json:"subtitles,omitempty"`
	// Enhance the quality and volume of the audio track. This is available for pre-recorded content only.
	//
	OptimizeAudio *bool `default:"true" json:"optimizeAudio"`
	// Determines the highest quality resolution available.
	//
	MaxResolution *MaxResolution `default:"1080p" json:"maxResolution"`
	// The sourceAccess parameter determines whether the original media file is accessible. Set to true to enable access or false to restrict it
	SourceAccess *bool `json:"sourceAccess,omitempty"`
	// Generates MP4 video up to 4K ("capped_4k"), m4a audio only ("audioOnly"), or both for offline viewing.
	//
	Mp4Support *DirectUploadVideoMediaMp4Support `json:"mp4Support,omitempty"`
	Summary    *Summary                          `json:"summary,omitempty"`
	// Enable or disable the chapters feature for the media. Set to `true` to enable chapters or `false` to disable.
	//
	Chapters *bool `json:"chapters,omitempty"`
	// Enable or disable named entity extraction. Set to `true` to enable or `false` to disable.
	//
	NamedEntities      *bool                                     `json:"namedEntities,omitempty"`
	Moderation         *DirectUploadVideoMediaModeration         `json:"moderation,omitempty"`
	AccessRestrictions *DirectUploadVideoMediaAccessRestrictions `json:"accessRestrictions,omitempty"`
}

func (p PushMediaSettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PushMediaSettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, []string{"accessPolicy"}); err != nil {
		return err
	}
	return nil
}

func (p *PushMediaSettings) GetAccessPolicy() components.BasicAccessPolicy {
	if p == nil {
		return components.BasicAccessPolicy("")
	}
	return p.AccessPolicy
}

func (p *PushMediaSettings) GetStartTime() *float64 {
	if p == nil {
		return nil
	}
	return p.StartTime
}

func (p *PushMediaSettings) GetEndTime() *float64 {
	if p == nil {
		return nil
	}
	return p.EndTime
}

func (p *PushMediaSettings) GetInputs() []Input {
	if p == nil {
		return nil
	}
	return p.Inputs
}

func (p *PushMediaSettings) GetMetadata() map[string]string {
	if p == nil {
		return nil
	}
	return p.Metadata
}

func (p *PushMediaSettings) GetSubtitles() *Subtitles {
	if p == nil {
		return nil
	}
	return p.Subtitles
}

func (p *PushMediaSettings) GetOptimizeAudio() *bool {
	if p == nil {
		return nil
	}
	return p.OptimizeAudio
}

func (p *PushMediaSettings) GetMaxResolution() *MaxResolution {
	if p == nil {
		return nil
	}
	return p.MaxResolution
}

func (p *PushMediaSettings) GetSourceAccess() *bool {
	if p == nil {
		return nil
	}
	return p.SourceAccess
}

func (p *PushMediaSettings) GetMp4Support() *DirectUploadVideoMediaMp4Support {
	if p == nil {
		return nil
	}
	return p.Mp4Support
}

func (p *PushMediaSettings) GetSummary() *Summary {
	if p == nil {
		return nil
	}
	return p.Summary
}

func (p *PushMediaSettings) GetChapters() *bool {
	if p == nil {
		return nil
	}
	return p.Chapters
}

func (p *PushMediaSettings) GetNamedEntities() *bool {
	if p == nil {
		return nil
	}
	return p.NamedEntities
}

func (p *PushMediaSettings) GetModeration() *DirectUploadVideoMediaModeration {
	if p == nil {
		return nil
	}
	return p.Moderation
}

func (p *PushMediaSettings) GetAccessRestrictions() *DirectUploadVideoMediaAccessRestrictions {
	if p == nil {
		return nil
	}
	return p.AccessRestrictions
}

// DirectUploadVideoMediaRequest - Request body for direct upload
type DirectUploadVideoMediaRequest struct {
	// Upload media directly from a device using the URL name or enter '*' to allow all.
	CorsOrigin string `json:"corsOrigin"`
	// Configuration settings for media upload.
	PushMediaSettings *PushMediaSettings `json:"pushMediaSettings,omitempty"`
}

func (d *DirectUploadVideoMediaRequest) GetCorsOrigin() string {
	if d == nil {
		return ""
	}
	return d.CorsOrigin
}

func (d *DirectUploadVideoMediaRequest) GetPushMediaSettings() *PushMediaSettings {
	if d == nil {
		return nil
	}
	return d.PushMediaSettings
}

// DirectUploadVideoMediaResponseBody - Direct upload created successfully
type DirectUploadVideoMediaResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data *components.DirectUpload `json:"data,omitempty"`
}

func (d *DirectUploadVideoMediaResponseBody) GetSuccess() *bool {
	if d == nil {
		return nil
	}
	return d.Success
}

func (d *DirectUploadVideoMediaResponseBody) GetData() *components.DirectUpload {
	if d == nil {
		return nil
	}
	return d.Data
}

type DirectUploadVideoMediaResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Direct upload created successfully
	Object *DirectUploadVideoMediaResponseBody
}

func (d *DirectUploadVideoMediaResponse) GetHTTPMeta() components.HTTPMetadata {
	if d == nil {
		return components.HTTPMetadata{}
	}
	return d.HTTPMeta
}

func (d *DirectUploadVideoMediaResponse) GetObject() *DirectUploadVideoMediaResponseBody {
	if d == nil {
		return nil
	}
	return d.Object
}
