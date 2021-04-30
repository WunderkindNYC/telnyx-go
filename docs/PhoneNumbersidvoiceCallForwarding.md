# PhoneNumbersidvoiceCallForwarding

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallForwardingEnabled** | **bool** | Indicates if call forwarding will be enabled for this number if forwards_to and forwarding_type are filled in. Defaults to true for backwards compatibility with APIV1 use of numbers endpoints. | [optional] [default to true]
**ForwardsTo** | **string** | The phone number to which inbound calls to this number are forwarded. Inbound calls will not be forwarded if this field is left blank. If set, must be a +E.164-formatted phone number. | [optional] [default to null]
**ForwardingType** | **string** | Call forwarding type. &#x27;forwards_to&#x27; must be set for this to have an effect. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

