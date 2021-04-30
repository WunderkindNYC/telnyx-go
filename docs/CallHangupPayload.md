# CallHangupPayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | **string** | Call ID used to issue commands via Call Control API. | [optional] [default to null]
**CallLegId** | **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] [default to null]
**CallSessionId** | **string** | ID that is unique to the call session and can be used to correlate webhook events. | [optional] [default to null]
**ClientState** | **string** | State received from a command. | [optional] [default to null]
**ConnectionId** | **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] [default to null]
**From** | **string** | Number or SIP URI placing the call. | [optional] [default to null]
**HangupCause** | **string** | The reason the call was ended (&#x60;call_rejected&#x60;, &#x60;normal_clearing&#x60;, &#x60;originator_cancel&#x60;, &#x60;timeout&#x60;, &#x60;time_limit&#x60;, &#x60;user_busy&#x60;, &#x60;not_found&#x60; or &#x60;unspecified&#x60;). | [optional] [default to null]
**HangupSource** | **string** | The party who ended the call (&#x60;callee&#x60;, &#x60;caller&#x60;, &#x60;unknown&#x60;). | [optional] [default to null]
**SipHangupCause** | **string** | The reason the call was ended (SIP response code). If the SIP response is unavailable (in inbound calls for example) this is set to &#x60;unspecified&#x60;. | [optional] [default to null]
**StartTime** | [**time.Time**](time.Time.md) | ISO 8601 datetime of when the call started. | [optional] [default to null]
**State** | **string** | State received from a command. | [optional] [default to null]
**To** | **string** | Destination number or SIP URI of the call. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

