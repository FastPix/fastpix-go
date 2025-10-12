package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
)

type InputType string

const (
	InputTypeVideoInput     InputType = "VideoInput"
	InputTypeWatermarkInput InputType = "WatermarkInput"
	InputTypeAudioInput     InputType = "AudioInput"
	InputTypeSubtitleInput  InputType = "SubtitleInput"
)

type Input struct {
	VideoInput     *VideoInput     `queryParam:"inline,name=input"`
	WatermarkInput *WatermarkInput `queryParam:"inline,name=input"`
	AudioInput     *AudioInput     `queryParam:"inline,name=input"`
	SubtitleInput  *SubtitleInput  `queryParam:"inline,name=input"`

	Type InputType
}

func CreateInputVideoInput(videoInput VideoInput) Input {
	typ := InputTypeVideoInput

	return Input{
		VideoInput: &videoInput,
		Type:       typ,
	}
}

func CreateInputWatermarkInput(watermarkInput WatermarkInput) Input {
	typ := InputTypeWatermarkInput

	return Input{
		WatermarkInput: &watermarkInput,
		Type:           typ,
	}
}

func CreateInputAudioInput(audioInput AudioInput) Input {
	typ := InputTypeAudioInput

	return Input{
		AudioInput: &audioInput,
		Type:       typ,
	}
}

func CreateInputSubtitleInput(subtitleInput SubtitleInput) Input {
	typ := InputTypeSubtitleInput

	return Input{
		SubtitleInput: &subtitleInput,
		Type:          typ,
	}
}

func (u *Input) UnmarshalJSON(data []byte) error {

	var subtitleInput SubtitleInput = SubtitleInput{}
	if err := utils.UnmarshalJSON(data, &subtitleInput, "", true, nil); err == nil {
		u.SubtitleInput = &subtitleInput
		u.Type = InputTypeSubtitleInput
		return nil
	}

	var videoInput VideoInput = VideoInput{}
	if err := utils.UnmarshalJSON(data, &videoInput, "", true, nil); err == nil {
		u.VideoInput = &videoInput
		u.Type = InputTypeVideoInput
		return nil
	}

	var watermarkInput WatermarkInput = WatermarkInput{}
	if err := utils.UnmarshalJSON(data, &watermarkInput, "", true, nil); err == nil {
		u.WatermarkInput = &watermarkInput
		u.Type = InputTypeWatermarkInput
		return nil
	}

	var audioInput AudioInput = AudioInput{}
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

// CreateMediaRequestLanguageCode - Language codes are concise, standardized symbols that denote languages, utilizing either two or three characters for identification. The language code must be compliant with the BCP 47 standard to ensure compatibility. (for text only).
type CreateMediaRequestLanguageCode string

const (
	CreateMediaRequestLanguageCodeEn CreateMediaRequestLanguageCode = "en"
	CreateMediaRequestLanguageCodeIt CreateMediaRequestLanguageCode = "it"
	CreateMediaRequestLanguageCodePl CreateMediaRequestLanguageCode = "pl"
	CreateMediaRequestLanguageCodeEs CreateMediaRequestLanguageCode = "es"
	CreateMediaRequestLanguageCodeFr CreateMediaRequestLanguageCode = "fr"
	CreateMediaRequestLanguageCodeRu CreateMediaRequestLanguageCode = "ru"
	CreateMediaRequestLanguageCodeNl CreateMediaRequestLanguageCode = "nl"
)

func (e CreateMediaRequestLanguageCode) ToPointer() *CreateMediaRequestLanguageCode {
	return &e
}
func (e *CreateMediaRequestLanguageCode) UnmarshalJSON(data []byte) error {
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
		*e = CreateMediaRequestLanguageCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateMediaRequestLanguageCode: %v", v)
	}
}

// Subtitles - Generates subtitle files for audio/video files.
type Subtitles struct {
	// Name of the language in which the subtitles will be generated.
	//
	LanguageName *string `json:"languageName,omitempty"`
	// You can search for videos with specific key value pairs using metadata, when you tag a video in "key" : "value" pairs. Dynamic Metadata allows you to define a key that allows any value pair. You can have maximum of 255 characters and upto 10 entries are allowed.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Language codes are concise, standardized symbols that denote languages, utilizing either two or three characters for identification. The language code must be compliant with the BCP 47 standard to ensure compatibility. (for text only).
	//
	LanguageCode *CreateMediaRequestLanguageCode `json:"languageCode,omitempty"`
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

func (s *Subtitles) GetLanguageCode() *CreateMediaRequestLanguageCode {
	if s == nil {
		return nil
	}
	return s.LanguageCode
}

// CreateMediaRequestAccessPolicy - Determines whether access to the streamed content is kept private or available to all.
type CreateMediaRequestAccessPolicy string

const (
	CreateMediaRequestAccessPolicyPublic  CreateMediaRequestAccessPolicy = "public"
	CreateMediaRequestAccessPolicyPrivate CreateMediaRequestAccessPolicy = "private"
	CreateMediaRequestAccessPolicyDrm     CreateMediaRequestAccessPolicy = "drm"
)

func (e CreateMediaRequestAccessPolicy) ToPointer() *CreateMediaRequestAccessPolicy {
	return &e
}
func (e *CreateMediaRequestAccessPolicy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "public":
		fallthrough
	case "private":
		fallthrough
	case "drm":
		*e = CreateMediaRequestAccessPolicy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateMediaRequestAccessPolicy: %v", v)
	}
}

// CreateMediaRequestMp4Support - "capped_4k": Generates an mp4 video file up to 4k resolution "audioOnly": Generates an m4a audio file of the media file "audioOnly,capped_4k": Generates both video and audio media files for offline viewing
type CreateMediaRequestMp4Support string

const (
	CreateMediaRequestMp4SupportCapped4k          CreateMediaRequestMp4Support = "capped_4k"
	CreateMediaRequestMp4SupportAudioOnly         CreateMediaRequestMp4Support = "audioOnly"
	CreateMediaRequestMp4SupportAudioOnlyCapped4k CreateMediaRequestMp4Support = "audioOnly,capped_4k"
)

func (e CreateMediaRequestMp4Support) ToPointer() *CreateMediaRequestMp4Support {
	return &e
}
func (e *CreateMediaRequestMp4Support) UnmarshalJSON(data []byte) error {
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
		*e = CreateMediaRequestMp4Support(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateMediaRequestMp4Support: %v", v)
	}
}

// CreateMediaRequestMaxResolution - The maximum resolution tier determines the highest quality your media will be available in.
type CreateMediaRequestMaxResolution string

const (
	CreateMediaRequestMaxResolutionTwoThousandOneHundredAndSixtyp  CreateMediaRequestMaxResolution = "2160p"
	CreateMediaRequestMaxResolutionOneThousandFourHundredAndFortyp CreateMediaRequestMaxResolution = "1440p"
	CreateMediaRequestMaxResolutionOneThousandAndEightyp           CreateMediaRequestMaxResolution = "1080p"
	CreateMediaRequestMaxResolutionSevenHundredAndTwentyp          CreateMediaRequestMaxResolution = "720p"
	CreateMediaRequestMaxResolutionFourHundredAndEightyp           CreateMediaRequestMaxResolution = "480p"
	CreateMediaRequestMaxResolutionThreeHundredAndSixtyp           CreateMediaRequestMaxResolution = "360p"
)

func (e CreateMediaRequestMaxResolution) ToPointer() *CreateMediaRequestMaxResolution {
	return &e
}
func (e *CreateMediaRequestMaxResolution) UnmarshalJSON(data []byte) error {
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
		*e = CreateMediaRequestMaxResolution(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateMediaRequestMaxResolution: %v", v)
	}
}

type Summary struct {
	// Enable or disable the summary feature for the media. Set to true to enable summary or false to disable.
	//
	Generate bool `json:"generate"`
	// Specifies the desired word count for the generated summary.
	// - The value must be between **30** and **250** words.
	//
	SummaryLength *int64 `default:"100" json:"summaryLength"`
}

func (s Summary) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Summary) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, []string{"generate"}); err != nil {
		return err
	}
	return nil
}

func (s *Summary) GetGenerate() bool {
	if s == nil {
		return false
	}
	return s.Generate
}

func (s *Summary) GetSummaryLength() *int64 {
	if s == nil {
		return nil
	}
	return s.SummaryLength
}

type Moderation struct {
	// Type of media content
	Type MediaType `json:"type"`
}

func (m *Moderation) GetType() MediaType {
	if m == nil {
		return MediaType("")
	}
	return m.Type
}

type CreateMediaRequestDomains struct {
	// Policy action type
	DefaultPolicy *PolicyAction `json:"defaultPolicy,omitempty"`
	// A list of domain names or patterns that are explicitly allowed access.
	// This list is only effective when the `defaultPolicy` is set to `deny`.
	//
	Allow []string `json:"allow,omitempty"`
	// A list of domain names or patterns that are explicitly denied access.
	// This list is only effective when the `defaultPolicy` is set to `allow`.
	//
	Deny []string `json:"deny,omitempty"`
}

func (c *CreateMediaRequestDomains) GetDefaultPolicy() *PolicyAction {
	if c == nil {
		return nil
	}
	return c.DefaultPolicy
}

func (c *CreateMediaRequestDomains) GetAllow() []string {
	if c == nil {
		return nil
	}
	return c.Allow
}

func (c *CreateMediaRequestDomains) GetDeny() []string {
	if c == nil {
		return nil
	}
	return c.Deny
}

type CreateMediaRequestUserAgents struct {
	// Policy action type
	DefaultPolicy *PolicyAction `json:"defaultPolicy,omitempty"`
	// A list of user agents (identified by string names or patterns) that are explicitly allowed access.
	// This list is only effective when the `defaultPolicy` is set to `deny`.
	//
	Allow []string `json:"allow,omitempty"`
	// A list of user agents (identified by string names or patterns) that are explicitly denied access.
	// This list is only effective when the `defaultPolicy` is set to `allow`.
	//
	Deny []string `json:"deny,omitempty"`
}

func (c *CreateMediaRequestUserAgents) GetDefaultPolicy() *PolicyAction {
	if c == nil {
		return nil
	}
	return c.DefaultPolicy
}

func (c *CreateMediaRequestUserAgents) GetAllow() []string {
	if c == nil {
		return nil
	}
	return c.Allow
}

func (c *CreateMediaRequestUserAgents) GetDeny() []string {
	if c == nil {
		return nil
	}
	return c.Deny
}

type CreateMediaRequestAccessRestrictions struct {
	Domains    *CreateMediaRequestDomains    `json:"domains,omitempty"`
	UserAgents *CreateMediaRequestUserAgents `json:"userAgents,omitempty"`
}

func (c *CreateMediaRequestAccessRestrictions) GetDomains() *CreateMediaRequestDomains {
	if c == nil {
		return nil
	}
	return c.Domains
}

func (c *CreateMediaRequestAccessRestrictions) GetUserAgents() *CreateMediaRequestUserAgents {
	if c == nil {
		return nil
	}
	return c.UserAgents
}

type CreateMediaRequest struct {
	Inputs []Input `json:"inputs"`
	// You can search for videos with specific key value pairs using metadata, when you tag a video in "key" : "value" pairs. Dynamic Metadata allows you to define a key that allows any value pair. You can have maximum of 255 characters and upto 10 entries are allowed.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Generates subtitle files for audio/video files.
	//
	Subtitles *Subtitles `json:"subtitles,omitempty"`
	// Determines whether access to the streamed content is kept private or available to all.
	//
	AccessPolicy CreateMediaRequestAccessPolicy `json:"accessPolicy"`
	// "capped_4k": Generates an mp4 video file up to 4k resolution "audioOnly": Generates an m4a audio file of the media file "audioOnly,capped_4k": Generates both video and audio media files for offline viewing
	//
	Mp4Support *CreateMediaRequestMp4Support `json:"mp4Support,omitempty"`
	// The sourceAccess parameter determines whether the original media file is accessible. Set to true to enable access or false to restrict it
	SourceAccess *bool `json:"sourceAccess,omitempty"`
	// normalize volume of the audio track. This is available for pre-recorded content only.
	//
	OptimizeAudio *bool `default:"false" json:"optimizeAudio"`
	// The maximum resolution tier determines the highest quality your media will be available in.
	//
	MaxResolution *CreateMediaRequestMaxResolution `default:"1080p" json:"maxResolution"`
	Summary       *Summary                         `json:"summary,omitempty"`
	// Enable or disable the chapters feature for the media. Set to `true` to enable chapters or `false` to disable.
	//
	Chapters *bool `json:"chapters,omitempty"`
	// Enable or disable named entity extraction. Set to `true` to enable or `false` to disable.
	//
	NamedEntities      *bool                                 `json:"namedEntities,omitempty"`
	Moderation         *Moderation                           `json:"moderation,omitempty"`
	AccessRestrictions *CreateMediaRequestAccessRestrictions `json:"accessRestrictions,omitempty"`
}

func (c CreateMediaRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateMediaRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, []string{"inputs", "accessPolicy"}); err != nil {
		return err
	}
	return nil
}

func (c *CreateMediaRequest) GetInputs() []Input {
	if c == nil {
		return []Input{}
	}
	return c.Inputs
}

func (c *CreateMediaRequest) GetMetadata() map[string]string {
	if c == nil {
		return nil
	}
	return c.Metadata
}

func (c *CreateMediaRequest) GetSubtitles() *Subtitles {
	if c == nil {
		return nil
	}
	return c.Subtitles
}

func (c *CreateMediaRequest) GetAccessPolicy() CreateMediaRequestAccessPolicy {
	if c == nil {
		return CreateMediaRequestAccessPolicy("")
	}
	return c.AccessPolicy
}

func (c *CreateMediaRequest) GetMp4Support() *CreateMediaRequestMp4Support {
	if c == nil {
		return nil
	}
	return c.Mp4Support
}

func (c *CreateMediaRequest) GetSourceAccess() *bool {
	if c == nil {
		return nil
	}
	return c.SourceAccess
}

func (c *CreateMediaRequest) GetOptimizeAudio() *bool {
	if c == nil {
		return nil
	}
	return c.OptimizeAudio
}

func (c *CreateMediaRequest) GetMaxResolution() *CreateMediaRequestMaxResolution {
	if c == nil {
		return nil
	}
	return c.MaxResolution
}

func (c *CreateMediaRequest) GetSummary() *Summary {
	if c == nil {
		return nil
	}
	return c.Summary
}

func (c *CreateMediaRequest) GetChapters() *bool {
	if c == nil {
		return nil
	}
	return c.Chapters
}

func (c *CreateMediaRequest) GetNamedEntities() *bool {
	if c == nil {
		return nil
	}
	return c.NamedEntities
}

func (c *CreateMediaRequest) GetModeration() *Moderation {
	if c == nil {
		return nil
	}
	return c.Moderation
}

func (c *CreateMediaRequest) GetAccessRestrictions() *CreateMediaRequestAccessRestrictions {
	if c == nil {
		return nil
	}
	return c.AccessRestrictions
}
