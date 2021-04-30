# BridgeRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | **string** | The Call Control ID of the call you want to bridge with. | [default to null]
**ClientState** | **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] [default to null]
**CommandId** | **string** | Use this field to avoid duplicate commands. Telnyx will ignore commands with the same &#x60;command_id&#x60;. | [optional] [default to null]
**ParkAfterUnbridge** | **string** | Specifies behavior after the bridge ends (i.e. the opposite leg either hangs up or is transferred). If supplied with the value &#x60;self&#x60;, the current leg will be parked after unbridge. If not set, the default behavior is to hang up the leg. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

