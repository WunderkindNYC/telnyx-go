# GatherUsingSpeakRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientState** | **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] [default to null]
**CommandId** | **string** | Use this field to avoid duplicate commands. Telnyx will ignore commands with the same &#x60;command_id&#x60;. | [optional] [default to null]
**InterDigitTimeoutMillis** | **int32** | The number of milliseconds to wait for input between digits. | [optional] [default to 5000]
**InvalidPayload** | **string** | The text or SSML to be converted into speech when digits don&#x27;t match the &#x60;valid_digits&#x60; parameter or the number of digits is not between &#x60;min&#x60; and &#x60;max&#x60;. There is a 5,000 character limit. | [optional] [default to null]
**Language** | **string** | The language you want spoken. | [default to null]
**MaximumDigits** | **int32** | The maximum number of digits to fetch. This parameter has a maximum value of 128. | [optional] [default to 128]
**MaximumTries** | **int32** | The maximum number of times that a file should be played back if there is no input from the user on the call. | [optional] [default to 3]
**MinimumDigits** | **int32** | The minimum number of digits to fetch. This parameter has a minimum value of 1. | [optional] [default to 1]
**Payload** | **string** | The text or SSML to be converted into speech. There is a 5,000 character limit. | [default to null]
**PayloadType** | **string** | The type of the provided payload. The payload can either be plain text, or Speech Synthesis Markup Language (SSML). | [optional] [default to PAYLOAD_TYPE.TEXT]
**ServiceLevel** | **string** | This parameter impacts speech quality, language options and payload types. When using &#x60;basic&#x60;, only the &#x60;en-US&#x60; language and payload type &#x60;text&#x60; are allowed. | [optional] [default to SERVICE_LEVEL.PREMIUM]
**TerminatingDigit** | **string** | The digit used to terminate input if fewer than &#x60;maximum_digits&#x60; digits have been gathered. | [optional] [default to #]
**TimeoutMillis** | **int32** | The number of milliseconds to wait for a DTMF response after speak ends before a replaying the sound file. | [optional] [default to 60000]
**ValidDigits** | **string** | A list of all digits accepted as valid. | [optional] [default to 0123456789#*]
**Voice** | **string** | The gender of the voice used to speak back the text. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

