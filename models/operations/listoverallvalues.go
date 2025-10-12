package operations

import (
	"encoding/json"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
	"github.com/FastPix/fastpix-go/models/components"
)

// ListOverallValuesMetricID - Pass metric Id
type ListOverallValuesMetricID string

const (
	ListOverallValuesMetricIDViews                         ListOverallValuesMetricID = "views"
	ListOverallValuesMetricIDUniqueViewers                 ListOverallValuesMetricID = "unique_viewers"
	ListOverallValuesMetricIDPlayingTime                   ListOverallValuesMetricID = "playing_time"
	ListOverallValuesMetricIDQualityOfExperienceScore      ListOverallValuesMetricID = "quality_of_experience_score"
	ListOverallValuesMetricIDPlaybackScore                 ListOverallValuesMetricID = "playback_score"
	ListOverallValuesMetricIDPlaybackFailurePercentage     ListOverallValuesMetricID = "playback_failure_percentage"
	ListOverallValuesMetricIDExitBeforeVideoStart          ListOverallValuesMetricID = "exit_before_video_start"
	ListOverallValuesMetricIDVideoStartupFailurePercentage ListOverallValuesMetricID = "video_startup_failure_percentage"
	ListOverallValuesMetricIDStartupScore                  ListOverallValuesMetricID = "startup_score"
	ListOverallValuesMetricIDVideoStartupTime              ListOverallValuesMetricID = "video_startup_time"
	ListOverallValuesMetricIDPlayerStartupTime             ListOverallValuesMetricID = "player_startup_time"
	ListOverallValuesMetricIDPageLoadTime                  ListOverallValuesMetricID = "page_load_time"
	ListOverallValuesMetricIDTotalStartupTime              ListOverallValuesMetricID = "total_startup_time"
	ListOverallValuesMetricIDLiveStreamLatency             ListOverallValuesMetricID = "live_stream_latency"
	ListOverallValuesMetricIDAverageBitrate                ListOverallValuesMetricID = "average_bitrate"
	ListOverallValuesMetricIDBufferCount                   ListOverallValuesMetricID = "buffer_count"
	ListOverallValuesMetricIDRenderQualityScore            ListOverallValuesMetricID = "render_quality_score"
	ListOverallValuesMetricIDAvgUpscaling                  ListOverallValuesMetricID = "avg_upscaling"
	ListOverallValuesMetricIDAvgDownscaling                ListOverallValuesMetricID = "avg_downscaling"
	ListOverallValuesMetricIDMaxUpscaling                  ListOverallValuesMetricID = "max_upscaling"
	ListOverallValuesMetricIDMaxDownscaling                ListOverallValuesMetricID = "max_downscaling"
	ListOverallValuesMetricIDJumpLatency                   ListOverallValuesMetricID = "jump_latency"
	ListOverallValuesMetricIDStabilityScore                ListOverallValuesMetricID = "stability_score"
	ListOverallValuesMetricIDBufferRatio                   ListOverallValuesMetricID = "buffer_ratio"
	ListOverallValuesMetricIDBufferFrequency               ListOverallValuesMetricID = "buffer_frequency"
	ListOverallValuesMetricIDBufferFill                    ListOverallValuesMetricID = "buffer_fill"
)

func (e ListOverallValuesMetricID) ToPointer() *ListOverallValuesMetricID {
	return &e
}
func (e *ListOverallValuesMetricID) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "views":
		fallthrough
	case "unique_viewers":
		fallthrough
	case "playing_time":
		fallthrough
	case "quality_of_experience_score":
		fallthrough
	case "playback_score":
		fallthrough
	case "playback_failure_percentage":
		fallthrough
	case "exit_before_video_start":
		fallthrough
	case "video_startup_failure_percentage":
		fallthrough
	case "startup_score":
		fallthrough
	case "video_startup_time":
		fallthrough
	case "player_startup_time":
		fallthrough
	case "page_load_time":
		fallthrough
	case "total_startup_time":
		fallthrough
	case "live_stream_latency":
		fallthrough
	case "average_bitrate":
		fallthrough
	case "buffer_count":
		fallthrough
	case "render_quality_score":
		fallthrough
	case "avg_upscaling":
		fallthrough
	case "avg_downscaling":
		fallthrough
	case "max_upscaling":
		fallthrough
	case "max_downscaling":
		fallthrough
	case "jump_latency":
		fallthrough
	case "stability_score":
		fallthrough
	case "buffer_ratio":
		fallthrough
	case "buffer_frequency":
		fallthrough
	case "buffer_fill":
		*e = ListOverallValuesMetricID(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListOverallValuesMetricID: %v", v)
	}
}

// ListOverallValuesTimespan - This parameter specifies the time span between which the video views list should be retrieved by. You can provide either from and to unix epoch timestamps or time duration. The scope of duration is between 60 minutes to 30 days.
type ListOverallValuesTimespan string

const (
	ListOverallValuesTimespanSixtyminutes    ListOverallValuesTimespan = "60:minutes"
	ListOverallValuesTimespanSixhours        ListOverallValuesTimespan = "6:hours"
	ListOverallValuesTimespanTwentyFourhours ListOverallValuesTimespan = "24:hours"
	ListOverallValuesTimespanThreedays       ListOverallValuesTimespan = "3:days"
	ListOverallValuesTimespanSevendays       ListOverallValuesTimespan = "7:days"
	ListOverallValuesTimespanThirtydays      ListOverallValuesTimespan = "30:days"
)

func (e ListOverallValuesTimespan) ToPointer() *ListOverallValuesTimespan {
	return &e
}
func (e *ListOverallValuesTimespan) UnmarshalJSON(data []byte) error {
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
		*e = ListOverallValuesTimespan(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListOverallValuesTimespan: %v", v)
	}
}

type ListOverallValuesRequest struct {
	// Pass metric Id
	//
	MetricID ListOverallValuesMetricID `pathParam:"style=simple,explode=false,name=metricId"`
	// The measurement for the given metrics.
	// Possible Values : [95th, median, avg, count or sum]
	//
	Measurement *string `default:"avg" queryParam:"style=form,explode=true,name=measurement"`
	// This parameter specifies the time span between which the video views list should be retrieved by. You can provide either from and to unix epoch timestamps or time duration. The scope of duration is between 60 minutes to 30 days.
	//
	Timespan ListOverallValuesTimespan `queryParam:"style=form,explode=true,name=timespan[]"`
	// Pass the dimensions and their corresponding values you want to filter the views by. For excluding the values in the filter we can pass '!' before the filter value. The list of filters can be obtained from list of dimensions endpoint.
	// Example Values : [ browser_name:Chrome , os_name:macOS , device_name:Galaxy ]
	//
	Filterby *string `queryParam:"style=form,explode=true,name=filterby[]"`
}

func (l ListOverallValuesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListOverallValuesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, []string{"metricId", "timespan[]"}); err != nil {
		return err
	}
	return nil
}

func (l *ListOverallValuesRequest) GetMetricID() ListOverallValuesMetricID {
	if l == nil {
		return ListOverallValuesMetricID("")
	}
	return l.MetricID
}

func (l *ListOverallValuesRequest) GetMeasurement() *string {
	if l == nil {
		return nil
	}
	return l.Measurement
}

func (l *ListOverallValuesRequest) GetTimespan() ListOverallValuesTimespan {
	if l == nil {
		return ListOverallValuesTimespan("")
	}
	return l.Timespan
}

func (l *ListOverallValuesRequest) GetFilterby() *string {
	if l == nil {
		return nil
	}
	return l.Filterby
}

// ListOverallValuesResponseBody - Get filter/ dimension value details by dimension name.
type ListOverallValuesResponseBody struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Metadata that has to be paased for metric calculations.
	MetaData *components.MetricsOverallMetaDataDetails `json:"metaData,omitempty"`
	// Retrieves overall values for a specified metric
	Data *components.MetricsOverallDataDetails `json:"data,omitempty"`
	// The timeframe from and to details displayed in the form of unix epoch timestamps.
	//
	Timespan []int64 `json:"timespan,omitempty"`
}

func (l *ListOverallValuesResponseBody) GetSuccess() *bool {
	if l == nil {
		return nil
	}
	return l.Success
}

func (l *ListOverallValuesResponseBody) GetMetaData() *components.MetricsOverallMetaDataDetails {
	if l == nil {
		return nil
	}
	return l.MetaData
}

func (l *ListOverallValuesResponseBody) GetData() *components.MetricsOverallDataDetails {
	if l == nil {
		return nil
	}
	return l.Data
}

func (l *ListOverallValuesResponseBody) GetTimespan() []int64 {
	if l == nil {
		return nil
	}
	return l.Timespan
}

type ListOverallValuesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Get filter/ dimension value details by dimension name.
	Object *ListOverallValuesResponseBody
}

func (l *ListOverallValuesResponse) GetHTTPMeta() components.HTTPMetadata {
	if l == nil {
		return components.HTTPMetadata{}
	}
	return l.HTTPMeta
}

func (l *ListOverallValuesResponse) GetObject() *ListOverallValuesResponseBody {
	if l == nil {
		return nil
	}
	return l.Object
}
