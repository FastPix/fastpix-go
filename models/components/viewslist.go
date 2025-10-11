

package components

import (
	"github.com/FastPix/fastpix-go/optionalnullable"
)

type ViewsList struct {
	// The unique identifier for the viewing session of the user.
	//
	ViewID string `json:"viewId"`
	// Operating System signifies the software platform utilized by the viewer
	//
	OperatingSystem *string `json:"operatingSystem"`
	// The browser name of the viewer.
	//
	Application *string `json:"application"`
	// The start timestamp of the video view.
	//
	ViewStartTime *string `json:"viewStartTime"`
	// The end timestamp of the video view.
	//
	ViewEndTime *string `json:"viewEndTime"`
	// The title of the Video.
	//
	VideoTitle *string `json:"videoTitle"`
	// The code which represents specific issues or failures that occur during playback. These can be implementation specific.
	//
	ErrorCode optionalnullable.OptionalNullable[string] `json:"errorCode,omitempty"`
	// The notifications or messages that inform users or developers about issues or failures that have occurred during the playback representing error codes.
	//
	ErrorMessage optionalnullable.OptionalNullable[string] `json:"errorMessage,omitempty"`
	// The unique identifier which identifies each type of error that occurs.
	//
	ErrorID optionalnullable.OptionalNullable[int64] `json:"errorId,omitempty"`
	// Country of the viewer.
	//
	Country *string `json:"country"`
	// The watch time represents the time spent watching the video including staruptime, playback time ,buffering time.
	//
	ViewWatchTime optionalnullable.OptionalNullable[float64] `json:"viewWatchTime,omitempty"`
	// The viewer experience encapsulated in the form of score while watching the video.
	//
	QoeScore optionalnullable.OptionalNullable[float64] `json:"QoeScore,omitempty"`
}

func (v *ViewsList) GetViewID() string {
	if v == nil {
		return ""
	}
	return v.ViewID
}

func (v *ViewsList) GetOperatingSystem() *string {
	if v == nil {
		return nil
	}
	return v.OperatingSystem
}

func (v *ViewsList) GetApplication() *string {
	if v == nil {
		return nil
	}
	return v.Application
}

func (v *ViewsList) GetViewStartTime() *string {
	if v == nil {
		return nil
	}
	return v.ViewStartTime
}

func (v *ViewsList) GetViewEndTime() *string {
	if v == nil {
		return nil
	}
	return v.ViewEndTime
}

func (v *ViewsList) GetVideoTitle() *string {
	if v == nil {
		return nil
	}
	return v.VideoTitle
}

func (v *ViewsList) GetErrorCode() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.ErrorCode
}

func (v *ViewsList) GetErrorMessage() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.ErrorMessage
}

func (v *ViewsList) GetErrorID() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.ErrorID
}

func (v *ViewsList) GetCountry() *string {
	if v == nil {
		return nil
	}
	return v.Country
}

func (v *ViewsList) GetViewWatchTime() optionalnullable.OptionalNullable[float64] {
	if v == nil {
		return nil
	}
	return v.ViewWatchTime
}

func (v *ViewsList) GetQoeScore() optionalnullable.OptionalNullable[float64] {
	if v == nil {
		return nil
	}
	return v.QoeScore
}
