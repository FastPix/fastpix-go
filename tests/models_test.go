package tests

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/models/operations"
)

// durationOf returns a fresh model plus an accessor for its Duration, one per media-bearing schema.
var durationModels = map[string]func() (any, func() *float64){
	"Media":                  func() (any, func() *float64) { m := &components.Media{}; return m, m.GetDuration },
	"GetAllMediaResponse":    func() (any, func() *float64) { m := &components.GetAllMediaResponse{}; return m, m.GetDuration },
	"GetMediaDetailResponse": func() (any, func() *float64) { m := &components.GetMediaDetailResponse{}; return m, m.GetDuration },
	"UpdateMedia":            func() (any, func() *float64) { m := &components.UpdateMedia{}; return m, m.GetDuration },
	"SourceAccessMedia":      func() (any, func() *float64) { m := &components.SourceAccessMedia{}; return m, m.GetDuration },
	"LiveMediaClips":         func() (any, func() *float64) { m := &components.LiveMediaClips{}; return m, m.GetDuration },
	"MediaClipResponseData":  func() (any, func() *float64) { m := &components.MediaClipResponseData{}; return m, m.GetDuration },
	"PlaylistByIDResponseMediaListItem": func() (any, func() *float64) {
		m := &components.PlaylistByIDResponseMediaListItem{}
		return m, m.GetDuration
	},
}

func TestDurationIsFloatSeconds(t *testing.T) {
	for name, mk := range durationModels {
		t.Run(name, func(t *testing.T) {
			m, get := mk()
			if err := json.Unmarshal([]byte(`{"duration": 145.821315}`), m); err != nil {
				t.Fatal(err)
			}
			if d := get(); d == nil || *d != 145.821315 {
				t.Fatalf("fractional: got %v", d)
			}

			m, get = mk()
			if err := json.Unmarshal([]byte(`{"duration": 10}`), m); err != nil {
				t.Fatal(err)
			}
			if d := get(); d == nil || *d != 10.0 {
				t.Fatalf("integer: got %v", d)
			}

			m, get = mk()
			if err := json.Unmarshal([]byte(`{}`), m); err != nil {
				t.Fatal(err)
			}
			if get() != nil {
				t.Fatal("absent duration should be nil")
			}

			m, _ = mk()
			err := json.Unmarshal([]byte(`{"duration": "00:02:25"}`), m)
			var typeErr *json.UnmarshalTypeError
			if !errors.As(err, &typeErr) || typeErr.Type == nil || typeErr.Type.Kind() != reflect.Float64 {
				t.Fatalf("legacy string should fail with UnmarshalTypeError into float64, got %v", err)
			}

			m, _ = mk()
			_ = json.Unmarshal([]byte(`{"duration": 145.82}`), m)
			out, err := json.Marshal(m)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(out), `"duration":145.82`) {
				t.Fatalf("round-trip should carry a JSON number, got %s", out)
			}
		})
	}
}

func TestEnableRecording(t *testing.T) {
	out, err := json.Marshal(components.CreateLiveStreamRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), `"enableRecording":true`) {
		t.Fatalf("default should serialize enableRecording=true, got %s", out)
	}

	var s components.InputMediaSettings
	if err := json.Unmarshal([]byte(`{"enableRecording": false}`), &s); err != nil {
		t.Fatal(err)
	}
	if s.GetEnableRecording() == nil || *s.GetEnableRecording() {
		t.Fatal("false should round-trip")
	}
	out, _ = json.Marshal(s)
	if !strings.Contains(string(out), `"enableRecording":false`) {
		t.Fatalf("got %s", out)
	}
}

const accessRestrictionsJSON = `{"domains":{"defaultPolicy":"deny","allow":["example.com"],"deny":[]},"userAgents":{"defaultPolicy":"allow","allow":[],"deny":[]}}`

func TestAccessRestrictionsWireShape(t *testing.T) {
	var ar components.AccessRestrictions
	if err := json.Unmarshal([]byte(accessRestrictionsJSON), &ar); err != nil {
		t.Fatal(err)
	}
	if ar.GetDomains().GetDefaultPolicy() == nil || *ar.Domains.DefaultPolicy != components.PolicyActionDeny || ar.Domains.Allow[0] != "example.com" {
		t.Fatalf("domains parsed wrong: %+v", ar.Domains)
	}
	if ar.GetUserAgents().GetDefaultPolicy() == nil || *ar.UserAgents.DefaultPolicy != components.PolicyActionAllow {
		t.Fatalf("userAgents parsed wrong: %+v", ar.UserAgents)
	}

	for name, v := range map[string]any{
		"PlaybackIDRequest": components.PlaybackIDRequest{AccessRestrictions: &ar},
		"PlaybackSettings":  components.PlaybackSettings{AccessRestrictions: &ar},
	} {
		out, err := json.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		for _, key := range []string{`"accessRestrictions":{`, `"domains":{`, `"userAgents":{`, `"defaultPolicy":"deny"`} {
			if !strings.Contains(string(out), key) {
				t.Fatalf("%s: missing %s in %s", name, key, out)
			}
		}
	}

	out, _ := json.Marshal(components.PlaybackIDRequest{})
	if strings.Contains(string(out), "accessRestrictions") {
		t.Fatalf("absent restrictions must be omitted, got %s", out)
	}

	var withData components.PlaybackIDSuccessResponse
	if err := json.Unmarshal([]byte(`{"success":true,"data":{"id":"p1","accessPolicy":"public","accessRestrictions":`+accessRestrictionsJSON+`}}`), &withData); err != nil {
		t.Fatal(err)
	}
	if withData.Data.GetAccessRestrictions().GetDomains().Allow[0] != "example.com" {
		t.Fatal("response data.accessRestrictions not parsed")
	}
	var withoutData components.PlaybackIDSuccessResponse
	if err := json.Unmarshal([]byte(`{"success":true,"data":{"id":"p1","accessPolicy":"public"}}`), &withoutData); err != nil {
		t.Fatal(err)
	}
	if withoutData.Data.GetAccessRestrictions() != nil {
		t.Fatal("absent data.accessRestrictions should be nil")
	}

	var item components.PlaybackIDResponse
	if err := json.Unmarshal([]byte(`{"id":"p1","accessRestrictions":`+accessRestrictionsJSON+`}`), &item); err != nil {
		t.Fatal(err)
	}
	if item.GetAccessRestrictions() == nil {
		t.Fatal("PlaybackIDResponse.accessRestrictions not parsed")
	}
}

func TestLiveRestrictionOperationModels(t *testing.T) {
	out, err := json.Marshal(operations.UpdateLiveStreamDomainRestrictionsRequestBody{Allow: []string{"yourdomain.com"}})
	if err != nil {
		t.Fatal(err)
	}
	if s := string(out); !strings.Contains(s, `"defaultPolicy":"allow"`) || !strings.Contains(s, `"allow":["yourdomain.com"]`) || strings.Contains(s, "deny") || strings.Contains(s, "domains") {
		t.Fatalf("domain body must be flat with default policy allow and no unset lists, got %s", s)
	}
	out, _ = json.Marshal(operations.UpdateLiveStreamUserAgentRestrictionsRequestBody{
		DefaultPolicy: operations.UpdateLiveStreamUserAgentRestrictionsDefaultPolicyDeny.ToPointer(),
		Deny:          []string{"PostmanRuntime/7.29.0"},
	})
	if s := string(out); s != `{"defaultPolicy":"deny","deny":["PostmanRuntime/7.29.0"]}` {
		t.Fatalf("user-agent body: got %s", s)
	}

	req := operations.UpdateLiveStreamDomainRestrictionsRequest{StreamID: "s1", PlaybackID: "p1"}
	if req.GetStreamID() != "s1" || req.GetPlaybackID() != "p1" {
		t.Fatal("request must carry streamId and playbackId")
	}

	var body operations.UpdateLiveStreamUserAgentRestrictionsResponseBody
	if err := json.Unmarshal([]byte(`{"success":true,"data":{"defaultPolicy":"allow","allow":["a"],"deny":["b"]}}`), &body); err != nil {
		t.Fatal(err)
	}
	if body.GetSuccess() == nil || !*body.Success || *body.GetData().GetDefaultPolicy() != "allow" || body.Data.Allow[0] != "a" || body.Data.Deny[0] != "b" {
		t.Fatalf("response body parsed wrong: %+v", body)
	}
}
