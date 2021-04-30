# ConferenceCallRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | **string** | Unique identifier and token for controlling the call | [default to null]
**ClientState** | **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] [default to null]
**CommandId** | **string** | Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same &#x60;command_id&#x60; as one that has already been executed. | [optional] [default to null]
**From** | **string** | The &#x60;from&#x60; number to be used as the caller id presented to the destination (&#x60;to&#x60; number). | [default to null]
**Hold** | **bool** | Whether the participant should be put on hold immediately after joining the conference. | [optional] [default to false]
**HoldAudioUrl** | **string** | The URL of an audio file to be played to the participant when they are put on hold after joining the conference. This property takes effect only if \&quot;hold\&quot; is set to \&quot;true\&quot;. | [optional] [default to null]
**Mute** | **bool** | Whether the participant should be muted immediately after joining the conference. | [optional] [default to false]
**StartConferenceOnEnter** | **bool** | Whether the conference should be started after the participant joins the conference. | [optional] [default to false]
**SupervisorRole** | **string** | Sets the joining participant as a supervisor for the conference. A conference can have multiple supervisors. \&quot;barge\&quot; means the supervisor enters the conference as a normal participant. This is the same as \&quot;none\&quot;. \&quot;monitor\&quot; means the supervisor is muted but can hear all participants. \&quot;whisper\&quot; means that only the specified \&quot;whisper_call_control_ids\&quot; can hear the supervisor. Defaults to \&quot;none\&quot;. | [optional] [default to null]
**To** | **string** | The DID or SIP URI to dial out and bridge to the given call. | [default to null]
**WhisperCallControlIds** | **[]string** | Array of unique call_control_ids the joining supervisor can whisper to. If none provided, the supervisor will join the conference as a monitoring participant only. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

