package components

import (
	"encoding/json"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
	"time"
)

// PlaylistCreatedSchemaType - Type will be either smart or manual, as sent in the request body.
type PlaylistCreatedSchemaType string

const (
	PlaylistCreatedSchemaTypeSmart  PlaylistCreatedSchemaType = "smart"
	PlaylistCreatedSchemaTypeManual PlaylistCreatedSchemaType = "manual"
)

func (e PlaylistCreatedSchemaType) ToPointer() *PlaylistCreatedSchemaType {
	return &e
}
func (e *PlaylistCreatedSchemaType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "manual":
		*e = PlaylistCreatedSchemaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlaylistCreatedSchemaType: %v", v)
	}
}

// PlaylistCreatedSchemaMetadata - date range filter used when creating the smart playlist
type PlaylistCreatedSchemaMetadata struct {
	// Date range with start and end dates.
	CreatedDate *DateRange `json:"createdDate,omitempty"`
	// Date range with start and end dates.
	UpdatedDate *DateRange `json:"updatedDate,omitempty"`
}

func (p *PlaylistCreatedSchemaMetadata) GetCreatedDate() *DateRange {
	if p == nil {
		return nil
	}
	return p.CreatedDate
}

func (p *PlaylistCreatedSchemaMetadata) GetUpdatedDate() *DateRange {
	if p == nil {
		return nil
	}
	return p.UpdatedDate
}

type PlaylistCreatedSchemaMediaList struct {
	// timestamp of media creation in the workspace
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// duration of the media in hh:mm:ss format
	Duration *string `json:"duration,omitempty"`
	// unique identifier of the media
	ID *string `json:"id,omitempty"`
	// The source resolution of the media
	SourceResolution *string `json:"sourceResolution,omitempty"`
	// The status of the video in the workspace. Only media which are in ready status are added into the playlist
	Status *string `json:"status,omitempty"`
	// Thumbnail to the particular media
	Thumbnail *string `json:"thumbnail,omitempty"`
}

func (p PlaylistCreatedSchemaMediaList) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PlaylistCreatedSchemaMediaList) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PlaylistCreatedSchemaMediaList) GetCreatedAt() *time.Time {
	if p == nil {
		return nil
	}
	return p.CreatedAt
}

func (p *PlaylistCreatedSchemaMediaList) GetDuration() *string {
	if p == nil {
		return nil
	}
	return p.Duration
}

func (p *PlaylistCreatedSchemaMediaList) GetID() *string {
	if p == nil {
		return nil
	}
	return p.ID
}

func (p *PlaylistCreatedSchemaMediaList) GetSourceResolution() *string {
	if p == nil {
		return nil
	}
	return p.SourceResolution
}

func (p *PlaylistCreatedSchemaMediaList) GetStatus() *string {
	if p == nil {
		return nil
	}
	return p.Status
}

func (p *PlaylistCreatedSchemaMediaList) GetThumbnail() *string {
	if p == nil {
		return nil
	}
	return p.Thumbnail
}

// PlaylistCreatedSchema - Displays the result of the request.
type PlaylistCreatedSchema struct {
	// Upon creating a new play,ist, FastPix assigns a unique identifier to the playlist.
	ID *string `json:"id,omitempty"`
	// The name to the playlist set by the user.
	Name *string `json:"name,omitempty"`
	// Unique string value assigned by user to the playlist.
	ReferenceID *string `json:"referenceId,omitempty"`
	// Type will be either smart or manual, as sent in the request body.
	Type *PlaylistCreatedSchemaType `json:"type,omitempty"`
	// The description to the playlist set by the user.
	Description *string `json:"description,omitempty"`
	// Determines the insertion order of media into playlist.
	PlayOrder *PlaylistOrder `json:"playOrder,omitempty"`
	// date range filter used when creating the smart playlist
	Metadata  *PlaylistCreatedSchemaMetadata   `json:"metadata,omitempty"`
	MediaList []PlaylistCreatedSchemaMediaList `json:"mediaList,omitempty"`
	// Id of the workspace
	WorkspaceID *string `json:"workspaceId,omitempty"`
	// Timestamp of playlist creation.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Playlist's most recent update timestamp.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// No. of media present in the playlist
	MediaCount *int64 `json:"mediaCount,omitempty"`
}

func (p PlaylistCreatedSchema) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PlaylistCreatedSchema) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PlaylistCreatedSchema) GetID() *string {
	if p == nil {
		return nil
	}
	return p.ID
}

func (p *PlaylistCreatedSchema) GetName() *string {
	if p == nil {
		return nil
	}
	return p.Name
}

func (p *PlaylistCreatedSchema) GetReferenceID() *string {
	if p == nil {
		return nil
	}
	return p.ReferenceID
}

func (p *PlaylistCreatedSchema) GetType() *PlaylistCreatedSchemaType {
	if p == nil {
		return nil
	}
	return p.Type
}

func (p *PlaylistCreatedSchema) GetDescription() *string {
	if p == nil {
		return nil
	}
	return p.Description
}

func (p *PlaylistCreatedSchema) GetPlayOrder() *PlaylistOrder {
	if p == nil {
		return nil
	}
	return p.PlayOrder
}

func (p *PlaylistCreatedSchema) GetMetadata() *PlaylistCreatedSchemaMetadata {
	if p == nil {
		return nil
	}
	return p.Metadata
}

func (p *PlaylistCreatedSchema) GetMediaList() []PlaylistCreatedSchemaMediaList {
	if p == nil {
		return nil
	}
	return p.MediaList
}

func (p *PlaylistCreatedSchema) GetWorkspaceID() *string {
	if p == nil {
		return nil
	}
	return p.WorkspaceID
}

func (p *PlaylistCreatedSchema) GetCreatedAt() *time.Time {
	if p == nil {
		return nil
	}
	return p.CreatedAt
}

func (p *PlaylistCreatedSchema) GetUpdatedAt() *time.Time {
	if p == nil {
		return nil
	}
	return p.UpdatedAt
}

func (p *PlaylistCreatedSchema) GetMediaCount() *int64 {
	if p == nil {
		return nil
	}
	return p.MediaCount
}
