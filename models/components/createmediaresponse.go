package components

import (
	"encoding/json"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
	"time"
)

// CreateMediaResponseMaxResolution - The maximum resolution tier determines the highest quality your media will be available in.
type CreateMediaResponseMaxResolution string

const (
	CreateMediaResponseMaxResolutionTwoThousandOneHundredAndSixtyp  CreateMediaResponseMaxResolution = "2160p"
	CreateMediaResponseMaxResolutionOneThousandFourHundredAndFortyp CreateMediaResponseMaxResolution = "1440p"
	CreateMediaResponseMaxResolutionOneThousandAndEightyp           CreateMediaResponseMaxResolution = "1080p"
	CreateMediaResponseMaxResolutionSevenHundredAndTwentyp          CreateMediaResponseMaxResolution = "720p"
	CreateMediaResponseMaxResolutionFourHundredAndEightyp           CreateMediaResponseMaxResolution = "480p"
	CreateMediaResponseMaxResolutionThreeHundredAndSixtyp           CreateMediaResponseMaxResolution = "360p"
)

func (e CreateMediaResponseMaxResolution) ToPointer() *CreateMediaResponseMaxResolution {
	return &e
}
func (e *CreateMediaResponseMaxResolution) UnmarshalJSON(data []byte) error {
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
		*e = CreateMediaResponseMaxResolution(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateMediaResponseMaxResolution: %v", v)
	}
}

type CreateMediaResponse struct {
	// The Media is assigned a universal unique identifier, which can contain a maximum of 255 characters.
	ID *string `json:"id,omitempty"`
	// FastPix allows for a free trial. Create as many media files as you like during the trial period. Remember, each clip can only be 10 seconds long and will be deleted after 24 hours. Also, all trial content will have the FastPix logo watermark.
	//
	Trial *bool `default:"true" json:"trial"`
	// Determines the media's status, which can be one of the possible values.
	Status *string `json:"status,omitempty"`
	// Time the media was created, defined as a localDateTime (UTC Time).
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Time the media was updated, defined as a localDateTime (UTC Time).
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// A collection of Playback ID objects utilized for crafting HLS playback URLs.
	PlaybackIds []PlaybackID `json:"playbackIds,omitempty"`
	// You can search for videos with specific key value pairs using metadata, when you tag a video in "key" : "value" pairs. Dynamic Metadata allows you to define a key that allows any value pair. You can have maximum of 255 characters and upto 10 entries are allowed.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The maximum resolution tier determines the highest quality your media will be available in.
	MaxResolution *CreateMediaResponseMaxResolution `default:"1080p" json:"maxResolution"`
}

func (c CreateMediaResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateMediaResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (c *CreateMediaResponse) GetID() *string {
	if c == nil {
		return nil
	}
	return c.ID
}

func (c *CreateMediaResponse) GetTrial() *bool {
	if c == nil {
		return nil
	}
	return c.Trial
}

func (c *CreateMediaResponse) GetStatus() *string {
	if c == nil {
		return nil
	}
	return c.Status
}

func (c *CreateMediaResponse) GetCreatedAt() *time.Time {
	if c == nil {
		return nil
	}
	return c.CreatedAt
}

func (c *CreateMediaResponse) GetUpdatedAt() *time.Time {
	if c == nil {
		return nil
	}
	return c.UpdatedAt
}

func (c *CreateMediaResponse) GetPlaybackIds() []PlaybackID {
	if c == nil {
		return nil
	}
	return c.PlaybackIds
}

func (c *CreateMediaResponse) GetMetadata() map[string]string {
	if c == nil {
		return nil
	}
	return c.Metadata
}

func (c *CreateMediaResponse) GetMaxResolution() *CreateMediaResponseMaxResolution {
	if c == nil {
		return nil
	}
	return c.MaxResolution
}
