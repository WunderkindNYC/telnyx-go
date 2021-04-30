# CreateConferenceRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | **string** | Unique identifier and token for controlling the call | [default to null]
**Name** | **string** | Name of the conference | [default to null]
**BeepEnabled** | **string** | Whether a beep sound should be played when participants join and/or leave the conference. | [optional] [default to null]
**ClientState** | **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] [default to null]
**CommandId** | **string** | Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same &#x60;command_id&#x60; as one that has already been executed. | [optional] [default to null]
**DurationMinutes** | **int32** | Time length (minutes) after which the conference will end. | [optional] [default to null]
**HoldAudioUrl** | **string** | The URL to an audio file to be played to participants joining the conference. Takes effect only when \&quot;start_conference_on_create\&quot; is set to \&quot;false\&quot;. | [optional] [default to null]
**StartConferenceOnCreate** | **bool** | Whether the conference should be started on creation. If the conference isn&#x27;t started all participants that join are automatically put on hold. Defaults to \&quot;true\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

