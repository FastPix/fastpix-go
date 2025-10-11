

package components

// Track - A media consists of different media tracks, like video, audio, and subtitle, all combined.
type Track struct {
	// FastPix generates a unique identifier for each track.
	ID *string `json:"id,omitempty"`
	// Defines the type of input. This option is mandatory.
	Type string `json:"type"`
	// Track width denotes the range of widths applicable to a specific track. Currently, this setting can be modified only for video tracks
	Width *float64 `json:"width,omitempty"`
	// Track height denotes the range of height applicable to a specific track. Currently, this setting can be modified only for video tracks.
	Height *float64 `json:"height,omitempty"`
	// Frame rate quantifies the speed at which frames are displayed per second. It represents the range of frames available for a specific track. If the frame rate of the input file is indeterminable, it will be indicated by a value of -1.
	FrameRate *string `json:"frameRate,omitempty"`
	// Indicates if the track contains closed captions.
	ClosedCaptions *bool `json:"closedCaptions,omitempty"`
}

func (t *Track) GetID() *string {
	if t == nil {
		return nil
	}
	return t.ID
}

func (t *Track) GetType() string {
	if t == nil {
		return ""
	}
	return t.Type
}

func (t *Track) GetWidth() *float64 {
	if t == nil {
		return nil
	}
	return t.Width
}

func (t *Track) GetHeight() *float64 {
	if t == nil {
		return nil
	}
	return t.Height
}

func (t *Track) GetFrameRate() *string {
	if t == nil {
		return nil
	}
	return t.FrameRate
}

func (t *Track) GetClosedCaptions() *bool {
	if t == nil {
		return nil
	}
	return t.ClosedCaptions
}
