# SpeakRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | **string** | The text or SSML to be converted into speech. There is a 5,000 character limit. | [default to null]
**PayloadType** | **string** | The type of the provided payload. The payload can either be plain text, or Speech Synthesis Markup Language (SSML). | [optional] [default to PAYLOAD_TYPE.TEXT]
**ServiceLevel** | **string** | This parameter impacts speech quality, language options and payload types. When using &#x60;basic&#x60;, only the &#x60;en-US&#x60; language and payload type &#x60;text&#x60; are allowed. | [optional] [default to SERVICE_LEVEL.PREMIUM]
**Stop** | **string** | When specified, it stops the current audio being played.  Specify &#x60;current&#x60; to stop the current audio being played, and to play the next file in the queue. Specify &#x60;all&#x60; to stop the current audio file being played and to also clear all audio files from the queue. | [optional] [default to null]
**Voice** | **string** | The gender of the voice used to speak back the text. | [default to null]
**Language** | **string** | The language you want spoken. | [default to null]
**ClientState** | **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] [default to null]
**CommandId** | **string** | Use this field to avoid duplicate commands. Telnyx will ignore commands with the same &#x60;command_id&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

