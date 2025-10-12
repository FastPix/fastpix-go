

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/fastpix/fastpix-go/internal/utils"
	"github.com/fastpix/fastpix-go/models/components"
)

// ListByTopContentTimespan - This parameter specifies the time span between which the video views list should be retrieved by. You can provide either from and to unix epoch timestamps or time duration. The scope of duration is between 60 minutes to 30 days.
type ListByTopContentTimespan string

const (
	ListByTopContentTimespanSixtyminutes    ListByTopContentTimespan = "60:minutes"
	ListByTopContentTimespanSixhours        ListByTopContentTimespan = "6:hours"
	ListByTopContentTimespanTwentyFourhours ListByTopContentTimespan = "24:hours"
	ListByTopContentTimespanThreedays       ListByTopContentTimespan = "3:days"
	ListByTopContentTimespanSevendays       ListByTopContentTimespan = "7:days"
	ListByTopContentTimespanThirtydays      ListByTopContentTimespan = "30:days"
)

func (e ListByTopContentTimespan) ToPointer() *ListByTopContentTimespan {
	return &e
}
func (e *ListByTopContentTimespan) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "60:minutes":
		fallthrough
	case "6:hours":
		fallthrough
	case "24:hours":
		fallthrough
	case "3:days":
		fallthrough
	case "7:days":
		fallthrough
	case "30:days":
		*e = ListByTopContentTimespan(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListByTopContentTimespan: %v", v)
	}
}

type ListByTopContentRequest struct {
	// This parameter specifies the time span between which the video views list should be retrieved by. You can provide either from and to unix epoch timestamps or time duration. The scope of duration is between 60 minutes to 30 days.
	//
	Timespan ListByTopContentTimespan `queryParam:"style=form,explode=true,name=timespan[]"`
	// Pass the dimensions and their corresponding values you want to filter the views by. For excluding the values in the filter we can pass '!' before the filter value. The list of filters can be obtained from list of dimensions endpoint.
	// Example Values : [ browser_name:Chrome , os_name:macOS , device_name:Galaxy ]
	//
	Filterby *string `queryParam:"style=form,explode=true,name=filterby[]"`
	// Pass the limit to display only the rows specified by the value.
	//
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
}

func (l ListByTopContentRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListByTopContentRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, []string{"timespan[]"}); err != nil {
		return err
	}
	return nil
}

func (l *ListByTopContentRequest) GetTimespan() ListByTopContentTimespan {
	if l == nil {
		return ListByTopContentTimespan("")
	}
	return l.Timespan
}

func (l *ListByTopContentRequest) GetFilterby() *string {
	if l == nil {
		return nil
	}
	return l.Filterby
}

func (l *ListByTopContentRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

// ListByTopContentResponseBody - Get the list of Views
type ListByTopContentResponseBody struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data []components.ViewsByTopContentDetails `json:"data,omitempty"`
}

func (l *ListByTopContentResponseBody) GetSuccess() *bool {
	if l == nil {
		return nil
	}
	return l.Success
}

func (l *ListByTopContentResponseBody) GetData() []components.ViewsByTopContentDetails {
	if l == nil {
		return nil
	}
	return l.Data
}

type ListByTopContentResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Get the list of Views
	Object *ListByTopContentResponseBody
}

func (l *ListByTopContentResponse) GetHTTPMeta() components.HTTPMetadata {
	if l == nil {
		return components.HTTPMetadata{}
	}
	return l.HTTPMeta
}

func (l *ListByTopContentResponse) GetObject() *ListByTopContentResponseBody {
	if l == nil {
		return nil
	}
	return l.Object
}
