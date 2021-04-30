# PlayAudioUrlRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioUrl** | **string** | The URL of the file to be played back on the call. The URL can point to either a WAV or MP3 file. | [default to null]
**Loop** | **string** | The number of times the audio file should be played. If supplied, the value must be an integer between 1 and 100, or the special string &#x60;infinity&#x60; for an endless loop. | [optional] [default to 1]
**Overlay** | **bool** | When enabled, audio will be mixed on top of any other audio that is actively being played back. Note that &#x60;overlay: true&#x60; will only work if there is another audio file already being played on the call. | [optional] [default to false]
**Stop** | **string** | When specified, it stops the current audio being played.  Specify &#x60;current&#x60; to stop the current audio being played, and to play the next file in the queue. Specify &#x60;all&#x60; to stop the current audio file being played and to also clear all audio files from the queue. | [optional] [default to null]
**TargetLegs** | **string** | Specifies the leg or legs on which audio will be played. If supplied, the value must be either &#x60;self&#x60;, &#x60;opposite&#x60; or &#x60;both&#x60;. | [optional] [default to self]
**ClientState** | **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] [default to null]
**CommandId** | **string** | Use this field to avoid duplicate commands. Telnyx will ignore commands with the same &#x60;command_id&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

