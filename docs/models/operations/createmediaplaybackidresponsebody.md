# CreateMediaPlaybackIDResponseBody

Playback ID for a media content.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Success`                                                                    | **bool*                                                                      | :heavy_minus_sign:                                                           | Shows the request status. Returns true for success and false for failure.    | true                                                                         |
| `Data`                                                                       | [*components.CreatePlaybackID](../../models/components/createplaybackid.md)  | :heavy_minus_sign:                                                           | A collection of Playback ID objects utilized for crafting HLS playback urls. |                                                                              |