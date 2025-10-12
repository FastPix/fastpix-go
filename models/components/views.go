

package components

import (
	"errors"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
	"github.com/FastPix/fastpix-go/optionalnullable"
)

type EventTimeType string

const (
	EventTimeTypeStr     EventTimeType = "str"
	EventTimeTypeInteger EventTimeType = "integer"
)

// EventTime - The unix epoch timestamp when the event was captured.
type EventTime struct {
	Str     *string `queryParam:"inline,name=event_time"`
	Integer *int64  `queryParam:"inline,name=event_time"`

	Type EventTimeType
}

func CreateEventTimeStr(str string) EventTime {
	typ := EventTimeTypeStr

	return EventTime{
		Str:  &str,
		Type: typ,
	}
}

func CreateEventTimeInteger(integer int64) EventTime {
	typ := EventTimeTypeInteger

	return EventTime{
		Integer: &integer,
		Type:    typ,
	}
}

func (u *EventTime) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, nil); err == nil {
		u.Str = &str
		u.Type = EventTimeTypeStr
		return nil
	}

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, nil); err == nil {
		u.Integer = &integer
		u.Type = EventTimeTypeInteger
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for EventTime", string(data))
}

func (u EventTime) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	return nil, errors.New("could not marshal union type EventTime: all fields are null")
}

type Details struct {
	// The player_source_bitrate represents the bitrate of the video stream that is being played, measured in bits per second (bps). This value indicates the quality of the video being streamed, with higher bitrates typically corresponding to better video quality but requiring more bandwidth.
	//
	PlayerSourceBitrate optionalnullable.OptionalNullable[int64] `json:"player_source_bitrate,omitempty"`
	// The player_source_codec represents the video or audio codec being used to decode and play the media. A codec is a technology used to compress and decompress digital media files, enabling efficient transmission and storage while maintaining quality.
	//
	PlayerSourceCodec optionalnullable.OptionalNullable[string] `json:"player_source_codec,omitempty"`
	// The player_source_height refers to the vertical resolution of the video being played, measured in pixels. This value represents the height dimension of the video frame and is part of the overall resolution of the video (e.g., 1920x1080, where the height is 1080 pixels).
	//
	PlayerSourceHeight optionalnullable.OptionalNullable[int64] `json:"playerSourceHeight,omitempty"`
	// The player_source_width refers to the horizontal resolution of the video being played, measured in pixels. This value represents the width dimension of the video frame and is part of the overall video resolution (e.g., 1920x1080, where the width is 1920 pixels).
	//
	PlayerSourceWidth optionalnullable.OptionalNullable[int64] `json:"playerSourceWidth,omitempty"`
}

func (d *Details) GetPlayerSourceBitrate() optionalnullable.OptionalNullable[int64] {
	if d == nil {
		return nil
	}
	return d.PlayerSourceBitrate
}

func (d *Details) GetPlayerSourceCodec() optionalnullable.OptionalNullable[string] {
	if d == nil {
		return nil
	}
	return d.PlayerSourceCodec
}

func (d *Details) GetPlayerSourceHeight() optionalnullable.OptionalNullable[int64] {
	if d == nil {
		return nil
	}
	return d.PlayerSourceHeight
}

func (d *Details) GetPlayerSourceWidth() optionalnullable.OptionalNullable[int64] {
	if d == nil {
		return nil
	}
	return d.PlayerSourceWidth
}

type Event struct {
	// Name of the event.
	//
	EventName optionalnullable.OptionalNullable[string] `json:"event_name,omitempty"`
	// The unix epoch timestamp when the event was captured.
	//
	EventTime optionalnullable.OptionalNullable[EventTime] `json:"event_time,omitempty"`
	// The unix epoch timestamp which represents the actual time the event has occured.
	//
	ViewerTime optionalnullable.OptionalNullable[int64] `json:"viewer_time,omitempty"`
	// The player_playhead_time represents the current position of the playhead (the point in the video that is being watched) on the video seekbar, measured in milliseconds. This value indicates how far into the video playback has progressed at any given moment.
	//
	PlayerPlayheadTime optionalnullable.OptionalNullable[int64] `json:"player_playhead_time,omitempty"`
	Details            *Details                                 `json:"details,omitempty"`
}

func (e *Event) GetEventName() optionalnullable.OptionalNullable[string] {
	if e == nil {
		return nil
	}
	return e.EventName
}

func (e *Event) GetEventTime() optionalnullable.OptionalNullable[EventTime] {
	if e == nil {
		return nil
	}
	return e.EventTime
}

func (e *Event) GetViewerTime() optionalnullable.OptionalNullable[int64] {
	if e == nil {
		return nil
	}
	return e.ViewerTime
}

func (e *Event) GetPlayerPlayheadTime() optionalnullable.OptionalNullable[int64] {
	if e == nil {
		return nil
	}
	return e.PlayerPlayheadTime
}

func (e *Event) GetDetails() *Details {
	if e == nil {
		return nil
	}
	return e.Details
}

type PlayerHeightType string

const (
	PlayerHeightTypeStr    PlayerHeightType = "str"
	PlayerHeightTypeNumber PlayerHeightType = "number"
)

// PlayerHeight - Player Height refers to the vertical dimension, measured in pixels, of the video player as it appears on the webpage.
type PlayerHeight struct {
	Str    *string  `queryParam:"inline,name=playerHeight"`
	Number *float64 `queryParam:"inline,name=playerHeight"`

	Type PlayerHeightType
}

func CreatePlayerHeightStr(str string) PlayerHeight {
	typ := PlayerHeightTypeStr

	return PlayerHeight{
		Str:  &str,
		Type: typ,
	}
}

func CreatePlayerHeightNumber(number float64) PlayerHeight {
	typ := PlayerHeightTypeNumber

	return PlayerHeight{
		Number: &number,
		Type:   typ,
	}
}

func (u *PlayerHeight) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, nil); err == nil {
		u.Str = &str
		u.Type = PlayerHeightTypeStr
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, nil); err == nil {
		u.Number = &number
		u.Type = PlayerHeightTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for PlayerHeight", string(data))
}

func (u PlayerHeight) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type PlayerHeight: all fields are null")
}

type PlayerWidthType string

const (
	PlayerWidthTypeStr    PlayerWidthType = "str"
	PlayerWidthTypeNumber PlayerWidthType = "number"
)

// PlayerWidth - Player Width refers to the width of the player displayed within the webpage, measured in pixels.
type PlayerWidth struct {
	Str    *string  `queryParam:"inline,name=playerWidth"`
	Number *float64 `queryParam:"inline,name=playerWidth"`

	Type PlayerWidthType
}

func CreatePlayerWidthStr(str string) PlayerWidth {
	typ := PlayerWidthTypeStr

	return PlayerWidth{
		Str:  &str,
		Type: typ,
	}
}

func CreatePlayerWidthNumber(number float64) PlayerWidth {
	typ := PlayerWidthTypeNumber

	return PlayerWidth{
		Number: &number,
		Type:   typ,
	}
}

func (u *PlayerWidth) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, nil); err == nil {
		u.Str = &str
		u.Type = PlayerWidthTypeStr
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, nil); err == nil {
		u.Number = &number
		u.Type = PlayerWidthTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for PlayerWidth", string(data))
}

func (u PlayerWidth) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type PlayerWidth: all fields are null")
}

// Views - Displays the result of the request.
type Views struct {
	// The Name associated with the asnId.
	//
	AsnName optionalnullable.OptionalNullable[string] `json:"asnName,omitempty"`
	// The unique identifier assigned to an Autonomous System (AS) on the Internet. The ASN is used to identify and exchange routing information between different networks.
	//
	AsnID optionalnullable.OptionalNullable[int64] `json:"asnId,omitempty"`
	// The media Id value if the video asset is internal to FastPix.
	//
	MediaID optionalnullable.OptionalNullable[string] `json:"mediaId,omitempty"`
	// Buffer Count represents the number of rebuffering events occurring during the video view.
	//
	BufferCount optionalnullable.OptionalNullable[int64] `json:"bufferCount,omitempty"`
	// Buffer Fill indicates the total time, in milliseconds, that viewers wait for rebuffering per video view.
	//
	BufferFill optionalnullable.OptionalNullable[int64] `json:"bufferFill,omitempty"`
	// Buffer Frequency measures the rate at which rebuffering events occur, expressed as events per millisecond.
	//
	BufferFrequency optionalnullable.OptionalNullable[float64] `json:"BufferFrequency,omitempty"`
	// Content Delivery Network (CDN) refers to the network infrastructure responsible for delivering the video content to the viewer.
	//
	Cdn optionalnullable.OptionalNullable[string] `json:"cdn,omitempty"`
	// City indicates the geographical location of the viewer accessing the video content.
	//
	City optionalnullable.OptionalNullable[string] `json:"city,omitempty"`
	// Continent represents the continent name of the viewer accessing the video content.
	//
	Continent optionalnullable.OptionalNullable[string] `json:"continent,omitempty"`
	// Country Code denotes the two-letter ISO code representing the country of origin for the viewer accessing the video content.
	//
	CountryCode optionalnullable.OptionalNullable[string] `json:"countryCode,omitempty"`
	// Country represents the coded text that represents the country name of viewer accessing the video content.
	//
	Country optionalnullable.OptionalNullable[string] `json:"country,omitempty"`
	// User defined metadata. Only accessible once it is enabled in the organization settings.
	//
	Custom1 optionalnullable.OptionalNullable[string] `json:"custom1,omitempty"`
	// User defined metadata. Only accessible once it is enabled in the organization settings.
	//
	Custom2 optionalnullable.OptionalNullable[string] `json:"custom2,omitempty"`
	// User defined metadata. Only accessible once it is enabled in the organization settings.
	//
	Custom3 optionalnullable.OptionalNullable[string] `json:"custom3,omitempty"`
	// User defined metadata. Only accessible once it is enabled in the organization settings.
	//
	Custom4 optionalnullable.OptionalNullable[string] `json:"custom4,omitempty"`
	// User defined metadata. Only accessible once it is enabled in the organization settings.
	//
	Custom5 optionalnullable.OptionalNullable[string] `json:"custom5,omitempty"`
	// User defined metadata. Only accessible once it is enabled in the organization settings.
	//
	Custom6 optionalnullable.OptionalNullable[string] `json:"custom6,omitempty"`
	// User defined metadata. Only accessible once it is enabled in the organization settings.
	//
	Custom7 optionalnullable.OptionalNullable[string] `json:"custom7,omitempty"`
	// User defined metadata. Only accessible once it is enabled in the organization settings.
	//
	Custom8 optionalnullable.OptionalNullable[string] `json:"custom8,omitempty"`
	// User defined metadata. Only accessible once it is enabled in the organization settings.
	//
	Custom9 optionalnullable.OptionalNullable[string] `json:"custom9,omitempty"`
	// User defined metadata. Only accessible once it is enabled in the organization settings.
	//
	Custom10 optionalnullable.OptionalNullable[string] `json:"custom10,omitempty"`
	// It is a unique identifier associated with a specific workspace within the FastPix platform.
	//
	WorkspaceID *string `json:"workspaceId,omitempty"`
	// Events specifies the order of events journey of the video playback
	//
	Events []Event `json:"events,omitempty"`
	// Exit Before Video Start indicates whether a viewer abandoned the video before it started playing, typically due to long loading times.
	//
	ExitBeforeVideoStart *bool `json:"exitBeforeVideoStart,omitempty"`
	// Experiment Name is used in A/B testing scenarios to categorize video views into different experiments.
	//
	ExperimentName *string `json:"experimentName,omitempty"`
	// Insert Timestamp refers to the time instance when the view is started.
	//
	InsertTimestamp *string `json:"insertTimestamp,omitempty"`
	// Latitude refers to the geographical coordinate representing the north-south position of the viewer's location, truncated to one decimal place.
	//
	Latitude optionalnullable.OptionalNullable[string] `json:"latitude,omitempty"`
	// FastPix Live Stream ID is the unique identifier associated with a live stream video media within the FastPix Video platform.
	//
	FpLiveStreamID optionalnullable.OptionalNullable[string] `json:"fpLiveStreamId,omitempty"`
	// Live Stream Latency measures the average time taken from the point of ingest to the point of display for live stream video views.
	//
	LiveStreamLatency optionalnullable.OptionalNullable[float64] `json:"liveStreamLatency,omitempty"`
	// Longitude denotes the geographical coordinate representing the east-west position of the viewer's location, truncated to one decimal place.
	//
	Longitude optionalnullable.OptionalNullable[string] `json:"longitude,omitempty"`
	// Page Load Time measures the time from when the user initiates loading the page to when all resources are loaded on the page.
	//
	PageLoadTime optionalnullable.OptionalNullable[int64] `json:"pageLoadTime,omitempty"`
	// Page Context provides contextual information about the type of page being accessed.
	//
	PageContext optionalnullable.OptionalNullable[string] `json:"pageContext,omitempty"`
	// View Page URL denotes the URL address of the web page where the video content is being accessed.
	//
	ViewPageURL optionalnullable.OptionalNullable[string] `json:"viewPageUrl,omitempty"`
	// FastPix Playback ID refers to the unique identifier associated with the playback instance of a video, particularly used in FastPix Video platform.
	//
	FpPlaybackID optionalnullable.OptionalNullable[string] `json:"fpPlaybackId,omitempty"`
	// Playback Success Score represents a numerical value indicating the success or quality of the video playback experience.
	//
	PlaybackScore optionalnullable.OptionalNullable[float32] `json:"playbackScore,omitempty"`
	// Player Autoplay On indicates whether the video player automatically initiated playback of the video content.
	//
	PlayerAutoplayOn *bool `json:"playerAutoplayOn,omitempty"`
	// Error Code is an identifier representing a specific type of error that occurred during video playback, potentially leading to playback failure.
	//
	ErrorCode optionalnullable.OptionalNullable[string] `json:"errorCode,omitempty"`
	// Error Message is a descriptive message generated by the video player when an error occurs during playback, associated with an error code.
	//
	ErrorMessage optionalnullable.OptionalNullable[string] `json:"errorMessage,omitempty"`
	// Player Height refers to the vertical dimension, measured in pixels, of the video player as it appears on the webpage.
	//
	PlayerHeight optionalnullable.OptionalNullable[PlayerHeight] `json:"playerHeight,omitempty"`
	// Player Instance ID is a unique identifier that distinguishes each instance of the Player class created when initializing a video.
	//
	PlayerInstanceID optionalnullable.OptionalNullable[string] `json:"playerInstanceId,omitempty"`
	// Player Language indicates the language used for text elements within the video player interface.
	//
	PlayerLanguage optionalnullable.OptionalNullable[string] `json:"playerLanguage,omitempty"`
	// FastPix SDK Name identifies the name of the FastPix Player SDK utilized within the player workspace.
	//
	FpSDK optionalnullable.OptionalNullable[string] `json:"fpSdk,omitempty"`
	// FastPix SDK Version specifies the version of the FastPix Player SDK integrated into the player.
	//
	FpSDKVersion optionalnullable.OptionalNullable[string] `json:"fpSdkVersion,omitempty"`
	// Player Name serves to differentiate various configurations or types of players used across the website or application.
	//
	PlayerName optionalnullable.OptionalNullable[string] `json:"playerName,omitempty"`
	// Player Poster refers to the image displayed as a preview before the video playback begins.
	//
	PlayerPoster optionalnullable.OptionalNullable[string] `json:"playerPoster,omitempty"`
	// Player Preload On indicates whether the player is configured to preload the video content upon page load.
	//
	PlayerPreloadOn *bool `json:"playerPreloadOn,omitempty"`
	// Player Remote Played specifies if the video is being remotely played to devices such as AirPlay or Chromecast, obtained from the SDK.
	//
	PlayerRemotePlayed *bool `json:"playerRemotePlayed,omitempty"`
	// Player Software Version indicates the version number of the player software installed.
	//
	PlayerSoftwareVersion optionalnullable.OptionalNullable[string] `json:"playerSoftwareVersion,omitempty"`
	// Player Software Name denotes the software utilized for video playback within the player workspace.
	//
	PlayerSoftwareName optionalnullable.OptionalNullable[string] `json:"playerSoftwareName,omitempty"`
	// Video Source Domain identifies the domain from which the video source originates.
	//
	VideoSourceDomain optionalnullable.OptionalNullable[string] `json:"videoSourceDomain,omitempty"`
	// Video Source Duration represents the duration of the video source content, measured in milliseconds.
	//
	VideoSourceDuration optionalnullable.OptionalNullable[int64] `json:"videoSourceDuration,omitempty"`
	// Player Source Height denotes the vertical dimension, measured in pixels, of the source video content being transmitted to the player.
	//
	PlayerSourceHeight optionalnullable.OptionalNullable[int64] `json:"playerSourceHeight,omitempty"`
	// Video Source Hostname represents the hostname of the video
	//
	VideoSourceHostname optionalnullable.OptionalNullable[string] `json:"videoSourceHostname,omitempty"`
	// Video Source Stream Type denotes the type of stream used by the player, although it is currently unused.
	//
	VideoSourceStreamType optionalnullable.OptionalNullable[string] `json:"videoSourceStreamType,omitempty"`
	// Video Source Type denotes the format of the video source as determined by the player, including formats
	//
	VideoSourceType optionalnullable.OptionalNullable[string] `json:"videoSourceType,omitempty"`
	// Player Source URL refers to the URL of the video source accessed by the player.
	//
	VideoSourceURL optionalnullable.OptionalNullable[string] `json:"videoSourceUrl,omitempty"`
	// Source Width represents the width of the source video as perceived by the player, typically measured in pixels.
	//
	PlayerSourceWidth optionalnullable.OptionalNullable[int64] `json:"playerSourceWidth,omitempty"`
	// Player Initialisation Time measures the duration, in milliseconds, from the initialization of the player within the webpage to its readiness to receive further instructions.
	//
	PlayerInitializationTime optionalnullable.OptionalNullable[int64] `json:"playerInitializationTime,omitempty"`
	// Player Version indicates the version of the player used to render the video content. It is often utilized for performance comparison between different player versions.
	//
	PlayerVersion optionalnullable.OptionalNullable[string] `json:"playerVersion,omitempty"`
	// Player Width refers to the width of the player displayed within the webpage, measured in pixels.
	//
	PlayerWidth optionalnullable.OptionalNullable[PlayerWidth] `json:"playerWidth,omitempty"`
	// Render Quality Score is a decimal value representing the score indicating the perceived quality of the video.
	//
	RenderQualityScore optionalnullable.OptionalNullable[float32] `json:"renderQualityScore,omitempty"`
	// Buffer Ratio refers to the percentage of time during video playback where the viewer experiences buffering or rebuffering events.
	//
	BufferRatio optionalnullable.OptionalNullable[float32] `json:"bufferRatio,omitempty"`
	// Stability Score quantifies the smoothness of video playback, typically represented as a decimal value.
	//
	StabilityScore optionalnullable.OptionalNullable[float32] `json:"stabilityScore,omitempty"`
	// Region denotes the geographical region of the viewer accessing the video content.
	//
	Region optionalnullable.OptionalNullable[string] `json:"region,omitempty"`
	// Session ID refers to the unique identifier tracking a viewer's session within the FastPix platform.
	//
	SessionID optionalnullable.OptionalNullable[string] `json:"sessionId,omitempty"`
	// Startup Time Score evaluates the startup performance of the player, usually represented as a decimal value
	//
	StartupScore optionalnullable.OptionalNullable[float32] `json:"startupScore,omitempty"`
	// Sub Property ID denotes the unique identifier assigned to FastPix properties, previously linked with a specific workspace.
	//
	SubPropertyID optionalnullable.OptionalNullable[string] `json:"subPropertyId,omitempty"`
	// Video Startup Time measures the duration, in milliseconds, from the initialization of the player within the webpage to its readiness to receive further instructions.
	//
	VideoStartupTime optionalnullable.OptionalNullable[int64] `json:"videoStartupTime,omitempty"`
	// Updated Timestamp refers to when the record is updated to a particular Video.
	//
	UpdatedTimestamp optionalnullable.OptionalNullable[string] `json:"updatedTimestamp,omitempty"`
	// Used Fullscreen denotes whether the viewer utilized the full-screen mode while watching the video.
	//
	UsedFullScreen *bool `json:"usedFullScreen,omitempty"`
	// Video Content Type specifies the classification of the video content.
	//
	VideoContentType optionalnullable.OptionalNullable[string] `json:"videoContentType,omitempty"`
	// Video Duration represents the length of the video, provided in milliseconds, typically supplied to FastPix via custom metadata.
	//
	VideoDuration optionalnullable.OptionalNullable[int64] `json:"videoDuration,omitempty"`
	// Video ID refers to an internal identifier assigned by the user or system to uniquely identify a particular video.
	//
	VideoID optionalnullable.OptionalNullable[string] `json:"videoId,omitempty"`
	// Video Language denotes the primary audio language of the video content, assuming it remains unchanged after playback initiation.
	//
	VideoLanguage optionalnullable.OptionalNullable[string] `json:"videoLanguage,omitempty"`
	// Video Series denotes the name of a series to which the video content belongs.
	//
	VideoSeries optionalnullable.OptionalNullable[string] `json:"videoSeries,omitempty"`
	// Video Startup Failure is a boolean metric indicating whether a viewer encountered an error before the first frame of the video commenced playback.
	//
	VideoStartupFailed *bool `json:"videoStartupFailed,omitempty"`
	// Video Title refers to the title of the video content being viewed.
	//
	VideoTitle optionalnullable.OptionalNullable[string] `json:"videoTitle,omitempty"`
	// Average Request Latency average time it takes for a request to be made and processed during video playback
	//
	AvgRequestLatency optionalnullable.OptionalNullable[float32] `json:"avgRequestLatency,omitempty"`
	// Average Request Throughput refers to the average throughput or data transfer rate of HTTP requests made during video playback
	//
	AvgRequestThroughput optionalnullable.OptionalNullable[float32] `json:"avgRequestThroughput,omitempty"`
	// DRM Type indicates the type of Digital Rights Management (DRM) utilized during video playback
	//
	DrmType optionalnullable.OptionalNullable[string] `json:"drmType,omitempty"`
	// Dropped Frame Count represents the number of frames dropped by the video player during playback.
	//
	DroppedFrameCount optionalnullable.OptionalNullable[int64] `json:"droppedFrameCount,omitempty"`
	// View End refers to the date and time, in Coordinated Universal Time (UTC), when the video viewing session concluded.
	//
	ViewEnd optionalnullable.OptionalNullable[string] `json:"viewEnd,omitempty"`
	// View Has Ad is a boolean metric indicating whether an advertisement played or attempted to play during the video view.
	//
	ViewHasAd *bool `json:"viewHasAd,omitempty"`
	// View ID is a unique identifier assigned to each individual video viewing session.
	//
	ViewID *string `json:"viewId,omitempty"`
	// Maximum Downscale Percentage represents the highest percentage of downscaling applied to the video during the view.
	//
	MaxDownscaling optionalnullable.OptionalNullable[float32] `json:"maxDownscaling,omitempty"`
	// View Max Playhead Position represents the furthest point reached by the playhead during the video view, measured in milliseconds.
	//
	ViewMaxPlayheadPosition optionalnullable.OptionalNullable[int64] `json:"viewMaxPlayheadPosition,omitempty"`
	// Max request Latency refers to the maximum rate of data transfer (throughput) during requests made by the playback.
	//
	MaxRequestLatency optionalnullable.OptionalNullable[float32] `json:"maxRequestLatency,omitempty"`
	// Maximum Upscale Percentage represents the highest percentage of upscaling applied to the video during the view.
	//
	MaxUpscaling optionalnullable.OptionalNullable[float32] `json:"maxUpscaling,omitempty"`
	// Playing Time denotes the total duration of time the video content was actively playing during the view, excluding time spent buffering, seeking, or joining.
	//
	ViewPlayingTime optionalnullable.OptionalNullable[int64] `json:"viewPlayingTime,omitempty"`
	// View Seeked Count signifies the number of times the viewer attempted to seek to a new location within the video.
	//
	ViewSeekedCount optionalnullable.OptionalNullable[int64] `json:"viewSeekedCount,omitempty"`
	// View Seeked Duration indicates the total duration of time spent waiting for playback to resume after the viewer seeks to a new location. Seek Latency metric in the Dashboard is derived by dividing this value by the view_seek_count.
	//
	ViewSeekedDuration optionalnullable.OptionalNullable[int64] `json:"viewSeekedDuration,omitempty"`
	// View Start refers to the date and time, in Coordinated Universal Time (UTC), when the video viewing session commenced.
	//
	ViewStart optionalnullable.OptionalNullable[string] `json:"viewStart,omitempty"`
	// View Total content Playback Time represents the cumulative duration of video content watched by the viewer, measured in milliseconds. This metric is internally utilized to calculate upscale and downscale percentages.
	//
	ViewTotalContentPlaybackTime optionalnullable.OptionalNullable[int64] `json:"viewTotalContentPlaybackTime,omitempty"`
	// Average Downscaling refers to the average reduction in video resolution or quality during the playback of video content.
	//
	AvgDownscaling optionalnullable.OptionalNullable[float32] `json:"avgDownscaling,omitempty"`
	// Average Upscaling refers to the average resolution of the video source is lower than the resolution of the playback device or screen.
	//
	AvgUpscaling optionalnullable.OptionalNullable[float32] `json:"avgUpscaling,omitempty"`
	// Browser denotes the software application utilized by the viewer to access and watch the video content
	//
	BrowserName optionalnullable.OptionalNullable[string] `json:"browserName,omitempty"`
	// Browser version signifies the specific version of the browser software employed by the viewer
	//
	BrowserVersion optionalnullable.OptionalNullable[string] `json:"browserVersion,omitempty"`
	// Connection Type signifies the type of network connection utilized by the viewer's device
	//
	Connectiontype optionalnullable.OptionalNullable[string] `json:"connectiontype,omitempty"`
	// Device Type denotes the classification of the device used by the viewer
	//
	DeviceType optionalnullable.OptionalNullable[string] `json:"deviceType,omitempty"`
	// Device Manufacturer indicates the brand or manufacturer of the device used by the viewer.
	//
	DeviceManufacturer optionalnullable.OptionalNullable[string] `json:"deviceManufacturer,omitempty"`
	// Device Model represents the specific model of the device used by the viewer.
	//
	DeviceModel optionalnullable.OptionalNullable[string] `json:"deviceModel,omitempty"`
	// Device Name refers to the name or label assigned to the device used by the viewer.
	//
	DeviceName optionalnullable.OptionalNullable[string] `json:"deviceName,omitempty"`
	// Quality Of Experience Score quantifies the overall viewer experience based on various metrics, providing a decimal score to assess the quality of the viewing experience.
	//
	QualityOfExperienceScore optionalnullable.OptionalNullable[float32] `json:"qualityOfExperienceScore,omitempty"`
	// Operating System signifies the name of software platform utilized by the viewer.
	//
	OsName optionalnullable.OptionalNullable[string] `json:"osName,omitempty"`
	// Operating System Version specifies the specific version of the operating system being used by the viewer
	//
	OsVersion *string `json:"osVersion,omitempty"`
	// User Agent represents the user agent string transmitted by the viewer's device to identify itself to the server, typically including information about the device and browser.
	//
	UserAgent optionalnullable.OptionalNullable[string] `json:"userAgent,omitempty"`
	// Viewer ID refers to a customer-defined identifier representing the viewer who is watching the video stream. It should be anonymized and not contain any personally identifiable information.
	//
	ViewerID optionalnullable.OptionalNullable[string] `json:"viewerId,omitempty"`
	// Total Watch Time denotes the total duration of video content watched by the viewer, encompassing startup time, playing time, and potential rebuffering time, measured in milliseconds.
	//
	TotalWatchTime optionalnullable.OptionalNullable[int64] `json:"totalWatchTime,omitempty"`
	// Average Bitrate represents the average bitrate of the video content watched by the viewer, expressed in bits per second (bps). This metric provides insight into the quality of the video stream.
	//
	AverageBitrate optionalnullable.OptionalNullable[float32] `json:"averageBitrate,omitempty"`
	// Jump Latency refers to the delay or latency experienced when there is a jump or seek action performed by the viewer while watching a video.
	//
	JumpLatency optionalnullable.OptionalNullable[float32] `json:"jumpLatency,omitempty"`
	// Player Resolution refers to the resolution of the video player window or viewport where the video content is being displayed.
	//
	PlayerResolution optionalnullable.OptionalNullable[string] `json:"playerResolution,omitempty"`
	// videoResolution refers to the resolution of the video being played.
	//
	VideoResolution optionalnullable.OptionalNullable[string] `json:"videoResolution,omitempty"`
}

func (v *Views) GetAsnName() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.AsnName
}

func (v *Views) GetAsnID() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.AsnID
}

func (v *Views) GetMediaID() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.MediaID
}

func (v *Views) GetBufferCount() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.BufferCount
}

func (v *Views) GetBufferFill() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.BufferFill
}

func (v *Views) GetBufferFrequency() optionalnullable.OptionalNullable[float64] {
	if v == nil {
		return nil
	}
	return v.BufferFrequency
}

func (v *Views) GetCdn() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Cdn
}

func (v *Views) GetCity() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.City
}

func (v *Views) GetContinent() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Continent
}

func (v *Views) GetCountryCode() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.CountryCode
}

func (v *Views) GetCountry() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Country
}

func (v *Views) GetCustom1() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Custom1
}

func (v *Views) GetCustom2() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Custom2
}

func (v *Views) GetCustom3() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Custom3
}

func (v *Views) GetCustom4() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Custom4
}

func (v *Views) GetCustom5() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Custom5
}

func (v *Views) GetCustom6() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Custom6
}

func (v *Views) GetCustom7() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Custom7
}

func (v *Views) GetCustom8() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Custom8
}

func (v *Views) GetCustom9() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Custom9
}

func (v *Views) GetCustom10() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Custom10
}

func (v *Views) GetWorkspaceID() *string {
	if v == nil {
		return nil
	}
	return v.WorkspaceID
}

func (v *Views) GetEvents() []Event {
	if v == nil {
		return nil
	}
	return v.Events
}

func (v *Views) GetExitBeforeVideoStart() *bool {
	if v == nil {
		return nil
	}
	return v.ExitBeforeVideoStart
}

func (v *Views) GetExperimentName() *string {
	if v == nil {
		return nil
	}
	return v.ExperimentName
}

func (v *Views) GetInsertTimestamp() *string {
	if v == nil {
		return nil
	}
	return v.InsertTimestamp
}

func (v *Views) GetLatitude() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Latitude
}

func (v *Views) GetFpLiveStreamID() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.FpLiveStreamID
}

func (v *Views) GetLiveStreamLatency() optionalnullable.OptionalNullable[float64] {
	if v == nil {
		return nil
	}
	return v.LiveStreamLatency
}

func (v *Views) GetLongitude() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Longitude
}

func (v *Views) GetPageLoadTime() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.PageLoadTime
}

func (v *Views) GetPageContext() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.PageContext
}

func (v *Views) GetViewPageURL() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.ViewPageURL
}

func (v *Views) GetFpPlaybackID() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.FpPlaybackID
}

func (v *Views) GetPlaybackScore() optionalnullable.OptionalNullable[float32] {
	if v == nil {
		return nil
	}
	return v.PlaybackScore
}

func (v *Views) GetPlayerAutoplayOn() *bool {
	if v == nil {
		return nil
	}
	return v.PlayerAutoplayOn
}

func (v *Views) GetErrorCode() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.ErrorCode
}

func (v *Views) GetErrorMessage() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.ErrorMessage
}

func (v *Views) GetPlayerHeight() optionalnullable.OptionalNullable[PlayerHeight] {
	if v == nil {
		return nil
	}
	return v.PlayerHeight
}

func (v *Views) GetPlayerInstanceID() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.PlayerInstanceID
}

func (v *Views) GetPlayerLanguage() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.PlayerLanguage
}

func (v *Views) GetFpSDK() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.FpSDK
}

func (v *Views) GetFpSDKVersion() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.FpSDKVersion
}

func (v *Views) GetPlayerName() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.PlayerName
}

func (v *Views) GetPlayerPoster() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.PlayerPoster
}

func (v *Views) GetPlayerPreloadOn() *bool {
	if v == nil {
		return nil
	}
	return v.PlayerPreloadOn
}

func (v *Views) GetPlayerRemotePlayed() *bool {
	if v == nil {
		return nil
	}
	return v.PlayerRemotePlayed
}

func (v *Views) GetPlayerSoftwareVersion() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.PlayerSoftwareVersion
}

func (v *Views) GetPlayerSoftwareName() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.PlayerSoftwareName
}

func (v *Views) GetVideoSourceDomain() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.VideoSourceDomain
}

func (v *Views) GetVideoSourceDuration() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.VideoSourceDuration
}

func (v *Views) GetPlayerSourceHeight() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.PlayerSourceHeight
}

func (v *Views) GetVideoSourceHostname() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.VideoSourceHostname
}

func (v *Views) GetVideoSourceStreamType() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.VideoSourceStreamType
}

func (v *Views) GetVideoSourceType() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.VideoSourceType
}

func (v *Views) GetVideoSourceURL() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.VideoSourceURL
}

func (v *Views) GetPlayerSourceWidth() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.PlayerSourceWidth
}

func (v *Views) GetPlayerInitializationTime() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.PlayerInitializationTime
}

func (v *Views) GetPlayerVersion() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.PlayerVersion
}

func (v *Views) GetPlayerWidth() optionalnullable.OptionalNullable[PlayerWidth] {
	if v == nil {
		return nil
	}
	return v.PlayerWidth
}

func (v *Views) GetRenderQualityScore() optionalnullable.OptionalNullable[float32] {
	if v == nil {
		return nil
	}
	return v.RenderQualityScore
}

func (v *Views) GetBufferRatio() optionalnullable.OptionalNullable[float32] {
	if v == nil {
		return nil
	}
	return v.BufferRatio
}

func (v *Views) GetStabilityScore() optionalnullable.OptionalNullable[float32] {
	if v == nil {
		return nil
	}
	return v.StabilityScore
}

func (v *Views) GetRegion() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Region
}

func (v *Views) GetSessionID() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.SessionID
}

func (v *Views) GetStartupScore() optionalnullable.OptionalNullable[float32] {
	if v == nil {
		return nil
	}
	return v.StartupScore
}

func (v *Views) GetSubPropertyID() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.SubPropertyID
}

func (v *Views) GetVideoStartupTime() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.VideoStartupTime
}

func (v *Views) GetUpdatedTimestamp() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.UpdatedTimestamp
}

func (v *Views) GetUsedFullScreen() *bool {
	if v == nil {
		return nil
	}
	return v.UsedFullScreen
}

func (v *Views) GetVideoContentType() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.VideoContentType
}

func (v *Views) GetVideoDuration() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.VideoDuration
}

func (v *Views) GetVideoID() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.VideoID
}

func (v *Views) GetVideoLanguage() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.VideoLanguage
}

func (v *Views) GetVideoSeries() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.VideoSeries
}

func (v *Views) GetVideoStartupFailed() *bool {
	if v == nil {
		return nil
	}
	return v.VideoStartupFailed
}

func (v *Views) GetVideoTitle() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.VideoTitle
}

func (v *Views) GetAvgRequestLatency() optionalnullable.OptionalNullable[float32] {
	if v == nil {
		return nil
	}
	return v.AvgRequestLatency
}

func (v *Views) GetAvgRequestThroughput() optionalnullable.OptionalNullable[float32] {
	if v == nil {
		return nil
	}
	return v.AvgRequestThroughput
}

func (v *Views) GetDrmType() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.DrmType
}

func (v *Views) GetDroppedFrameCount() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.DroppedFrameCount
}

func (v *Views) GetViewEnd() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.ViewEnd
}

func (v *Views) GetViewHasAd() *bool {
	if v == nil {
		return nil
	}
	return v.ViewHasAd
}

func (v *Views) GetViewID() *string {
	if v == nil {
		return nil
	}
	return v.ViewID
}

func (v *Views) GetMaxDownscaling() optionalnullable.OptionalNullable[float32] {
	if v == nil {
		return nil
	}
	return v.MaxDownscaling
}

func (v *Views) GetViewMaxPlayheadPosition() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.ViewMaxPlayheadPosition
}

func (v *Views) GetMaxRequestLatency() optionalnullable.OptionalNullable[float32] {
	if v == nil {
		return nil
	}
	return v.MaxRequestLatency
}

func (v *Views) GetMaxUpscaling() optionalnullable.OptionalNullable[float32] {
	if v == nil {
		return nil
	}
	return v.MaxUpscaling
}

func (v *Views) GetViewPlayingTime() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.ViewPlayingTime
}

func (v *Views) GetViewSeekedCount() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.ViewSeekedCount
}

func (v *Views) GetViewSeekedDuration() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.ViewSeekedDuration
}

func (v *Views) GetViewStart() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.ViewStart
}

func (v *Views) GetViewTotalContentPlaybackTime() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.ViewTotalContentPlaybackTime
}

func (v *Views) GetAvgDownscaling() optionalnullable.OptionalNullable[float32] {
	if v == nil {
		return nil
	}
	return v.AvgDownscaling
}

func (v *Views) GetAvgUpscaling() optionalnullable.OptionalNullable[float32] {
	if v == nil {
		return nil
	}
	return v.AvgUpscaling
}

func (v *Views) GetBrowserName() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.BrowserName
}

func (v *Views) GetBrowserVersion() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.BrowserVersion
}

func (v *Views) GetConnectiontype() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.Connectiontype
}

func (v *Views) GetDeviceType() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.DeviceType
}

func (v *Views) GetDeviceManufacturer() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.DeviceManufacturer
}

func (v *Views) GetDeviceModel() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.DeviceModel
}

func (v *Views) GetDeviceName() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.DeviceName
}

func (v *Views) GetQualityOfExperienceScore() optionalnullable.OptionalNullable[float32] {
	if v == nil {
		return nil
	}
	return v.QualityOfExperienceScore
}

func (v *Views) GetOsName() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.OsName
}

func (v *Views) GetOsVersion() *string {
	if v == nil {
		return nil
	}
	return v.OsVersion
}

func (v *Views) GetUserAgent() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.UserAgent
}

func (v *Views) GetViewerID() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.ViewerID
}

func (v *Views) GetTotalWatchTime() optionalnullable.OptionalNullable[int64] {
	if v == nil {
		return nil
	}
	return v.TotalWatchTime
}

func (v *Views) GetAverageBitrate() optionalnullable.OptionalNullable[float32] {
	if v == nil {
		return nil
	}
	return v.AverageBitrate
}

func (v *Views) GetJumpLatency() optionalnullable.OptionalNullable[float32] {
	if v == nil {
		return nil
	}
	return v.JumpLatency
}

func (v *Views) GetPlayerResolution() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.PlayerResolution
}

func (v *Views) GetVideoResolution() optionalnullable.OptionalNullable[string] {
	if v == nil {
		return nil
	}
	return v.VideoResolution
}
