

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/fastpix/fastpix-go/models/components"
)

type CreateMediaPlaybackIDAccessRestrictions struct {
	// Restrictions based on the originating domain of a request
	Domains *components.DomainRestrictions `json:"domains,omitempty"`
	// Restrictions based on the user agent
	UserAgents *components.UserAgentRestrictions `json:"userAgents,omitempty"`
}

func (c *CreateMediaPlaybackIDAccessRestrictions) GetDomains() *components.DomainRestrictions {
	if c == nil {
		return nil
	}
	return c.Domains
}

func (c *CreateMediaPlaybackIDAccessRestrictions) GetUserAgents() *components.UserAgentRestrictions {
	if c == nil {
		return nil
	}
	return c.UserAgents
}

// Resolution - The maximum resolution for the playback ID.
type Resolution string

const (
	ResolutionFourHundredAndEightyp           Resolution = "480p"
	ResolutionSevenHundredAndTwentyp          Resolution = "720p"
	ResolutionOneThousandAndEightyp           Resolution = "1080p"
	ResolutionOneThousandFourHundredAndFortyp Resolution = "1440p"
	ResolutionTwoThousandOneHundredAndSixtyp  Resolution = "2160p"
)

func (e Resolution) ToPointer() *Resolution {
	return &e
}
func (e *Resolution) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "480p":
		fallthrough
	case "720p":
		fallthrough
	case "1080p":
		fallthrough
	case "1440p":
		fallthrough
	case "2160p":
		*e = Resolution(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Resolution: %v", v)
	}
}

// CreateMediaPlaybackIDRequestBody - Request body for creating playback id for an media
type CreateMediaPlaybackIDRequestBody struct {
	// Access policy for media content
	AccessPolicy       components.AccessPolicy                  `json:"accessPolicy"`
	AccessRestrictions *CreateMediaPlaybackIDAccessRestrictions `json:"accessRestrictions,omitempty"`
	// DRM configuration ID (required if accessPolicy is 'drm')
	DrmConfigurationID *string `json:"drmConfigurationId,omitempty"`
	// The maximum resolution for the playback ID.
	Resolution *Resolution `json:"resolution,omitempty"`
}

func (c *CreateMediaPlaybackIDRequestBody) GetAccessPolicy() components.AccessPolicy {
	if c == nil {
		return components.AccessPolicy("")
	}
	return c.AccessPolicy
}

func (c *CreateMediaPlaybackIDRequestBody) GetAccessRestrictions() *CreateMediaPlaybackIDAccessRestrictions {
	if c == nil {
		return nil
	}
	return c.AccessRestrictions
}

func (c *CreateMediaPlaybackIDRequestBody) GetDrmConfigurationID() *string {
	if c == nil {
		return nil
	}
	return c.DrmConfigurationID
}

func (c *CreateMediaPlaybackIDRequestBody) GetResolution() *Resolution {
	if c == nil {
		return nil
	}
	return c.Resolution
}

type CreateMediaPlaybackIDRequest struct {
	// When creating the media, FastPix assigns a universally unique identifier with a maximum length of 255 characters.
	MediaID string `pathParam:"style=simple,explode=false,name=mediaId"`
	// Request body for creating playback id for an media
	RequestBody *CreateMediaPlaybackIDRequestBody `request:"mediaType=application/json"`
}

func (c *CreateMediaPlaybackIDRequest) GetMediaID() string {
	if c == nil {
		return ""
	}
	return c.MediaID
}

func (c *CreateMediaPlaybackIDRequest) GetRequestBody() *CreateMediaPlaybackIDRequestBody {
	if c == nil {
		return nil
	}
	return c.RequestBody
}

// CreateMediaPlaybackIDData - Displays the result of the request.
type CreateMediaPlaybackIDData struct {
	// A collection of Playback ID objects utilized for crafting HLS playback URLs.
	PlaybackIds []components.PlaybackID `json:"playbackIds,omitempty"`
}

func (c *CreateMediaPlaybackIDData) GetPlaybackIds() []components.PlaybackID {
	if c == nil {
		return nil
	}
	return c.PlaybackIds
}

// CreateMediaPlaybackIDResponseBody - Playback id for an media
type CreateMediaPlaybackIDResponseBody struct {
	// Demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data *CreateMediaPlaybackIDData `json:"data,omitempty"`
}

func (c *CreateMediaPlaybackIDResponseBody) GetSuccess() *bool {
	if c == nil {
		return nil
	}
	return c.Success
}

func (c *CreateMediaPlaybackIDResponseBody) GetData() *CreateMediaPlaybackIDData {
	if c == nil {
		return nil
	}
	return c.Data
}

type CreateMediaPlaybackIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Playback id for an media
	Object *CreateMediaPlaybackIDResponseBody
}

func (c *CreateMediaPlaybackIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CreateMediaPlaybackIDResponse) GetObject() *CreateMediaPlaybackIDResponseBody {
	if c == nil {
		return nil
	}
	return c.Object
}
