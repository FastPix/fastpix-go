

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/fastpix/fastpix-go/internal/utils"
	"github.com/fastpix/fastpix-go/models/components"
)

// ListBreakdownValuesMetricID - Pass metric Id
type ListBreakdownValuesMetricID string

const (
	ListBreakdownValuesMetricIDViews                         ListBreakdownValuesMetricID = "views"
	ListBreakdownValuesMetricIDUniqueViewers                 ListBreakdownValuesMetricID = "unique_viewers"
	ListBreakdownValuesMetricIDPlayingTime                   ListBreakdownValuesMetricID = "playing_time"
	ListBreakdownValuesMetricIDQualityOfExperienceScore      ListBreakdownValuesMetricID = "quality_of_experience_score"
	ListBreakdownValuesMetricIDPlaybackScore                 ListBreakdownValuesMetricID = "playback_score"
	ListBreakdownValuesMetricIDPlaybackFailurePercentage     ListBreakdownValuesMetricID = "playback_failure_percentage"
	ListBreakdownValuesMetricIDExitBeforeVideoStart          ListBreakdownValuesMetricID = "exit_before_video_start"
	ListBreakdownValuesMetricIDVideoStartupFailurePercentage ListBreakdownValuesMetricID = "video_startup_failure_percentage"
	ListBreakdownValuesMetricIDStartupScore                  ListBreakdownValuesMetricID = "startup_score"
	ListBreakdownValuesMetricIDVideoStartupTime              ListBreakdownValuesMetricID = "video_startup_time"
	ListBreakdownValuesMetricIDPlayerStartupTime             ListBreakdownValuesMetricID = "player_startup_time"
	ListBreakdownValuesMetricIDPageLoadTime                  ListBreakdownValuesMetricID = "page_load_time"
	ListBreakdownValuesMetricIDTotalStartupTime              ListBreakdownValuesMetricID = "total_startup_time"
	ListBreakdownValuesMetricIDLiveStreamLatency             ListBreakdownValuesMetricID = "live_stream_latency"
	ListBreakdownValuesMetricIDAverageBitrate                ListBreakdownValuesMetricID = "average_bitrate"
	ListBreakdownValuesMetricIDBufferCount                   ListBreakdownValuesMetricID = "buffer_count"
	ListBreakdownValuesMetricIDRenderQualityScore            ListBreakdownValuesMetricID = "render_quality_score"
	ListBreakdownValuesMetricIDAvgUpscaling                  ListBreakdownValuesMetricID = "avg_upscaling"
	ListBreakdownValuesMetricIDAvgDownscaling                ListBreakdownValuesMetricID = "avg_downscaling"
	ListBreakdownValuesMetricIDMaxUpscaling                  ListBreakdownValuesMetricID = "max_upscaling"
	ListBreakdownValuesMetricIDMaxDownscaling                ListBreakdownValuesMetricID = "max_downscaling"
	ListBreakdownValuesMetricIDJumpLatency                   ListBreakdownValuesMetricID = "jump_latency"
	ListBreakdownValuesMetricIDStabilityScore                ListBreakdownValuesMetricID = "stability_score"
	ListBreakdownValuesMetricIDBufferRatio                   ListBreakdownValuesMetricID = "buffer_ratio"
	ListBreakdownValuesMetricIDBufferFrequency               ListBreakdownValuesMetricID = "buffer_frequency"
	ListBreakdownValuesMetricIDBufferFill                    ListBreakdownValuesMetricID = "buffer_fill"
)

func (e ListBreakdownValuesMetricID) ToPointer() *ListBreakdownValuesMetricID {
	return &e
}
func (e *ListBreakdownValuesMetricID) UnmarshalJSON(data []byte) error {
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
		*e = ListBreakdownValuesMetricID(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListBreakdownValuesMetricID: %v", v)
	}
}

// ListBreakdownValuesTimespan - This parameter specifies the time span between which the video views list should be retrieved by. You can provide either from and to unix epoch timestamps or time duration. The scope of duration is between 60 minutes to 30 days.
type ListBreakdownValuesTimespan string

const (
	ListBreakdownValuesTimespanSixtyminutes    ListBreakdownValuesTimespan = "60:minutes"
	ListBreakdownValuesTimespanSixhours        ListBreakdownValuesTimespan = "6:hours"
	ListBreakdownValuesTimespanTwentyFourhours ListBreakdownValuesTimespan = "24:hours"
	ListBreakdownValuesTimespanThreedays       ListBreakdownValuesTimespan = "3:days"
	ListBreakdownValuesTimespanSevendays       ListBreakdownValuesTimespan = "7:days"
	ListBreakdownValuesTimespanThirtydays      ListBreakdownValuesTimespan = "30:days"
)

func (e ListBreakdownValuesTimespan) ToPointer() *ListBreakdownValuesTimespan {
	return &e
}
func (e *ListBreakdownValuesTimespan) UnmarshalJSON(data []byte) error {
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
		*e = ListBreakdownValuesTimespan(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListBreakdownValuesTimespan: %v", v)
	}
}

// ListBreakdownValuesSortOrder - The order direction to sort the metrics list by.
type ListBreakdownValuesSortOrder string

const (
	ListBreakdownValuesSortOrderAsc  ListBreakdownValuesSortOrder = "asc"
	ListBreakdownValuesSortOrderDesc ListBreakdownValuesSortOrder = "desc"
)

func (e ListBreakdownValuesSortOrder) ToPointer() *ListBreakdownValuesSortOrder {
	return &e
}
func (e *ListBreakdownValuesSortOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = ListBreakdownValuesSortOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListBreakdownValuesSortOrder: %v", v)
	}
}

type ListBreakdownValuesRequest struct {
	// Pass metric Id
	//
	MetricID ListBreakdownValuesMetricID `pathParam:"style=simple,explode=false,name=metricId"`
	// This parameter specifies the time span between which the video views list should be retrieved by. You can provide either from and to unix epoch timestamps or time duration. The scope of duration is between 60 minutes to 30 days.
	//
	Timespan ListBreakdownValuesTimespan `queryParam:"style=form,explode=true,name=timespan[]"`
	// Pass the dimensions and their corresponding values you want to filter the views by. For excluding the values in the filter we can pass '!' before the filter value. The list of filters can be obtained from list of dimensions endpoint.
	// Example Values : [ browser_name:Chrome , os_name:macOS , device_name:Galaxy ]
	//
	Filterby *string `queryParam:"style=form,explode=true,name=filterby[]"`
	// Pass the limit to display only the rows specified by the value.
	//
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	// Pass the offset value to indicate the page number.
	//
	Offset *int64 `default:"1" queryParam:"style=form,explode=true,name=offset"`
	// Pass this value to group the metrics list by.
	// Possible Values : ["browser_name", "browser_version", "os_name","os_version" , "device_name", "device_model", "device_type", "device_manufacturer", "player_remote_played",player_name", "player_version", "player_software_name", "player_software_version", "player_resolution", "fp_sdk","fp_sdk_version", "player_autoplay_on", "player_preload_on","video_title",  "video_id", "video_series" ,  "fp_playback_id","fp_live_stream_id", "media_id","video_source_stream_type", "video_source_type", "video_encoding_variant", "experiment_name", "sub_property_id", "drm_type","asn_name", "cdn", "video_source_hostname", "connection_type", "view_session_id","continent","country", "region","viewer_id", "error_code", "exit_before_video_start", "view_has_ad", "video_startup_failed" , "page_context", "playback_failed".]
	//
	GroupBy *string `queryParam:"style=form,explode=true,name=groupBy"`
	// Pass this value to order the metrics list by.
	//
	OrderBy *string `default:"views" queryParam:"style=form,explode=true,name=orderBy"`
	// The order direction to sort the metrics list by.
	//
	SortOrder *ListBreakdownValuesSortOrder `default:"asc" queryParam:"style=form,explode=true,name=sortOrder"`
	// The measurement for the given metrics.
	//
	Measurement *string `default:"avg" queryParam:"style=form,explode=true,name=measurement"`
}

func (l ListBreakdownValuesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListBreakdownValuesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, []string{"metricId", "timespan[]"}); err != nil {
		return err
	}
	return nil
}

func (l *ListBreakdownValuesRequest) GetMetricID() ListBreakdownValuesMetricID {
	if l == nil {
		return ListBreakdownValuesMetricID("")
	}
	return l.MetricID
}

func (l *ListBreakdownValuesRequest) GetTimespan() ListBreakdownValuesTimespan {
	if l == nil {
		return ListBreakdownValuesTimespan("")
	}
	return l.Timespan
}

func (l *ListBreakdownValuesRequest) GetFilterby() *string {
	if l == nil {
		return nil
	}
	return l.Filterby
}

func (l *ListBreakdownValuesRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListBreakdownValuesRequest) GetOffset() *int64 {
	if l == nil {
		return nil
	}
	return l.Offset
}

func (l *ListBreakdownValuesRequest) GetGroupBy() *string {
	if l == nil {
		return nil
	}
	return l.GroupBy
}

func (l *ListBreakdownValuesRequest) GetOrderBy() *string {
	if l == nil {
		return nil
	}
	return l.OrderBy
}

func (l *ListBreakdownValuesRequest) GetSortOrder() *ListBreakdownValuesSortOrder {
	if l == nil {
		return nil
	}
	return l.SortOrder
}

func (l *ListBreakdownValuesRequest) GetMeasurement() *string {
	if l == nil {
		return nil
	}
	return l.Measurement
}

// ListBreakdownValuesResponseBody - Get filter/ dimension value details by dimension name.
type ListBreakdownValuesResponseBody struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Retrieves breakdown values for a specified metric and timespan
	MetaData *components.MetricsTimeseriesMetaDataDetails `json:"metaData,omitempty"`
	// Retrieves breakdown values for a specified metric and timespan
	Data []components.MetricsBreakdownDetails `json:"data,omitempty"`
	// The timeframe from and to details displayed in the form of unix epoch timestamps.
	//
	Timespan []int64 `json:"timespan,omitempty"`
}

func (l *ListBreakdownValuesResponseBody) GetSuccess() *bool {
	if l == nil {
		return nil
	}
	return l.Success
}

func (l *ListBreakdownValuesResponseBody) GetMetaData() *components.MetricsTimeseriesMetaDataDetails {
	if l == nil {
		return nil
	}
	return l.MetaData
}

func (l *ListBreakdownValuesResponseBody) GetData() []components.MetricsBreakdownDetails {
	if l == nil {
		return nil
	}
	return l.Data
}

func (l *ListBreakdownValuesResponseBody) GetTimespan() []int64 {
	if l == nil {
		return nil
	}
	return l.Timespan
}

type ListBreakdownValuesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Get filter/ dimension value details by dimension name.
	Object *ListBreakdownValuesResponseBody
}

func (l *ListBreakdownValuesResponse) GetHTTPMeta() components.HTTPMetadata {
	if l == nil {
		return components.HTTPMetadata{}
	}
	return l.HTTPMeta
}

func (l *ListBreakdownValuesResponse) GetObject() *ListBreakdownValuesResponseBody {
	if l == nil {
		return nil
	}
	return l.Object
}
