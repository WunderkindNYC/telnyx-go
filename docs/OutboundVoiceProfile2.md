# OutboundVoiceProfile2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifies the resource. | [optional] [default to null]
**RecordType** | **string** | Identifies the type of the resource. | [optional] [default to null]
**Name** | **string** | A user-supplied name to help with organization. | [optional] 
**ConnectionsCount** | **int32** | Amount of connections associated with this outbound voice profile. | [optional] [default to null]
**TrafficType** | **string** | Specifies the type of traffic allowed in this profile. | [optional] [default to TRAFFIC_TYPE.CONVERSATIONAL]
**ServicePlan** | **string** |  | [optional] [default to SERVICE_PLAN.US]
**ConcurrentCallLimit** | **int32** | Must be no more than your global concurrent call limit. Null means no limit. | [optional] [default to null]
**Enabled** | **bool** | Specifies whether the outbound voice profile can be used. Disabled profiles will result in outbound calls being blocked for the associated Connections. | [optional] [default to true]
**Tags** | **[]string** |  | [optional] [default to null]
**UsagePaymentMethod** | **string** |  | [optional] [default to null]
**WhitelistedDestinations** | **[]string** | The list of destinations you want to be able to call using this outbound voice profile formatted in alpha2. | [optional] [default to null]
**MaxDestinationRate** | **float64** | Maximum rate (price per minute) for a Destination to be allowed when making outbound calls. | [optional] [default to null]
**DailySpendLimit** | **string** | The maximum amount of usage charges, in USD, you want Telnyx to allow on this outbound voice profile in a day before disallowing new calls. | [optional] [default to null]
**DailySpendLimitEnabled** | **bool** | Specifies whether to enforce the daily_spend_limit on this outbound voice profile. | [optional] [default to false]
**CallRecording** | [***InlineResponse20037CallRecording**](inline_response_200_37_call_recording.md) |  | [optional] [default to null]
**BillingGroupId** | **string** | The ID of the billing group associated with the outbound proflile. Defaults to null (for no group assigned). | [optional] [default to null]
**CreatedAt** | **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [default to null]
**UpdatedAt** | **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

