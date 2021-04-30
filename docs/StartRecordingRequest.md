# StartRecordingRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** | The audio file format used when storing the call recording. Can be either &#x60;mp3&#x60; or &#x60;wav&#x60;. | [default to null]
**Channels** | **string** | When &#x60;dual&#x60;, final audio file will be stereo recorded with the first leg on channel A, and the rest on channel B. | [default to null]
**ClientState** | **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] [default to null]
**CommandId** | **string** | Use this field to avoid duplicate commands. Telnyx will ignore commands with the same &#x60;command_id&#x60;. | [optional] [default to null]
**PlayBeep** | **bool** | If enabled, a beep sound will be played at the start of a recording. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

