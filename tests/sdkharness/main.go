// Command sdkharness is invoked as a subprocess by the TypeScript endpoint
// validators (validate-get-endpoints.ts / validate-non-get-endpoints.ts).
//
// It reads a single JSON payload from stdin:
//
//	{ "operationId": "...", "request": {...}, "baseUrl": "...",
//	  "username": "...", "password": "..." }
//
// dispatches to the matching fastpix-go SDK method, and prints a JSON result to
// stdout following the contract the validators expect:
//
//	success: { "ok": true, "value": <sdk response body>, "statusCode": <int|null>, "rawBody": <parsed json|string|null> }
//	failure: { "ok": false, "error": { "name": "...", "message": "...", "statusCode": <int>, "bodyJson": <...> } }
//
// The raw wire body is captured via a wrapping HTTP client so the non-GET
// validator can still validate the on-the-wire response against the OpenAPI
// schema (the SDK otherwise consumes the body during deserialization).
//
// This file is hand-written (not generated). It lives under tests/ so it builds
// as part of the module and can resolve the local SDK packages directly.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/apierrors"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/models/operations"
)

// ---------------------------------------------------------------------------
// stdin / stdout contract
// ---------------------------------------------------------------------------

type payload struct {
	OperationID string         `json:"operationId"`
	Request     map[string]any `json:"request"`
	BaseURL     string         `json:"baseUrl"`
	Username    string         `json:"username"`
	Password    string         `json:"password"`
}

type errOut struct {
	Name       string `json:"name"`
	Message    string `json:"message"`
	StatusCode *int   `json:"statusCode,omitempty"`
	BodyJSON   any    `json:"bodyJson,omitempty"`
}

func emitErr(e errOut) {
	out, _ := json.Marshal(map[string]any{"ok": false, "error": e})
	fmt.Println(string(out))
}

func emitOK(value any, statusCode *int, rawBody any) {
	out, _ := json.Marshal(map[string]any{
		"ok":         true,
		"value":      value,
		"statusCode": statusCode,
		"rawBody":    rawBody,
	})
	fmt.Println(string(out))
}

// ---------------------------------------------------------------------------
// capturing HTTP client — retains the last response's status + raw body so the
// non-GET validator can compare against the real wire JSON.
// ---------------------------------------------------------------------------

type capturingClient struct {
	inner      *http.Client
	lastStatus int
	lastBody   []byte
}

func (c *capturingClient) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.inner.Do(req)
	if err != nil || resp == nil {
		return resp, err
	}
	c.lastStatus = resp.StatusCode
	if resp.Body != nil {
		b, readErr := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if readErr == nil {
			c.lastBody = b
			resp.Body = io.NopCloser(bytes.NewReader(b))
		}
	}
	return resp, err
}

// ---------------------------------------------------------------------------
// request-argument helpers (the JSON `request` map is loosely typed)
// ---------------------------------------------------------------------------

func str(req map[string]any, key string) *string {
	v, ok := req[key]
	if !ok || v == nil {
		return nil
	}
	s, ok := v.(string)
	if !ok {
		return nil
	}
	return &s
}

func strVal(req map[string]any, key string) string {
	if p := str(req, key); p != nil {
		return *p
	}
	return ""
}

func i64(req map[string]any, key string) *int64 {
	v, ok := req[key]
	if !ok || v == nil {
		return nil
	}
	switch n := v.(type) {
	case float64:
		x := int64(n)
		return &x
	case int64:
		return &n
	case json.Number:
		if x, err := n.Int64(); err == nil {
			return &x
		}
	}
	return nil
}

func sortOrder(req map[string]any) *components.SortOrder {
	if str(req, "orderBy") == nil {
		return nil
	}
	v := components.SortOrderAsc
	if strVal(req, "orderBy") == "desc" {
		v = components.SortOrderDesc
	}
	return &v
}

func opOrderBy(req map[string]any) *operations.OrderBy {
	if str(req, "orderBy") == nil {
		return nil
	}
	v := operations.OrderByAsc
	if strVal(req, "orderBy") == "desc" {
		v = operations.OrderByDesc
	}
	return &v
}

// ptr returns a pointer to v — used for the string-backed enum casts below.
func ptr[T any](v T) *T { return &v }

// ---------------------------------------------------------------------------
// response extraction — find the populated body field on the SDK response
// struct (skipping HTTPMeta and other metadata) via reflection, mirroring the
// PHP harness's get_object_vars walk.
// ---------------------------------------------------------------------------

var metadataFields = map[string]bool{
	"HTTPMeta":     true,
	"ContentType":  true,
	"StatusCode":   true,
	"RawResponse":  true,
	"DefaultError": true,
	"Headers":      true,
}

func extractBody(res any) any {
	if res == nil {
		return nil
	}
	rv := reflect.ValueOf(res)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return res
	}
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		f := rt.Field(i)
		if f.PkgPath != "" { // unexported
			continue
		}
		if metadataFields[f.Name] {
			continue
		}
		fv := rv.Field(i)
		switch fv.Kind() {
		case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Interface:
			if !fv.IsNil() {
				return fv.Interface()
			}
		default:
			// non-nilable populated body (rare); return as-is
			return fv.Interface()
		}
	}
	return res
}

func httpStatusOf(res any) *int {
	rv := reflect.ValueOf(res)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return nil
	}
	meta := rv.FieldByName("HTTPMeta")
	if !meta.IsValid() {
		return nil
	}
	respField := meta.FieldByName("Response")
	if !respField.IsValid() || respField.IsNil() {
		return nil
	}
	resp, ok := respField.Interface().(*http.Response)
	if !ok || resp == nil {
		return nil
	}
	return &resp.StatusCode
}

// ---------------------------------------------------------------------------
// main
// ---------------------------------------------------------------------------

func main() {
	raw, err := io.ReadAll(os.Stdin)
	if err != nil {
		emitErr(errOut{Name: "StdinError", Message: err.Error()})
		return
	}
	var p payload
	if err := json.Unmarshal(raw, &p); err != nil {
		emitErr(errOut{Name: "PayloadParseError", Message: err.Error()})
		return
	}

	cc := &capturingClient{inner: &http.Client{}}
	opts := []fastpixgo.SDKOption{
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.String(p.Username),
			Password: fastpixgo.String(p.Password),
		}),
		fastpixgo.WithClient(cc),
	}
	if p.BaseURL != "" {
		opts = append(opts, fastpixgo.WithServerURL(p.BaseURL))
	}
	s := fastpixgo.New(opts...)

	res, callErr := dispatch(s, p.OperationID, p.Request)
	if callErr != nil {
		// Surface API error details (status + parsed body) when available.
		var apiErr *apierrors.APIError
		eo := errOut{Name: "SDKError", Message: callErr.Error()}
		if errors.As(callErr, &apiErr) {
			eo.Name = "APIError"
			sc := apiErr.StatusCode
			eo.StatusCode = &sc
			var bodyJSON any
			if json.Unmarshal([]byte(apiErr.Body), &bodyJSON) == nil {
				eo.BodyJSON = bodyJSON
			}
		} else if cc.lastStatus != 0 {
			sc := cc.lastStatus
			eo.StatusCode = &sc
		}
		emitErr(eo)
		return
	}

	value := extractBody(res)
	status := httpStatusOf(res)
	if status == nil && cc.lastStatus != 0 {
		status = &cc.lastStatus
	}

	var rawBody any
	if len(cc.lastBody) > 0 {
		if json.Unmarshal(cc.lastBody, &rawBody) != nil {
			rawBody = string(cc.lastBody)
		}
	}

	emitOK(value, status, rawBody)
}

// dispatch maps an operationId to the corresponding SDK call. It returns the
// SDK response object (so reflection can pull out the body) or an error.
func dispatch(s *fastpixgo.Fastpixgo, op string, req map[string]any) (any, error) {
	ctx := context.Background()

	switch op {
	// ----------------------------- GET -----------------------------
	case "list-media":
		return s.ManageVideos.List(ctx, i64(req, "limit"), i64(req, "offset"), sortOrder(req))
	case "get-media":
		return s.Videos.Get(ctx, strVal(req, "mediaId"))
	case "get-media-summary":
		return s.ManageVideos.GetMediaSummary(ctx, strVal(req, "mediaId"))
	case "retrieveMediaInputInfo":
		return s.ManageVideos.GetInputInfo(ctx, strVal(req, "mediaId"))
	case "list-uploads":
		return s.ManageVideos.ListUnusedUploadUrls(ctx, i64(req, "limit"), i64(req, "offset"), sortOrder(req))
	case "get-media-clips":
		return s.ManageVideos.GetMediaClips(ctx, strVal(req, "mediaId"), i64(req, "offset"), i64(req, "limit"), sortOrder(req))
	case "list-live-clips":
		return s.Videos.ListLiveClips(ctx, strVal(req, "livestreamId"), i64(req, "limit"), i64(req, "offset"), sortOrder(req))
	case "get-all-playlists":
		return s.Playlists.List(ctx, i64(req, "limit"), i64(req, "offset"))
	case "get-playlist-by-id":
		return s.Playlist.Get(ctx, strVal(req, "playlistId"))
	case "list-playback-ids":
		return s.Playback.List(ctx, strVal(req, "mediaId"))
	case "get-playback-id":
		return s.Playback.GetByID(ctx, strVal(req, "mediaId"), strVal(req, "playbackId"))
	case "getDrmConfiguration":
		return s.DrmConfigurations.List(ctx, i64(req, "offset"), i64(req, "limit"))
	case "getDrmConfigurationById":
		return s.DrmConfigurations.GetByID(ctx, strVal(req, "drmConfigurationId"))
	case "get-all-streams":
		return s.ManageLiveStream.List(ctx, i64(req, "limit"), i64(req, "offset"), opOrderBy(req))
	case "get-live-stream-by-id":
		return s.LiveStreams.GetByID(ctx, strVal(req, "streamId"))
	case "get-live-stream-viewer-count-by-id":
		return s.LiveStreams.GetViewerCount(ctx, strVal(req, "streamId"))
	case "get-live-stream-playback-id":
		return s.LivePlayback.GetPlaybackIDDetails(ctx, strVal(req, "streamId"), strVal(req, "playbackId"))
	case "get-specific-simulcast-of-stream":
		return s.SimulcastStreams.GetSpecific(ctx, strVal(req, "streamId"), strVal(req, "simulcastId"))
	case "list_signing_keys":
		return s.SigningKeys.List(ctx, i64(req, "limit"), i64(req, "offset"))
	case "get-signing_key_by_id":
		return s.SigningKeys.GetByID(ctx, strVal(req, "signingKeyId"))
	case "list_video_views":
		r := operations.ListVideoViewsRequest{Limit: i64(req, "limit"), Offset: i64(req, "offset")}
		if str(req, "timespan") != nil {
			r.Timespan = ptr(operations.ListVideoViewsTimespan(strVal(req, "timespan")))
		}
		return s.Views.ListVideoViews(ctx, r)
	case "get_video_view_details":
		return s.Views.GetDetails(ctx, strVal(req, "viewId"))
	case "list_by_top_content":
		var ts *operations.ListByTopContentTimespan
		if str(req, "timespan") != nil {
			ts = ptr(operations.ListByTopContentTimespan(strVal(req, "timespan")))
		}
		return s.Views.ListByTopContent(ctx, ts, nil, i64(req, "limit"))
	case "list_dimensions":
		return s.Dimensions.List(ctx)
	case "list_filter_values_for_dimension":
		var ts *operations.ListFilterValuesForDimensionTimespan
		if str(req, "timespan") != nil {
			ts = ptr(operations.ListFilterValuesForDimensionTimespan(strVal(req, "timespan")))
		}
		return s.Dimensions.ListFilterValues(ctx, operations.DimensionsID(strVal(req, "dimensionsId")), ts, nil)
	case "list_breakdown_values":
		r := operations.ListBreakdownValuesRequest{
			MetricID: operations.ListBreakdownValuesMetricID(strVal(req, "metricId")),
			GroupBy:  str(req, "groupBy"),
		}
		if str(req, "timespan") != nil {
			r.Timespan = ptr(operations.ListBreakdownValuesTimespan(strVal(req, "timespan")))
		}
		return s.Metrics.ListBreakdownValues(ctx, r)
	case "list_overall_values":
		var ts *operations.ListOverallValuesTimespan
		if str(req, "timespan") != nil {
			ts = ptr(operations.ListOverallValuesTimespan(strVal(req, "timespan")))
		}
		return s.Metrics.ListOverallValues(ctx, operations.ListOverallValuesMetricID(strVal(req, "metricId")), nil, ts, nil)
	case "get_timeseries_data":
		r := operations.GetTimeseriesDataRequest{
			MetricID: operations.GetTimeseriesDataMetricID(strVal(req, "metricId")),
		}
		if str(req, "timespan") != nil {
			r.Timespan = ptr(operations.GetTimeseriesDataTimespan(strVal(req, "timespan")))
		}
		if str(req, "groupBy") != nil {
			r.GroupBy = ptr(operations.GroupBy(strVal(req, "groupBy")))
		}
		return s.Metrics.GetTimeseriesData(ctx, r)
	case "list_comparison_values":
		var ts *operations.ListComparisonValuesTimespan
		if str(req, "timespan") != nil {
			ts = ptr(operations.ListComparisonValuesTimespan(strVal(req, "timespan")))
		}
		var dim *operations.Dimension
		if str(req, "dimension") != nil {
			dim = ptr(operations.Dimension(strVal(req, "dimension")))
		}
		return s.Metrics.ListComparisonValues(ctx, ts, nil, dim, str(req, "value"))
	case "list_errors":
		var ts *operations.ListErrorsTimespan
		if str(req, "timespan") != nil {
			ts = ptr(operations.ListErrorsTimespan(strVal(req, "timespan")))
		}
		return s.Errors.List(ctx, ts, nil, i64(req, "limit"))

	// --------------------------- POST (create) ---------------------------
	case "create-media":
		return s.InputVideo.Create(ctx, components.CreateMediaRequest{
			Inputs:   []components.Input{components.CreateInputPullVideoInput(components.PullVideoInput{})},
			Metadata: map[string]string{"source": "sdk-validate"},
		})
	case "create_signing_key":
		return s.SigningKeys.Create(ctx)
	case "create-a-playlist":
		return s.Playlists.Create(ctx, components.CreateCreatePlaylistRequestManual(components.CreatePlaylistRequestManual{
			Name: "sdk-validate-playlist",
			// The API requires referenceId to be alphanumeric (no separators).
			ReferenceID: fmt.Sprintf("sdkvalidate%d", os.Getpid()),
			Type:        components.CreatePlaylistRequestManualTypeManual,
		}))
	case "create-new-stream":
		return s.StartLiveStream.Create(ctx, components.CreateLiveStreamRequest{
			PlaybackSettings:   components.PlaybackSettings{},
			InputMediaSettings: components.InputMediaSettings{Metadata: map[string]string{"name": "sdk-validate"}},
		})
	case "create-media-playback-id":
		// Request an explicit resolution so the API returns a non-null
		// `resolution` in the response. Omitting it makes the API return
		// `resolution: null`, which the OpenAPI spec's enum does not model.
		res := operations.ResolutionOneThousandAndEightyp
		return s.Playback.Create(ctx, strVal(req, "mediaId"), &operations.CreateMediaPlaybackIDRequestBody{
			AccessPolicy: components.AccessPolicyPublic,
			Resolution:   &res,
		})
	case "Add-media-track":
		return s.Videos.AddMediaTrack(ctx, strVal(req, "mediaId"), operations.AddMediaTrackRequestBody{
			Tracks: components.AddTrackRequest{},
		})
	case "Generate-subtitle-track":
		return s.ManageVideos.GenerateSubtitleTrack(ctx, strVal(req, "mediaId"), strVal(req, "trackId"), components.TrackSubtitlesGenerateRequest{})
	case "create-playbackId-of-stream":
		return s.LivePlayback.Create(ctx, strVal(req, "streamId"), components.PlaybackIDRequest{})
	case "create-simulcast-of-stream":
		return s.SimulcastStreams.Create(ctx, strVal(req, "streamId"), components.SimulcastRequest{
			URL:       fastpixgo.String("rtmp://example.com/live"),
			StreamKey: fastpixgo.String(fmt.Sprintf("sk-%d", os.Getpid())),
		})
	case "direct-upload-video-media":
		return s.InputVideo.DirectUploadMedia(ctx, &operations.DirectUploadVideoMediaRequest{
			PushMediaSettings: &operations.PushMediaSettings{Metadata: map[string]string{"source": "sdk-validate"}},
		})

	// --------------------------- PUT / PATCH (update) ---------------------------
	case "updated-media":
		return s.Videos.Update(ctx, strVal(req, "mediaId"), operations.UpdatedMediaRequestBody{
			Metadata: map[string]string{"updated": "true"},
			Title:    fastpixgo.String("SDK Validate Title"),
		})
	case "updated-source-access":
		return s.Videos.UpdateSourceAccess(ctx, strVal(req, "mediaId"), operations.UpdatedSourceAccessRequestBody{
			SourceAccess: true,
		})
	case "updated-mp4Support":
		return s.Videos.UpdateMp4Support(ctx, strVal(req, "mediaId"), operations.UpdatedMp4SupportRequestBody{})
	case "update-media-summary":
		return s.InVideoAI.GenerateSummary(ctx, strVal(req, "mediaId"), operations.UpdateMediaSummaryRequestBody{Generate: true})
	case "update-media-chapters":
		return s.InVideoAIFeatures.UpdateChapters(ctx, strVal(req, "mediaId"), operations.UpdateMediaChaptersRequestBody{Chapters: fastpixgo.Bool(true)})
	case "update-media-named-entities":
		return s.InVideoAIFeatures.UpdateMediaNamedEntities(ctx, strVal(req, "mediaId"), operations.UpdateMediaNamedEntitiesRequestBody{NamedEntities: true})
	case "update-media-moderation":
		return s.InVideoAIFeatures.UpdateModeration(ctx, strVal(req, "mediaId"), operations.UpdateMediaModerationRequestBody{
			Moderation: &operations.UpdateMediaModerationModeration{},
		})
	case "update-media-track":
		return s.ManageVideos.UpdateTrack(ctx, strVal(req, "trackId"), strVal(req, "mediaId"), components.UpdateTrackRequest{})
	case "update-domain-restrictions":
		return s.Playback.UpdateDomainRestrictions(ctx, strVal(req, "mediaId"), strVal(req, "playbackId"), operations.UpdateDomainRestrictionsRequestBody{
			Allow: []string{"example.com"},
		})
	case "update-user-agent-restrictions":
		return s.PlaybackIds.UpdateUserAgentRestrictions(ctx, strVal(req, "mediaId"), strVal(req, "playbackId"), operations.UpdateUserAgentRestrictionsRequestBody{
			Allow: []string{"Mozilla"},
		})
	case "update-a-playlist":
		return s.Playlists.Update(ctx, strVal(req, "playlistId"), components.UpdatePlaylistRequest{
			Name:        "SDK Validate Updated",
			Description: "updated by validator",
		})
	case "add-media-to-playlist":
		return s.Playlist.AddMedia(ctx, strVal(req, "playlistId"), components.MediaIdsRequest{MediaIds: []string{strVal(req, "mediaId")}})
	case "change-media-order-in-playlist":
		return s.Playlist.ChangeMediaOrder(ctx, strVal(req, "playlistId"), components.MediaIdsRequest{MediaIds: []string{strVal(req, "mediaId")}})
	case "update-live-stream":
		return s.ManageLiveStream.UpdateLiveStream(ctx, strVal(req, "streamId"), components.PatchLiveStreamRequest{
			Metadata:        map[string]string{"updated": "true"},
			ReconnectWindow: fastpixgo.Int64(120),
		})
	case "update-specific-simulcast-of-stream":
		return s.SimulcastStreams.Update(ctx, strVal(req, "streamId"), strVal(req, "simulcastId"), components.SimulcastUpdateRequest{
			IsEnabled: fastpixgo.Bool(false),
		})
	case "enable-live-stream":
		return s.ManageLiveStream.Enable(ctx, strVal(req, "streamId"))
	case "disable-live-stream":
		return s.LiveStreams.Disable(ctx, strVal(req, "streamId"))
	case "complete-live-stream":
		return s.ManageLiveStream.Complete(ctx, strVal(req, "streamId"))
	case "cancel-upload":
		return s.ManageVideos.CancelUpload(ctx, strVal(req, "uploadId"))

	// ----------------------------- DELETE -----------------------------
	case "delete-media-from-playlist":
		return s.Playlists.DeleteMedia(ctx, strVal(req, "playlistId"), &components.MediaIdsRequest{MediaIds: []string{strVal(req, "mediaId")}})
	case "delete-a-playlist":
		return s.Playlists.Delete(ctx, strVal(req, "playlistId"))
	case "delete-media-track":
		return s.Videos.DeleteTrack(ctx, strVal(req, "mediaId"), strVal(req, "trackId"))
	case "delete-media-playback-id":
		return s.Playback.Delete(ctx, strVal(req, "mediaId"), strVal(req, "playbackId"))
	case "delete-simulcast-of-stream":
		return s.SimulcastStream.Delete(ctx, strVal(req, "streamId"), strVal(req, "simulcastId"))
	case "delete-playbackId-of-stream":
		return s.LivePlayback.DeletePlaybackID(ctx, strVal(req, "streamId"), strVal(req, "playbackId"))
	case "delete-live-stream":
		return s.LiveStreams.Delete(ctx, strVal(req, "streamId"))
	case "delete-media":
		return s.ManageVideos.Delete(ctx, strVal(req, "mediaId"))
	case "delete_signing_key":
		return s.SigningKeys.Delete(ctx, strVal(req, "signingKeyId"))
	}

	return nil, fmt.Errorf("no Go SDK mapping for operationId %q", op)
}
