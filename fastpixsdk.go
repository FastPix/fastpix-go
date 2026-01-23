package fastpixgo

import (
	"github.com/FastPix/fastpix-go/internal/config"
	"github.com/FastPix/fastpix-go/internal/hooks"
)

// LiveStreamServerList contains the list of servers available for live streaming
var LiveStreamServerList = []string{
	// LIVE STREAM
	"https://api.fastpix.io/live",
}

// FastPixSDK - LIVE STREAMING API: The Live Stream APIs in FastPix simplifies the process of creating,managing, and distributing live content. This set of API endpoints is designed to help developers initiate live broadcasts, configure stream settings, and extend streams to external platforms (via simulcasting). By integrating FastPix's live streaming capabilities into your applications, you can provide users with seamless and high-quality live video experiences, whether for events, webinars, gaming, or live content creation.
//
//	Live streams can be customized with various parameters, such as stream metadata, privacy settings, and playback configurations. Additionally, the API supports real-time interaction with streams, including updating stream details, managing playback IDs, and extending the reach of a stream through simulcasting to platforms like YouTube or Facebook.
//
// <h3>Use case scenarios</h3>
//
//	**Event Broadcasting**: A developer integrates FastPix live streaming APIs into an event management platform. By leveraging these APIs, the platform can enable event organizers to set up live streams for conferences, concerts, or webinars, allowing viewers to tune in from multiple platforms simultaneously via simulcasting.
//
//	**Live Content Platforms**: Developers working on a live content platform for creators can use the live stream APIs to allow users to broadcast gaming, vlogs, or tutorials. Creators can manage their streams in real time, control playback options, and extend their reach by simulcasting to popular platforms like Twitch or YouTube.
//
//	**Corporate Streaming Services**: A corporate communication tool can integrate live streaming functionality for internal town halls or global employee meetings. The live streams can be made accessible to different employee groups through privacy settings and playback control, ensuring secure and efficient internal communication.
type FastPixSDK struct {
	StartLiveStream  *StartLiveStream
	ManageLiveStream *ManageLiveStream
	Playback         *Playback
	SimulcastStream  *SimulcastStream
	InputVideo       *InputVideo
	ManageVideos     *ManageVideos

	// Internal Fastpixgo instance
	fastpixgo *Fastpixgo
}

type FastPixSDKOption func(*FastPixSDK)

// Note: FastPixSDKOption functions are placeholders for backward compatibility.
// The actual options should be passed when creating the internal Fastpixgo instance.

// NewFastPixSDK creates a new instance of the FastPixSDK with the provided options
// Note: This is a legacy wrapper. Consider using New() which returns *Fastpixgo instead.
// The opts parameter is currently ignored - use New() with WithServerURL(LiveStreamServerList[0]) instead.
func NewFastPixSDK(opts ...FastPixSDKOption) *FastPixSDK {
	// Create internal Fastpixgo instance with live stream server
	fastpixOpts := []SDKOption{
		WithServerURL(LiveStreamServerList[0]),
	}
	
	fg := createFastpixgoWithLiveStreamServer(fastpixOpts...)
	
	// Create wrapper
	return NewFastPixSDKFromFastpixgo(fg)
}

// NewFastPixSDKFromFastpixgo creates a FastPixSDK wrapper from an existing Fastpixgo instance
func NewFastPixSDKFromFastpixgo(fg *Fastpixgo) *FastPixSDK {
	sdk := &FastPixSDK{
		fastpixgo:        fg,
		StartLiveStream:  fg.StartLiveStream,
		ManageLiveStream: fg.ManageLiveStream,
		Playback:         fg.Playback,
		SimulcastStream:  fg.SimulcastStream,
		InputVideo:       fg.InputVideo,
		ManageVideos:     fg.ManageVideos,
	}
	return sdk
}

// Helper function to create Fastpixgo with live stream server
func createFastpixgoWithLiveStreamServer(opts ...SDKOption) *Fastpixgo {
	// Create options with live stream server URL if not already set
	hasServerURL := false
	for _, opt := range opts {
		// Check if any option sets ServerURL by applying to a temp SDK
		tempSDK := &Fastpixgo{
			sdkConfiguration: config.SDKConfiguration{
				ServerList: ServerList,
			},
			hooks: hooks.New(),
		}
		opt(tempSDK)
		if tempSDK.sdkConfiguration.ServerURL != "" {
			hasServerURL = true
			break
		}
	}
	
	// If no server URL is set, use live stream server
	if !hasServerURL {
		opts = append([]SDKOption{WithServerURL(LiveStreamServerList[0])}, opts...)
	}
	
	return New(opts...)
}
