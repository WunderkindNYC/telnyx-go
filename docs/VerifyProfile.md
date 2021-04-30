# VerifyProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** |  | [default to null]
**DefaultTimeoutSecs** | **int32** | For every request that is initiated via this Verify profile, this sets the number of seconds before a verification request code expires. Once the verification request expires, the user cannot use the code to verify their identity. | [default to null]
**Id** | **string** |  | [default to null]
**MessagingEnabled** | **bool** | Enables SMS text messaging for the Verify profile. | [optional] [default to null]
**MessagingTemplate** | **string** | Optionally sets a messaging text template when sending the verification code. Uses &#x60;{code}&#x60; to template in the actual verification code. | [optional] [default to null]
**Name** | **string** | The name of the Verify profile. | [default to null]
**RcsEnabled** | **bool** | Enables RCS messaging for the Verify profile. | [default to null]
**RecordType** | **string** | Identifies the type of the resource. | [optional] [default to null]
**UpdatedAt** | **string** |  | [default to null]
**VsmsEnabled** | **bool** | Enables VSMS for the Verify profile. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

