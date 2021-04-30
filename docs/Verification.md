# Verification

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** |  | [default to null]
**Id** | **string** |  | [default to null]
**PhoneNumber** | **string** | +E164 formatted phone number. | [default to null]
**RecordType** | **string** | Identifies the type of the resource. | [optional] [default to null]
**Status** | [***VerificationStatus**](VerificationStatus.md) |  | [default to null]
**TimeoutSecs** | **int32** | This is the number of seconds before the code of the request is expired. Once this request has expired, the code will no longer verify the user. Note: this will override the &#x60;default_timeout_secs&#x60; on the Verify profile. | [optional] [default to null]
**Type_** | [***VerificationType**](VerificationType.md) |  | [default to null]
**UpdatedAt** | **string** |  | [default to null]
**VerifyProfileId** | **string** | The identifier of the associated Verify profile. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

