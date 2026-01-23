# UnusedUploadsPlaybackIDDomains

Restrictions based on the originating domain of a request (for example, whether requests from certain websites must be allowed or blocked).


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `DefaultPolicy`                                                            | [*components.PolicyAction](../../models/components/policyaction.md)        | :heavy_minus_sign:                                                         | Policy action type                                                         |
| `Allow`                                                                    | []*string*                                                                 | :heavy_minus_sign:                                                         | A list of domains that are explicitly allowed access.                      |
| `Deny`                                                                     | []*string*                                                                 | :heavy_minus_sign:                                                         | A list of domains that are explicitly blocked from accessing the resource. |