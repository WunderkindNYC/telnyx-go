# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTelephonyCredentialToken**](AccessTokensApi.md#CreateTelephonyCredentialToken) | **Post** /telephony_credentials/{id}/token | Create an Access Token.

# **CreateTelephonyCredentialToken**
> string CreateTelephonyCredentialToken(ctx, id)
Create an Access Token.

Create an Access Token (JWT) for the credential.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

