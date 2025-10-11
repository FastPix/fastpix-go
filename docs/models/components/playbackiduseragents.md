# PlaybackIDUserAgents

Restrictions based on the user agent (which is typically a string sent by browsers or bots identifying themselves).


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `DefaultPolicy`                                                         | [*components.PolicyAction](../../models/components/policyaction.md)     | :heavy_minus_sign:                                                      | Policy action type                                                      |
| `Allow`                                                                 | []*string*                                                              | :heavy_minus_sign:                                                      | A list of specific user agents that are allowed to access the resource. |
| `Deny`                                                                  | []*string*                                                              | :heavy_minus_sign:                                                      | A list of specific user agents that are blocked.                        |