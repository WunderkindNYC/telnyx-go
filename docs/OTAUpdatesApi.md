# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OTAUpdateGET**](OTAUpdatesApi.md#OTAUpdateGET) | **Get** /ota_updates/{id} | Get OTA update
[**OTAUpdatesList**](OTAUpdatesApi.md#OTAUpdatesList) | **Get** /ota_updates | List OTA updates

# **OTAUpdateGET**
> InlineResponse20025 OTAUpdateGET(ctx, id)
Get OTA update

This API returns the details of an Over the Air (OTA) update.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OTAUpdatesList**
> InlineResponse20024 OTAUpdatesList(ctx, optional)
List OTA updates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OTAUpdatesApiOTAUpdatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OTAUpdatesApiOTAUpdatesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterSimCardId** | **optional.String**| The SIM card identification UUID. | 
 **filterStatus** | **optional.String**| Filter by possible OTA updates statuses. | 
 **filterType** | **optional.String**| Filter by type. | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

