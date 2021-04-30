# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveCallStatus**](CallInformationApi.md#RetrieveCallStatus) | **Get** /calls/{call_control_id} | Retrieve a call status

# **RetrieveCallStatus**
> RetrieveCallStatusResponse RetrieveCallStatus(ctx, callControlId)
Retrieve a call status

Returns the status of a call (data is available 10 minutes after call ended).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**RetrieveCallStatusResponse**](Retrieve Call Status Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

