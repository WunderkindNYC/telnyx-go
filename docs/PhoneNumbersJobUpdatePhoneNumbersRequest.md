# PhoneNumbersJobUpdatePhoneNumbersRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingGroupId** | **string** | Identifies the billing group associated with the phone number. | [optional] [default to null]
**ConnectionId** | **string** | Identifies the connection associated with the phone number. | [optional] [default to null]
**CustomerReference** | **string** | A customer reference string for customer look ups. | [optional] [default to null]
**ExternalPin** | **string** | If someone attempts to port your phone number away from Telnyx and your phone number has an external PIN set, we will attempt to verify that you provided the correct external PIN to the winning carrier. Note that not all carriers cooperate with this security mechanism. | [optional] [default to null]
**PhoneNumbers** | **[]string** | Array of phone number ids and/or phone numbers in E164 format to update | [default to null]
**Tags** | **[]string** | A list of user-assigned tags to help organize phone numbers. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

