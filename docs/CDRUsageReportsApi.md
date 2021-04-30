# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsageReportSync**](CDRUsageReportsApi.md#GetUsageReportSync) | **Get** /reports/cdr_usage_reports/sync | 

# **GetUsageReportSync**
> CdrGetSyncUsageReportResponse GetUsageReportSync(ctx, aggregationType, productBreakdown, optional)


Generate and fetch voice usage report synchronously. This endpoint will both generate and fetch the voice report over a specified time period. No polling is necessary but the response may take up to a couple of minutes. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aggregationType** | **string**|  | 
  **productBreakdown** | **string**|  | 
 **optional** | ***CDRUsageReportsApiGetUsageReportSyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CDRUsageReportsApiGetUsageReportSyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **optional.Time**|  | 
 **endDate** | **optional.Time**|  | 
 **connections** | [**optional.Interface of []float64**](float64.md)|  | 

### Return type

[**CdrGetSyncUsageReportResponse**](CdrGetSyncUsageReportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

