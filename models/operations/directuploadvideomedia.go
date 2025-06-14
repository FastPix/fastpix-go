// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
	"github.com/FastPix/fastpix-go/models/components"
)

// DirectUploadVideoMediaAccessPolicy - Determines if access to the streamed content is kept private or available to all.
type DirectUploadVideoMediaAccessPolicy string

const (
	DirectUploadVideoMediaAccessPolicyPublic  DirectUploadVideoMediaAccessPolicy = "public"
	DirectUploadVideoMediaAccessPolicyPrivate DirectUploadVideoMediaAccessPolicy = "private"
)

func (e DirectUploadVideoMediaAccessPolicy) ToPointer() *DirectUploadVideoMediaAccessPolicy {
	return &e
}
func (e *DirectUploadVideoMediaAccessPolicy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "public":
		fallthrough
	case "private":
		*e = DirectUploadVideoMediaAccessPolicy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DirectUploadVideoMediaAccessPolicy: %v", v)
	}
}

type InputType string

const (
	InputTypeVideoInput     InputType = "VideoInput"
	InputTypeWatermarkInput InputType = "WatermarkInput"
	InputTypeAudioInput     InputType = "AudioInput"
	InputTypeSubtitleInput  InputType = "SubtitleInput"
)

type Input struct {
	VideoInput     *components.VideoInput     `queryParam:"inline"`
	WatermarkInput *components.WatermarkInput `queryParam:"inline"`
	AudioInput     *components.AudioInput     `queryParam:"inline"`
	SubtitleInput  *components.SubtitleInput  `queryParam:"inline"`

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

	var audioInput components.AudioInput = components.AudioInput{}
	if err := utils.UnmarshalJSON(data, &audioInput, "", true, true); err == nil {
		u.AudioInput = &audioInput
		u.Type = InputTypeAudioInput
		return nil
	}

	var subtitleInput components.SubtitleInput = components.SubtitleInput{}
	if err := utils.UnmarshalJSON(data, &subtitleInput, "", true, true); err == nil {
		u.SubtitleInput = &subtitleInput
		u.Type = InputTypeSubtitleInput
		return nil
	}

	var watermarkInput components.WatermarkInput = components.WatermarkInput{}
	if err := utils.UnmarshalJSON(data, &watermarkInput, "", true, true); err == nil {
		u.WatermarkInput = &watermarkInput
		u.Type = InputTypeWatermarkInput
		return nil
	}

	var videoInput components.VideoInput = components.VideoInput{}
	if err := utils.UnmarshalJSON(data, &videoInput, "", true, true); err == nil {
		u.VideoInput = &videoInput
		u.Type = InputTypeVideoInput
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

// DirectUploadVideoMediaMetadata - Tag a video in "key" : "value" pairs for searchable metadata. Maximum 10 entries, 255 characters each.
type DirectUploadVideoMediaMetadata struct {
}

// SubtitlesMetadata - Searchable metadata tags for the video in key-value pairs.
type SubtitlesMetadata struct {
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
	// Searchable metadata tags for the video in key-value pairs.
	Metadata *SubtitlesMetadata `json:"metadata,omitempty"`
	// Language codes (BCP 47 compliant) used for text files.
	//
	LanguageCode *LanguageCode `json:"languageCode,omitempty"`
}

func (o *Subtitles) GetLanguageName() *string {
	if o == nil {
		return nil
	}
	return o.LanguageName
}

func (o *Subtitles) GetMetadata() *SubtitlesMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *Subtitles) GetLanguageCode() *LanguageCode {
	if o == nil {
		return nil
	}
	return o.LanguageCode
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

// Mp4Support - Generates MP4 video up to 4K ("capped_4k"), m4a audio only ("audioOnly"), or both for offline viewing.
type Mp4Support string

const (
	Mp4SupportCapped4k          Mp4Support = "capped_4k"
	Mp4SupportAudioOnly         Mp4Support = "audioOnly"
	Mp4SupportAudioOnlyCapped4k Mp4Support = "audioOnly,capped_4k"
)

func (e Mp4Support) ToPointer() *Mp4Support {
	return &e
}
func (e *Mp4Support) UnmarshalJSON(data []byte) error {
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
		*e = Mp4Support(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Mp4Support: %v", v)
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
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Summary) GetGenerate() *bool {
	if o == nil {
		return nil
	}
	return o.Generate
}

func (o *Summary) GetSummaryLength() *int64 {
	if o == nil {
		return nil
	}
	return o.SummaryLength
}

// Type - Defines the type of input. Possible values include video, audio, av.
type Type string

const (
	TypeVideo Type = "video"
	TypeAudio Type = "audio"
	TypeAv    Type = "av"
)

func (e Type) ToPointer() *Type {
	return &e
}
func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "video":
		fallthrough
	case "audio":
		fallthrough
	case "av":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type Moderation struct {
	// Defines the type of input. Possible values include video, audio, av.
	//
	Type *Type `json:"type,omitempty"`
}

func (o *Moderation) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}

// DirectUploadVideoMediaDomainsDefaultPolicy - Specifies the default access policy for domains.
// If set to `allow`, all domains are allowed access unless otherwise specified in the `deny` list.
// If set to `deny`, all domains are denied access unless otherwise specified in the `allow` list.
type DirectUploadVideoMediaDomainsDefaultPolicy string

const (
	DirectUploadVideoMediaDomainsDefaultPolicyAllow DirectUploadVideoMediaDomainsDefaultPolicy = "allow"
	DirectUploadVideoMediaDomainsDefaultPolicyDeny  DirectUploadVideoMediaDomainsDefaultPolicy = "deny"
)

func (e DirectUploadVideoMediaDomainsDefaultPolicy) ToPointer() *DirectUploadVideoMediaDomainsDefaultPolicy {
	return &e
}
func (e *DirectUploadVideoMediaDomainsDefaultPolicy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "allow":
		fallthrough
	case "deny":
		*e = DirectUploadVideoMediaDomainsDefaultPolicy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DirectUploadVideoMediaDomainsDefaultPolicy: %v", v)
	}
}

type DirectUploadVideoMediaDomains struct {
	// Specifies the default access policy for domains.
	// If set to `allow`, all domains are allowed access unless otherwise specified in the `deny` list.
	// If set to `deny`, all domains are denied access unless otherwise specified in the `allow` list.
	//
	DefaultPolicy *DirectUploadVideoMediaDomainsDefaultPolicy `json:"defaultPolicy,omitempty"`
	// A list of domain names or patterns that are explicitly allowed access.
	// This list is only effective when the `defaultPolicy` is set to `deny`.
	//
	Allow []string `json:"allow,omitempty"`
	// A list of domain names or patterns that are explicitly denied access.
	// This list is only effective when the `defaultPolicy` is set to `allow`.
	//
	Deny []string `json:"deny,omitempty"`
}

func (o *DirectUploadVideoMediaDomains) GetDefaultPolicy() *DirectUploadVideoMediaDomainsDefaultPolicy {
	if o == nil {
		return nil
	}
	return o.DefaultPolicy
}

func (o *DirectUploadVideoMediaDomains) GetAllow() []string {
	if o == nil {
		return nil
	}
	return o.Allow
}

func (o *DirectUploadVideoMediaDomains) GetDeny() []string {
	if o == nil {
		return nil
	}
	return o.Deny
}

// DirectUploadVideoMediaUserAgentsDefaultPolicy - Specifies the default access policy for user agents (browsers, bots, etc.).
// If set to `allow`, all user agents are allowed access unless otherwise specified in the `deny` list.
// If set to `deny`, all user agents are denied access unless otherwise specified in the `allow` list.
type DirectUploadVideoMediaUserAgentsDefaultPolicy string

const (
	DirectUploadVideoMediaUserAgentsDefaultPolicyAllow DirectUploadVideoMediaUserAgentsDefaultPolicy = "allow"
	DirectUploadVideoMediaUserAgentsDefaultPolicyDeny  DirectUploadVideoMediaUserAgentsDefaultPolicy = "deny"
)

func (e DirectUploadVideoMediaUserAgentsDefaultPolicy) ToPointer() *DirectUploadVideoMediaUserAgentsDefaultPolicy {
	return &e
}
func (e *DirectUploadVideoMediaUserAgentsDefaultPolicy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "allow":
		fallthrough
	case "deny":
		*e = DirectUploadVideoMediaUserAgentsDefaultPolicy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DirectUploadVideoMediaUserAgentsDefaultPolicy: %v", v)
	}
}

type DirectUploadVideoMediaUserAgents struct {
	// Specifies the default access policy for user agents (browsers, bots, etc.).
	// If set to `allow`, all user agents are allowed access unless otherwise specified in the `deny` list.
	// If set to `deny`, all user agents are denied access unless otherwise specified in the `allow` list.
	//
	DefaultPolicy *DirectUploadVideoMediaUserAgentsDefaultPolicy `json:"defaultPolicy,omitempty"`
	// A list of user agents (identified by string names or patterns) that are explicitly allowed access.
	// This list is only effective when the `defaultPolicy` is set to `deny`.
	//
	Allow []string `json:"allow,omitempty"`
	// A list of user agents (identified by string names or patterns) that are explicitly denied access.
	// This list is only effective when the `defaultPolicy` is set to `allow`.
	//
	Deny []string `json:"deny,omitempty"`
}

func (o *DirectUploadVideoMediaUserAgents) GetDefaultPolicy() *DirectUploadVideoMediaUserAgentsDefaultPolicy {
	if o == nil {
		return nil
	}
	return o.DefaultPolicy
}

func (o *DirectUploadVideoMediaUserAgents) GetAllow() []string {
	if o == nil {
		return nil
	}
	return o.Allow
}

func (o *DirectUploadVideoMediaUserAgents) GetDeny() []string {
	if o == nil {
		return nil
	}
	return o.Deny
}

type DirectUploadVideoMediaAccessRestrictions struct {
	Domains    *DirectUploadVideoMediaDomains    `json:"domains,omitempty"`
	UserAgents *DirectUploadVideoMediaUserAgents `json:"userAgents,omitempty"`
}

func (o *DirectUploadVideoMediaAccessRestrictions) GetDomains() *DirectUploadVideoMediaDomains {
	if o == nil {
		return nil
	}
	return o.Domains
}

func (o *DirectUploadVideoMediaAccessRestrictions) GetUserAgents() *DirectUploadVideoMediaUserAgents {
	if o == nil {
		return nil
	}
	return o.UserAgents
}

// PushMediaSettings - Configuration settings for media upload.
type PushMediaSettings struct {
	// Determines if access to the streamed content is kept private or available to all.
	AccessPolicy DirectUploadVideoMediaAccessPolicy `json:"accessPolicy"`
	// Start time indicates where encoding should begin within the video file, in seconds.
	StartTime *float64 `json:"startTime,omitempty"`
	// End time indicates where encoding should end within the video file, in seconds.
	EndTime *float64 `json:"endTime,omitempty"`
	Inputs  []Input  `json:"inputs,omitempty"`
	// Tag a video in "key" : "value" pairs for searchable metadata. Maximum 10 entries, 255 characters each.
	Metadata *DirectUploadVideoMediaMetadata `json:"metadata,omitempty"`
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
	Mp4Support *Mp4Support `json:"mp4Support,omitempty"`
	Summary    *Summary    `json:"summary,omitempty"`
	// Enable or disable the chapters feature for the media. Set to `true` to enable chapters or `false` to disable.
	//
	Chapters *bool `json:"chapters,omitempty"`
	// Enable or disable named entity extraction. Set to `true` to enable or `false` to disable.
	//
	NamedEntities      *bool                                     `json:"namedEntities,omitempty"`
	Moderation         *Moderation                               `json:"moderation,omitempty"`
	AccessRestrictions *DirectUploadVideoMediaAccessRestrictions `json:"accessRestrictions,omitempty"`
}

func (p PushMediaSettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PushMediaSettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PushMediaSettings) GetAccessPolicy() DirectUploadVideoMediaAccessPolicy {
	if o == nil {
		return DirectUploadVideoMediaAccessPolicy("")
	}
	return o.AccessPolicy
}

func (o *PushMediaSettings) GetStartTime() *float64 {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *PushMediaSettings) GetEndTime() *float64 {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *PushMediaSettings) GetInputs() []Input {
	if o == nil {
		return nil
	}
	return o.Inputs
}

func (o *PushMediaSettings) GetMetadata() *DirectUploadVideoMediaMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *PushMediaSettings) GetSubtitles() *Subtitles {
	if o == nil {
		return nil
	}
	return o.Subtitles
}

func (o *PushMediaSettings) GetOptimizeAudio() *bool {
	if o == nil {
		return nil
	}
	return o.OptimizeAudio
}

func (o *PushMediaSettings) GetMaxResolution() *MaxResolution {
	if o == nil {
		return nil
	}
	return o.MaxResolution
}

func (o *PushMediaSettings) GetSourceAccess() *bool {
	if o == nil {
		return nil
	}
	return o.SourceAccess
}

func (o *PushMediaSettings) GetMp4Support() *Mp4Support {
	if o == nil {
		return nil
	}
	return o.Mp4Support
}

func (o *PushMediaSettings) GetSummary() *Summary {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *PushMediaSettings) GetChapters() *bool {
	if o == nil {
		return nil
	}
	return o.Chapters
}

func (o *PushMediaSettings) GetNamedEntities() *bool {
	if o == nil {
		return nil
	}
	return o.NamedEntities
}

func (o *PushMediaSettings) GetModeration() *Moderation {
	if o == nil {
		return nil
	}
	return o.Moderation
}

func (o *PushMediaSettings) GetAccessRestrictions() *DirectUploadVideoMediaAccessRestrictions {
	if o == nil {
		return nil
	}
	return o.AccessRestrictions
}

// DirectUploadVideoMediaRequest - Request body for direct upload
type DirectUploadVideoMediaRequest struct {
	// Upload media directly from a device using the URL name or enter '*' to allow all.
	CorsOrigin string `json:"corsOrigin"`
	// Configuration settings for media upload.
	PushMediaSettings *PushMediaSettings `json:"pushMediaSettings,omitempty"`
}

func (o *DirectUploadVideoMediaRequest) GetCorsOrigin() string {
	if o == nil {
		return ""
	}
	return o.CorsOrigin
}

func (o *DirectUploadVideoMediaRequest) GetPushMediaSettings() *PushMediaSettings {
	if o == nil {
		return nil
	}
	return o.PushMediaSettings
}

// DirectUploadVideoMediaResponseBody - Direct upload created successfully
type DirectUploadVideoMediaResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data *components.DirectUpload `json:"data,omitempty"`
}

func (o *DirectUploadVideoMediaResponseBody) GetSuccess() *bool {
	if o == nil {
		return nil
	}
	return o.Success
}

func (o *DirectUploadVideoMediaResponseBody) GetData() *components.DirectUpload {
	if o == nil {
		return nil
	}
	return o.Data
}

type DirectUploadVideoMediaResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Direct upload created successfully
	Object *DirectUploadVideoMediaResponseBody
}

func (o *DirectUploadVideoMediaResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DirectUploadVideoMediaResponse) GetObject() *DirectUploadVideoMediaResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
