# CallInitiatedPayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | **string** | Call ID used to issue commands via Call Control API. | [optional] [default to null]
**CallLegId** | **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] [default to null]
**CallSessionId** | **string** | ID that is unique to the call session and can be used to correlate webhook events. | [optional] [default to null]
**ClientState** | **string** | State received from a command. | [optional] [default to null]
**ConnectionId** | **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] [default to null]
**Direction** | **string** | Whether the call is &#x60;incoming&#x60; or &#x60;outgoing&#x60;. | [optional] [default to null]
**From** | **string** | Number or SIP URI placing the call. | [optional] [default to null]
**State** | **string** | State received from a command. | [optional] [default to null]
**To** | **string** | Destination number or SIP URI of the call. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
