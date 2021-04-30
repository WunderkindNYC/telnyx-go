# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCredentialAction**](BulkCredentialsApi.md#BulkCredentialAction) | **Post** /actions/{action}/telephony_credentials | Perform activate or deactivate action on all credentials filtered by the provided tag.
[**CreateBulkTelephonyCredential**](BulkCredentialsApi.md#CreateBulkTelephonyCredential) | **Post** /actions/bulk/telephony_credentials | Creates several credentials
[**DeleteBulkTelephonyCredential**](BulkCredentialsApi.md#DeleteBulkTelephonyCredential) | **Delete** /actions/bulk/telephony_credentials | Delete several credentials
[**UpdateBulkTelephonyCredential**](BulkCredentialsApi.md#UpdateBulkTelephonyCredential) | **Patch** /actions/bulk/telephony_credentials | Update several credentials

# **BulkCredentialAction**
> InlineResponse201 BulkCredentialAction(ctx, action, filterTag)
Perform activate or deactivate action on all credentials filtered by the provided tag.

Perform activate or deactivate action on all credentials filtered by the provided tag. Activate action will change the status to active, making it possible to connect calls with the credential. Deactivate action will change the status to inactive, making it impossible to connect calls with the credential.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **action** | **string**| Identifies the action to be taken. Activate will change the status to active. Deactivate will change the status to inactive. | 
  **filterTag** | **string**| Filter by tag, required by bulk operations. | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBulkTelephonyCredential**
> InlineResponse201 CreateBulkTelephonyCredential(ctx, body)
Creates several credentials

Creates several credentials in bulk.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkCredentialRequest**](BulkCredentialRequest.md)| Requested parameters to create credentials on bulk | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBulkTelephonyCredential**
> InlineResponse201 DeleteBulkTelephonyCredential(ctx, filterTag)
Delete several credentials

Delete several credentials in bulk.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filterTag** | **string**| Filter by tag, required by bulk operations. | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBulkTelephonyCredential**
> InlineResponse201 UpdateBulkTelephonyCredential(ctx, body, filterTag)
Update several credentials

Update several credentials in bulk.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkCredentialRequest**](BulkCredentialRequest.md)| Parameters to update credentials on bulk | 
  **filterTag** | **string**| Filter by tag, required by bulk operations. | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

