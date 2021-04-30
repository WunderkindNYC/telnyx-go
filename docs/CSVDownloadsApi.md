# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCsvDownload**](CSVDownloadsApi.md#CreateCsvDownload) | **Post** /phone_numbers/csv_downloads | Create a CSV download
[**ListCsvDownloads**](CSVDownloadsApi.md#ListCsvDownloads) | **Get** /phone_numbers/csv_downloads | List CSV downloads
[**RetrieveCsvDownload**](CSVDownloadsApi.md#RetrieveCsvDownload) | **Get** /phone_numbers/csv_downloads/{id} | Retrieve a CSV download

# **CreateCsvDownload**
> CsvDownloadResponse CreateCsvDownload(ctx, )
Create a CSV download

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CsvDownloadResponse**](CSV Download Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCsvDownloads**
> ListCsvDownloadsResponse ListCsvDownloads(ctx, optional)
List CSV downloads

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CSVDownloadsApiListCsvDownloadsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CSVDownloadsApiListCsvDownloadsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**ListCsvDownloadsResponse**](List CSV Downloads Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCsvDownload**
> CsvDownloadResponse RetrieveCsvDownload(ctx, id)
Retrieve a CSV download

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the CSV download. | 

### Return type

[**CsvDownloadResponse**](CSV Download Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

