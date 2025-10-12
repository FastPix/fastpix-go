package components

import (
	"encoding/json"
	"fmt"
)

// PlaylistOrder - Determines the insertion order of media into playlist.
type PlaylistOrder string

const (
	PlaylistOrderCreatedDateAsc  PlaylistOrder = "createdDate ASC"
	PlaylistOrderCreatedDateDesc PlaylistOrder = "createdDate DESC"
)

func (e PlaylistOrder) ToPointer() *PlaylistOrder {
	return &e
}
func (e *PlaylistOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "createdDate ASC":
		fallthrough
	case "createdDate DESC":
		*e = PlaylistOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlaylistOrder: %v", v)
	}
}
