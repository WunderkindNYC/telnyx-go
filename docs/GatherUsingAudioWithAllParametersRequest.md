# GatherUsingAudioWithAllParametersRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioUrl** | **string** | The URL of a file to be played back at the beginning of each prompt. The URL can point to either a WAV or MP3 file. | [default to null]
**InvalidAudioUrl** | **string** | The URL of a file to play when digits don&#x27;t match the &#x60;valid_digits&#x60; parameter or the number of digits is not between &#x60;min&#x60; and &#x60;max&#x60;. The URL can point to either a WAV or MP3 file. | [optional] [default to null]
**MinimumDigits** | **int32** | The minimum number of digits to fetch. This parameter has a minimum value of 1. | [optional] [default to 1]
**MaximumDigits** | **int32** | The maximum number of digits to fetch. This parameter has a maximum value of 128. | [optional] [default to 128]
**MaximumTries** | **int32** | The maximum number of times the file should be played if there is no input from the user on the call. | [optional] [default to 3]
**TimeoutMillis** | **int32** | The number of milliseconds to wait for a DTMF response after file playback ends before a replaying the sound file. | [optional] [default to 60000]
**TerminatingDigit** | **string** | The digit used to terminate input if fewer than &#x60;maximum_digits&#x60; digits have been gathered. | [optional] [default to #]
**ValidDigits** | **string** | A list of all digits accepted as valid. | [optional] [default to 0123456789#*]
**InterDigitTimeoutMillis** | **int32** | The number of milliseconds to wait for input between digits. | [optional] [default to 5000]
**ClientState** | **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] [default to null]
**CommandId** | **string** | Use this field to avoid duplicate commands. Telnyx will ignore commands with the same &#x60;command_id&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

