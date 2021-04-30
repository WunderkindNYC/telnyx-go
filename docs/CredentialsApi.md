# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTelephonyCredential**](CredentialsApi.md#CreateTelephonyCredential) | **Post** /telephony_credentials | Create a credential
[**DeleteTelephonyCredential**](CredentialsApi.md#DeleteTelephonyCredential) | **Delete** /telephony_credentials/{id} | Delete a credential
[**FindTelephonyCredentials**](CredentialsApi.md#FindTelephonyCredentials) | **Get** /telephony_credentials | List all credentials
[**GetTelephonyCredential**](CredentialsApi.md#GetTelephonyCredential) | **Get** /telephony_credentials/{id} | Get a credential
[**ListTags**](CredentialsApi.md#ListTags) | **Get** /telephony_credentials/tags | List all tags
[**TelephonyCredentialAction**](CredentialsApi.md#TelephonyCredentialAction) | **Post** /telephony_credentials/{id}/actions/{action} | Perform activate or deactivate action on provided Credential
[**UpdateTelephonyCredential**](CredentialsApi.md#UpdateTelephonyCredential) | **Patch** /telephony_credentials/{id} | Update a credential

# **CreateTelephonyCredential**
> InlineResponse2018 CreateTelephonyCredential(ctx, body)
Create a credential

Create a credential.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TelephonyCredentialCreate**](TelephonyCredentialCreate.md)| Parameters that can be defined during credential creation | 

### Return type

[**InlineResponse2018**](inline_response_201_8.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTelephonyCredential**
> InlineResponse2018 DeleteTelephonyCredential(ctx, id)
Delete a credential

Delete an existing credential.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse2018**](inline_response_201_8.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTelephonyCredentials**
> InlineResponse20042 FindTelephonyCredentials(ctx, optional)
List all credentials

List all On-demand Credentials.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CredentialsApiFindTelephonyCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiFindTelephonyCredentialsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterTag** | **optional.String**| Filter by tag | 
 **filterName** | **optional.String**| Filter by name | 
 **filterStatus** | **optional.String**| Filter by status | 
 **filterResourceId** | **optional.String**| Filter by resource_id | 
 **filterSipUsername** | **optional.String**| Filter by sip_username | 

### Return type

[**InlineResponse20042**](inline_response_200_42.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTelephonyCredential**
> InlineResponse2018 GetTelephonyCredential(ctx, id)
Get a credential

Get the details of an existing On-demand Credential.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse2018**](inline_response_201_8.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTags**
> InlineResponse20043 ListTags(ctx, optional)
List all tags

Returns a list of tags used on Credentials

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CredentialsApiListTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiListTagsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**InlineResponse20043**](inline_response_200_43.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TelephonyCredentialAction**
> InlineResponse2018 TelephonyCredentialAction(ctx, id, action)
Perform activate or deactivate action on provided Credential

Perform activate or deactivate action on provided Credential. Activate action will change the status to active, making it possible to connect calls with the credential. Deactivate action will change the status to inactive, making it impossible to connect calls with the credential.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 
  **action** | **string**| Identifies the action to be taken. | 

### Return type

[**InlineResponse2018**](inline_response_201_8.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTelephonyCredential**
> InlineResponse2018 UpdateTelephonyCredential(ctx, body, id)
Update a credential

Update an existing credential.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TelephonyCredentialUpdate**](TelephonyCredentialUpdate.md)| Parameters that can be updated in a credential | 
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse2018**](inline_response_201_8.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

