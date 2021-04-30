# PlaybackStopRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientState** | **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] [default to null]
**CommandId** | **string** | Use this field to avoid duplicate commands. Telnyx will ignore commands with the same &#x60;command_id&#x60;. | [optional] [default to null]
**Stop** | **string** | Use &#x60;current&#x60; to stop only the current audio or &#x60;all&#x60; to stop all audios in the queue. | [optional] [default to all]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

