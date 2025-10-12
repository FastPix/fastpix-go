package components

import (
	"encoding/json"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
	"time"
)

type PlaylistItemType string

const (
	PlaylistItemTypeManual PlaylistItemType = "manual"
	PlaylistItemTypeSmart  PlaylistItemType = "smart"
)

func (e PlaylistItemType) ToPointer() *PlaylistItemType {
	return &e
}
func (e *PlaylistItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "manual":
		fallthrough
	case "smart":
		*e = PlaylistItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlaylistItemType: %v", v)
	}
}

type PlaylistItem struct {
	ID          *string           `json:"id,omitempty"`
	Name        *string           `json:"name,omitempty"`
	Type        *PlaylistItemType `json:"type,omitempty"`
	ReferenceID *string           `json:"referenceId,omitempty"`
	CreatedAt   *time.Time        `json:"createdAt,omitempty"`
	MediaCount  *int64            `json:"mediaCount,omitempty"`
}

func (p PlaylistItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PlaylistItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PlaylistItem) GetID() *string {
	if p == nil {
		return nil
	}
	return p.ID
}

func (p *PlaylistItem) GetName() *string {
	if p == nil {
		return nil
	}
	return p.Name
}

func (p *PlaylistItem) GetType() *PlaylistItemType {
	if p == nil {
		return nil
	}
	return p.Type
}

func (p *PlaylistItem) GetReferenceID() *string {
	if p == nil {
		return nil
	}
	return p.ReferenceID
}

func (p *PlaylistItem) GetCreatedAt() *time.Time {
	if p == nil {
		return nil
	}
	return p.CreatedAt
}

func (p *PlaylistItem) GetMediaCount() *int64 {
	if p == nil {
		return nil
	}
	return p.MediaCount
}
