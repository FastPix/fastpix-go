

package operations

import (
	"github.com/fastpix/fastpix-go/models/components"
)

// ListDimensionsResponseBody - Get the list of Views
type ListDimensionsResponseBody struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data []string `json:"data,omitempty"`
}

func (l *ListDimensionsResponseBody) GetSuccess() *bool {
	if l == nil {
		return nil
	}
	return l.Success
}

func (l *ListDimensionsResponseBody) GetData() []string {
	if l == nil {
		return nil
	}
	return l.Data
}

type ListDimensionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Get the list of Views
	Object *ListDimensionsResponseBody
}

func (l *ListDimensionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if l == nil {
		return components.HTTPMetadata{}
	}
	return l.HTTPMeta
}

func (l *ListDimensionsResponse) GetObject() *ListDimensionsResponseBody {
	if l == nil {
		return nil
	}
	return l.Object
}
