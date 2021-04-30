# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMessagingHostedNumber**](MessagingHostedNumbersApi.md#DeleteMessagingHostedNumber) | **Delete** /messaging_hosted_numbers/{id} | Delete Messaging Hosted Number
[**GetMessagingHostedNumberOrder**](MessagingHostedNumbersApi.md#GetMessagingHostedNumberOrder) | **Get** /messaging_hosted_number_orders/{id} | Get Messaging Hosted Numbers Order Information
[**ListMessagingHostedNumberOrder**](MessagingHostedNumbersApi.md#ListMessagingHostedNumberOrder) | **Get** /messaging_hosted_number_orders | List All Messaging Hosted Number Orders
[**NewMessagingHostedNumberOrder**](MessagingHostedNumbersApi.md#NewMessagingHostedNumberOrder) | **Post** /messaging_hosted_number_orders | New Messaging Hosted Numbers Order
[**UploadFilesMessagingHostedNumberOrder**](MessagingHostedNumbersApi.md#UploadFilesMessagingHostedNumberOrder) | **Post** /messaging_hosted_number_orders/{id}/actions/file_upload | Upload LOA and Bill required for a Messaging Hosted Number Order

# **DeleteMessagingHostedNumber**
> InlineResponse20024 DeleteMessagingHostedNumber(ctx, id)
Delete Messaging Hosted Number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the type of resource. | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagingHostedNumberOrder**
> InlineResponse20023 GetMessagingHostedNumberOrder(ctx, id)
Get Messaging Hosted Numbers Order Information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the type of resource. | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMessagingHostedNumberOrder**
> InlineResponse20022 ListMessagingHostedNumberOrder(ctx, )
List All Messaging Hosted Number Orders

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewMessagingHostedNumberOrder**
> InlineResponse20023 NewMessagingHostedNumberOrder(ctx, optional)
New Messaging Hosted Numbers Order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagingHostedNumbersApiNewMessagingHostedNumberOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagingHostedNumbersApiNewMessagingHostedNumberOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body17**](Body17.md)| Message payload | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFilesMessagingHostedNumberOrder**
> InlineResponse20023 UploadFilesMessagingHostedNumberOrder(ctx, id, optional)
Upload LOA and Bill required for a Messaging Hosted Number Order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the type of resource. | 
 **optional** | ***MessagingHostedNumbersApiUploadFilesMessagingHostedNumberOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagingHostedNumbersApiUploadFilesMessagingHostedNumberOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loa** | **optional.Interface of *os.File****optional.**|  | 
 **bill** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

