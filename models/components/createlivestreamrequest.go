package components

import (
	"encoding/json"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
)

// CreateLiveStreamRequestMaxResolution - Max resolution can be used to control the maximum resolution your media is encoded, stored, and streamed at.
type CreateLiveStreamRequestMaxResolution string

const (
	CreateLiveStreamRequestMaxResolutionOneThousandAndEightyp  CreateLiveStreamRequestMaxResolution = "1080p"
	CreateLiveStreamRequestMaxResolutionSevenHundredAndTwentyp CreateLiveStreamRequestMaxResolution = "720p"
	CreateLiveStreamRequestMaxResolutionFourHundredAndEightyp  CreateLiveStreamRequestMaxResolution = "480p"
)

func (e CreateLiveStreamRequestMaxResolution) ToPointer() *CreateLiveStreamRequestMaxResolution {
	return &e
}
func (e *CreateLiveStreamRequestMaxResolution) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "1080p":
		fallthrough
	case "720p":
		fallthrough
	case "480p":
		*e = CreateLiveStreamRequestMaxResolution(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateLiveStreamRequestMaxResolution: %v", v)
	}
}

// InputMediaSettings - Displays the result of the input Media settings.
type InputMediaSettings struct {
	// Max resolution can be used to control the maximum resolution your media is encoded, stored, and streamed at.
	MaxResolution *CreateLiveStreamRequestMaxResolution `default:"1080p" json:"maxResolution"`
	// In case the software streaming the live, gets disrupted for any reason and gets disconnected from FastPix, the reconnect window specifies the time span FastPix will wait before ending the stream. Before starting the stream, you can set the reconnect window time which is up to 1800 seconds.
	ReconnectWindow *int64 `default:"60" json:"reconnectWindow"`
	// Basic access policy for media content
	MediaPolicy *BasicAccessPolicy `json:"mediaPolicy,omitempty"`
	// You can search for videos with specific key value pairs using metadata, when you tag a video in "key":"value"s pairs. Dynamic Metadata allows you to define a key that allows any value pair. You can have maximum of 255 characters and upto 10 entries are allowed.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Enables DVR (Digital Video Recorder) functionality for the live stream. When set to true, viewers can pause, rewind, and resume playback during the live broadcast. This allows time-shifted viewing of the stream while it is still ongoing.
	EnableDvrMode *bool `json:"enableDvrMode,omitempty"`
}

func (i InputMediaSettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputMediaSettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (i *InputMediaSettings) GetMaxResolution() *CreateLiveStreamRequestMaxResolution {
	if i == nil {
		return nil
	}
	return i.MaxResolution
}

func (i *InputMediaSettings) GetReconnectWindow() *int64 {
	if i == nil {
		return nil
	}
	return i.ReconnectWindow
}

func (i *InputMediaSettings) GetMediaPolicy() *BasicAccessPolicy {
	if i == nil {
		return nil
	}
	return i.MediaPolicy
}

func (i *InputMediaSettings) GetMetadata() map[string]string {
	if i == nil {
		return nil
	}
	return i.Metadata
}

func (i *InputMediaSettings) GetEnableDvrMode() *bool {
	if i == nil {
		return nil
	}
	return i.EnableDvrMode
}

type CreateLiveStreamRequest struct {
	// Displays the result of the playback settings.
	PlaybackSettings PlaybackSettings `json:"playbackSettings"`
	// Displays the result of the input Media settings.
	InputMediaSettings InputMediaSettings `json:"inputMediaSettings"`
}

func (c *CreateLiveStreamRequest) GetPlaybackSettings() PlaybackSettings {
	if c == nil {
		return PlaybackSettings{}
	}
	return c.PlaybackSettings
}

func (c *CreateLiveStreamRequest) GetInputMediaSettings() InputMediaSettings {
	if c == nil {
		return InputMediaSettings{}
	}
	return c.InputMediaSettings
}
