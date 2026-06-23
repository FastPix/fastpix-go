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
	// BodyJSON holds the API error body as raw JSON bytes (not decoded into a
	// map) so the field order from the wire is preserved verbatim. Decoding into
	// `any` would yield a map[string]any, and encoding/json sorts map keys
	// alphabetically — reordering e.g. {success, error} into {error, success}.
	BodyJSON json.RawMessage `json:"bodyJson,omitempty"`
}

// Request key constants — eliminates S1192 duplicate string literals.
const (
	keyMediaID      = "mediaId"
	keyStreamID     = "streamId"
	keyPlaylistID   = "playlistId"
	keyPlaybackID   = "playbackId"
	keySimulcastID  = "simulcastId"
	keyTrackID      = "trackId"
	keyMetricID     = "metricId"
	keyTimespan     = "timespan"
	keyGroupBy      = "groupBy"
	keyDimension    = "dimension"
	keyLimit        = "limit"
	keyOffset       = "offset"
	keySigningKeyID = "signingKeyId"
	keyUploadID     = "uploadId"
	keyViewID       = "viewId"
	keyDimID        = "dimensionsId"
	keyLivestreamID = "livestreamId"
	keyDrmConfigID  = "drmConfigurationId"
	keyValue        = "value"
	keyOrderBy      = "orderBy"
	keyName         = "name"
	errNoSDKMapping = "no Go SDK mapping for operationId %q"
)

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
	if str(req, keyOrderBy) == nil {
		return nil
	}
	v := components.SortOrderAsc
	if strVal(req, keyOrderBy) == "desc" {
		v = components.SortOrderDesc
	}
	return &v
}

func opOrderBy(req map[string]any) *operations.OrderBy {
	if str(req, keyOrderBy) == nil {
		return nil
	}
	v := operations.OrderByAsc
	if strVal(req, keyOrderBy) == "desc" {
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
			// Preserve the wire byte order: keep the raw body as-is when it is
			// valid JSON instead of decoding into a map (which would sort keys).
			if json.Valid([]byte(apiErr.Body)) {
				eo.BodyJSON = json.RawMessage(apiErr.Body)
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
	const sdkValidate = "sdk-validate"
	_, isGet := map[string]struct{}{
		"list-media": {}, "get-media": {}, "get-media-summary": {}, "retrieveMediaInputInfo": {},
		"list-uploads": {}, "get-media-clips": {}, "list-live-clips": {}, "get-all-playlists": {},
		"get-playlist-by-id": {}, "list-playback-ids": {}, "get-playback-id": {},
		"getDrmConfiguration": {}, "getDrmConfigurationById": {}, "get-all-streams": {},
		"get-live-stream-by-id": {}, "get-live-stream-viewer-count-by-id": {},
		"get-live-stream-playback-id": {}, "get-specific-simulcast-of-stream": {},
		"list_signing_keys": {}, "get-signing_key_by_id": {}, "list_video_views": {},
		"get_video_view_details": {}, "list_by_top_content": {}, "list_dimensions": {},
		"list_filter_values_for_dimension": {}, "list_breakdown_values": {},
		"list_overall_values": {}, "get_timeseries_data": {}, "list_comparison_values": {},
		"list_errors": {},
	}[op]
	if isGet {
		return dispatchGet(s, op, req)
	}
	_, isCreate := map[string]struct{}{
		"create-media": {}, "create_signing_key": {}, "create-a-playlist": {},
		"create-new-stream": {}, "create-media-playback-id": {}, "Add-media-track": {},
		"create-playbackId-of-stream": {}, "create-simulcast-of-stream": {},
		"direct-upload-video-media": {}, "Generate-subtitle-track": {},
	}[op]
	if isCreate {
		return dispatchCreate(s, op, req, sdkValidate)
	}
	_, isDelete := map[string]struct{}{
		"delete-media-from-playlist": {}, "delete-a-playlist": {}, "delete-media-track": {},
		"delete-media-playback-id": {}, "delete-simulcast-of-stream": {},
		"delete-playbackId-of-stream": {}, "delete-live-stream": {},
		"delete-media": {}, "delete_signing_key": {},
	}[op]
	if isDelete {
		return dispatchDelete(s, op, req)
	}
	return dispatchUpdate(s, op, req)
}

// ptrTimespan returns a typed timespan pointer if the "timespan" key is present.
func ptrStr[T ~string](req map[string]any, key string) *T {
	if s := str(req, key); s != nil {
		v := T(*s)
		return &v
	}
	return nil
}

func dispatchGet(s *fastpixgo.Fastpixgo, op string, req map[string]any) (any, error) {
	ctx := context.Background()
	switch op {
	case "list-media":
		return s.ManageVideos.List(ctx, i64(req, keyLimit), i64(req, keyOffset), sortOrder(req))
	case "get-media":
		return s.Videos.Get(ctx, strVal(req, keyMediaID))
	case "get-media-summary":
		return s.ManageVideos.GetMediaSummary(ctx, strVal(req, keyMediaID))
	case "retrieveMediaInputInfo":
		return s.ManageVideos.GetInputInfo(ctx, strVal(req, keyMediaID))
	case "list-uploads":
		return s.ManageVideos.ListUnusedUploadUrls(ctx, i64(req, keyLimit), i64(req, keyOffset), sortOrder(req))
	case "get-media-clips":
		return s.ManageVideos.GetMediaClips(ctx, strVal(req, keyMediaID), i64(req, keyOffset), i64(req, keyLimit), sortOrder(req))
	case "list-live-clips":
		return s.Videos.ListLiveClips(ctx, strVal(req, keyLivestreamID), i64(req, keyLimit), i64(req, keyOffset), sortOrder(req))
	case "get-all-playlists":
		return s.Playlists.List(ctx, i64(req, keyLimit), i64(req, keyOffset))
	case "get-playlist-by-id":
		return s.Playlist.Get(ctx, strVal(req, keyPlaylistID))
	case "list-playback-ids":
		return s.Playback.List(ctx, strVal(req, keyMediaID))
	case "get-playback-id":
		return s.Playback.GetByID(ctx, strVal(req, keyMediaID), strVal(req, keyPlaybackID))
	case "getDrmConfiguration":
		return s.DrmConfigurations.List(ctx, i64(req, keyOffset), i64(req, keyLimit))
	case "getDrmConfigurationById":
		return s.DrmConfigurations.GetByID(ctx, strVal(req, keyDrmConfigID))
	case "get-all-streams":
		return s.ManageLiveStream.List(ctx, i64(req, keyLimit), i64(req, keyOffset), opOrderBy(req))
	case "get-live-stream-by-id":
		return s.LiveStreams.GetByID(ctx, strVal(req, keyStreamID))
	case "get-live-stream-viewer-count-by-id":
		return s.LiveStreams.GetViewerCount(ctx, strVal(req, keyStreamID))
	case "get-live-stream-playback-id":
		return s.LivePlayback.GetPlaybackIDDetails(ctx, strVal(req, keyStreamID), strVal(req, keyPlaybackID))
	case "get-specific-simulcast-of-stream":
		return s.SimulcastStreams.GetSpecific(ctx, strVal(req, keyStreamID), strVal(req, keySimulcastID))
	case "list_signing_keys":
		return s.SigningKeys.List(ctx, i64(req, keyLimit), i64(req, keyOffset))
	case "get-signing_key_by_id":
		return s.SigningKeys.GetByID(ctx, strVal(req, keySigningKeyID))
	case "list_video_views":
		r := operations.ListVideoViewsRequest{Limit: i64(req, keyLimit), Offset: i64(req, keyOffset)}
		if str(req, keyTimespan) != nil {
			r.Timespan = ptr(operations.ListVideoViewsTimespan(strVal(req, keyTimespan)))
		}
		return s.Views.ListVideoViews(ctx, r)
	case "get_video_view_details":
		return s.Views.GetDetails(ctx, strVal(req, keyViewID))
	case "list_by_top_content":
		ts := ptrStr[operations.ListByTopContentTimespan](req, keyTimespan)
		return s.Views.ListByTopContent(ctx, ts, nil, i64(req, keyLimit))
	case "list_dimensions":
		return s.Dimensions.List(ctx)
	case "list_filter_values_for_dimension":
		ts := ptrStr[operations.ListFilterValuesForDimensionTimespan](req, keyTimespan)
		return s.Dimensions.ListFilterValues(ctx, operations.DimensionsID(strVal(req, keyDimID)), ts, nil)
	case "list_breakdown_values":
		r := operations.ListBreakdownValuesRequest{
			MetricID: operations.ListBreakdownValuesMetricID(strVal(req, keyMetricID)),
			GroupBy:  str(req, keyGroupBy),
		}
		if str(req, keyTimespan) != nil {
			r.Timespan = ptr(operations.ListBreakdownValuesTimespan(strVal(req, keyTimespan)))
		}
		return s.Metrics.ListBreakdownValues(ctx, r)
	case "list_overall_values":
		ts := ptrStr[operations.ListOverallValuesTimespan](req, keyTimespan)
		return s.Metrics.ListOverallValues(ctx, operations.ListOverallValuesMetricID(strVal(req, keyMetricID)), nil, ts, nil)
	case "get_timeseries_data":
		r := operations.GetTimeseriesDataRequest{
			MetricID: operations.GetTimeseriesDataMetricID(strVal(req, keyMetricID)),
		}
		if str(req, keyTimespan) != nil {
			r.Timespan = ptr(operations.GetTimeseriesDataTimespan(strVal(req, keyTimespan)))
		}
		if str(req, keyGroupBy) != nil {
			r.GroupBy = ptr(operations.GroupBy(strVal(req, keyGroupBy)))
		}
		return s.Metrics.GetTimeseriesData(ctx, r)
	case "list_comparison_values":
		ts := ptrStr[operations.ListComparisonValuesTimespan](req, keyTimespan)
		var dim *operations.Dimension
		if str(req, keyDimension) != nil {
			dim = ptr(operations.Dimension(strVal(req, keyDimension)))
		}
		return s.Metrics.ListComparisonValues(ctx, ts, nil, dim, str(req, keyValue))
	case "list_errors":
		ts := ptrStr[operations.ListErrorsTimespan](req, keyTimespan)
		return s.Errors.List(ctx, ts, nil, i64(req, keyLimit))

	}
	return nil, fmt.Errorf(errNoSDKMapping, op)
}

func dispatchCreate(s *fastpixgo.Fastpixgo, op string, req map[string]any, sdkValidate string) (any, error) {
	ctx := context.Background()
	switch op {
	case "create-media":
		return s.InputVideo.Create(ctx, components.CreateMediaRequest{
			Inputs:   []components.Input{components.CreateInputPullVideoInput(components.PullVideoInput{})},
			Metadata: map[string]string{"source": sdkValidate},
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
			InputMediaSettings: components.InputMediaSettings{Metadata: map[string]string{"name": sdkValidate}},
		})
	case "create-media-playback-id":
		// Request an explicit resolution so the API returns a non-null
		// `resolution` in the response. Omitting it makes the API return
		// `resolution: null`, which the OpenAPI spec's enum does not model.
		res := operations.ResolutionOneThousandAndEightyp
		return s.Playback.Create(ctx, strVal(req, keyMediaID), &operations.CreateMediaPlaybackIDRequestBody{
			AccessPolicy: components.AccessPolicyPublic,
			Resolution:   &res,
		})
	case "Add-media-track":
		return s.Videos.AddMediaTrack(ctx, strVal(req, keyMediaID), operations.AddMediaTrackRequestBody{
			Tracks: components.AddTrackRequest{},
		})
	case "Generate-subtitle-track":
		return s.ManageVideos.GenerateSubtitleTrack(ctx, strVal(req, keyMediaID), strVal(req, keyTrackID), components.TrackSubtitlesGenerateRequest{})
	case "create-playbackId-of-stream":
		return s.LivePlayback.Create(ctx, strVal(req, keyStreamID), components.PlaybackIDRequest{})
	case "create-simulcast-of-stream":
		return s.SimulcastStreams.Create(ctx, strVal(req, keyStreamID), components.SimulcastRequest{
			URL:       fastpixgo.String("rtmp://example.com/live"),
			StreamKey: fastpixgo.String(fmt.Sprintf("sk-%d", os.Getpid())),
		})
	case "direct-upload-video-media":
		return s.InputVideo.DirectUploadMedia(ctx, &operations.DirectUploadVideoMediaRequest{
			PushMediaSettings: &operations.PushMediaSettings{Metadata: map[string]string{"source": "sdk-validate"}},
		})

		// --------------------------- PUT / PATCH (update) ---------------------------
	}
	return nil, fmt.Errorf(errNoSDKMapping, op)
}

func dispatchUpdate(s *fastpixgo.Fastpixgo, op string, req map[string]any) (any, error) {
	ctx := context.Background()
	switch op {
	case "updated-media":
		return s.Videos.Update(ctx, strVal(req, keyMediaID), operations.UpdatedMediaRequestBody{
			Metadata: map[string]string{"updated": "true"},
			Title:    fastpixgo.String("SDK Validate Title"),
		})
	case "updated-source-access":
		return s.Videos.UpdateSourceAccess(ctx, strVal(req, keyMediaID), operations.UpdatedSourceAccessRequestBody{
			SourceAccess: true,
		})
	case "updated-mp4Support":
		return s.Videos.UpdateMp4Support(ctx, strVal(req, keyMediaID), operations.UpdatedMp4SupportRequestBody{})
	case "update-media-summary":
		return s.InVideoAI.GenerateSummary(ctx, strVal(req, keyMediaID), operations.UpdateMediaSummaryRequestBody{Generate: true})
	case "update-media-chapters":
		return s.InVideoAIFeatures.UpdateChapters(ctx, strVal(req, keyMediaID), operations.UpdateMediaChaptersRequestBody{Chapters: fastpixgo.Bool(true)})
	case "update-media-named-entities":
		return s.InVideoAIFeatures.UpdateMediaNamedEntities(ctx, strVal(req, keyMediaID), operations.UpdateMediaNamedEntitiesRequestBody{NamedEntities: true})
	case "update-media-moderation":
		return s.InVideoAIFeatures.UpdateModeration(ctx, strVal(req, keyMediaID), operations.UpdateMediaModerationRequestBody{
			Moderation: &operations.UpdateMediaModerationModeration{},
		})
	case "update-media-track":
		return s.ManageVideos.UpdateTrack(ctx, strVal(req, keyTrackID), strVal(req, keyMediaID), components.UpdateTrackRequest{})
	case "update-domain-restrictions":
		return s.Playback.UpdateDomainRestrictions(ctx, strVal(req, keyMediaID), strVal(req, keyPlaybackID), operations.UpdateDomainRestrictionsRequestBody{
			Allow: []string{"example.com"},
		})
	case "update-user-agent-restrictions":
		return s.PlaybackIds.UpdateUserAgentRestrictions(ctx, strVal(req, keyMediaID), strVal(req, keyPlaybackID), operations.UpdateUserAgentRestrictionsRequestBody{
			Allow: []string{"Mozilla"},
		})
	case "update-a-playlist":
		return s.Playlists.Update(ctx, strVal(req, keyPlaylistID), components.UpdatePlaylistRequest{
			Name:        "SDK Validate Updated",
			Description: "updated by validator",
		})
	case "add-media-to-playlist":
		return s.Playlist.AddMedia(ctx, strVal(req, keyPlaylistID), components.MediaIdsRequest{MediaIds: []string{strVal(req, keyMediaID)}})
	case "change-media-order-in-playlist":
		return s.Playlist.ChangeMediaOrder(ctx, strVal(req, keyPlaylistID), components.MediaIdsRequest{MediaIds: []string{strVal(req, keyMediaID)}})
	case "update-live-stream":
		return s.ManageLiveStream.UpdateLiveStream(ctx, strVal(req, keyStreamID), components.PatchLiveStreamRequest{
			Metadata:        map[string]string{"updated": "true"},
			ReconnectWindow: fastpixgo.Int64(120),
		})
	case "update-specific-simulcast-of-stream":
		return s.SimulcastStreams.Update(ctx, strVal(req, keyStreamID), strVal(req, keySimulcastID), components.SimulcastUpdateRequest{
			IsEnabled: fastpixgo.Bool(false),
		})
	case "enable-live-stream":
		return s.ManageLiveStream.Enable(ctx, strVal(req, keyStreamID))
	case "disable-live-stream":
		return s.LiveStreams.Disable(ctx, strVal(req, keyStreamID))
	case "complete-live-stream":
		return s.ManageLiveStream.Complete(ctx, strVal(req, keyStreamID))
	case "cancel-upload":
		return s.ManageVideos.CancelUpload(ctx, strVal(req, keyUploadID))

		// ----------------------------- DELETE -----------------------------
	}
	return nil, fmt.Errorf(errNoSDKMapping, op)
}

func dispatchDelete(s *fastpixgo.Fastpixgo, op string, req map[string]any) (any, error) {
	ctx := context.Background()
	switch op {
	case "delete-media-from-playlist":
		return s.Playlists.DeleteMedia(ctx, strVal(req, keyPlaylistID), &components.MediaIdsRequest{MediaIds: []string{strVal(req, keyMediaID)}})
	case "delete-a-playlist":
		return s.Playlists.Delete(ctx, strVal(req, keyPlaylistID))
	case "delete-media-track":
		return s.Videos.DeleteTrack(ctx, strVal(req, keyMediaID), strVal(req, keyTrackID))
	case "delete-media-playback-id":
		return s.Playback.Delete(ctx, strVal(req, keyMediaID), strVal(req, keyPlaybackID))
	case "delete-simulcast-of-stream":
		return s.SimulcastStream.Delete(ctx, strVal(req, keyStreamID), strVal(req, keySimulcastID))
	case "delete-playbackId-of-stream":
		return s.LivePlayback.DeletePlaybackID(ctx, strVal(req, keyStreamID), strVal(req, keyPlaybackID))
	case "delete-live-stream":
		return s.LiveStreams.Delete(ctx, strVal(req, keyStreamID))
	case "delete-media":
		return s.ManageVideos.Delete(ctx, strVal(req, keyMediaID))
	case "delete_signing_key":
		return s.SigningKeys.Delete(ctx, strVal(req, keySigningKeyID))
	}
	return nil, fmt.Errorf(errNoSDKMapping, op)
}
