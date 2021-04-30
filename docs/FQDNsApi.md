# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFqdn**](FQDNsApi.md#CreateFqdn) | **Post** /fqdns | Create an Fqdn
[**DeleteFqdn**](FQDNsApi.md#DeleteFqdn) | **Delete** /fqdns/{id} | Delete an Fqdn
[**ListFqdns**](FQDNsApi.md#ListFqdns) | **Get** /fqdns | List Fqdns
[**RetrieveFqdn**](FQDNsApi.md#RetrieveFqdn) | **Get** /fqdns/{id} | Retrieve an Fqdn
[**UpdateFqdn**](FQDNsApi.md#UpdateFqdn) | **Patch** /fqdns/{id} | Update an Fqdn

# **CreateFqdn**
> FqdnResponse CreateFqdn(ctx, optional)
Create an Fqdn

Create a new FQDN object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FQDNsApiCreateFqdnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FQDNsApiCreateFqdnOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateFqdnRequest**](CreateFqdnRequest.md)|  | 

### Return type

[**FqdnResponse**](Fqdn Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFqdn**
> FqdnResponse DeleteFqdn(ctx, id)
Delete an Fqdn

Delete an FQDN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**FqdnResponse**](Fqdn Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFqdns**
> ListFqdnsResponse ListFqdns(ctx, optional)
List Fqdns

Get all FQDNs belonging to the user that match the given filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FQDNsApiListFqdnsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FQDNsApiListFqdnsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterConnectionId** | **optional.String**| ID of the FQDN connection to which the FQDN belongs. | 
 **filterFqdn** | **optional.String**| FQDN represented by the resource. | 
 **filterPort** | **optional.Int32**| Port to use when connecting to the FQDN. | 
 **filterDnsRecordType** | **optional.String**| DNS record type used by the FQDN. | 

### Return type

[**ListFqdnsResponse**](List Fqdns Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveFqdn**
> FqdnResponse RetrieveFqdn(ctx, id)
Retrieve an Fqdn

Return the details regarding a specific FQDN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**FqdnResponse**](Fqdn Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFqdn**
> FqdnResponse UpdateFqdn(ctx, id, optional)
Update an Fqdn

Update the details of a specific FQDN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 
 **optional** | ***FQDNsApiUpdateFqdnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FQDNsApiUpdateFqdnOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateFqdnRequest**](UpdateFqdnRequest.md)|  | 

### Return type

[**FqdnResponse**](Fqdn Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

