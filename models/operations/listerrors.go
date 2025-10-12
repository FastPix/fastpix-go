

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
	"github.com/FastPix/fastpix-go/models/components"
)

// ListErrorsTimespan - This parameter specifies the time span between which the video views list should be retrieved by. You can provide either from and to unix epoch timestamps or time duration. The scope of duration is between 60 minutes to 30 days.
type ListErrorsTimespan string

const (
	ListErrorsTimespanSixtyminutes    ListErrorsTimespan = "60:minutes"
	ListErrorsTimespanSixhours        ListErrorsTimespan = "6:hours"
	ListErrorsTimespanTwentyFourhours ListErrorsTimespan = "24:hours"
	ListErrorsTimespanThreedays       ListErrorsTimespan = "3:days"
	ListErrorsTimespanSevendays       ListErrorsTimespan = "7:days"
	ListErrorsTimespanThirtydays      ListErrorsTimespan = "30:days"
)

func (e ListErrorsTimespan) ToPointer() *ListErrorsTimespan {
	return &e
}
func (e *ListErrorsTimespan) UnmarshalJSON(data []byte) error {
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
		*e = ListErrorsTimespan(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListErrorsTimespan: %v", v)
	}
}

type ListErrorsRequest struct {
	// This parameter specifies the time span between which the video views list should be retrieved by. You can provide either from and to unix epoch timestamps or time duration. The scope of duration is between 60 minutes to 30 days.
	//
	Timespan ListErrorsTimespan `queryParam:"style=form,explode=true,name=timespan[]"`
	// Pass the dimensions and their corresponding values you want to filter the views by. For excluding the values in the filter we can pass '!' before the filter value. The list of filters can be obtained from list of dimensions endpoint.
	// Example Values : [ browser_name:Chrome , os_name:macOS , device_name:Galaxy ]
	//
	Filterby *string `queryParam:"style=form,explode=true,name=filterby[]"`
	// Pass the limit to display only the rows specified by the value for top errors.
	//
	Limit *int64 `default:"1" queryParam:"style=form,explode=true,name=limit"`
}

func (l ListErrorsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListErrorsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, []string{"timespan[]"}); err != nil {
		return err
	}
	return nil
}

func (l *ListErrorsRequest) GetTimespan() ListErrorsTimespan {
	if l == nil {
		return ListErrorsTimespan("")
	}
	return l.Timespan
}

func (l *ListErrorsRequest) GetFilterby() *string {
	if l == nil {
		return nil
	}
	return l.Filterby
}

func (l *ListErrorsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

// ListErrorsData - Displays the result of the request.
type ListErrorsData struct {
	// Retrieves a list of errors that have occurred in the system.
	Errors []components.ErrorDetails `json:"errors,omitempty"`
	// Retrieves a list of errors that have occurred most frequently in the system, ranked by their count of occurrences.
	TopErrors []components.TopErrorDetails `json:"topErrors,omitempty"`
}

func (l *ListErrorsData) GetErrors() []components.ErrorDetails {
	if l == nil {
		return nil
	}
	return l.Errors
}

func (l *ListErrorsData) GetTopErrors() []components.TopErrorDetails {
	if l == nil {
		return nil
	}
	return l.TopErrors
}

// ListErrorsResponseBody - Get filter/ dimension value details by dimension name.
type ListErrorsResponseBody struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	Data *ListErrorsData `json:"data,omitempty"`
	// The timeframe from and to details displayed in the form of unix epoch timestamps.
	//
	Timespan []int64 `json:"timespan,omitempty"`
}

func (l *ListErrorsResponseBody) GetSuccess() *bool {
	if l == nil {
		return nil
	}
	return l.Success
}

func (l *ListErrorsResponseBody) GetData() *ListErrorsData {
	if l == nil {
		return nil
	}
	return l.Data
}

func (l *ListErrorsResponseBody) GetTimespan() []int64 {
	if l == nil {
		return nil
	}
	return l.Timespan
}

type ListErrorsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Get filter/ dimension value details by dimension name.
	Object *ListErrorsResponseBody
}

func (l *ListErrorsResponse) GetHTTPMeta() components.HTTPMetadata {
	if l == nil {
		return components.HTTPMetadata{}
	}
	return l.HTTPMeta
}

func (l *ListErrorsResponse) GetObject() *ListErrorsResponseBody {
	if l == nil {
		return nil
	}
	return l.Object
}
