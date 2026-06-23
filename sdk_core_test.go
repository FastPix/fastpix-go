package fastpixgo_test

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/apierrors"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/retry"
)

const (
	testUser = "test-user"
	testPass = "test-pass"
)

// successBody is a minimal-but-valid JSON payload matching
// components.GetAllSigningKeysResponse.
const successBody = `{
	"success": true,
	"data": [
		{ "id": "key-123", "createdAt": "2026-06-23T00:00:00Z" }
	],
	"pagination": { "totalRecords": 1, "currentPage": 1, "perPage": 10 }
}`

func newTestClient(t *testing.T, serverURL string, opts ...fastpixgo.SDKOption) *fastpixgo.Fastpixgo {
	t.Helper()
	base := []fastpixgo.SDKOption{
		fastpixgo.WithServerURL(serverURL),
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.String(testUser),
			Password: fastpixgo.String(testPass),
		}),
	}
	return fastpixgo.New(append(base, opts...)...)
}

// TestSDKCore_Success verifies the happy path: a 200 with a valid JSON body is
// parsed into the typed response and reachable via the getters.
func TestSDKCore_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(successBody))
	}))
	defer srv.Close()

	client := newTestClient(t, srv.URL)

	limit := int64(10)
	offset := int64(1)
	res, err := client.SigningKeys.List(context.Background(), &limit, &offset)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil response")
	}
	if res.GetAllSigningKeysResponse == nil {
		t.Fatal("expected parsed GetAllSigningKeysResponse, got nil")
	}
	if res.HTTPMeta.Response == nil || res.HTTPMeta.Response.StatusCode != http.StatusOK {
		t.Fatalf("expected HTTPMeta with status 200, got %#v", res.HTTPMeta.Response)
	}
	if got := res.GetAllSigningKeysResponse.GetSuccess(); got == nil || !*got {
		t.Fatalf("expected success=true, got %v", got)
	}
	data := res.GetAllSigningKeysResponse.GetData()
	if len(data) != 1 {
		t.Fatalf("expected 1 data item, got %d", len(data))
	}
}

// TestSDKCore_AuthHeader verifies HTTP Basic auth is applied with the
// configured username/password.
func TestSDKCore_AuthHeader(t *testing.T) {
	var gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(successBody))
	}))
	defer srv.Close()

	client := newTestClient(t, srv.URL)
	if _, err := client.SigningKeys.List(context.Background(), nil, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotAuth == "" {
		t.Fatal("expected non-empty Authorization header")
	}
	if !strings.HasPrefix(gotAuth, "Basic ") {
		t.Fatalf("expected Basic auth, got %q", gotAuth)
	}
	wantToken := base64.StdEncoding.EncodeToString([]byte(testUser + ":" + testPass))
	if got := strings.TrimPrefix(gotAuth, "Basic "); got != wantToken {
		t.Fatalf("auth token mismatch: got %q want %q", got, wantToken)
	}
}

// TestSDKCore_RequestShape verifies the method and URL path of the issued
// request.
func TestSDKCore_RequestShape(t *testing.T) {
	var gotMethod, gotPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(successBody))
	}))
	defer srv.Close()

	client := newTestClient(t, srv.URL)
	if _, err := client.SigningKeys.List(context.Background(), nil, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodGet {
		t.Fatalf("expected GET, got %s", gotMethod)
	}
	if gotPath != "/iam/signing-keys" {
		t.Fatalf("expected path /iam/signing-keys, got %s", gotPath)
	}
}

// TestSDKCore_ErrorPath verifies 4xx responses surface an *apierrors.APIError
// with the right status code.
func TestSDKCore_ErrorPath(t *testing.T) {
	cases := []struct {
		name   string
		status int
	}{
		{"unauthorized", http.StatusUnauthorized},
		{"notFound", http.StatusNotFound},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tc.status)
				_, _ = w.Write([]byte(`{"success":false,"error":{"code":` +
					itoa(tc.status) + `,"message":"boom"}}`))
			}))
			defer srv.Close()

			client := newTestClient(t, srv.URL)
			res, err := client.SigningKeys.List(context.Background(), nil, nil)
			if err == nil {
				t.Fatal("expected an error, got nil")
			}
			if res != nil {
				t.Fatalf("expected nil response on error, got %#v", res)
			}

			var apiErr *apierrors.APIError
			if !errors.As(err, &apiErr) {
				t.Fatalf("expected *apierrors.APIError, got %T: %v", err, err)
			}
			if apiErr.StatusCode != tc.status {
				t.Fatalf("expected status %d, got %d", tc.status, apiErr.StatusCode)
			}
		})
	}
}

// TestSDKCore_Retry verifies that with a backoff retry config the SDK retries
// on a retryable 5xx and ultimately succeeds.
func TestSDKCore_Retry(t *testing.T) {
	var hits int32
	const failFirst = 2 // first two calls fail with 503, third succeeds

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&hits, 1)
		if n <= failFirst {
			w.WriteHeader(http.StatusServiceUnavailable) // 503 is retryable
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(successBody))
	}))
	defer srv.Close()

	retryCfg := retry.Config{
		Strategy: "backoff",
		Backoff: &retry.BackoffStrategy{
			InitialInterval: 1, // 1ms
			MaxInterval:     5, // 5ms
			Exponent:        1.1,
			MaxElapsedTime:  5000, // 5s, plenty for 3 fast attempts
		},
		RetryConnectionErrors: false,
	}

	client := newTestClient(t, srv.URL, fastpixgo.WithRetryConfig(retryCfg))

	res, err := client.SigningKeys.List(context.Background(), nil, nil)
	if err != nil {
		t.Fatalf("expected success after retries, got %v", err)
	}
	if res == nil || res.GetAllSigningKeysResponse == nil {
		t.Fatalf("expected parsed response after retries, got %#v", res)
	}

	got := atomic.LoadInt32(&hits)
	if got <= 1 {
		t.Fatalf("expected SDK to retry (hits > 1), got %d", got)
	}
	if int(got) != failFirst+1 {
		t.Fatalf("expected exactly %d hits, got %d", failFirst+1, got)
	}
}

// itoa is a tiny dependency-free int->string helper for building error bodies.
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
