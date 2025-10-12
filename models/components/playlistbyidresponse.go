

package components

import (
	"encoding/json"
	"fmt"
	"github.com/fastpix/fastpix-go/internal/utils"
	"time"
)

// PlaylistByIDResponseType - type of the playlist, when it was created
type PlaylistByIDResponseType string

const (
	PlaylistByIDResponseTypeManual PlaylistByIDResponseType = "manual"
	PlaylistByIDResponseTypeSmart  PlaylistByIDResponseType = "smart"
)

func (e PlaylistByIDResponseType) ToPointer() *PlaylistByIDResponseType {
	return &e
}
func (e *PlaylistByIDResponseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "manual":
		fallthrough
	case "smart":
		*e = PlaylistByIDResponseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlaylistByIDResponseType: %v", v)
	}
}

type PlaylistByIDResponseMediaList struct {
	// Timestamp of media creation in the workspace.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Duration of the media in hh:mm:ss format.
	Duration *string `json:"duration,omitempty"`
	// unique id of the particular media.
	ID *string `json:"id,omitempty"`
	// source resolution of the media.
	SourceResolution *string `json:"sourceResolution,omitempty"`
	// status of the media, only media with ready status will be added to playlist.
	Status *string `json:"status,omitempty"`
	// thumbnail for the particular media.
	Thumbnail *string `json:"thumbnail,omitempty"`
}

func (p PlaylistByIDResponseMediaList) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PlaylistByIDResponseMediaList) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PlaylistByIDResponseMediaList) GetCreatedAt() *time.Time {
	if p == nil {
		return nil
	}
	return p.CreatedAt
}

func (p *PlaylistByIDResponseMediaList) GetDuration() *string {
	if p == nil {
		return nil
	}
	return p.Duration
}

func (p *PlaylistByIDResponseMediaList) GetID() *string {
	if p == nil {
		return nil
	}
	return p.ID
}

func (p *PlaylistByIDResponseMediaList) GetSourceResolution() *string {
	if p == nil {
		return nil
	}
	return p.SourceResolution
}

func (p *PlaylistByIDResponseMediaList) GetStatus() *string {
	if p == nil {
		return nil
	}
	return p.Status
}

func (p *PlaylistByIDResponseMediaList) GetThumbnail() *string {
	if p == nil {
		return nil
	}
	return p.Thumbnail
}

type PlaylistByIDResponseData struct {
	// The unique id of the playlist
	ID *string `json:"id,omitempty"`
	// The name of the playlist set by the user
	Name *string `json:"name,omitempty"`
	// Unique string value assigned by user to the playlist.
	ReferenceID *string `json:"referenceId,omitempty"`
	// type of the playlist, when it was created
	Type *PlaylistByIDResponseType `json:"type,omitempty"`
	// Description of the playlist set by the user.
	Description *string                         `json:"description,omitempty"`
	MediaList   []PlaylistByIDResponseMediaList `json:"mediaList,omitempty"`
	// The unique id of the workspace in which the playlist is present.
	WorkspaceID *string `json:"workspaceId,omitempty"`
	// Timestamp of playlist creation.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Playlist's most recent update timestamp.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// No. of media present in the playlist
	MediaCount *int64 `json:"mediaCount,omitempty"`
}

func (p PlaylistByIDResponseData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PlaylistByIDResponseData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PlaylistByIDResponseData) GetID() *string {
	if p == nil {
		return nil
	}
	return p.ID
}

func (p *PlaylistByIDResponseData) GetName() *string {
	if p == nil {
		return nil
	}
	return p.Name
}

func (p *PlaylistByIDResponseData) GetReferenceID() *string {
	if p == nil {
		return nil
	}
	return p.ReferenceID
}

func (p *PlaylistByIDResponseData) GetType() *PlaylistByIDResponseType {
	if p == nil {
		return nil
	}
	return p.Type
}

func (p *PlaylistByIDResponseData) GetDescription() *string {
	if p == nil {
		return nil
	}
	return p.Description
}

func (p *PlaylistByIDResponseData) GetMediaList() []PlaylistByIDResponseMediaList {
	if p == nil {
		return nil
	}
	return p.MediaList
}

func (p *PlaylistByIDResponseData) GetWorkspaceID() *string {
	if p == nil {
		return nil
	}
	return p.WorkspaceID
}

func (p *PlaylistByIDResponseData) GetCreatedAt() *time.Time {
	if p == nil {
		return nil
	}
	return p.CreatedAt
}

func (p *PlaylistByIDResponseData) GetUpdatedAt() *time.Time {
	if p == nil {
		return nil
	}
	return p.UpdatedAt
}

func (p *PlaylistByIDResponseData) GetMediaCount() *int64 {
	if p == nil {
		return nil
	}
	return p.MediaCount
}

type PlaylistByIDResponse struct {
	Success *bool                     `json:"success,omitempty"`
	Data    *PlaylistByIDResponseData `json:"data,omitempty"`
}

func (p *PlaylistByIDResponse) GetSuccess() *bool {
	if p == nil {
		return nil
	}
	return p.Success
}

func (p *PlaylistByIDResponse) GetData() *PlaylistByIDResponseData {
	if p == nil {
		return nil
	}
	return p.Data
}
