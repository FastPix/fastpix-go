# DomainRestrictions

Restrictions based on the originating domain of a request


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `DefaultPolicy`                                                       | [*components.PolicyAction](../../models/components/policyaction.md)   | :heavy_minus_sign:                                                    | Policy action type                                                    |
| `Allow`                                                               | []*string*                                                            | :heavy_minus_sign:                                                    | A list of domain names or patterns that are explicitly allowed access |
| `Deny`                                                                | []*string*                                                            | :heavy_minus_sign:                                                    | A list of domain names or patterns that are explicitly denied access  |