# Subtitles

Generates subtitle files for audio/video files.



## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   | Example                                                                       |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `LanguageName`                                                                | **string*                                                                     | :heavy_minus_sign:                                                            | Name of the language for the subtitles.                                       | english                                                                       |
| `Metadata`                                                                    | [*operations.SubtitlesMetadata](../../models/operations/subtitlesmetadata.md) | :heavy_minus_sign:                                                            | Searchable metadata tags for the video in key-value pairs.                    | {<br/>"key1": "value1"<br/>}                                                  |
| `LanguageCode`                                                                | [*operations.LanguageCode](../../models/operations/languagecode.md)           | :heavy_minus_sign:                                                            | Language codes (BCP 47 compliant) used for text files.<br/>                   | en                                                                            |