# UserAgentRestrictions

Restrictions based on the user agent


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `DefaultPolicy`                                                     | [*components.PolicyAction](../../models/components/policyaction.md) | :heavy_minus_sign:                                                  | Policy action type                                                  |
| `Allow`                                                             | []*string*                                                          | :heavy_minus_sign:                                                  | A list of user agents that are explicitly allowed access            |
| `Deny`                                                              | []*string*                                                          | :heavy_minus_sign:                                                  | A list of user agents that are explicitly denied access             |