package operations

import (
	"encoding/json"
	"fmt"
	"github.com/FastPix/fastpix-go/models/components"
)

// ListComparisonValuesTimespan - This parameter specifies the time span between which the video views list should be retrieved by. You can provide either from and to unix epoch timestamps or time duration. The scope of duration is between 60 minutes to 30 days.
type ListComparisonValuesTimespan string

const (
	ListComparisonValuesTimespanSixtyminutes    ListComparisonValuesTimespan = "60:minutes"
	ListComparisonValuesTimespanSixhours        ListComparisonValuesTimespan = "6:hours"
	ListComparisonValuesTimespanTwentyFourhours ListComparisonValuesTimespan = "24:hours"
	ListComparisonValuesTimespanThreedays       ListComparisonValuesTimespan = "3:days"
	ListComparisonValuesTimespanSevendays       ListComparisonValuesTimespan = "7:days"
	ListComparisonValuesTimespanThirtydays      ListComparisonValuesTimespan = "30:days"
)

func (e ListComparisonValuesTimespan) ToPointer() *ListComparisonValuesTimespan {
	return &e
}
func (e *ListComparisonValuesTimespan) UnmarshalJSON(data []byte) error {
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
		*e = ListComparisonValuesTimespan(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListComparisonValuesTimespan: %v", v)
	}
}

// ListComparisonValuesDimension - The dimension id in which the views are watched.
type ListComparisonValuesDimension string

const (
	ListComparisonValuesDimensionBrowserName           ListComparisonValuesDimension = "browser_name"
	ListComparisonValuesDimensionBrowserVersion        ListComparisonValuesDimension = "browser_version"
	ListComparisonValuesDimensionOsName                ListComparisonValuesDimension = "os_name"
	ListComparisonValuesDimensionOsVersion             ListComparisonValuesDimension = "os_version"
	ListComparisonValuesDimensionDeviceName            ListComparisonValuesDimension = "device_name"
	ListComparisonValuesDimensionDeviceModel           ListComparisonValuesDimension = "device_model"
	ListComparisonValuesDimensionDeviceType            ListComparisonValuesDimension = "device_type"
	ListComparisonValuesDimensionDeviceManufacturer    ListComparisonValuesDimension = "device_manufacturer"
	ListComparisonValuesDimensionPlayerRemotePlayed    ListComparisonValuesDimension = "player_remote_played"
	ListComparisonValuesDimensionPlayerName            ListComparisonValuesDimension = "player_name"
	ListComparisonValuesDimensionPlayerVersion         ListComparisonValuesDimension = "player_version"
	ListComparisonValuesDimensionPlayerSoftwareName    ListComparisonValuesDimension = "player_software_name"
	ListComparisonValuesDimensionPlayerSoftwareVersion ListComparisonValuesDimension = "player_software_version"
	ListComparisonValuesDimensionPlayerResolution      ListComparisonValuesDimension = "player_resolution"
	ListComparisonValuesDimensionFpSDK                 ListComparisonValuesDimension = "fp_sdk"
	ListComparisonValuesDimensionFpSDKVersion          ListComparisonValuesDimension = "fp_sdk_version"
	ListComparisonValuesDimensionPlayerAutoplayOn      ListComparisonValuesDimension = "player_autoplay_on"
	ListComparisonValuesDimensionPlayerPreloadOn       ListComparisonValuesDimension = "player_preload_on"
	ListComparisonValuesDimensionVideoTitle            ListComparisonValuesDimension = "video_title"
	ListComparisonValuesDimensionVideoID               ListComparisonValuesDimension = "video_id"
	ListComparisonValuesDimensionVideoSeries           ListComparisonValuesDimension = "video_series"
	ListComparisonValuesDimensionFpPlaybackID          ListComparisonValuesDimension = "fp_playback_id"
	ListComparisonValuesDimensionFpLiveStreamID        ListComparisonValuesDimension = "fp_live_stream_id"
	ListComparisonValuesDimensionMediaID               ListComparisonValuesDimension = "media_id"
	ListComparisonValuesDimensionVideoSourceStreamType ListComparisonValuesDimension = "video_source_stream_type"
	ListComparisonValuesDimensionVideoSourceType       ListComparisonValuesDimension = "video_source_type"
	ListComparisonValuesDimensionVideoEncodingVariant  ListComparisonValuesDimension = "video_encoding_variant"
	ListComparisonValuesDimensionExperimentName        ListComparisonValuesDimension = "experiment_name"
	ListComparisonValuesDimensionSubPropertyID         ListComparisonValuesDimension = "sub_property_id"
	ListComparisonValuesDimensionDrmType               ListComparisonValuesDimension = "drm_type"
	ListComparisonValuesDimensionAsnName               ListComparisonValuesDimension = "asn_name"
	ListComparisonValuesDimensionCdn                   ListComparisonValuesDimension = "cdn"
	ListComparisonValuesDimensionVideoSourceHostname   ListComparisonValuesDimension = "video_source_hostname"
	ListComparisonValuesDimensionConnectionType        ListComparisonValuesDimension = "connection_type"
	ListComparisonValuesDimensionViewSessionID         ListComparisonValuesDimension = "view_session_id"
	ListComparisonValuesDimensionContinent             ListComparisonValuesDimension = "continent"
	ListComparisonValuesDimensionCountry               ListComparisonValuesDimension = "country"
	ListComparisonValuesDimensionRegion                ListComparisonValuesDimension = "region"
	ListComparisonValuesDimensionViewerID              ListComparisonValuesDimension = "viewer_id"
	ListComparisonValuesDimensionErrorCode             ListComparisonValuesDimension = "error_code"
	ListComparisonValuesDimensionExitBeforeVideoStart  ListComparisonValuesDimension = "exit_before_video_start"
	ListComparisonValuesDimensionViewHasAd             ListComparisonValuesDimension = "view_has_ad"
	ListComparisonValuesDimensionVideoStartupFailed    ListComparisonValuesDimension = "video_startup_failed"
	ListComparisonValuesDimensionPageContext           ListComparisonValuesDimension = "page_context"
	ListComparisonValuesDimensionPlaybackFailed        ListComparisonValuesDimension = "playback_failed"
	ListComparisonValuesDimensionCustom1               ListComparisonValuesDimension = "custom_1"
	ListComparisonValuesDimensionCustom2               ListComparisonValuesDimension = "custom_2"
	ListComparisonValuesDimensionCustom3               ListComparisonValuesDimension = "custom_3"
	ListComparisonValuesDimensionCustom4               ListComparisonValuesDimension = "custom_4"
	ListComparisonValuesDimensionCustom5               ListComparisonValuesDimension = "custom_5"
	ListComparisonValuesDimensionCustom6               ListComparisonValuesDimension = "custom_6"
	ListComparisonValuesDimensionCustom7               ListComparisonValuesDimension = "custom_7"
	ListComparisonValuesDimensionCustom8               ListComparisonValuesDimension = "custom_8"
	ListComparisonValuesDimensionCustom9               ListComparisonValuesDimension = "custom_9"
	ListComparisonValuesDimensionCustom10              ListComparisonValuesDimension = "custom_10"
)

func (e ListComparisonValuesDimension) ToPointer() *ListComparisonValuesDimension {
	return &e
}
func (e *ListComparisonValuesDimension) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "browser_name":
		fallthrough
	case "browser_version":
		fallthrough
	case "os_name":
		fallthrough
	case "os_version":
		fallthrough
	case "device_name":
		fallthrough
	case "device_model":
		fallthrough
	case "device_type":
		fallthrough
	case "device_manufacturer":
		fallthrough
	case "player_remote_played":
		fallthrough
	case "player_name":
		fallthrough
	case "player_version":
		fallthrough
	case "player_software_name":
		fallthrough
	case "player_software_version":
		fallthrough
	case "player_resolution":
		fallthrough
	case "fp_sdk":
		fallthrough
	case "fp_sdk_version":
		fallthrough
	case "player_autoplay_on":
		fallthrough
	case "player_preload_on":
		fallthrough
	case "video_title":
		fallthrough
	case "video_id":
		fallthrough
	case "video_series":
		fallthrough
	case "fp_playback_id":
		fallthrough
	case "fp_live_stream_id":
		fallthrough
	case "media_id":
		fallthrough
	case "video_source_stream_type":
		fallthrough
	case "video_source_type":
		fallthrough
	case "video_encoding_variant":
		fallthrough
	case "experiment_name":
		fallthrough
	case "sub_property_id":
		fallthrough
	case "drm_type":
		fallthrough
	case "asn_name":
		fallthrough
	case "cdn":
		fallthrough
	case "video_source_hostname":
		fallthrough
	case "connection_type":
		fallthrough
	case "view_session_id":
		fallthrough
	case "continent":
		fallthrough
	case "country":
		fallthrough
	case "region":
		fallthrough
	case "viewer_id":
		fallthrough
	case "error_code":
		fallthrough
	case "exit_before_video_start":
		fallthrough
	case "view_has_ad":
		fallthrough
	case "video_startup_failed":
		fallthrough
	case "page_context":
		fallthrough
	case "playback_failed":
		fallthrough
	case "custom_1":
		fallthrough
	case "custom_2":
		fallthrough
	case "custom_3":
		fallthrough
	case "custom_4":
		fallthrough
	case "custom_5":
		fallthrough
	case "custom_6":
		fallthrough
	case "custom_7":
		fallthrough
	case "custom_8":
		fallthrough
	case "custom_9":
		fallthrough
	case "custom_10":
		*e = ListComparisonValuesDimension(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListComparisonValuesDimension: %v", v)
	}
}

type ListComparisonValuesRequest struct {
	// This parameter specifies the time span between which the video views list should be retrieved by. You can provide either from and to unix epoch timestamps or time duration. The scope of duration is between 60 minutes to 30 days.
	//
	Timespan ListComparisonValuesTimespan `queryParam:"style=form,explode=true,name=timespan[]"`
	// Pass the dimensions and their corresponding values you want to filter the views by. For excluding the values in the filter we can pass '!' before the filter value. The list of filters can be obtained from list of dimensions endpoint.
	// Example Values : [ browser_name:Chrome , os_name:macOS , device_name:Galaxy ]
	//
	Filterby *string `queryParam:"style=form,explode=true,name=filterby[]"`
	// The dimension id in which the views are watched.
	//
	Dimension *ListComparisonValuesDimension `queryParam:"style=form,explode=true,name=dimension"`
	// The value for the selected dimension.
	// For example:
	//  If `dimension` is `browser_name`, the value could be  `Chrome` `,` `Firefox` `etc` .
	//  If `dimension` is `os_name`, the value could be `macOS` `,` `Windows` `etc` .
	//
	Value *string `queryParam:"style=form,explode=true,name=value"`
}

func (l *ListComparisonValuesRequest) GetTimespan() ListComparisonValuesTimespan {
	if l == nil {
		return ListComparisonValuesTimespan("")
	}
	return l.Timespan
}

func (l *ListComparisonValuesRequest) GetFilterby() *string {
	if l == nil {
		return nil
	}
	return l.Filterby
}

func (l *ListComparisonValuesRequest) GetDimension() *ListComparisonValuesDimension {
	if l == nil {
		return nil
	}
	return l.Dimension
}

func (l *ListComparisonValuesRequest) GetValue() *string {
	if l == nil {
		return nil
	}
	return l.Value
}

// ListComparisonValuesResponseBody - Get filter/ dimension value details by dimension name.
type ListComparisonValuesResponseBody struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// Displays the result of the request.
	//
	Data [][]components.MetricsComparisonDetails `json:"data,omitempty"`
	// The timeframe from and to details displayed in the form of unix epoch timestamps.
	//
	Timespan []int64 `json:"timespan,omitempty"`
}

func (l *ListComparisonValuesResponseBody) GetSuccess() *bool {
	if l == nil {
		return nil
	}
	return l.Success
}

func (l *ListComparisonValuesResponseBody) GetData() [][]components.MetricsComparisonDetails {
	if l == nil {
		return nil
	}
	return l.Data
}

func (l *ListComparisonValuesResponseBody) GetTimespan() []int64 {
	if l == nil {
		return nil
	}
	return l.Timespan
}

type ListComparisonValuesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Get filter/ dimension value details by dimension name.
	Object *ListComparisonValuesResponseBody
}

func (l *ListComparisonValuesResponse) GetHTTPMeta() components.HTTPMetadata {
	if l == nil {
		return components.HTTPMetadata{}
	}
	return l.HTTPMeta
}

func (l *ListComparisonValuesResponse) GetObject() *ListComparisonValuesResponseBody {
	if l == nil {
		return nil
	}
	return l.Object
}
