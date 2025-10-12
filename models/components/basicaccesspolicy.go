package components

import (
	"encoding/json"
	"fmt"
)

// BasicAccessPolicy - Basic access policy for media content
type BasicAccessPolicy string

const (
	BasicAccessPolicyPublic  BasicAccessPolicy = "public"
	BasicAccessPolicyPrivate BasicAccessPolicy = "private"
)

func (e BasicAccessPolicy) ToPointer() *BasicAccessPolicy {
	return &e
}
func (e *BasicAccessPolicy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "public":
		fallthrough
	case "private":
		*e = BasicAccessPolicy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BasicAccessPolicy: %v", v)
	}
}
