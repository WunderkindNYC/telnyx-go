# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessagingHostedNumberOrder**](MessagingHostedNumberApi.md#CreateMessagingHostedNumberOrder) | **Post** /messaging_hosted_number_orders | Create a messaging hosted number order
[**DeleteMessagingHostedNumber**](MessagingHostedNumberApi.md#DeleteMessagingHostedNumber) | **Delete** /messaging_hosted_numbers/{id} | Delete a messaging hosted number
[**ListMessagingHostedNumberOrder**](MessagingHostedNumberApi.md#ListMessagingHostedNumberOrder) | **Get** /messaging_hosted_number_orders | List messaging hosted number orders
[**RetrieveMessagingHostedNumberOrder**](MessagingHostedNumberApi.md#RetrieveMessagingHostedNumberOrder) | **Get** /messaging_hosted_number_orders/{id} | Retrieve a messaging hosted number order
[**UploadFileMessagingHostedNumberOrder**](MessagingHostedNumberApi.md#UploadFileMessagingHostedNumberOrder) | **Post** /messaging_hosted_number_orders/{id}/actions/file_upload | Upload file required for a messaging hosted number order

# **CreateMessagingHostedNumberOrder**
> RetrieveMessagingHostedNumberOrderResponse CreateMessagingHostedNumberOrder(ctx, optional)
Create a messaging hosted number order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagingHostedNumberApiCreateMessagingHostedNumberOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagingHostedNumberApiCreateMessagingHostedNumberOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateMessagingHostedNumberOrderRequest**](CreateMessagingHostedNumberOrderRequest.md)| Message payload | 

### Return type

[**RetrieveMessagingHostedNumberOrderResponse**](Retrieve Messaging Hosted Number Order Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMessagingHostedNumber**
> RetrieveMessagingHostedNumberResponse DeleteMessagingHostedNumber(ctx, id)
Delete a messaging hosted number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the type of resource. | 

### Return type

[**RetrieveMessagingHostedNumberResponse**](Retrieve Messaging Hosted Number Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMessagingHostedNumberOrder**
> ListMessagingHostedNumberOrderResponse ListMessagingHostedNumberOrder(ctx, optional)
List messaging hosted number orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagingHostedNumberApiListMessagingHostedNumberOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagingHostedNumberApiListMessagingHostedNumberOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**ListMessagingHostedNumberOrderResponse**](List Messaging Hosted Number Order Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveMessagingHostedNumberOrder**
> RetrieveMessagingHostedNumberOrderResponse RetrieveMessagingHostedNumberOrder(ctx, id)
Retrieve a messaging hosted number order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the type of resource. | 

### Return type

[**RetrieveMessagingHostedNumberOrderResponse**](Retrieve Messaging Hosted Number Order Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFileMessagingHostedNumberOrder**
> RetrieveMessagingHostedNumberOrderResponse UploadFileMessagingHostedNumberOrder(ctx, id, optional)
Upload file required for a messaging hosted number order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the type of resource. | 
 **optional** | ***MessagingHostedNumberApiUploadFileMessagingHostedNumberOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagingHostedNumberApiUploadFileMessagingHostedNumberOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bill** | **optional.Interface of *os.File****optional.**|  | 
 **loa** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**RetrieveMessagingHostedNumberOrderResponse**](Retrieve Messaging Hosted Number Order Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

