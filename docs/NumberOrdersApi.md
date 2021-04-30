# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNumberOrder**](NumberOrdersApi.md#CreateNumberOrder) | **Post** /number_orders | Create Phone Number Order
[**ListNumberOrders**](NumberOrdersApi.md#ListNumberOrders) | **Get** /number_orders | Retrieve multiple Number Orders
[**RetrieveNumberOrder**](NumberOrdersApi.md#RetrieveNumberOrder) | **Get** /number_orders/{number_order_id} | Retrieve a single phone number order
[**UpdateNumberOrder**](NumberOrdersApi.md#UpdateNumberOrder) | **Patch** /number_orders/{number_order_id} | Update phone number order

# **CreateNumberOrder**
> InlineResponse20034 CreateNumberOrder(ctx, body)
Create Phone Number Order

Creates a Phone Number Order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body24**](Body24.md)|  | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNumberOrders**
> InlineResponse20033 ListNumberOrders(ctx, optional)
Retrieve multiple Number Orders

Retrieve a paginated list of Number Orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberOrdersApiListNumberOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberOrdersApiListNumberOrdersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStatus** | **optional.String**| Filter number orders by status | 
 **filterCreatedAtGt** | **optional.String**| Filter number orders later than this value | 
 **filterCreatedAtLt** | **optional.String**| Filter number orders earlier than this value | 
 **filterPhoneNumbersPhoneNumber** | **optional.String**| Filter number orders having these phone numbers | 
 **filterCustomerReference** | **optional.String**| Filter number orders via the customer reference set | 
 **filterRequirementsMet** | **optional.Bool**| Filter number orders by requirements met | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveNumberOrder**
> InlineResponse20034 RetrieveNumberOrder(ctx, numberOrderId)
Retrieve a single phone number order

Retrieve an existing single phone number order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **numberOrderId** | **string**| The number order id | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNumberOrder**
> InlineResponse20034 UpdateNumberOrder(ctx, body, numberOrderId)
Update phone number order

Updates a Phone Number Order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body25**](Body25.md)|  | 
  **numberOrderId** | **string**| The number order id | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

