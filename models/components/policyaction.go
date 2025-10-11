

package components

import (
	"encoding/json"
	"fmt"
)

// PolicyAction - Policy action type
type PolicyAction string

const (
	PolicyActionAllow PolicyAction = "allow"
	PolicyActionDeny  PolicyAction = "deny"
)

func (e PolicyAction) ToPointer() *PolicyAction {
	return &e
}
func (e *PolicyAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "allow":
		fallthrough
	case "deny":
		*e = PolicyAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PolicyAction: %v", v)
	}
}
