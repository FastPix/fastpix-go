

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/FastPix/fastpix-go/models/components"
)

// UpdatedMp4SupportMp4Support - Determines the type of MP4 support for the media.   - **none**: Disables MP4 support.   - **capped_4k**: Enables MP4 downloads with resolutions up to 4K.   - **audioOnly**: Provides an MP4 stream containing only the audio.   - **audioOnly,capped_4k**: Enables both MP4 video downloads (up to 4K) and an audio-only stream.
type UpdatedMp4SupportMp4Support string

const (
	UpdatedMp4SupportMp4SupportNone              UpdatedMp4SupportMp4Support = "none"
	UpdatedMp4SupportMp4SupportCapped4k          UpdatedMp4SupportMp4Support = "capped_4k"
	UpdatedMp4SupportMp4SupportAudioOnly         UpdatedMp4SupportMp4Support = "audioOnly"
	UpdatedMp4SupportMp4SupportAudioOnlyCapped4k UpdatedMp4SupportMp4Support = "audioOnly,capped_4k"
)

func (e UpdatedMp4SupportMp4Support) ToPointer() *UpdatedMp4SupportMp4Support {
	return &e
}
func (e *UpdatedMp4SupportMp4Support) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "capped_4k":
		fallthrough
	case "audioOnly":
		fallthrough
	case "audioOnly,capped_4k":
		*e = UpdatedMp4SupportMp4Support(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdatedMp4SupportMp4Support: %v", v)
	}
}

type UpdatedMp4SupportRequestBody struct {
	// Determines the type of MP4 support for the media.   - **none**: Disables MP4 support.   - **capped_4k**: Enables MP4 downloads with resolutions up to 4K.   - **audioOnly**: Provides an MP4 stream containing only the audio.   - **audioOnly,capped_4k**: Enables both MP4 video downloads (up to 4K) and an audio-only stream.
	Mp4Support *UpdatedMp4SupportMp4Support `json:"mp4Support,omitempty"`
}

func (u *UpdatedMp4SupportRequestBody) GetMp4Support() *UpdatedMp4SupportMp4Support {
	if u == nil {
		return nil
	}
	return u.Mp4Support
}

type UpdatedMp4SupportRequest struct {
	// When creating the media, FastPix assigns a universally unique identifier with a maximum length of 255 characters.
	//
	MediaID     string                       `pathParam:"style=simple,explode=false,name=mediaId"`
	RequestBody UpdatedMp4SupportRequestBody `request:"mediaType=application/json"`
}

func (u *UpdatedMp4SupportRequest) GetMediaID() string {
	if u == nil {
		return ""
	}
	return u.MediaID
}

func (u *UpdatedMp4SupportRequest) GetRequestBody() UpdatedMp4SupportRequestBody {
	if u == nil {
		return UpdatedMp4SupportRequestBody{}
	}
	return u.RequestBody
}

// UpdatedMp4SupportResponseBody - Media details updated successfully
type UpdatedMp4SupportResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool             `json:"success,omitempty"`
	Data    *components.Media `json:"data,omitempty"`
}

func (u *UpdatedMp4SupportResponseBody) GetSuccess() *bool {
	if u == nil {
		return nil
	}
	return u.Success
}

func (u *UpdatedMp4SupportResponseBody) GetData() *components.Media {
	if u == nil {
		return nil
	}
	return u.Data
}

type UpdatedMp4SupportResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Media details updated successfully
	Object *UpdatedMp4SupportResponseBody
}

func (u *UpdatedMp4SupportResponse) GetHTTPMeta() components.HTTPMetadata {
	if u == nil {
		return components.HTTPMetadata{}
	}
	return u.HTTPMeta
}

func (u *UpdatedMp4SupportResponse) GetObject() *UpdatedMp4SupportResponseBody {
	if u == nil {
		return nil
	}
	return u.Object
}
