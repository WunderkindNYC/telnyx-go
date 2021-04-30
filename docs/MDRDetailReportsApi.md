# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMdrRequest**](MDRDetailReportsApi.md#DeleteMdrRequest) | **Delete** /reports/batch_mdr_reports/{id} | 
[**GetCdrRequests**](MDRDetailReportsApi.md#GetCdrRequests) | **Get** /reports/batch_mdr_reports | 
[**GetMdrRequest**](MDRDetailReportsApi.md#GetMdrRequest) | **Get** /reports/batch_mdr_reports/{id} | 
[**GetPaginatedMdrs**](MDRDetailReportsApi.md#GetPaginatedMdrs) | **Get** /reports/mdrs | 
[**SubmitMdrRequest**](MDRDetailReportsApi.md#SubmitMdrRequest) | **Post** /reports/batch_mdr_reports | 

# **DeleteMdrRequest**
> MdrDeleteDetailReportResponse DeleteMdrRequest(ctx, id)


Delete generated messaging detail report by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)|  | 

### Return type

[**MdrDeleteDetailReportResponse**](MdrDeleteDetailReportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCdrRequests**
> MdrGetDetailReportResponse GetCdrRequests(ctx, optional)


Fetch all previous requests for messaging detail reports. Messaging detail reports are reports for pulling all messaging records. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MDRDetailReportsApiGetCdrRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MDRDetailReportsApiGetCdrRequestsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Page number | [default to 1]
 **pageSize** | **optional.Int32**| Size of the page | [default to 20]

### Return type

[**MdrGetDetailReportResponse**](MdrGetDetailReportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMdrRequest**
> MdrGetDetailReportByIdResponse GetMdrRequest(ctx, id)


Fetch single messaging detail report by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)|  | 

### Return type

[**MdrGetDetailReportByIdResponse**](MdrGetDetailReportByIdResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaginatedMdrs**
> MdrGetDetailResponse GetPaginatedMdrs(ctx, optional)


Fetch all Mdr records 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MDRDetailReportsApiGetPaginatedMdrsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MDRDetailReportsApiGetPaginatedMdrsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **optional.String**| Pagination start date | 
 **endDate** | **optional.String**| Pagination end date | 
 **id** | **optional.String**|  | 
 **direction** | **optional.String**|  | 
 **profile** | **optional.String**|  | 
 **cld** | **optional.String**|  | 
 **cli** | **optional.String**|  | 
 **status** | **optional.String**|  | 
 **messageType** | **optional.String**|  | 

### Return type

[**MdrGetDetailResponse**](MdrGetDetailResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitMdrRequest**
> MdrPostDetailReportResponse SubmitMdrRequest(ctx, body)


Submit a request for new messaging detail report. Messaging detail report pulls all raw messaging data according to defined filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MdrPostDetailReportRequest**](MdrPostDetailReportRequest.md)| Mdr detail request data | 

### Return type

[**MdrPostDetailReportResponse**](MdrPostDetailReportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

