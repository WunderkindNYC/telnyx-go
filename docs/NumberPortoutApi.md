# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindPortoutComments**](NumberPortoutApi.md#FindPortoutComments) | **Get** /portouts/{id}/comments | List all comments for a portout request
[**FindPortoutRequest**](NumberPortoutApi.md#FindPortoutRequest) | **Get** /portouts/{id} | Retrieve a portout request
[**ListPortoutRequest**](NumberPortoutApi.md#ListPortoutRequest) | **Get** /portouts | Retrieve a list of portout requests
[**PostPortRequestComment**](NumberPortoutApi.md#PostPortRequestComment) | **Post** /portouts/{id}/comments | Create a comment on a portout request
[**UpdatePortoutRequest**](NumberPortoutApi.md#UpdatePortoutRequest) | **Patch** /portouts/{id}/{status} | Update Status

# **FindPortoutComments**
> InlineResponse20033 FindPortoutComments(ctx, id)
List all comments for a portout request

Returns a list of comments for a portout request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Portout id | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindPortoutRequest**
> InlineResponse20032 FindPortoutRequest(ctx, id)
Retrieve a portout request

Returns the portout request based on the ID provided

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Portout id | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPortoutRequest**
> InlineResponse20031 ListPortoutRequest(ctx, optional)
Retrieve a list of portout requests

Returns the portout requests according to filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberPortoutApiListPortoutRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberPortoutApiListPortoutRequestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCarrierName** | **optional.String**| Filter by new carrier name. | 
 **filterSpid** | **optional.String**| Filter by new carrier spid. | 
 **filterStatus** | **optional.String**| Filter by portout status. | 
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPortRequestComment**
> InlineResponse2014 PostPortRequestComment(ctx, body, id)
Create a comment on a portout request

Creates a comment on a portout request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body3**](Body3.md)|  | 
  **id** | [**string**](.md)| Portout id | 

### Return type

[**InlineResponse2014**](inline_response_201_4.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePortoutRequest**
> InlineResponse20032 UpdatePortoutRequest(ctx, id, status)
Update Status

Authorize or reject portout request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Portout id | 
  **status** | **string**| Updated portout status | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

