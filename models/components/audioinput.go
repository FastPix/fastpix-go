

package components

import (
	"encoding/json"
	"fmt"
	"github.com/fastpix/fastpix-go/internal/utils"
)

// AudioInputType - Type of overlay (currently only supports 'audio').
type AudioInputType string

const (
	AudioInputTypeAudio AudioInputType = "audio"
)

func (e AudioInputType) ToPointer() *AudioInputType {
	return &e
}
func (e *AudioInputType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "audio":
		*e = AudioInputType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AudioInputType: %v", v)
	}
}

type ImposeTrack struct {
	// URL of the audio track to impose on the video.
	URL *string `json:"url,omitempty"`
	// Start time (in seconds) of the imposed audio in the video.
	StartTime *int64 `json:"startTime,omitempty"`
	// End time (in seconds) of the imposed audio in the video.
	EndTime *int64 `json:"endTime,omitempty"`
	// Level of fade-in effect (in seconds) at the start of the imposed audio.
	FadeInLevel *int64 `json:"fadeInLevel,omitempty"`
	// Level of fade-out effect (in seconds) at the end of the imposed audio.
	FadeOutLevel *int64 `json:"fadeOutLevel,omitempty"`
}

func (i ImposeTrack) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *ImposeTrack) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (i *ImposeTrack) GetURL() *string {
	if i == nil {
		return nil
	}
	return i.URL
}

func (i *ImposeTrack) GetStartTime() *int64 {
	if i == nil {
		return nil
	}
	return i.StartTime
}

func (i *ImposeTrack) GetEndTime() *int64 {
	if i == nil {
		return nil
	}
	return i.EndTime
}

func (i *ImposeTrack) GetFadeInLevel() *int64 {
	if i == nil {
		return nil
	}
	return i.FadeInLevel
}

func (i *ImposeTrack) GetFadeOutLevel() *int64 {
	if i == nil {
		return nil
	}
	return i.FadeOutLevel
}

type AudioInput struct {
	// Type of overlay (currently only supports 'audio').
	Type *AudioInputType `json:"type,omitempty"`
	// URL of the audio track to replace the existing audio in the video.
	SwapTrackURL *string `json:"swapTrackUrl,omitempty"`
	// List of additional audio tracks to overlay on the video.
	ImposeTracks []ImposeTrack `json:"imposeTracks,omitempty"`
}

func (a AudioInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AudioInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AudioInput) GetType() *AudioInputType {
	if a == nil {
		return nil
	}
	return a.Type
}

func (a *AudioInput) GetSwapTrackURL() *string {
	if a == nil {
		return nil
	}
	return a.SwapTrackURL
}

func (a *AudioInput) GetImposeTracks() []ImposeTrack {
	if a == nil {
		return nil
	}
	return a.ImposeTracks
}
