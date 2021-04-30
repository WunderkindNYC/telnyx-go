# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCsvDownload**](CSVDownloadsApi.md#CreateCsvDownload) | **Post** /phone_numbers/csv_downloads | create a new CSV download request
[**FindCsvDownloads**](CSVDownloadsApi.md#FindCsvDownloads) | **Get** /phone_numbers/csv_downloads | List your submitted CSV download requests
[**RetrieveCsvDownload**](CSVDownloadsApi.md#RetrieveCsvDownload) | **Get** /phone_numbers/csv_downloads/{id} | Get a single submitted CSV download request.

# **CreateCsvDownload**
> InlineResponse20041 CreateCsvDownload(ctx, )
create a new CSV download request

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindCsvDownloads**
> InlineResponse20041 FindCsvDownloads(ctx, optional)
List your submitted CSV download requests

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CSVDownloadsApiFindCsvDownloadsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CSVDownloadsApiFindCsvDownloadsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCsvDownload**
> InlineResponse20041 RetrieveCsvDownload(ctx, id)
Get a single submitted CSV download request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the CSV download. | 

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

