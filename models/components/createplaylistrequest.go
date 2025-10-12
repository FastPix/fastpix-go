

package components

import (
	"encoding/json"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
)

// CreatePlaylistRequestType - For a smart playlist metadata is required.
type CreatePlaylistRequestType string

const (
	CreatePlaylistRequestTypeSmart  CreatePlaylistRequestType = "smart"
	CreatePlaylistRequestTypeManual CreatePlaylistRequestType = "manual"
)

func (e CreatePlaylistRequestType) ToPointer() *CreatePlaylistRequestType {
	return &e
}
func (e *CreatePlaylistRequestType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "manual":
		*e = CreatePlaylistRequestType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatePlaylistRequestType: %v", v)
	}
}

// CreatePlaylistRequestMetadata - Required when playlist type is smart - media created between startDate and endDate of createdDate will be add, similarily updatedDate (Optional)
type CreatePlaylistRequestMetadata struct {
	// Date range with start and end dates.
	CreatedDate *DateRange `json:"createdDate,omitempty"`
	// Date range with start and end dates.
	UpdatedDate *DateRange `json:"updatedDate,omitempty"`
}

func (c *CreatePlaylistRequestMetadata) GetCreatedDate() *DateRange {
	if c == nil {
		return nil
	}
	return c.CreatedDate
}

func (c *CreatePlaylistRequestMetadata) GetUpdatedDate() *DateRange {
	if c == nil {
		return nil
	}
	return c.UpdatedDate
}

type CreatePlaylistRequest struct {
	// Name of the playlist.
	Name string `json:"name"`
	// Unique string value assigned by user to the playlist.
	ReferenceID string `json:"referenceId"`
	// For a smart playlist metadata is required.
	Type CreatePlaylistRequestType `json:"type"`
	// Description for a playlist (Optional).
	Description *string `json:"description,omitempty"`
	// Determines the insertion order of media into playlist.
	PlayOrder *PlaylistOrder `json:"playOrder,omitempty"`
	// Optional parameter to limit no. of media in a playlist.
	Limit *int64 `default:"1000" json:"limit"`
	// Required when playlist type is smart - media created between startDate and endDate of createdDate will be add, similarily updatedDate (Optional)
	Metadata *CreatePlaylistRequestMetadata `json:"metadata,omitempty"`
}

func (c CreatePlaylistRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePlaylistRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, []string{"name", "referenceId", "type"}); err != nil {
		return err
	}
	return nil
}

func (c *CreatePlaylistRequest) GetName() string {
	if c == nil {
		return ""
	}
	return c.Name
}

func (c *CreatePlaylistRequest) GetReferenceID() string {
	if c == nil {
		return ""
	}
	return c.ReferenceID
}

func (c *CreatePlaylistRequest) GetType() CreatePlaylistRequestType {
	if c == nil {
		return CreatePlaylistRequestType("")
	}
	return c.Type
}

func (c *CreatePlaylistRequest) GetDescription() *string {
	if c == nil {
		return nil
	}
	return c.Description
}

func (c *CreatePlaylistRequest) GetPlayOrder() *PlaylistOrder {
	if c == nil {
		return nil
	}
	return c.PlayOrder
}

func (c *CreatePlaylistRequest) GetLimit() *int64 {
	if c == nil {
		return nil
	}
	return c.Limit
}

func (c *CreatePlaylistRequest) GetMetadata() *CreatePlaylistRequestMetadata {
	if c == nil {
		return nil
	}
	return c.Metadata
}
