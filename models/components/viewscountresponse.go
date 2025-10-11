

package components

// ViewsCountResponseData - Contains the view count details.
type ViewsCountResponseData struct {
	// Number of views for the stream or resource.
	Views *int64 `json:"views,omitempty"`
}

func (v *ViewsCountResponseData) GetViews() *int64 {
	if v == nil {
		return nil
	}
	return v.Views
}

type ViewsCountResponse struct {
	// Indicates whether the request was successful or not.
	Success *bool `json:"success,omitempty"`
	// Contains the view count details.
	Data *ViewsCountResponseData `json:"data,omitempty"`
}

func (v *ViewsCountResponse) GetSuccess() *bool {
	if v == nil {
		return nil
	}
	return v.Success
}

func (v *ViewsCountResponse) GetData() *ViewsCountResponseData {
	if v == nil {
		return nil
	}
	return v.Data
}
