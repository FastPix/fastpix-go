// Verify a FastPix webhook signature before trusting the payload.
//
// FastPix signs the raw request body with your webhook Signing Secret
// (Dashboard > Webhooks) and sends it as a Base64 HMAC-SHA256 in the
// "FastPix-Signature" header. The Signing Secret is itself Base64-encoded, so
// sign with its decoded bytes as the key. Verify the body exactly as received:
// parsing and re-serializing changes the bytes and the signature won't match.
//
// Run:
//
//	go run ./examples/verify-webhook
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
)

// isValidSignature reports whether sig is a valid FastPix-Signature for rawBody.
func isValidSignature(rawBody []byte, sig, secret string) bool {
	if secret == "" || sig == "" {
		return false
	}
	key, err := base64.StdEncoding.DecodeString(secret) // Signing Secret is Base64
	if err != nil {
		return false
	}
	mac := hmac.New(sha256.New, key)
	mac.Write(rawBody)
	expected := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expected), []byte(sig)) // constant-time compare
}

// In an HTTP handler you would read the raw body and the header, then verify:
//
//	func handler(w http.ResponseWriter, r *http.Request) {
//		raw, _ := io.ReadAll(r.Body)
//		if !isValidSignature(raw, r.Header.Get("FastPix-Signature"), os.Getenv("FASTPIX_WEBHOOK_SECRET")) {
//			http.Error(w, "invalid signature", http.StatusUnauthorized)
//			return
//		}
//		// ... decode raw and handle the event ...
//	}

func main() {
	secret := os.Getenv("FASTPIX_WEBHOOK_SECRET")
	if secret == "" {
		secret = base64.StdEncoding.EncodeToString([]byte("demo-secret"))
	}

	// Demo: build a signature the way FastPix would, then verify it.
	rawBody := []byte(`{"type":"video.media.ready","data":{"id":"abc-123"}}`)
	key, _ := base64.StdEncoding.DecodeString(secret)
	mac := hmac.New(sha256.New, key)
	mac.Write(rawBody)
	sig := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	if isValidSignature(rawBody, sig, secret) {
		fmt.Println("verified")
	} else {
		fmt.Println("rejected")
	}
}
