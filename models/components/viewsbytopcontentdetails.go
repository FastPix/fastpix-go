

package components

// ViewsByTopContentDetails - Retrieves a list of the top video views
type ViewsByTopContentDetails struct {
	// Title of the video
	VideoTitle *string `json:"videoTitle,omitempty"`
	// Total count of view sessions for a paricular video content.
	Views *int64 `json:"views,omitempty"`
	// Total count of unique video viewers for particular video content.
	UniqueViews *int64 `json:"uniqueViews,omitempty"`
}

func (v *ViewsByTopContentDetails) GetVideoTitle() *string {
	if v == nil {
		return nil
	}
	return v.VideoTitle
}

func (v *ViewsByTopContentDetails) GetViews() *int64 {
	if v == nil {
		return nil
	}
	return v.Views
}

func (v *ViewsByTopContentDetails) GetUniqueViews() *int64 {
	if v == nil {
		return nil
	}
	return v.UniqueViews
}
