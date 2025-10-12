

package components

import (
	"github.com/fastpix/fastpix-go/internal/utils"
	"time"
)

// PatchResponseData - Displays the result of the request.
type PatchResponseData struct {
	// Upon creating a new live stream, FastPix assigns a unique identifier to the stream.
	StreamID *string `json:"streamId,omitempty"`
	// A unique stream key is generated for streaming, allowing the user to start streaming on any third-party platform using this key.
	StreamKey *string `json:"streamKey,omitempty"`
	// A secret used for securing the SRT stream. This ensures that only authorized users can access the stream.
	SrtSecret *string `json:"srtSecret,omitempty"`
	// FastPix allows for a to trial the live stream for free. The duration of trial streams is five minutes. After five minutes of activity, the trial stream is turned off, and the recorded asset is removed after a day.
	Trial *bool `json:"trial,omitempty"`
	// The current live stream status can be one of four values:Idle, Preparing, Active or Disabled.The Idle status signifies that there isn't a broadcast in progress.The preparing status indicates that the stream is getting prepared. while, the Active status indicates that a broadcast is currently in progress. The Disabled status means that no more RTMPS streams can be published.
	Status *string `json:"status,omitempty"`
	// Max resolution can be used to control the maximum resolution your media is encoded, stored, and streamed at.
	MaxResolution *string `json:"maxResolution,omitempty"`
	// The maximum duration in seconds that a live stream can have before it ends the stream.
	MaxDuration *int64 `json:"maxDuration,omitempty"`
	// It is the moment when the stream was created Time the media was generated, defined as a localDateTime (UTC Time).
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// In case the software streaming the live, gets disrupted for any reason and gets disconnected from FastPix, the reconnect window specifies the time span FastPix will wait before ending the stream. Before starting the stream, you can set the reconnect window time which is up to 1800 seconds.
	ReconnectWindow *int64 `json:"reconnectWindow,omitempty"`
	// When set to true, the livestream will be recorded and stored for later viewing purposes. If set to false, the livestream will not be recorded.
	EnableRecording *bool `json:"enableRecording,omitempty"`
	// Determines whether the recorded stream should be publicly accessible or private in Live to VOD (Video on Demand).
	MediaPolicy *string `json:"mediaPolicy,omitempty"`
	// You can search for videos with specific key value pairs using metadata, when you tag a video in "key":"value"s pairs. Dynamic Metadata allows you to define a key that allows any value pair. You can have maximum of 255 characters and upto 10 entries are allowed.
	Metadata    map[string]string    `json:"metadata,omitempty"`
	PlaybackIds []PlaybackIDResponse `json:"playbackIds,omitempty"`
	// This object contains the livestream playback response details for SRT Protocol
	SrtPlaybackResponse *SrtPlaybackResponse `json:"srtPlaybackResponse,omitempty"`
}

func (p PatchResponseData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PatchResponseData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PatchResponseData) GetStreamID() *string {
	if p == nil {
		return nil
	}
	return p.StreamID
}

func (p *PatchResponseData) GetStreamKey() *string {
	if p == nil {
		return nil
	}
	return p.StreamKey
}

func (p *PatchResponseData) GetSrtSecret() *string {
	if p == nil {
		return nil
	}
	return p.SrtSecret
}

func (p *PatchResponseData) GetTrial() *bool {
	if p == nil {
		return nil
	}
	return p.Trial
}

func (p *PatchResponseData) GetStatus() *string {
	if p == nil {
		return nil
	}
	return p.Status
}

func (p *PatchResponseData) GetMaxResolution() *string {
	if p == nil {
		return nil
	}
	return p.MaxResolution
}

func (p *PatchResponseData) GetMaxDuration() *int64 {
	if p == nil {
		return nil
	}
	return p.MaxDuration
}

func (p *PatchResponseData) GetCreatedAt() *time.Time {
	if p == nil {
		return nil
	}
	return p.CreatedAt
}

func (p *PatchResponseData) GetReconnectWindow() *int64 {
	if p == nil {
		return nil
	}
	return p.ReconnectWindow
}

func (p *PatchResponseData) GetEnableRecording() *bool {
	if p == nil {
		return nil
	}
	return p.EnableRecording
}

func (p *PatchResponseData) GetMediaPolicy() *string {
	if p == nil {
		return nil
	}
	return p.MediaPolicy
}

func (p *PatchResponseData) GetMetadata() map[string]string {
	if p == nil {
		return nil
	}
	return p.Metadata
}

func (p *PatchResponseData) GetPlaybackIds() []PlaybackIDResponse {
	if p == nil {
		return nil
	}
	return p.PlaybackIds
}

func (p *PatchResponseData) GetSrtPlaybackResponse() *SrtPlaybackResponse {
	if p == nil {
		return nil
	}
	return p.SrtPlaybackResponse
}
