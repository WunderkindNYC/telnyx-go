# InlineResponse20045Data

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifies the type of resource. | [optional] [default to null]
**RecordType** | **string** | Identifies the type of the resource. | [optional] [default to null]
**PhoneNumber** | **string** | The phone number in +E164 format. | [optional] [default to null]
**ConnectionId** | **string** | Identifies the connection associated with this phone number. | [optional] [default to null]
**TechPrefixEnabled** | **bool** | Controls whether a tech prefix is enabled for this phone number. | [optional] [default to false]
**TranslatedNumber** | **string** | This field allows you to rewrite the destination number of an inbound call before the call is routed to you. The value of this field may be any alphanumeric value, and the value will replace the number originally dialed. | [optional] 
**CallForwarding** | [***PhoneNumbersidvoiceCallForwarding**](phone_numbersidvoice_call_forwarding.md) |  | [optional] [default to null]
**CnamListing** | [***PhoneNumbersidvoiceCnamListing**](phone_numbersidvoice_cnam_listing.md) |  | [optional] [default to null]
**Emergency** | [***InlineResponse20045Emergency**](inline_response_200_45_emergency.md) |  | [optional] [default to null]
**UsagePaymentMethod** | **string** | Controls whether a number is billed per minute or uses your concurrent channels. | [optional] [default to USAGE_PAYMENT_METHOD.PAY_PER_MINUTE]
**MediaFeatures** | [***PhoneNumbersidvoiceMediaFeatures**](phone_numbersidvoice_media_features.md) |  | [optional] [default to null]
**CallRecording** | [***PhoneNumbersidvoiceCallRecording**](phone_numbersidvoice_call_recording.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

