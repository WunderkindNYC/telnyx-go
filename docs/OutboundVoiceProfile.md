# OutboundVoiceProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingGroupId** | **string** | The ID of the billing group associated with the outbound proflile. Defaults to null (for no group assigned). | [optional] [default to null]
**CallRecording** | [***OutboundCallRecording**](OutboundCallRecording.md) |  | [optional] [default to null]
**ConcurrentCallLimit** | **int32** | Must be no more than your global concurrent call limit. Null means no limit. | [optional] [default to null]
**ConnectionsCount** | **int32** | Amount of connections associated with this outbound voice profile. | [optional] [default to null]
**CreatedAt** | **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [default to null]
**DailySpendLimit** | **string** | The maximum amount of usage charges, in USD, you want Telnyx to allow on this outbound voice profile in a day before disallowing new calls. | [optional] [default to null]
**DailySpendLimitEnabled** | **bool** | Specifies whether to enforce the daily_spend_limit on this outbound voice profile. | [optional] [default to false]
**Enabled** | **bool** | Specifies whether the outbound voice profile can be used. Disabled profiles will result in outbound calls being blocked for the associated Connections. | [optional] [default to true]
**Id** | **string** | Identifies the resource. | [optional] [default to null]
**MaxDestinationRate** | **float64** | Maximum rate (price per minute) for a Destination to be allowed when making outbound calls. | [optional] [default to null]
**Name** | **string** | A user-supplied name to help with organization. | [default to null]
**RecordType** | **string** | Identifies the type of the resource. | [optional] [default to null]
**ServicePlan** | [***ServicePlan**](ServicePlan.md) |  | [optional] [default to null]
**Tags** | **[]string** |  | [optional] [default to null]
**TrafficType** | [***TrafficType**](TrafficType.md) |  | [optional] [default to null]
**UpdatedAt** | **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [default to null]
**UsagePaymentMethod** | [***UsagePaymentMethod**](UsagePaymentMethod.md) |  | [optional] [default to null]
**WhitelistedDestinations** | **[]string** | The list of destinations you want to be able to call using this outbound voice profile formatted in alpha2. | [optional] [default to ["US","CA"]]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

