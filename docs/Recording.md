# Recording

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallLegId** | **string** | If a call recording, uniquely identifies the recorded call leg | [optional] [default to null]
**CallSessionId** | **string** | If a call recording, uniquely identifies the recorded call session | [optional] [default to null]
**Channels** | **string** | When &#x60;dual&#x60;, final audio file has the first leg on channel A, and the rest on channel B. | [default to null]
**ConferenceId** | **string** | If a conference recording, uniquely identifies the recorded conference | [optional] [default to null]
**CreatedAt** | **string** | ISO 8601 formatted date of when the recording was created | [default to null]
**DownloadUrls** | [***RecordingDownloadUrls**](Recording_download_urls.md) |  | [default to null]
**DurationMillis** | **int32** | The duration of the recording in milliseconds | [default to null]
**Id** | **string** | Uniquely identifies the recording | [default to null]
**RecordType** | **string** |  | [default to null]
**RecordingEndedAt** | **string** | ISO 8601 formatted date of when the recording ended | [default to null]
**RecordingStartedAt** | **string** | ISO 8601 formatted date of when the recording started | [default to null]
**Source** | **string** | The kind of event that led to this recording being created | [default to null]
**Status** | **string** | The status of the recording. Only resources for &#x60;completed&#x60; recordings are currently supported | [default to null]
**UpdatedAt** | **string** | ISO 8601 formatted date of when the recording was last updated | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

