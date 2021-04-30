# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFQDN**](FQDNsApi.md#AddFQDN) | **Post** /fqdns | Create an FQDN
[**DeleteFQDN**](FQDNsApi.md#DeleteFQDN) | **Delete** /fqdns/{id} | Delete an FQDN
[**FQDNsGet**](FQDNsApi.md#FQDNsGet) | **Get** /fqdns | Get all FQDNs
[**GetFQDNDetails**](FQDNsApi.md#GetFQDNDetails) | **Get** /fqdns/{id} | Get FQDN
[**UpdateFQDN**](FQDNsApi.md#UpdateFQDN) | **Patch** /fqdns/{id} | Update FQDN

# **AddFQDN**
> InlineResponse2013 AddFQDN(ctx, optional)
Create an FQDN

Create a new FQDN object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FQDNsApiAddFQDNOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FQDNsApiAddFQDNOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body6**](Body6.md)|  | 

### Return type

[**InlineResponse2013**](inline_response_201_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFQDN**
> InlineResponse2013 DeleteFQDN(ctx, id)
Delete an FQDN

Delete an FQDN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse2013**](inline_response_201_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FQDNsGet**
> InlineResponse20017 FQDNsGet(ctx, optional)
Get all FQDNs

Get all FQDNs belonging to the user that match the given filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FQDNsApiFQDNsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FQDNsApiFQDNsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterConnectionId** | **optional.String**| ID of the FQDN connection to which the FQDN belongs. | 
 **filterFqdn** | **optional.String**| FQDN represented by the resource. | 
 **filterPort** | **optional.Int32**| Port to use when connecting to the FQDN. | 
 **filterDnsRecordType** | **optional.String**| DNS record type used by the FQDN. | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFQDNDetails**
> InlineResponse2013 GetFQDNDetails(ctx, id)
Get FQDN

Return the details regarding a specific FQDN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse2013**](inline_response_201_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFQDN**
> InlineResponse2013 UpdateFQDN(ctx, id, optional)
Update FQDN

Update the details of a specific FQDN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 
 **optional** | ***FQDNsApiUpdateFQDNOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FQDNsApiUpdateFQDNOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body7**](Body7.md)|  | 

### Return type

[**InlineResponse2013**](inline_response_201_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

