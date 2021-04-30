# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIp**](IPsApi.md#CreateIp) | **Post** /ips | Create an Ip
[**DeleteIp**](IPsApi.md#DeleteIp) | **Delete** /ips/{id} | Delete an Ip
[**ListIps**](IPsApi.md#ListIps) | **Get** /ips | List Ips
[**RetrieveIp**](IPsApi.md#RetrieveIp) | **Get** /ips/{id} | Retrieve an Ip
[**UpdateIp**](IPsApi.md#UpdateIp) | **Patch** /ips/{id} | Update an Ip

# **CreateIp**
> IpResponse CreateIp(ctx, optional)
Create an Ip

Create a new IP object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPsApiCreateIpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPsApiCreateIpOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateIpRequest**](CreateIpRequest.md)|  | 

### Return type

[**IpResponse**](Ip Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIp**
> IpResponse DeleteIp(ctx, id)
Delete an Ip

Delete an IP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the type of resource. | 

### Return type

[**IpResponse**](Ip Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIps**
> ListIpsResponse ListIps(ctx, optional)
List Ips

Get all IPs belonging to the user that match the given filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPsApiListIpsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPsApiListIpsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterConnectionId** | **optional.String**| ID of the IP Connection to which this IP should be attached. | 
 **filterIpAddress** | **optional.String**| IP adddress represented by this resource. | 
 **filterPort** | **optional.Int32**| Port to use when connecting to this IP. | 

### Return type

[**ListIpsResponse**](List Ips Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveIp**
> IpResponse RetrieveIp(ctx, id)
Retrieve an Ip

Return the details regarding a specific IP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the type of resource. | 

### Return type

[**IpResponse**](Ip Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIp**
> IpResponse UpdateIp(ctx, id, optional)
Update an Ip

Update the details of a specific IP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the type of resource. | 
 **optional** | ***IPsApiUpdateIpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPsApiUpdateIpOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateIpRequest**](UpdateIpRequest.md)|  | 

### Return type

[**IpResponse**](Ip Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

