

package components

type DrmIDResponse struct {
	// The unique identifier of the DRM configuration.
	ID *string `json:"id,omitempty"`
}

func (d *DrmIDResponse) GetID() *string {
	if d == nil {
		return nil
	}
	return d.ID
}
