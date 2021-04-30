# ConferenceSpeakRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlIds** | **[]string** | Call Control IDs of participants who will hear the spoken text. When empty all participants will hear the spoken text. | [optional] [default to null]
**CommandId** | **string** | Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same &#x60;command_id&#x60; as one that has already been executed. | [optional] [default to null]
**Language** | **string** | The language used to speak the text. | [default to null]
**Payload** | **string** | The text or SSML to be converted into speech. There is a 5,000 character limit. | [default to null]
**PayloadType** | **string** | The type of the provided payload. The payload can either be plain text, or Speech Synthesis Markup Language (SSML). | [optional] [default to PAYLOAD_TYPE.TEXT]
**Voice** | **string** | The gender of the voice used to speak the text. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

