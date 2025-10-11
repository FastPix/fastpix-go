

package components

// UserAgentRestrictions - Restrictions based on the user agent
type UserAgentRestrictions struct {
	// Policy action type
	DefaultPolicy *PolicyAction `json:"defaultPolicy,omitempty"`
	// A list of user agents that are explicitly allowed access
	Allow []string `json:"allow,omitempty"`
	// A list of user agents that are explicitly denied access
	Deny []string `json:"deny,omitempty"`
}

func (u *UserAgentRestrictions) GetDefaultPolicy() *PolicyAction {
	if u == nil {
		return nil
	}
	return u.DefaultPolicy
}

func (u *UserAgentRestrictions) GetAllow() []string {
	if u == nil {
		return nil
	}
	return u.Allow
}

func (u *UserAgentRestrictions) GetDeny() []string {
	if u == nil {
		return nil
	}
	return u.Deny
}
