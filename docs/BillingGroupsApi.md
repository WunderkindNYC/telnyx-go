# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingGroup**](BillingGroupsApi.md#CreateBillingGroup) | **Post** /billing_groups | Create a billing group
[**DeleteBillingGroup**](BillingGroupsApi.md#DeleteBillingGroup) | **Delete** /billing_groups/{id} | Delete a billing group
[**ListBillingGroups**](BillingGroupsApi.md#ListBillingGroups) | **Get** /billing_groups | List all billing groups
[**RetrieveBillingGroup**](BillingGroupsApi.md#RetrieveBillingGroup) | **Get** /billing_groups/{id} | Retrieve a billing group
[**UpdateBillingGroup**](BillingGroupsApi.md#UpdateBillingGroup) | **Patch** /billing_groups/{id} | Update a billing group

# **CreateBillingGroup**
> InlineResponse2005 CreateBillingGroup(ctx, body)
Create a billing group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body2**](Body2.md)| New billing group object | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBillingGroup**
> InlineResponse2005 DeleteBillingGroup(ctx, id)
Delete a billing group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the billing group | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBillingGroups**
> InlineResponse2004 ListBillingGroups(ctx, optional)
List all billing groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BillingGroupsApiListBillingGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingGroupsApiListBillingGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveBillingGroup**
> InlineResponse2005 RetrieveBillingGroup(ctx, id)
Retrieve a billing group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the billing group | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBillingGroup**
> InlineResponse2005 UpdateBillingGroup(ctx, body, id)
Update a billing group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body3**](Body3.md)| Update billing group object | 
  **id** | [**string**](.md)| The id of the billing group | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

