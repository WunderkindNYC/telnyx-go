# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveCallStatus**](CallInformationApi.md#RetrieveCallStatus) | **Get** /calls/{call_control_id} | Retrieve call status

# **RetrieveCallStatus**
> InlineResponse2008 RetrieveCallStatus(ctx, callControlId)
Retrieve call status

Returns the status of a call.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

