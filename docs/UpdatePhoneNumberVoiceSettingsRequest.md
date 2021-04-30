# UpdatePhoneNumberVoiceSettingsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallForwarding** | [***CallForwarding**](CallForwarding.md) |  | [optional] [default to null]
**CallRecording** | [***CallRecording**](CallRecording.md) |  | [optional] [default to null]
**CnamListing** | [***CnamListing**](CnamListing.md) |  | [optional] [default to null]
**MediaFeatures** | [***MediaFeatures**](MediaFeatures.md) |  | [optional] [default to null]
**TechPrefixEnabled** | **bool** | Controls whether a tech prefix is enabled for this phone number. | [optional] [default to false]
**TranslatedNumber** | **string** | This field allows you to rewrite the destination number of an inbound call before the call is routed to you. The value of this field may be any alphanumeric value, and the value will replace the number originally dialed. | [optional] [default to null]
**UsagePaymentMethod** | **string** | Controls whether a number is billed per minute or uses your concurrent channels. | [optional] [default to USAGE_PAYMENT_METHOD.PAY_PER_MINUTE]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

