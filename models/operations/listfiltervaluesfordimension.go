package operations

import (
	"encoding/json"
	"fmt"
	"github.com/FastPix/fastpix-go/models/components"
)

// DimensionsID - Pass Dimensions id
type DimensionsID string

const (
	DimensionsIDBrowserName           DimensionsID = "browser_name"
	DimensionsIDBrowserVersion        DimensionsID = "browser_version"
	DimensionsIDOsName                DimensionsID = "os_name"
	DimensionsIDOsVersion             DimensionsID = "os_version"
	DimensionsIDDeviceName            DimensionsID = "device_name"
	DimensionsIDDeviceModel           DimensionsID = "device_model"
	DimensionsIDDeviceType            DimensionsID = "device_type"
	DimensionsIDDeviceManufacturer    DimensionsID = "device_manufacturer"
	DimensionsIDPlayerRemotePlayed    DimensionsID = "player_remote_played"
	DimensionsIDPlayerName            DimensionsID = "player_name"
	DimensionsIDPlayerVersion         DimensionsID = "player_version"
	DimensionsIDPlayerSoftwareName    DimensionsID = "player_software_name"
	DimensionsIDPlayerSoftwareVersion DimensionsID = "player_software_version"
	DimensionsIDPlayerResolution      DimensionsID = "player_resolution"
	DimensionsIDFpSDK                 DimensionsID = "fp_sdk"
	DimensionsIDFpSDKVersion          DimensionsID = "fp_sdk_version"
	DimensionsIDPlayerAutoplayOn      DimensionsID = "player_autoplay_on"
	DimensionsIDPlayerPreloadOn       DimensionsID = "player_preload_on"
	DimensionsIDVideoTitle            DimensionsID = "video_title"
	DimensionsIDVideoID               DimensionsID = "video_id"
	DimensionsIDVideoSeries           DimensionsID = "video_series"
	DimensionsIDFpPlaybackID          DimensionsID = "fp_playback_id"
	DimensionsIDFpLiveStreamID        DimensionsID = "fp_live_stream_id"
	DimensionsIDMediaID               DimensionsID = "media_id"
	DimensionsIDVideoSourceStreamType DimensionsID = "video_source_stream_type"
	DimensionsIDVideoSourceType       DimensionsID = "video_source_type"
	DimensionsIDVideoEncodingVariant  DimensionsID = "video_encoding_variant"
	DimensionsIDExperimentName        DimensionsID = "experiment_name"
	DimensionsIDSubPropertyID         DimensionsID = "sub_property_id"
	DimensionsIDDrmType               DimensionsID = "drm_type"
	DimensionsIDAsnName               DimensionsID = "asn_name"
	DimensionsIDCdn                   DimensionsID = "cdn"
	DimensionsIDVideoSourceHostname   DimensionsID = "video_source_hostname"
	DimensionsIDConnectionType        DimensionsID = "connection_type"
	DimensionsIDViewSessionID         DimensionsID = "view_session_id"
	DimensionsIDContinent             DimensionsID = "continent"
	DimensionsIDCountry               DimensionsID = "country"
	DimensionsIDRegion                DimensionsID = "region"
	DimensionsIDViewerID              DimensionsID = "viewer_id"
	DimensionsIDErrorCode             DimensionsID = "error_code"
	DimensionsIDExitBeforeVideoStart  DimensionsID = "exit_before_video_start"
	DimensionsIDViewHasAd             DimensionsID = "view_has_ad"
	DimensionsIDVideoStartupFailed    DimensionsID = "video_startup_failed"
	DimensionsIDPageContext           DimensionsID = "page_context"
	DimensionsIDPlaybackFailed        DimensionsID = "playback_failed"
	DimensionsIDCustom1               DimensionsID = "custom_1"
	DimensionsIDCustom2               DimensionsID = "custom_2"
	DimensionsIDCustom3               DimensionsID = "custom_3"
	DimensionsIDCustom4               DimensionsID = "custom_4"
	DimensionsIDCustom5               DimensionsID = "custom_5"
	DimensionsIDCustom6               DimensionsID = "custom_6"
	DimensionsIDCustom7               DimensionsID = "custom_7"
	DimensionsIDCustom8               DimensionsID = "custom_8"
	DimensionsIDCustom9               DimensionsID = "custom_9"
	DimensionsIDCustom10              DimensionsID = "custom_10"
)

func (e DimensionsID) ToPointer() *DimensionsID {
	return &e
}
func (e *DimensionsID) UnmarshalJSON(data []byte) error {
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
		*e = DimensionsID(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DimensionsID: %v", v)
	}
}

// ListFilterValuesForDimensionTimespan - This parameter specifies the time span between which the video views list should be retrieved by. You can provide either from and to unix epoch timestamps or time duration. The scope of duration is between 60 minutes to 30 days.
type ListFilterValuesForDimensionTimespan string

const (
	ListFilterValuesForDimensionTimespanSixtyminutes    ListFilterValuesForDimensionTimespan = "60:minutes"
	ListFilterValuesForDimensionTimespanSixhours        ListFilterValuesForDimensionTimespan = "6:hours"
	ListFilterValuesForDimensionTimespanTwentyFourhours ListFilterValuesForDimensionTimespan = "24:hours"
	ListFilterValuesForDimensionTimespanThreedays       ListFilterValuesForDimensionTimespan = "3:days"
	ListFilterValuesForDimensionTimespanSevendays       ListFilterValuesForDimensionTimespan = "7:days"
	ListFilterValuesForDimensionTimespanThirtydays      ListFilterValuesForDimensionTimespan = "30:days"
)

func (e ListFilterValuesForDimensionTimespan) ToPointer() *ListFilterValuesForDimensionTimespan {
	return &e
}
func (e *ListFilterValuesForDimensionTimespan) UnmarshalJSON(data []byte) error {
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
		*e = ListFilterValuesForDimensionTimespan(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListFilterValuesForDimensionTimespan: %v", v)
	}
}

type ListFilterValuesForDimensionRequest struct {
	// Pass Dimensions id
	//
	DimensionsID DimensionsID `pathParam:"style=simple,explode=false,name=dimensionsId"`
	// This parameter specifies the time span between which the video views list should be retrieved by. You can provide either from and to unix epoch timestamps or time duration. The scope of duration is between 60 minutes to 30 days.
	//
	Timespan ListFilterValuesForDimensionTimespan `queryParam:"style=form,explode=true,name=timespan[]"`
	// Pass the dimensions and their corresponding values you want to filter the views by. For excluding the values in the filter we can pass '!' before the filter value. The list of filters can be obtained from list of dimensions endpoint.
	// Example Values : [ browser_name:Chrome , os_name:macOS , device_name:Galaxy ]
	//
	Filterby *string `queryParam:"style=form,explode=true,name=filterby[]"`
}

func (l *ListFilterValuesForDimensionRequest) GetDimensionsID() DimensionsID {
	if l == nil {
		return DimensionsID("")
	}
	return l.DimensionsID
}

func (l *ListFilterValuesForDimensionRequest) GetTimespan() ListFilterValuesForDimensionTimespan {
	if l == nil {
		return ListFilterValuesForDimensionTimespan("")
	}
	return l.Timespan
}

func (l *ListFilterValuesForDimensionRequest) GetFilterby() *string {
	if l == nil {
		return nil
	}
	return l.Filterby
}

// ListFilterValuesForDimensionResponseBody - Get filter / dimension value details by dimension name.
type ListFilterValuesForDimensionResponseBody struct {
	// It demonstrates whether the request is successful or not.
	Success *bool `json:"success,omitempty"`
	// filter values associated with a specific dimension
	Data []components.BrowserNameDimensiondetails `json:"data,omitempty"`
	// The timeframe from and to details displayed in the form of unix epoch timestamps.
	//
	Timespan []int64 `json:"timespan,omitempty"`
}

func (l *ListFilterValuesForDimensionResponseBody) GetSuccess() *bool {
	if l == nil {
		return nil
	}
	return l.Success
}

func (l *ListFilterValuesForDimensionResponseBody) GetData() []components.BrowserNameDimensiondetails {
	if l == nil {
		return nil
	}
	return l.Data
}

func (l *ListFilterValuesForDimensionResponseBody) GetTimespan() []int64 {
	if l == nil {
		return nil
	}
	return l.Timespan
}

type ListFilterValuesForDimensionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Get filter / dimension value details by dimension name.
	Object *ListFilterValuesForDimensionResponseBody
}

func (l *ListFilterValuesForDimensionResponse) GetHTTPMeta() components.HTTPMetadata {
	if l == nil {
		return components.HTTPMetadata{}
	}
	return l.HTTPMeta
}

func (l *ListFilterValuesForDimensionResponse) GetObject() *ListFilterValuesForDimensionResponseBody {
	if l == nil {
		return nil
	}
	return l.Object
}
