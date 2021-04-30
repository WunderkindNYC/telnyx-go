# SendDtmfRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientState** | **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] [default to null]
**CommandId** | **string** | Use this field to avoid duplicate commands. Telnyx will ignore commands with the same &#x60;command_id&#x60;. | [optional] [default to null]
**Digits** | **string** | DTMF digits to send. Valid digits are 0-9, A-D, *, and #. Pauses can be added using w (0.5s) and W (1s). | [default to null]
**DurationMillis** | **int32** | Specifies for how many milliseconds each digit will be played in the audio stream. Ranges from 100 to 500ms | [optional] [default to 250]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

