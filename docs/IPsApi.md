# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIP**](IPsApi.md#AddIP) | **Post** /ips | Create an IP
[**DeleteIP**](IPsApi.md#DeleteIP) | **Delete** /ips/{id} | Delete an IP
[**GetIPDetails**](IPsApi.md#GetIPDetails) | **Get** /ips/{id} | Get IP
[**IPsGet**](IPsApi.md#IPsGet) | **Get** /ips | Get all IPs
[**UpdateIP**](IPsApi.md#UpdateIP) | **Patch** /ips/{id} | Update IP

# **AddIP**
> InlineResponse2015 AddIP(ctx, optional)
Create an IP

Create a new IP object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPsApiAddIPOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPsApiAddIPOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body10**](Body10.md)|  | 

### Return type

[**InlineResponse2015**](inline_response_201_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIP**
> InlineResponse2015 DeleteIP(ctx, id)
Delete an IP

Delete an IP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the type of resource. | 

### Return type

[**InlineResponse2015**](inline_response_201_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPDetails**
> InlineResponse2015 GetIPDetails(ctx, id)
Get IP

Return the details regarding a specific IP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the type of resource. | 

### Return type

[**InlineResponse2015**](inline_response_201_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IPsGet**
> InlineResponse20019 IPsGet(ctx, optional)
Get all IPs

Get all IPs belonging to the user that match the given filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPsApiIPsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPsApiIPsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterConnectionId** | **optional.String**| ID of the IPConnection to which this IP should be attached. | 
 **filterIpAddress** | **optional.String**| IP adddress represented by this resource. | 
 **filterPort** | **optional.Int32**| Port to use when connecting to this IP. | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIP**
> InlineResponse2015 UpdateIP(ctx, id, optional)
Update IP

Update the details of a specific IP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the type of resource. | 
 **optional** | ***IPsApiUpdateIPOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPsApiUpdateIPOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body11**](Body11.md)|  | 

### Return type

[**InlineResponse2015**](inline_response_201_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

