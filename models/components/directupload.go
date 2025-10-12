

package components

import (
	"github.com/fastpix/fastpix-go/internal/utils"
)

// DirectUpload - Displays the result of the request.
type DirectUpload struct {
	// When creating the upload, FastPix assigns a universally unique identifier with a maximum length of 255 characters.
	ID *string `json:"id,omitempty"`
	// When creating the media, FastPix assigns a universally unique identifier with a maximum length of 255 characters.
	MediaID *string `json:"mediaId,omitempty"`
	// Determines the media's status, which can be one of the possible values.
	Status *string `json:"status,omitempty"`
	// The url hosts the media file for FastPix, which needs to be download to use further.  It supports formats like MP3, MP4, MOV, MKV, or TS, and includes text tracks for subtitles/CC (SRT file/VTT file). While FastPix can handle various audio and video formats and codecs, using standard inputs can help with optimal processing speed.
	URL *string `json:"url,omitempty"`
	// The duration set for the validity of the upload URL. If the upload isn't completed within this timeframe, it's marked as timed out.
	//
	Timeout *float64 `default:"14400" json:"timeout"`
	// Upload media directly from a device using the url name or enter '*' to allow all.
	CorsOrigin        *string               `json:"corsOrigin,omitempty"`
	PushMediaSettings *DirectUploadResponse `json:"pushMediaSettings,omitempty"`
}

func (d DirectUpload) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DirectUpload) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (d *DirectUpload) GetID() *string {
	if d == nil {
		return nil
	}
	return d.ID
}

func (d *DirectUpload) GetMediaID() *string {
	if d == nil {
		return nil
	}
	return d.MediaID
}

func (d *DirectUpload) GetStatus() *string {
	if d == nil {
		return nil
	}
	return d.Status
}

func (d *DirectUpload) GetURL() *string {
	if d == nil {
		return nil
	}
	return d.URL
}

func (d *DirectUpload) GetTimeout() *float64 {
	if d == nil {
		return nil
	}
	return d.Timeout
}

func (d *DirectUpload) GetCorsOrigin() *string {
	if d == nil {
		return nil
	}
	return d.CorsOrigin
}

func (d *DirectUpload) GetPushMediaSettings() *DirectUploadResponse {
	if d == nil {
		return nil
	}
	return d.PushMediaSettings
}
