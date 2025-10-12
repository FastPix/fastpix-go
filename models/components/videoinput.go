

package components

import (
	"errors"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
)

type Segment2 struct {
	// URL of the segment to be added.
	URL string `json:"url"`
	// Flag indicating the segment should be inserted at the end.
	InsertAtEnd bool `json:"insertAtEnd"`
}

func (s Segment2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Segment2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, []string{"url", "insertAtEnd"}); err != nil {
		return err
	}
	return nil
}

func (s *Segment2) GetURL() string {
	if s == nil {
		return ""
	}
	return s.URL
}

func (s *Segment2) GetInsertAtEnd() bool {
	if s == nil {
		return false
	}
	return s.InsertAtEnd
}

type Segment1 struct {
	// URL of the segment to be added.
	URL string `json:"url"`
	// The timestamp at which the segment should be inserted.
	InsertAt int64 `json:"insertAt"`
}

func (s Segment1) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Segment1) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, []string{"url", "insertAt"}); err != nil {
		return err
	}
	return nil
}

func (s *Segment1) GetURL() string {
	if s == nil {
		return ""
	}
	return s.URL
}

func (s *Segment1) GetInsertAt() int64 {
	if s == nil {
		return 0
	}
	return s.InsertAt
}

type SegmentUnionType string

const (
	SegmentUnionTypeSegment1 SegmentUnionType = "segment_1"
	SegmentUnionTypeSegment2 SegmentUnionType = "segment_2"
)

type SegmentUnion struct {
	Segment1 *Segment1 `queryParam:"inline,name=segment"`
	Segment2 *Segment2 `queryParam:"inline,name=segment"`

	Type SegmentUnionType
}

func CreateSegmentUnionSegment1(segment1 Segment1) SegmentUnion {
	typ := SegmentUnionTypeSegment1

	return SegmentUnion{
		Segment1: &segment1,
		Type:     typ,
	}
}

func CreateSegmentUnionSegment2(segment2 Segment2) SegmentUnion {
	typ := SegmentUnionTypeSegment2

	return SegmentUnion{
		Segment2: &segment2,
		Type:     typ,
	}
}

func (u *SegmentUnion) UnmarshalJSON(data []byte) error {

	var segment1 Segment1 = Segment1{}
	if err := utils.UnmarshalJSON(data, &segment1, "", true, nil); err == nil {
		u.Segment1 = &segment1
		u.Type = SegmentUnionTypeSegment1
		return nil
	}

	var segment2 Segment2 = Segment2{}
	if err := utils.UnmarshalJSON(data, &segment2, "", true, nil); err == nil {
		u.Segment2 = &segment2
		u.Type = SegmentUnionTypeSegment2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SegmentUnion", string(data))
}

func (u SegmentUnion) MarshalJSON() ([]byte, error) {
	if u.Segment1 != nil {
		return utils.MarshalJSON(u.Segment1, "", true)
	}

	if u.Segment2 != nil {
		return utils.MarshalJSON(u.Segment2, "", true)
	}

	return nil, errors.New("could not marshal union type SegmentUnion: all fields are null")
}

type VideoInput struct {
	// Defines the type of input.
	//
	Type string `json:"type"`
	// The url hosts the media file for FastPix, which needs to be downloaded to use further. It supports formats like MP3, MP4, MOV, MKV, or TS, and includes text tracks for subtitles/CC (SRT file/VTT file). While FastPix can handle various audio and video formats and codecs, using standard inputs can help with optimal processing speed.
	//
	URL string `json:"url"`
	// Start time indicates where encoding should begin within the video file. For example, if you want to encode a segment from 3 minutes (180 seconds) to 6 minutes (360 seconds) in a 10-minute (600 seconds) video, the start time is 3 minutes (180 seconds). Note: Start time is always mentioned in seconds.
	//
	StartTime *float64 `json:"startTime,omitempty"`
	// End time indicates where encoding should end within the video file. For example, if you want to encode a segment from 3 minutes (180 seconds) to 6 minutes (360 seconds) in a 10-minute (600 seconds) video, the end time is 6 minutes (360 seconds). Note: End time is always mentioned in seconds.
	//
	EndTime *float64 `json:"endTime,omitempty"`
	// The url of the intro video which is to be added at the start of the video.
	//
	IntroURL *string `json:"introUrl,omitempty"`
	// The url of the outro video which is to be added at the end of the video.
	//
	OutroURL *string `json:"outroUrl,omitempty"`
	// The list of the startTime-endTime of the segments to be removed from the actual video.
	//
	ExpungeSegments []string `json:"expungeSegments,omitempty"`
	// A list of media segments to be added or processed. Each segment includes details such as the URL of the media file and instructions on where it should be inserted in the final media composition. A segment can either specify an exact timestamp  (`insertAt`) or indicate that it should be added at the end (`insertAtEnd`).
	Segments []SegmentUnion `json:"segments,omitempty"`
}

func (v VideoInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *VideoInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, []string{"type", "url"}); err != nil {
		return err
	}
	return nil
}

func (v *VideoInput) GetType() string {
	if v == nil {
		return ""
	}
	return v.Type
}

func (v *VideoInput) GetURL() string {
	if v == nil {
		return ""
	}
	return v.URL
}

func (v *VideoInput) GetStartTime() *float64 {
	if v == nil {
		return nil
	}
	return v.StartTime
}

func (v *VideoInput) GetEndTime() *float64 {
	if v == nil {
		return nil
	}
	return v.EndTime
}

func (v *VideoInput) GetIntroURL() *string {
	if v == nil {
		return nil
	}
	return v.IntroURL
}

func (v *VideoInput) GetOutroURL() *string {
	if v == nil {
		return nil
	}
	return v.OutroURL
}

func (v *VideoInput) GetExpungeSegments() []string {
	if v == nil {
		return nil
	}
	return v.ExpungeSegments
}

func (v *VideoInput) GetSegments() []SegmentUnion {
	if v == nil {
		return nil
	}
	return v.Segments
}
