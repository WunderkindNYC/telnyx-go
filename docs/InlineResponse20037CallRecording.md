# InlineResponse20037CallRecording

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallRecordingType** | **string** | Specifies which calls are recorded. | [optional] [default to null]
**CallRecordingCallerPhoneNumbers** | **[]string** | When call_recording_type is &#x27;by_caller_phone_number&#x27;, only outbound calls using one of these numbers will be recorded. Numbers must be specified in E164 format. | [optional] [default to null]
**CallRecordingChannels** | **string** | When using &#x27;dual&#x27; channels, the final audio file will be a stereo recording with the first leg on channel A, and the rest on channel B. | [optional] [default to CALL_RECORDING_CHANNELS.SINGLE]
**CallRecordingFormat** | **string** | The audio file format for calls being recorded. | [optional] [default to CALL_RECORDING_FORMAT.WAV]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

