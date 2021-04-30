# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVerification**](VerifyApi.md#CreateVerification) | **Post** /verifications | Trigger a verification
[**CreateVerifyProfile**](VerifyApi.md#CreateVerifyProfile) | **Post** /verify_profiles | Create a Verify profile
[**DeleteVerifyProfile**](VerifyApi.md#DeleteVerifyProfile) | **Delete** /verify_profiles/{verify_profile_id} | Delete a Verify profile
[**ListVerifications**](VerifyApi.md#ListVerifications) | **Get** /verifications/by_phone_number/{phone_number} | List verifications by phone number
[**ListVerifyProfiles**](VerifyApi.md#ListVerifyProfiles) | **Get** /verify_profiles | List all Verify profiles
[**RetrieveVerification**](VerifyApi.md#RetrieveVerification) | **Get** /verifications/{verification_id} | Retrieve a verification
[**RetrieveVerifyProfile**](VerifyApi.md#RetrieveVerifyProfile) | **Get** /verify_profiles/{verify_profile_id} | Retrieve a Verify profile
[**UpdateVerifyProfile**](VerifyApi.md#UpdateVerifyProfile) | **Patch** /verify_profiles/{verify_profile_id} | Update a Verify profile
[**VerifyVerificationCode**](VerifyApi.md#VerifyVerificationCode) | **Post** /verifications/by_phone_number/{phone_number}/actions/verify | Submit a verification code

# **CreateVerification**
> CreateVerificationResponse CreateVerification(ctx, body)
Trigger a verification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVerificationRequest**](CreateVerificationRequest.md)|  | 

### Return type

[**CreateVerificationResponse**](CreateVerificationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVerifyProfile**
> VerifyProfileResponse CreateVerifyProfile(ctx, body)
Create a Verify profile

Creates a new Verify profile to associate verifications with.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVerifyProfileRequest**](CreateVerifyProfileRequest.md)|  | 

### Return type

[**VerifyProfileResponse**](VerifyProfileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVerifyProfile**
> VerifyProfileResponse DeleteVerifyProfile(ctx, verifyProfileId)
Delete a Verify profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **verifyProfileId** | [**string**](.md)| The identifier of the Verify profile to delete. | 

### Return type

[**VerifyProfileResponse**](VerifyProfileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVerifications**
> ListVerificationsResponse ListVerifications(ctx, phoneNumber)
List verifications by phone number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **phoneNumber** | **string**| The phone number associated with the verifications to retrieve. | 

### Return type

[**ListVerificationsResponse**](ListVerificationsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVerifyProfiles**
> ListVerifyProfilesResponse ListVerifyProfiles(ctx, optional)
List all Verify profiles

Gets a paginated list of Verify profiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VerifyApiListVerifyProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VerifyApiListVerifyProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterName** | **optional.String**|  | 
 **pageSize** | **optional.Int32**|  | [default to 25]
 **pageNumber** | **optional.Int32**|  | [default to 1]

### Return type

[**ListVerifyProfilesResponse**](ListVerifyProfilesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveVerification**
> RetrieveVerificationResponse RetrieveVerification(ctx, verificationId)
Retrieve a verification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **verificationId** | [**string**](.md)| The identifier of the verification to retrieve. | 

### Return type

[**RetrieveVerificationResponse**](RetrieveVerificationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveVerifyProfile**
> VerifyProfileResponse RetrieveVerifyProfile(ctx, verifyProfileId)
Retrieve a Verify profile

Gets a single Verify profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **verifyProfileId** | [**string**](.md)| The identifier of the Verify profile to retrieve. | 

### Return type

[**VerifyProfileResponse**](VerifyProfileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVerifyProfile**
> VerifyProfileResponse UpdateVerifyProfile(ctx, body, verifyProfileId)
Update a Verify profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVerifyProfileRequest**](UpdateVerifyProfileRequest.md)|  | 
  **verifyProfileId** | [**string**](.md)| The identifier of the Verify profile to update. | 

### Return type

[**VerifyProfileResponse**](VerifyProfileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyVerificationCode**
> VerifyVerificationCodeResponse VerifyVerificationCode(ctx, body, phoneNumber)
Submit a verification code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VerifyVerificationCodeRequest**](VerifyVerificationCodeRequest.md)|  | 
  **phoneNumber** | **string**| The phone number associated with the verification code being verified. | 

### Return type

[**VerifyVerificationCodeResponse**](VerifyVerificationCodeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

