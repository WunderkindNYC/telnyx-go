# Body25

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**RecordType** | **string** |  | [optional] [default to null]
**PhoneNumbers** | [**[]NumberOrdersPhoneNumbers**](number_orders_phone_numbers.md) |  | [optional] [default to null]
**PhoneNumbersCount** | **int32** | The count of phone numbers in the number order | [optional] [default to null]
**Status** | **string** | The status of the order | [optional] [default to null]
**CustomerReference** | **string** | A customer reference string for customer look ups | [optional] [default to null]
**ConnectionId** | **string** | The connection id to set the number to upon acquiring the number | [optional] [default to null]
**MessagingProfileId** | **string** | The messaging profile id to set the number to upon acquiring the number | [optional] [default to null]
**WebhookUrl** | **string** | A webhook URL for number order status updates | [optional] [default to null]
**WebhookFailoverUrl** | **string** | If above webhook URL fails, this URL will be used as a fail over | [optional] [default to null]
**CreatedAt** | **string** | An ISO 8901 datetime string denoting when the number order was created | [optional] [default to null]
**UpdatedAt** | **string** | An ISO 8901 datetime string for when the number order was updated | [optional] [default to null]
**RequirementsMet** | **bool** | True if all requirements are met for every phone number, false otherwise | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

