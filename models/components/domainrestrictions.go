package components

// DomainRestrictions - Restrictions based on the originating domain of a request
type DomainRestrictions struct {
	// Policy action type
	DefaultPolicy *PolicyAction `json:"defaultPolicy,omitempty"`
	// A list of domain names or patterns that are explicitly allowed access
	Allow []string `json:"allow,omitempty"`
	// A list of domain names or patterns that are explicitly denied access
	Deny []string `json:"deny,omitempty"`
}

func (d *DomainRestrictions) GetDefaultPolicy() *PolicyAction {
	if d == nil {
		return nil
	}
	return d.DefaultPolicy
}

func (d *DomainRestrictions) GetAllow() []string {
	if d == nil {
		return nil
	}
	return d.Allow
}

func (d *DomainRestrictions) GetDeny() []string {
	if d == nil {
		return nil
	}
	return d.Deny
}
