package components

import (
	"encoding/json"
	"fmt"
)

// MediaType - Type of media content
type MediaType string

const (
	MediaTypeVideo MediaType = "video"
	MediaTypeAudio MediaType = "audio"
	MediaTypeAv    MediaType = "av"
)

func (e MediaType) ToPointer() *MediaType {
	return &e
}
func (e *MediaType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "video":
		fallthrough
	case "audio":
		fallthrough
	case "av":
		*e = MediaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MediaType: %v", v)
	}
}
