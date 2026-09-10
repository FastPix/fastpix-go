# AccessRestrictions

Controls access based on domains and user agents. Defines a default policy (either "allow" or "deny") and provides lists for explicitly allowed or denied domains and user agents.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `Domains`                                                                             | [*components.DomainRestrictions](../../models/components/domainrestrictions.md)       | :heavy_minus_sign:                                                                    | Restrictions based on the originating domain of a request                             |
| `UserAgents`                                                                          | [*components.UserAgentRestrictions](../../models/components/useragentrestrictions.md) | :heavy_minus_sign:                                                                    | Restrictions based on the user agent                                                  |