# ConferenceParticipantPlaybackEndedPayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | **string** | Participant&#x27;s call ID used to issue commands via Call Control API. | [optional] [default to null]
**CallLegId** | **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] [default to null]
**CallSessionId** | **string** | ID that is unique to the call session and can be used to correlate webhook events. | [optional] [default to null]
**ClientState** | **string** | State received from a command. | [optional] [default to null]
**ConferenceId** | **string** | ID of the conference the text was spoken in. | [optional] [default to null]
**ConnectionId** | **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] [default to null]
**CreatorCallSessionId** | **string** | ID that is unique to the call session that started the conference. | [optional] [default to null]
**MediaUrl** | **string** | The URL to the audio file being played. | [optional] [default to null]
**OccurredAt** | [**time.Time**](time.Time.md) | ISO 8601 datetime of when the event occurred. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

