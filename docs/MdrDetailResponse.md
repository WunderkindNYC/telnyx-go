# MdrDetailResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cld** | **string** | The destination number for a call, or the callee | [optional] [default to null]
**Cli** | **string** | The number associated with the person initiating the call, or the caller | [optional] [default to null]
**Cost** | **string** | Final cost. Cost is calculated as rate * parts | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Message sent time | [optional] [default to null]
**Currency** | **string** | Currency of the rate and cost | [optional] [default to null]
**Direction** | **string** | Direction of message - inbound or outbound. | [optional] [default to null]
**Id** | **string** | Id of message detail record | [optional] [default to null]
**MessageType** | **string** | Type of message | [optional] [default to null]
**Parts** | **float64** | Number of parts this message has. Max number of character is 160. If message contains more characters then that it will be broken down in multiple parts | [optional] [default to null]
**ProfileName** | **string** | Configured profile name. New profiles can be created and configured on Telnyx portal | [optional] [default to null]
**Rate** | **string** | Rate applied to the message | [optional] [default to null]
**RecordType** | **string** |  | [optional] [default to null]
**Status** | **string** | Message status | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

