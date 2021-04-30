# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsageReport**](MDRUsageReportsApi.md#DeleteUsageReport) | **Delete** /reports/mdr_usage_reports/{id} | 
[**GetUsageReport**](MDRUsageReportsApi.md#GetUsageReport) | **Get** /reports/mdr_usage_reports/{id} | 
[**GetUsageReportSync1**](MDRUsageReportsApi.md#GetUsageReportSync1) | **Get** /reports/mdr_usage_reports/sync | 
[**GetUsageReports**](MDRUsageReportsApi.md#GetUsageReports) | **Get** /reports/mdr_usage_reports | 
[**SubmitUsageReport**](MDRUsageReportsApi.md#SubmitUsageReport) | **Post** /reports/mdr_usage_reports | 

# **DeleteUsageReport**
> MdrDeleteUsageReportsResponse DeleteUsageReport(ctx, id)


Delete messaging usage report by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)|  | 

### Return type

[**MdrDeleteUsageReportsResponse**](MdrDeleteUsageReportsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageReport**
> MdrGetUsageReportsByIdResponse GetUsageReport(ctx, id)


Fetch a single messaging usage report by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)|  | 

### Return type

[**MdrGetUsageReportsByIdResponse**](MdrGetUsageReportsByIdResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageReportSync1**
> MdrGetSyncUsageReportResponse GetUsageReportSync1(ctx, aggregationType, optional)


Generate and fetch messaging usage report synchronously. This endpoint will both generate and fetch the messaging report over a specified time period. No polling is necessary but the response may take up to a couple of minutes. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aggregationType** | **string**|  | 
 **optional** | ***MDRUsageReportsApiGetUsageReportSync1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MDRUsageReportsApiGetUsageReportSync1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **optional.Time**|  | 
 **endDate** | **optional.Time**|  | 
 **profiles** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**MdrGetSyncUsageReportResponse**](MdrGetSyncUsageReportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageReports**
> MdrGetUsageReportsResponse GetUsageReports(ctx, optional)


Fetch all messaging usage reports. Usage reports are aggregated messaging data for specified time period and breakdown

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MDRUsageReportsApiGetUsageReportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MDRUsageReportsApiGetUsageReportsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Page number | [default to 1]
 **pageSize** | **optional.Int32**| Size of the page | [default to 20]

### Return type

[**MdrGetUsageReportsResponse**](MdrGetUsageReportsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitUsageReport**
> MdrPostUsageReportsResponse SubmitUsageReport(ctx, body)


Submit request for new new messaging usage report. This endpoint will pull and aggregate messaging data in specified time period. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MdrPostUsageReportRequest**](MdrPostUsageReportRequest.md)| Mdr usage report data | 

### Return type

[**MdrPostUsageReportsResponse**](MdrPostUsageReportsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

