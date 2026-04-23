// Hook for the get_video_view_details operation.
//
// The live FastPix analytics API emits abbreviated keys for event elements
// ("pt", "e", "d", "vt", "et", and inside "d": "br", "h", "w", ...). The
// OpenAPI spec — and therefore the generated models — use fully spelled-out
// names. This AfterSuccess hook rewrites the response body in place so the
// downstream generated deserializer sees spec-shaped JSON.
//
// Field order from the wire is preserved end-to-end: the hook walks the
// response with json.Decoder (which yields tokens in input order) and
// re-emits each object as an ordered slice of {key,value} pairs rather
// than routing through map[string]any (which would sort keys alphabetically
// on re-marshal). Only the subtrees that actually need rewriting
// (`data.events[]` elements, and `event_details` inside each event) are
// touched; every other byte passes through as json.RawMessage, preserving
// formatting-insensitive content exactly.
//
// This file is NOT generated. It is registered in registration.go, which the
// Speakeasy generator explicitly marks as free-to-modify.

package hooks

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

const eventsFieldRemapOperationID = "get_video_view_details"

var eventsOuterKeyMap = map[string]string{
	"pt": "player_playhead_time",
	"e":  "event_name",
	"d":  "event_details",
	"vt": "viewer_time",
	"et": "event_time",
}

var eventDetailsInnerKeyMap = map[string]string{
	"br":   "bitrate",
	"h":    "height",
	"w":    "width",
	"cd":   "codec",
	"host": "hostName",
	"txt":  "text",
	"c":    "code",
	"err":  "error",
	"t":    "type",
	"u":    "url",
}

type eventsFieldRemapHook struct{}

var _ afterSuccessHook = (*eventsFieldRemapHook)(nil)

// orderedPair preserves the original key/value order of a JSON object while
// leaving the value as raw bytes for byte-identical passthrough.
type orderedPair struct {
	key string
	val json.RawMessage
}

func (h *eventsFieldRemapHook) AfterSuccess(hookCtx AfterSuccessContext, res *http.Response) (*http.Response, error) {
	if hookCtx.OperationID != eventsFieldRemapOperationID {
		return res, nil
	}
	if res == nil || res.StatusCode < 200 || res.StatusCode >= 300 {
		return res, nil
	}
	if !isJSONContentType(res.Header.Get("Content-Type")) {
		return res, nil
	}
	if res.Body == nil {
		return res, nil
	}

	original, err := io.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return res, nil
	}

	restore := func() *http.Response {
		res.Body = io.NopCloser(bytes.NewReader(original))
		res.ContentLength = int64(len(original))
		return res
	}

	rewritten, ok := remapTopLevel(original)
	if !ok {
		return restore(), nil
	}

	res.Body = io.NopCloser(bytes.NewReader(rewritten))
	res.ContentLength = int64(len(rewritten))
	return res, nil
}

// remapTopLevel parses the full response body as an ordered object, replaces
// the `data.events[]` subtree with a remapped version, and re-emits the
// result preserving the original key order everywhere else.
func remapTopLevel(body []byte) ([]byte, bool) {
	topPairs, err := decodeOrderedObjectBytes(body)
	if err != nil {
		return nil, false
	}

	dataIdx := indexOfKey(topPairs, "data")
	if dataIdx < 0 {
		return nil, false
	}

	dataPairs, err := decodeOrderedObjectBytes(topPairs[dataIdx].val)
	if err != nil {
		return nil, false
	}

	eventsIdx := indexOfKey(dataPairs, "events")
	if eventsIdx < 0 {
		return nil, false
	}

	var events []json.RawMessage
	if err := json.Unmarshal(dataPairs[eventsIdx].val, &events); err != nil {
		return nil, false
	}

	for i, rawEvent := range events {
		newEvent, ok := remapEventOrdered(rawEvent)
		if !ok {
			// Non-object or malformed event — leave untouched.
			continue
		}
		events[i] = newEvent
	}

	eventsBytes, err := marshalRawArray(events)
	if err != nil {
		return nil, false
	}
	dataPairs[eventsIdx].val = eventsBytes

	dataBytes, err := encodeOrderedObject(dataPairs)
	if err != nil {
		return nil, false
	}
	topPairs[dataIdx].val = dataBytes

	out, err := encodeOrderedObject(topPairs)
	if err != nil {
		return nil, false
	}
	return out, true
}

// remapEventOrdered rebuilds one event object, preserving the wire order of
// its keys. If a key maps (per the OUTER map) to `event_details` and the
// value is an object, the inner keys are also remapped per the INNER map,
// again preserving order.
func remapEventOrdered(raw json.RawMessage) (json.RawMessage, bool) {
	pairs, err := decodeOrderedObjectBytes(raw)
	if err != nil {
		return nil, false
	}
	for i := range pairs {
		if mapped, ok := eventsOuterKeyMap[pairs[i].key]; ok {
			pairs[i].key = mapped
		}
		if pairs[i].key == "event_details" {
			if inner, ok := remapEventDetailsOrdered(pairs[i].val); ok {
				pairs[i].val = inner
			}
		}
	}
	out, err := encodeOrderedObject(pairs)
	if err != nil {
		return nil, false
	}
	return out, true
}

// remapEventDetailsOrdered rebuilds an `event_details` object, preserving
// key order. Non-object values (null/string/number/array) are rejected so
// the caller keeps the original bytes.
func remapEventDetailsOrdered(raw json.RawMessage) (json.RawMessage, bool) {
	pairs, err := decodeOrderedObjectBytes(raw)
	if err != nil {
		return nil, false
	}
	for i := range pairs {
		if mapped, ok := eventDetailsInnerKeyMap[pairs[i].key]; ok {
			pairs[i].key = mapped
		}
	}
	out, err := encodeOrderedObject(pairs)
	if err != nil {
		return nil, false
	}
	return out, true
}

// decodeOrderedObjectBytes parses a JSON object into an ordered slice of
// {key,value} pairs. Values stay as raw bytes. Returns an error if the
// input is not a JSON object.
func decodeOrderedObjectBytes(b []byte) ([]orderedPair, error) {
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()

	openTok, err := dec.Token()
	if err != nil {
		return nil, err
	}
	if d, ok := openTok.(json.Delim); !ok || d != '{' {
		return nil, errors.New("expected JSON object")
	}

	var out []orderedPair
	for dec.More() {
		keyTok, err := dec.Token()
		if err != nil {
			return nil, err
		}
		key, ok := keyTok.(string)
		if !ok {
			return nil, errors.New("expected string key")
		}
		var val json.RawMessage
		if err := dec.Decode(&val); err != nil {
			return nil, err
		}
		out = append(out, orderedPair{key: key, val: val})
	}

	closeTok, err := dec.Token()
	if err != nil {
		return nil, err
	}
	if d, ok := closeTok.(json.Delim); !ok || d != '}' {
		return nil, errors.New("expected closing brace")
	}
	return out, nil
}

// encodeOrderedObject re-emits an ordered object in compact form. The raw
// values are written through verbatim, so any pre-existing formatting
// inside them is preserved. Consumers that want indentation can wrap the
// output with json.Indent.
func encodeOrderedObject(pairs []orderedPair) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, p := range pairs {
		if i > 0 {
			buf.WriteByte(',')
		}
		keyBytes, err := json.Marshal(p.key)
		if err != nil {
			return nil, err
		}
		buf.Write(keyBytes)
		buf.WriteByte(':')
		if len(p.val) == 0 {
			buf.WriteString("null")
		} else {
			buf.Write(p.val)
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// marshalRawArray emits a JSON array of pre-serialized element bytes
// without round-tripping through map/slice reflection that would alter
// their internal key order.
func marshalRawArray(elements []json.RawMessage) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, e := range elements {
		if i > 0 {
			buf.WriteByte(',')
		}
		if len(e) == 0 {
			buf.WriteString("null")
		} else {
			buf.Write(e)
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func indexOfKey(pairs []orderedPair, key string) int {
	for i, p := range pairs {
		if p.key == key {
			return i
		}
	}
	return -1
}

func isJSONContentType(ct string) bool {
	if ct == "" {
		return false
	}
	return strings.HasPrefix(strings.ToLower(strings.TrimSpace(ct)), "application/json")
}
