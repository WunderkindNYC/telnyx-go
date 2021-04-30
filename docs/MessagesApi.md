# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLongCodeMessage**](MessagesApi.md#CreateLongCodeMessage) | **Post** /messages/long_code | Send a Long Code message
[**CreateMessage**](MessagesApi.md#CreateMessage) | **Post** /messages | Send a message
[**CreateNumberPoolMessage**](MessagesApi.md#CreateNumberPoolMessage) | **Post** /messages/number_pool | Send a message using Number Pool
[**CreateShortCodeMessage**](MessagesApi.md#CreateShortCodeMessage) | **Post** /messages/short_code | Send a Short Code message
[**RetrieveMessage**](MessagesApi.md#RetrieveMessage) | **Get** /messages/{id} | Retrieve a message

# **CreateLongCodeMessage**
> InlineResponse20021 CreateLongCodeMessage(ctx, optional)
Send a Long Code message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesApiCreateLongCodeMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiCreateLongCodeMessageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body14**](Body14.md)| Message payload | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMessage**
> InlineResponse20021 CreateMessage(ctx, optional)
Send a message

Send a message with a Phone Number, Alphanumeric Sender ID, Short Code or Number Pool.  This endpoint allows you to send a message with any messaging resource. Current messaging resources include: long-code, short-code, number-pool, and alphanumeric-sender-id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesApiCreateMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiCreateMessageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body13**](Body13.md)| Message payload | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNumberPoolMessage**
> InlineResponse20021 CreateNumberPoolMessage(ctx, optional)
Send a message using Number Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesApiCreateNumberPoolMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiCreateNumberPoolMessageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body15**](Body15.md)| Message payload | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateShortCodeMessage**
> InlineResponse20021 CreateShortCodeMessage(ctx, optional)
Send a Short Code message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesApiCreateShortCodeMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiCreateShortCodeMessageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body16**](Body16.md)| Message payload | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveMessage**
> InlineResponse20021 RetrieveMessage(ctx, id)
Retrieve a message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the message | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

