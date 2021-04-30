# JoinAConferenceRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | **string** | Unique identifier and token for controlling the call | [default to null]
**ClientState** | **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] [default to null]
**CommandId** | **string** | Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same &#x60;command_id&#x60; as one that has already been executed. | [optional] [default to null]
**EndConferenceOnExit** | **bool** | Whether the conference should end and all remaining participant be hung up after the participant leaves the conference. Defaults to \&quot;false\&quot;. | [optional] [default to null]
**Hold** | **bool** | Whether the participant should be put on hold immediately after joining the conference. Defaults to \&quot;false\&quot;. | [optional] [default to null]
**HoldAudioUrl** | **string** | The URL of an audio file to be played to the participant when they are put on hold after joining the conference. This property takes effect only if \&quot;hold\&quot; is set to \&quot;true\&quot;. | [optional] [default to null]
**Mute** | **bool** | Whether the participant should be muted immediately after joining the conference. Defaults to \&quot;false\&quot;. | [optional] [default to null]
**StartConferenceOnEnter** | **bool** | Whether the conference should be started after the participant joins the conference. Defaults to \&quot;false\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

