# UpdateNumberOrderRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | An ISO 8901 datetime string denoting when the number order was created. | [optional] [default to null]
**CustomerReference** | **string** | A customer reference string for customer look ups. | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**PhoneNumbers** | [**[]PhoneNumber**](PhoneNumber.md) |  | [optional] [default to null]
**PhoneNumbersCount** | **int32** | The count of phone numbers in the number order. | [optional] [default to null]
**RecordType** | **string** |  | [optional] [default to null]
**RequirementsMet** | **bool** | True if all requirements are met for every phone number, false otherwise. | [optional] [default to null]
**Status** | **string** | The status of the order. | [optional] [default to null]
**UpdatedAt** | **string** | An ISO 8901 datetime string for when the number order was updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

