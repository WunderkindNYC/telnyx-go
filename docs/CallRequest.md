# CallRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnsweringMachineDetection** | **string** | Enables Answering Machine Detection. When a call is answered, Telnyx runs real-time detection to determine if it was picked up by a human or a machine and sends an &#x60;call.machine.detection.ended&#x60; webhook with the analysis result. If &#x27;greeting_end&#x27; or &#x27;detect_words&#x27; is used and a &#x27;machine&#x27; is detected, you will receive another &#x27;call.machine.greeting.ended&#x27; webhook when the answering machine greeting ends with a beep or silence. If &#x60;detect_beep&#x60; is used, you will only receive &#x27;call.machine.greeting.ended&#x27; if a beep is detected. | [optional] [default to ANSWERING_MACHINE_DETECTION.DISABLED]
**AnsweringMachineDetectionConfig** | [***CallRequestAnsweringMachineDetectionConfig**](CallRequest_answering_machine_detection_config.md) |  | [optional] [default to null]
**AudioUrl** | **string** | The URL of a file to be played back to the callee when the call is answered. The URL can point to either a WAV or MP3 file. | [optional] [default to null]
**BillingGroupId** | **string** | Use this field to set the Billing Group ID for the call. Must be a valid and existing Billing Group ID. | [optional] [default to null]
**ClientState** | **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] [default to null]
**CommandId** | **string** | Use this field to avoid duplicate commands. Telnyx will ignore commands with the same &#x60;command_id&#x60;. | [optional] [default to null]
**ConnectionId** | **string** | The ID of the Call Control App (formerly ID of the connection) to be used when dialing the destination. | [default to null]
**CustomHeaders** | [**[]CustomSipHeader**](CustomSipHeader.md) | Custom headers to be added to the SIP INVITE. | [optional] [default to null]
**From** | **string** | The &#x60;from&#x60; number to be used as the caller id presented to the destination (&#x60;to&#x60; number). The number should be in +E164 format. This attribute will default to the &#x60;from&#x60; number of the original call if omitted. | [default to null]
**FromDisplayName** | **string** | The &#x60;from_display_name&#x60; string to be used as the caller id name (SIP From Display Name) presented to the destination (&#x60;to&#x60; number). The string should have a maximum of 128 characters, containing only letters, numbers, spaces, and -_~!.+ special characters. If ommited, the display name will be the same as the number in the &#x60;from&#x60; field. | [optional] [default to null]
**LinkTo** | **string** | Use another call&#x27;s control id for sharing the same call session id | [optional] [default to null]
**SipAuthPassword** | **string** | SIP Authentication password used for SIP challenges. | [optional] [default to null]
**SipAuthUsername** | **string** | SIP Authentication username used for SIP challenges. | [optional] [default to null]
**TimeLimitSecs** | **int32** | Sets the maximum duration of a Call Control Leg in seconds. If the time limit is reached, the call will hangup and a &#x60;call.hangup&#x60; webhook with a &#x60;hangup_cause&#x60; of &#x60;time_limit&#x60; will be sent. For example, by setting a time limit of 120 seconds, a Call Leg will be automatically terminated two minutes after being answered. The default time limit is 14400 seconds or 4 hours and this is also the maximum allowed call length. | [optional] [default to 14400]
**TimeoutSecs** | **int32** | The number of seconds that Telnyx will wait for the call to be answered by the destination to which it is being called. If the timeout is reached before an answer is received, the call will hangup and a &#x60;call.hangup&#x60; webhook with a &#x60;hangup_cause&#x60; of &#x60;timeout&#x60; will be sent. Minimum value is 5 seconds. Maximum value is 120 seconds. | [optional] [default to 30]
**To** | **string** | The DID or SIP URI to dial out to. | [default to null]
**WebhookUrl** | **string** | Use this field to override the URL for which Telnyx will send subsequent webhooks to for this call. | [optional] [default to null]
**WebhookUrlMethod** | **string** | HTTP request type used for &#x60;webhook_url&#x60;. | [optional] [default to WEBHOOK_URL_METHOD.POST]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

