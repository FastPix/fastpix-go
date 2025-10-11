

package components

import (
	"encoding/json"
	"fmt"
)

// AccessPolicy - Access policy for media content
type AccessPolicy string

const (
	AccessPolicyPublic  AccessPolicy = "public"
	AccessPolicyPrivate AccessPolicy = "private"
	AccessPolicyDrm     AccessPolicy = "drm"
)

func (e AccessPolicy) ToPointer() *AccessPolicy {
	return &e
}
func (e *AccessPolicy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "public":
		fallthrough
	case "private":
		fallthrough
	case "drm":
		*e = AccessPolicy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccessPolicy: %v", v)
	}
}
