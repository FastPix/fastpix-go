package fastpixgo

import (
	"context"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/hooks"
	"github.com/FastPix/fastpix-go/internal/utils"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/retry"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	//On demand streaming
	"https://api.fastpix.io/v1/on-demand",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
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

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*FastPixSDK)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *FastPixSDK) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *FastPixSDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *FastPixSDK) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *FastPixSDK) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security components.Security) SDKOption {
	return func(sdk *FastPixSDK) {
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (components.Security, error)) SDKOption {
	return func(sdk *FastPixSDK) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *FastPixSDK) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *FastPixSDK) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *FastPixSDK {
	sdk := &FastPixSDK{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0.0",
			SDKVersion:        "0.1.0",
			GenVersion:        "2.605.6",
			UserAgent:         "speakeasy-sdk/go 0.1.0 2.605.6 1.0.0 github.com/FastPix/fastpix-go",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	if sdk.sdkConfiguration.Security == nil {
		var envVarSecurity components.Security
		if utils.PopulateSecurityFromEnv(&envVarSecurity) {
			sdk.sdkConfiguration.Security = utils.AsSecuritySource(envVarSecurity)
		}
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.StartLiveStream = newStartLiveStream(sdk.sdkConfiguration)

	sdk.ManageLiveStream = newManageLiveStream(sdk.sdkConfiguration)

	sdk.Playback = newPlayback(sdk.sdkConfiguration)

	sdk.SimulcastStream = newSimulcastStream(sdk.sdkConfiguration)

	sdk.InputVideo = newInputVideo(sdk.sdkConfiguration)

	sdk.ManageVideos = newManageVideos(sdk.sdkConfiguration)

	return sdk
}
