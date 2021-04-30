# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNumberOrder**](NumberOrdersApi.md#CreateNumberOrder) | **Post** /number_orders | Create a number order
[**ListNumberOrders**](NumberOrdersApi.md#ListNumberOrders) | **Get** /number_orders | List number orders
[**RetrieveNumberOrder**](NumberOrdersApi.md#RetrieveNumberOrder) | **Get** /number_orders/{number_order_id} | Retrieve a number order
[**UpdateNumberOrder**](NumberOrdersApi.md#UpdateNumberOrder) | **Patch** /number_orders/{number_order_id} | Update a number order

# **CreateNumberOrder**
> NumberOrderResponse CreateNumberOrder(ctx, body)
Create a number order

Creates a phone number order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateNumberOrderRequest**](CreateNumberOrderRequest.md)|  | 

### Return type

[**NumberOrderResponse**](Number Order Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNumberOrders**
> ListNumberOrdersResponse ListNumberOrders(ctx, optional)
List number orders

Get a paginated list of number orders.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberOrdersApiListNumberOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberOrdersApiListNumberOrdersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStatus** | **optional.String**| Filter number orders by status. | 
 **filterCreatedAtGt** | **optional.String**| Filter number orders later than this value. | 
 **filterCreatedAtLt** | **optional.String**| Filter number orders earlier than this value. | 
 **filterPhoneNumbersPhoneNumber** | **optional.String**| Filter number orders having these phone numbers. | 
 **filterCustomerReference** | **optional.String**| Filter number orders via the customer reference set. | 
 **filterRequirementsMet** | **optional.Bool**| Filter number orders by requirements met. | 
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**ListNumberOrdersResponse**](List Number Orders Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveNumberOrder**
> NumberOrderResponse RetrieveNumberOrder(ctx, numberOrderId)
Retrieve a number order

Get an existing phone number order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **numberOrderId** | **string**| The number order ID. | 

### Return type

[**NumberOrderResponse**](Number Order Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNumberOrder**
> NumberOrderResponse UpdateNumberOrder(ctx, body, numberOrderId)
Update a number order

Updates a phone number order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateNumberOrderRequest**](UpdateNumberOrderRequest.md)|  | 
  **numberOrderId** | **string**| The number order ID. | 

### Return type

[**NumberOrderResponse**](Number Order Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

