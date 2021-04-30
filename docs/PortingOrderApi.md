# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPortingOrder**](PortingOrderApi.md#CancelPortingOrder) | **Post** /porting_orders/{id}/actions/cancel | Cancel this porting order
[**ConfirmPortingOrder**](PortingOrderApi.md#ConfirmPortingOrder) | **Post** /porting_orders/{id}/actions/confirm | Confirms the porting order is ready to be actioned.
[**CreatePortingOrder**](PortingOrderApi.md#CreatePortingOrder) | **Post** /porting_orders | Create a porting order
[**DeletePortingOrder**](PortingOrderApi.md#DeletePortingOrder) | **Delete** /porting_orders/{id} | Request cancellation
[**GetPortingOrder**](PortingOrderApi.md#GetPortingOrder) | **Get** /porting_orders/{id} | Retrieve a porting order
[**GetPortingOrderLOATemplate**](PortingOrderApi.md#GetPortingOrderLOATemplate) | **Get** /porting_orders/{id}/loa_template | Download a porting order loa template
[**ListPhoneNumbers**](PortingOrderApi.md#ListPhoneNumbers) | **Get** /porting_phone_numbers | List all porting phone numbers
[**ListPortingOrders**](PortingOrderApi.md#ListPortingOrders) | **Get** /porting_orders | List all porting orders
[**UpdatePortingOrder**](PortingOrderApi.md#UpdatePortingOrder) | **Patch** /porting_orders/{id} | Edit a porting order

# **CancelPortingOrder**
> InlineResponse20029 CancelPortingOrder(ctx, id)
Cancel this porting order

Cancel this porting order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Porting Order id | 

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfirmPortingOrder**
> InlineResponse20029 ConfirmPortingOrder(ctx, id)
Confirms the porting order is ready to be actioned.

Confirm your porting order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Porting Order id | 

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePortingOrder**
> InlineResponse2013 CreatePortingOrder(ctx, body)
Create a porting order

Creates a new porting order object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreatePortingOrder**](CreatePortingOrder.md)|  | 

### Return type

[**InlineResponse2013**](inline_response_201_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePortingOrder**
> InlineResponse20029 DeletePortingOrder(ctx, id)
Request cancellation

Request the cancellation of a porting order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Porting Order id | 

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPortingOrder**
> InlineResponse20029 GetPortingOrder(ctx, id, optional)
Retrieve a porting order

Retrieves the details of an existing porting order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Porting Order id | 
 **optional** | ***PortingOrderApiGetPortingOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortingOrderApiGetPortingOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includePhoneNumbers** | **optional.Bool**| Include the first 50 phone number objects in the results | [default to true]

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPortingOrderLOATemplate**
> *os.File GetPortingOrderLOATemplate(ctx, id)
Download a porting order loa template

Download a porting order loa template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Porting Order id | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPhoneNumbers**
> InlineResponse20030 ListPhoneNumbers(ctx, optional)
List all porting phone numbers

Returns a list of your porting phone numbers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortingOrderApiListPhoneNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortingOrderApiListPhoneNumbersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterPortingOrderId** | [**optional.Interface of string**](.md)| Filter results by porting order id | 
 **filterPhoneNumber** | **optional.String**| Filter results by phone number | 
 **filterActivationStatus** | [**optional.Interface of ActivationStatus**](.md)| Filter results by activation status | 
 **filterPortabilityStatus** | [**optional.Interface of PortabilityStatus**](.md)| Filter results by portability status | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPortingOrders**
> InlineResponse20028 ListPortingOrders(ctx, optional)
List all porting orders

Returns a list of your porting order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortingOrderApiListPortingOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortingOrderApiListPortingOrdersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **includePhoneNumbers** | **optional.Bool**| Include the first 50 phone number objects in the results | [default to true]
 **filterStatus** | [**optional.Interface of PortingOrderStatus**](.md)| Filter results by status | 
 **filterCustomerReference** | **optional.String**| Filter results by user reference | 
 **filterPhoneNumbersCarrierName** | **optional.String**| Filter results by old service provider | 
 **filterMiscType** | [**optional.Interface of PortingOrderType**](.md)| Filter results by porting order type | 
 **filterEndUserAdminEntityName** | **optional.String**| Filter results by person or company name | 
 **filterEndUserAdminAuthPersonName** | **optional.String**| Filter results by authorized person | 
 **filterActivationSettingsFastPortEligible** | **optional.Bool**| Filter results by fast port eligible | 
 **filterActivationSettingsFocDatetimeRequestedGt** | **optional.String**| Filter results by foc date later than this value | 
 **filterActivationSettingsFocDatetimeRequestedLt** | **optional.String**| Filter results by foc date earlier than this value | 
 **sort** | **optional.String**| Specifies the sort order for results. If not given, results are sorted by created_at in descending order. | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePortingOrder**
> InlineResponse20029 UpdatePortingOrder(ctx, body, id)
Edit a porting order

Edits the details of an existing porting order.  Any or all of a porting order’s attributes may be included in the resource object included in a PATCH request.  If a request does not include all of the attributes for a resource, the system will interpret the missing attributes as if they were included with their current values. To explicitly set something to null, it must be included in the request with a null value.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePortingOrder**](UpdatePortingOrder.md)|  | 
  **id** | [**string**](.md)| Porting Order id | 

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

