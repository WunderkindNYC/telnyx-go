# CreateManagedAccountRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessName** | **string** | The name of the business for which the new managed account is being created, that will be used as the managed accounts&#x27;s organization&#x27;s name. | [default to null]
**Email** | **string** | The email address for the managed account. If not provided, the email address will be generated based on the email address of the manager account. | [optional] [default to null]
**Password** | **string** | Password for the managed account. If a password is not supplied, the account will not be able to be signed into directly. (A password reset may still be performed later to enable sign-in via password.) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

