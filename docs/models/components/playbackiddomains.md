# PlaybackIDDomains

Restrictions based on the originating domain of a request (e.g., whether requests from certain websites should be allowed or blocked).


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `DefaultPolicy`                                                            | [*components.PolicyAction](../../models/components/policyaction.md)        | :heavy_minus_sign:                                                         | Policy action type                                                         |
| `Allow`                                                                    | []*string*                                                                 | :heavy_minus_sign:                                                         | A list of domains that are explicitly allowed access.                      |
| `Deny`                                                                     | []*string*                                                                 | :heavy_minus_sign:                                                         | A list of domains that are explicitly blocked from accessing the resource. |