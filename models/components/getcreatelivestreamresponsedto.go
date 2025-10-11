

package components

import (
	"github.com/FastPix/fastpix-go/internal/utils"
	"time"
)

// GetCreateLiveStreamResponseDTO - Displays the result of the request.
type GetCreateLiveStreamResponseDTO struct {
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
	ReconnectWindow *int64 `default:"60" json:"reconnectWindow"`
	// When set to true, the livestream will be recorded and stored for later viewing purposes. If set to false, the livestream will not be recorded.
	EnableRecording *bool `json:"enableRecording,omitempty"`
	// Determines whether the recorded stream should be publicly accessible or private in Live to VOD (Video on Demand).
	MediaPolicy *string `json:"mediaPolicy,omitempty"`
	// You can search for videos with specific key value pairs using metadata, when you tag a video in "key":"value"s pairs. Dynamic Metadata allows you to define a key that allows any value pair. You can have maximum of 255 characters and upto 10 entries are allowed.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Enables DVR (Digital Video Recorder) functionality for the live stream. When set to true, viewers can pause, rewind, and resume playback during the live broadcast. This allows time-shifted viewing of the stream while it is still ongoing.
	EnableDvrMode *bool `json:"enableDvrMode,omitempty"`
	// A collection of Playback ID objects utilized for crafting HLS playback urls.
	PlaybackIds []PlaybackIDResponse `json:"playbackIds,omitempty"`
	// A set of media IDs created after the livestream ends. (live to VOD)
	MediaIds []string `json:"mediaIds,omitempty"`
	// This object contains the livestream playback response details for SRT Protocol
	SrtPlaybackResponse *SrtPlaybackResponse `json:"srtPlaybackResponse,omitempty"`
}

func (g GetCreateLiveStreamResponseDTO) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetCreateLiveStreamResponseDTO) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (g *GetCreateLiveStreamResponseDTO) GetStreamID() *string {
	if g == nil {
		return nil
	}
	return g.StreamID
}

func (g *GetCreateLiveStreamResponseDTO) GetStreamKey() *string {
	if g == nil {
		return nil
	}
	return g.StreamKey
}

func (g *GetCreateLiveStreamResponseDTO) GetSrtSecret() *string {
	if g == nil {
		return nil
	}
	return g.SrtSecret
}

func (g *GetCreateLiveStreamResponseDTO) GetTrial() *bool {
	if g == nil {
		return nil
	}
	return g.Trial
}

func (g *GetCreateLiveStreamResponseDTO) GetStatus() *string {
	if g == nil {
		return nil
	}
	return g.Status
}

func (g *GetCreateLiveStreamResponseDTO) GetMaxResolution() *string {
	if g == nil {
		return nil
	}
	return g.MaxResolution
}

func (g *GetCreateLiveStreamResponseDTO) GetMaxDuration() *int64 {
	if g == nil {
		return nil
	}
	return g.MaxDuration
}

func (g *GetCreateLiveStreamResponseDTO) GetCreatedAt() *time.Time {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetCreateLiveStreamResponseDTO) GetReconnectWindow() *int64 {
	if g == nil {
		return nil
	}
	return g.ReconnectWindow
}

func (g *GetCreateLiveStreamResponseDTO) GetEnableRecording() *bool {
	if g == nil {
		return nil
	}
	return g.EnableRecording
}

func (g *GetCreateLiveStreamResponseDTO) GetMediaPolicy() *string {
	if g == nil {
		return nil
	}
	return g.MediaPolicy
}

func (g *GetCreateLiveStreamResponseDTO) GetMetadata() map[string]string {
	if g == nil {
		return nil
	}
	return g.Metadata
}

func (g *GetCreateLiveStreamResponseDTO) GetEnableDvrMode() *bool {
	if g == nil {
		return nil
	}
	return g.EnableDvrMode
}

func (g *GetCreateLiveStreamResponseDTO) GetPlaybackIds() []PlaybackIDResponse {
	if g == nil {
		return nil
	}
	return g.PlaybackIds
}

func (g *GetCreateLiveStreamResponseDTO) GetMediaIds() []string {
	if g == nil {
		return nil
	}
	return g.MediaIds
}

func (g *GetCreateLiveStreamResponseDTO) GetSrtPlaybackResponse() *SrtPlaybackResponse {
	if g == nil {
		return nil
	}
	return g.SrtPlaybackResponse
}
