# CallPlaybackStartedPayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | **string** | Call ID used to issue commands via Call Control API. | [optional] [default to null]
**CallLegId** | **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] [default to null]
**CallSessionId** | **string** | ID that is unique to the call session and can be used to correlate webhook events. | [optional] [default to null]
**ClientState** | **string** | State received from a command. | [optional] [default to null]
**ConnectionId** | **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] [default to null]
**MediaUrl** | **string** | The audio URL being played back. | [optional] [default to null]
**Overlay** | **bool** | Whether the audio is going to be played in overlay mode or not. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

