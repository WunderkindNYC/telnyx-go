# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWdrReport**](ReportingApi.md#CreateWdrReport) | **Post** /wireless/detail_records_reports | Create a Wireless Detail Records (WDRs) Report
[**DeleteWdrReport**](ReportingApi.md#DeleteWdrReport) | **Delete** /wireless/detail_records_reports/{id} | Delete a Wireless Detail Record (WDR) Report
[**GetWdrReport**](ReportingApi.md#GetWdrReport) | **Get** /wireless/detail_records_reports/{id} | Get a Wireless Detail Record (WDR) Report
[**GetWdrReports**](ReportingApi.md#GetWdrReports) | **Get** /wireless/detail_records_reports | Get all Wireless Detail Records (WDRs) Reports

# **CreateWdrReport**
> InlineResponse2017 CreateWdrReport(ctx, body)
Create a Wireless Detail Records (WDRs) Report

Asynchronously create a report containing Wireless Detail Records (WDRs) for the SIM cards that consumed wireless data in the given time period. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body37**](Body37.md)|  | 

### Return type

[**InlineResponse2017**](inline_response_201_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWdrReport**
> InlineResponse2017 DeleteWdrReport(ctx, id)
Delete a Wireless Detail Record (WDR) Report

Deletes one specific WDR report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse2017**](inline_response_201_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWdrReport**
> InlineResponse2017 GetWdrReport(ctx, id)
Get a Wireless Detail Record (WDR) Report

Returns one specific WDR report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse2017**](inline_response_201_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWdrReports**
> InlineResponse20058 GetWdrReports(ctx, optional)
Get all Wireless Detail Records (WDRs) Reports

Returns the WDR Reports that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiGetWdrReportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiGetWdrReportsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**InlineResponse20058**](inline_response_200_58.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

