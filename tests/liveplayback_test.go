package tests

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/apierrors"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/models/operations"
)

type capturedRequest struct {
	method, path, contentType, body string
}

func newRestrictionServer(t *testing.T, status int, respBody string, got *capturedRequest) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		*got = capturedRequest{method: r.Method, path: r.URL.Path, contentType: r.Header.Get("Content-Type"), body: string(b)}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(respBody))
	}))
}

func newOfflineClient(url string) *fastpixgo.Fastpixgo {
	return fastpixgo.New(
		fastpixgo.WithServerURL(url),
		fastpixgo.WithSecurity(components.Security{Username: fastpixgo.String("u"), Password: fastpixgo.String("p")}),
	)
}

func TestLivePlaybackUpdateDomainRestrictions(t *testing.T) {
	var got capturedRequest
	srv := newRestrictionServer(t, 200, `{"success":true,"data":{"defaultPolicy":"allow","allow":["yourdomain.com"],"deny":["malicioussite.io"]}}`, &got)
	defer srv.Close()

	res, err := newOfflineClient(srv.URL).LivePlayback.UpdateDomainRestrictions(context.Background(), "s1", "p1",
		operations.UpdateLiveStreamDomainRestrictionsRequestBody{Allow: []string{"yourdomain.com"}, Deny: []string{"malicioussite.io"}})
	if err != nil {
		t.Fatal(err)
	}
	if got.method != http.MethodPatch || got.path != "/live/streams/s1/playback-ids/p1/domains" {
		t.Fatalf("got %s %s", got.method, got.path)
	}
	if got.contentType != "application/json" {
		t.Fatalf("content type %q", got.contentType)
	}
	if got.body != `{"defaultPolicy":"allow","allow":["yourdomain.com"],"deny":["malicioussite.io"]}` {
		t.Fatalf("body %s", got.body)
	}
	if res.Object == nil || res.Object.Data == nil || *res.Object.Data.DefaultPolicy != "allow" || res.Object.Data.Deny[0] != "malicioussite.io" {
		t.Fatalf("response %+v", res.Object)
	}
}

func TestLivePlaybackUpdateUserAgentRestrictions(t *testing.T) {
	var got capturedRequest
	srv := newRestrictionServer(t, 200, `{"success":true,"data":{"defaultPolicy":"deny","allow":["PostmanRuntime/7.29.0"],"deny":[]}}`, &got)
	defer srv.Close()

	res, err := newOfflineClient(srv.URL).LivePlayback.UpdateUserAgentRestrictions(context.Background(), "s1", "p1",
		operations.UpdateLiveStreamUserAgentRestrictionsRequestBody{
			DefaultPolicy: operations.UpdateLiveStreamUserAgentRestrictionsDefaultPolicyDeny.ToPointer(),
			Allow:         []string{"PostmanRuntime/7.29.0"},
		})
	if err != nil {
		t.Fatal(err)
	}
	if got.method != http.MethodPatch || got.path != "/live/streams/s1/playback-ids/p1/user-agents" {
		t.Fatalf("got %s %s", got.method, got.path)
	}
	if got.body != `{"defaultPolicy":"deny","allow":["PostmanRuntime/7.29.0"]}` {
		t.Fatalf("body %s", got.body)
	}
	if res.Object == nil || res.Object.Data == nil || res.Object.Data.Allow[0] != "PostmanRuntime/7.29.0" {
		t.Fatalf("response %+v", res.Object)
	}
}

func TestLivePlaybackRestrictionErrors(t *testing.T) {
	for _, status := range []int{404, 500} {
		var got capturedRequest
		srv := newRestrictionServer(t, status, `{"success":false,"error":{"code":`+http.StatusText(status)+`}}`, &got)
		res, err := newOfflineClient(srv.URL).LivePlayback.UpdateDomainRestrictions(context.Background(), "s1", "p1",
			operations.UpdateLiveStreamDomainRestrictionsRequestBody{})
		srv.Close()
		var apiErr *apierrors.APIError
		if res != nil || !errors.As(err, &apiErr) || apiErr.StatusCode != status {
			t.Fatalf("status %d: expected APIError, got res=%v err=%v", status, res, err)
		}
	}
}
