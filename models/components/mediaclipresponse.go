

package components

import (
	"encoding/json"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
	"time"
)

// MediaClipResponseMaxResolution - The maximum resolution specified for the media.
type MediaClipResponseMaxResolution string

const (
	MediaClipResponseMaxResolutionTwoThousandOneHundredAndSixtyp  MediaClipResponseMaxResolution = "2160p"
	MediaClipResponseMaxResolutionOneThousandFourHundredAndFortyp MediaClipResponseMaxResolution = "1440p"
	MediaClipResponseMaxResolutionOneThousandAndEightyp           MediaClipResponseMaxResolution = "1080p"
	MediaClipResponseMaxResolutionSevenHundredAndTwentyp          MediaClipResponseMaxResolution = "720p"
	MediaClipResponseMaxResolutionFourHundredAndEightyp           MediaClipResponseMaxResolution = "480p"
	MediaClipResponseMaxResolutionThreeHundredAndSixtyp           MediaClipResponseMaxResolution = "360p"
)

func (e MediaClipResponseMaxResolution) ToPointer() *MediaClipResponseMaxResolution {
	return &e
}
func (e *MediaClipResponseMaxResolution) UnmarshalJSON(data []byte) error {
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
		*e = MediaClipResponseMaxResolution(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MediaClipResponseMaxResolution: %v", v)
	}
}

// MediaClipResponseSourceResolution - The actual resolution of the uploaded media.
type MediaClipResponseSourceResolution string

const (
	MediaClipResponseSourceResolutionTwoThousandOneHundredAndSixtyp  MediaClipResponseSourceResolution = "2160p"
	MediaClipResponseSourceResolutionOneThousandFourHundredAndFortyp MediaClipResponseSourceResolution = "1440p"
	MediaClipResponseSourceResolutionOneThousandAndEightyp           MediaClipResponseSourceResolution = "1080p"
	MediaClipResponseSourceResolutionSevenHundredAndTwentyp          MediaClipResponseSourceResolution = "720p"
	MediaClipResponseSourceResolutionFourHundredAndEightyp           MediaClipResponseSourceResolution = "480p"
	MediaClipResponseSourceResolutionThreeHundredAndSixtyp           MediaClipResponseSourceResolution = "360p"
)

func (e MediaClipResponseSourceResolution) ToPointer() *MediaClipResponseSourceResolution {
	return &e
}
func (e *MediaClipResponseSourceResolution) UnmarshalJSON(data []byte) error {
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
		*e = MediaClipResponseSourceResolution(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MediaClipResponseSourceResolution: %v", v)
	}
}

// Status - The current processing status of the media.
type Status string

const (
	StatusPreparing Status = "preparing"
	StatusReady     Status = "ready"
	StatusFailed    Status = "failed"
	StatusCreated   Status = "created"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "preparing":
		fallthrough
	case "ready":
		fallthrough
	case "failed":
		fallthrough
	case "created":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type MediaClipResponseDomains struct {
	DefaultPolicy *string  `json:"defaultPolicy,omitempty"`
	Allow         []string `json:"allow,omitempty"`
	Deny          []string `json:"deny,omitempty"`
}

func (m *MediaClipResponseDomains) GetDefaultPolicy() *string {
	if m == nil {
		return nil
	}
	return m.DefaultPolicy
}

func (m *MediaClipResponseDomains) GetAllow() []string {
	if m == nil {
		return nil
	}
	return m.Allow
}

func (m *MediaClipResponseDomains) GetDeny() []string {
	if m == nil {
		return nil
	}
	return m.Deny
}

type MediaClipResponseUserAgents struct {
	DefaultPolicy *string  `json:"defaultPolicy,omitempty"`
	Allow         []string `json:"allow,omitempty"`
	Deny          []string `json:"deny,omitempty"`
}

func (m *MediaClipResponseUserAgents) GetDefaultPolicy() *string {
	if m == nil {
		return nil
	}
	return m.DefaultPolicy
}

func (m *MediaClipResponseUserAgents) GetAllow() []string {
	if m == nil {
		return nil
	}
	return m.Allow
}

func (m *MediaClipResponseUserAgents) GetDeny() []string {
	if m == nil {
		return nil
	}
	return m.Deny
}

type MediaClipResponseAccessRestrictions struct {
	Domains    *MediaClipResponseDomains    `json:"domains,omitempty"`
	UserAgents *MediaClipResponseUserAgents `json:"userAgents,omitempty"`
}

func (m *MediaClipResponseAccessRestrictions) GetDomains() *MediaClipResponseDomains {
	if m == nil {
		return nil
	}
	return m.Domains
}

func (m *MediaClipResponseAccessRestrictions) GetUserAgents() *MediaClipResponseUserAgents {
	if m == nil {
		return nil
	}
	return m.UserAgents
}

type MediaClipResponsePlaybackID struct {
	// The unique identifier for playback.
	ID *string `json:"id,omitempty"`
	// The access policy of the playback.
	AccessPolicy       *string                              `json:"accessPolicy,omitempty"`
	AccessRestrictions *MediaClipResponseAccessRestrictions `json:"accessRestrictions,omitempty"`
}

func (m *MediaClipResponsePlaybackID) GetID() *string {
	if m == nil {
		return nil
	}
	return m.ID
}

func (m *MediaClipResponsePlaybackID) GetAccessPolicy() *string {
	if m == nil {
		return nil
	}
	return m.AccessPolicy
}

func (m *MediaClipResponsePlaybackID) GetAccessRestrictions() *MediaClipResponseAccessRestrictions {
	if m == nil {
		return nil
	}
	return m.AccessRestrictions
}

// MediaClipResponseType - The type of media track.
type MediaClipResponseType string

const (
	MediaClipResponseTypeVideo    MediaClipResponseType = "video"
	MediaClipResponseTypeAudio    MediaClipResponseType = "audio"
	MediaClipResponseTypeSubtitle MediaClipResponseType = "subtitle"
)

func (e MediaClipResponseType) ToPointer() *MediaClipResponseType {
	return &e
}
func (e *MediaClipResponseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "video":
		fallthrough
	case "audio":
		fallthrough
	case "subtitle":
		*e = MediaClipResponseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MediaClipResponseType: %v", v)
	}
}

type MediaClipResponseTrack struct {
	// The unique identifier for the media track.
	ID *string `json:"id,omitempty"`
	// The type of media track.
	Type *MediaClipResponseType `json:"type,omitempty"`
	// The width of the video track (applicable to video only).
	Width *int64 `json:"width,omitempty"`
	// The height of the video track (applicable to video only).
	Height *int64 `json:"height,omitempty"`
	// The current processing status of the track.
	Status *string `json:"status,omitempty"`
	// The language code of the audio or subtitle track.
	LanguageCode *string `json:"languageCode,omitempty"`
	// The language name of the audio or subtitle track.
	LanguageName *string `json:"languageName,omitempty"`
}

func (m *MediaClipResponseTrack) GetID() *string {
	if m == nil {
		return nil
	}
	return m.ID
}

func (m *MediaClipResponseTrack) GetType() *MediaClipResponseType {
	if m == nil {
		return nil
	}
	return m.Type
}

func (m *MediaClipResponseTrack) GetWidth() *int64 {
	if m == nil {
		return nil
	}
	return m.Width
}

func (m *MediaClipResponseTrack) GetHeight() *int64 {
	if m == nil {
		return nil
	}
	return m.Height
}

func (m *MediaClipResponseTrack) GetStatus() *string {
	if m == nil {
		return nil
	}
	return m.Status
}

func (m *MediaClipResponseTrack) GetLanguageCode() *string {
	if m == nil {
		return nil
	}
	return m.LanguageCode
}

func (m *MediaClipResponseTrack) GetLanguageName() *string {
	if m == nil {
		return nil
	}
	return m.LanguageName
}

type GeneratedSubtitle struct {
}

type MediaClipResponseData struct {
	// A video thumbnail that acts as a preview image for the video.
	Thumbnail *string `json:"thumbnail,omitempty"`
	// The unique identifier assigned to the media by FastPix.
	ID *string `json:"id,omitempty"`
	// The ID of the original source media.
	SourceMediaID *string `json:"sourceMediaId,omitempty"`
	// The unique identifier for the workspace associated with the media.
	WorkspaceID *string `json:"workspaceId,omitempty"`
	// Tag a video in "key" : "value" pairs for searchable metadata. Maximum 10 entries, 255 characters each.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The maximum resolution specified for the media.
	MaxResolution *MediaClipResponseMaxResolution `json:"maxResolution,omitempty"`
	// The actual resolution of the uploaded media.
	SourceResolution *MediaClipResponseSourceResolution `json:"sourceResolution,omitempty"`
	// The current processing status of the media.
	Status *Status `json:"status,omitempty"`
	// Indicates whether the original media file is accessible.
	SourceAccess *bool                         `json:"sourceAccess,omitempty"`
	PlaybackIds  []MediaClipResponsePlaybackID `json:"playbackIds,omitempty"`
	Tracks       []MediaClipResponseTrack      `json:"tracks,omitempty"`
	// Generated subtitle tracks associated with the media.
	GeneratedSubtitles []GeneratedSubtitle `json:"generatedSubtitles,omitempty"`
	// Indicates whether the media contains only audio.
	IsAudioOnly *bool `json:"isAudioOnly,omitempty"`
	// Indicates whether subtitles are available for the media.
	SubtitleAvailable *bool `json:"subtitleAvailable,omitempty"`
	// The total duration of the media.
	Duration *string `json:"duration,omitempty"`
	// The aspect ratio of the media.
	AspectRatio *string `json:"aspectRatio,omitempty"`
	// Timestamp of when the media was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Timestamp of when the media was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

func (m MediaClipResponseData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MediaClipResponseData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (m *MediaClipResponseData) GetThumbnail() *string {
	if m == nil {
		return nil
	}
	return m.Thumbnail
}

func (m *MediaClipResponseData) GetID() *string {
	if m == nil {
		return nil
	}
	return m.ID
}

func (m *MediaClipResponseData) GetSourceMediaID() *string {
	if m == nil {
		return nil
	}
	return m.SourceMediaID
}

func (m *MediaClipResponseData) GetWorkspaceID() *string {
	if m == nil {
		return nil
	}
	return m.WorkspaceID
}

func (m *MediaClipResponseData) GetMetadata() map[string]string {
	if m == nil {
		return nil
	}
	return m.Metadata
}

func (m *MediaClipResponseData) GetMaxResolution() *MediaClipResponseMaxResolution {
	if m == nil {
		return nil
	}
	return m.MaxResolution
}

func (m *MediaClipResponseData) GetSourceResolution() *MediaClipResponseSourceResolution {
	if m == nil {
		return nil
	}
	return m.SourceResolution
}

func (m *MediaClipResponseData) GetStatus() *Status {
	if m == nil {
		return nil
	}
	return m.Status
}

func (m *MediaClipResponseData) GetSourceAccess() *bool {
	if m == nil {
		return nil
	}
	return m.SourceAccess
}

func (m *MediaClipResponseData) GetPlaybackIds() []MediaClipResponsePlaybackID {
	if m == nil {
		return nil
	}
	return m.PlaybackIds
}

func (m *MediaClipResponseData) GetTracks() []MediaClipResponseTrack {
	if m == nil {
		return nil
	}
	return m.Tracks
}

func (m *MediaClipResponseData) GetGeneratedSubtitles() []GeneratedSubtitle {
	if m == nil {
		return nil
	}
	return m.GeneratedSubtitles
}

func (m *MediaClipResponseData) GetIsAudioOnly() *bool {
	if m == nil {
		return nil
	}
	return m.IsAudioOnly
}

func (m *MediaClipResponseData) GetSubtitleAvailable() *bool {
	if m == nil {
		return nil
	}
	return m.SubtitleAvailable
}

func (m *MediaClipResponseData) GetDuration() *string {
	if m == nil {
		return nil
	}
	return m.Duration
}

func (m *MediaClipResponseData) GetAspectRatio() *string {
	if m == nil {
		return nil
	}
	return m.AspectRatio
}

func (m *MediaClipResponseData) GetCreatedAt() *time.Time {
	if m == nil {
		return nil
	}
	return m.CreatedAt
}

func (m *MediaClipResponseData) GetUpdatedAt() *time.Time {
	if m == nil {
		return nil
	}
	return m.UpdatedAt
}

type MediaClipResponsePagination struct {
	// Total number of records available.
	TotalRecords *int64 `json:"totalRecords,omitempty"`
	// The starting offset of the current result set.
	CurrentOffset *int64 `json:"currentOffset,omitempty"`
	// The number of items returned in the current response.
	OffsetCount *int64 `json:"offsetCount,omitempty"`
}

func (m *MediaClipResponsePagination) GetTotalRecords() *int64 {
	if m == nil {
		return nil
	}
	return m.TotalRecords
}

func (m *MediaClipResponsePagination) GetCurrentOffset() *int64 {
	if m == nil {
		return nil
	}
	return m.CurrentOffset
}

func (m *MediaClipResponsePagination) GetOffsetCount() *int64 {
	if m == nil {
		return nil
	}
	return m.OffsetCount
}

type MediaClipResponse struct {
	Success    bool                        `json:"success"`
	Data       []MediaClipResponseData     `json:"data"`
	Pagination MediaClipResponsePagination `json:"pagination"`
}

func (m *MediaClipResponse) GetSuccess() bool {
	if m == nil {
		return false
	}
	return m.Success
}

func (m *MediaClipResponse) GetData() []MediaClipResponseData {
	if m == nil {
		return []MediaClipResponseData{}
	}
	return m.Data
}

func (m *MediaClipResponse) GetPagination() MediaClipResponsePagination {
	if m == nil {
		return MediaClipResponsePagination{}
	}
	return m.Pagination
}
