# CallRecordingSavedPayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallLegId** | **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] [default to null]
**CallSessionId** | **string** | ID that is unique to the call session and can be used to correlate webhook events. | [optional] [default to null]
**Channels** | **string** | Whether recording was recorded in &#x60;single&#x60; or &#x60;dual&#x60; channel. | [optional] [default to null]
**ClientState** | **string** | State received from a command. | [optional] [default to null]
**ConnectionId** | **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] [default to null]
**PublicRecordingUrls** | [***CallRecordingSavedPayloadPublicRecordingUrls**](CallRecordingSaved_payload_public_recording_urls.md) |  | [optional] [default to null]
**RecordingEndedAt** | [**time.Time**](time.Time.md) | ISO 8601 datetime of when recording ended. | [optional] [default to null]
**RecordingStartedAt** | [**time.Time**](time.Time.md) | ISO 8601 datetime of when recording started. | [optional] [default to null]
**RecordingUrls** | [***CallRecordingSavedPayloadRecordingUrls**](CallRecordingSaved_payload_recording_urls.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

