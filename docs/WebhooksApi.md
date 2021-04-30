# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebhookDeliveries**](WebhooksApi.md#GetWebhookDeliveries) | **Get** /webhook_deliveries | List webhook deliveries
[**GetWebhookDelivery**](WebhooksApi.md#GetWebhookDelivery) | **Get** /webhook_deliveries/{id} | Find webhook_delivery details by ID

# **GetWebhookDeliveries**
> InlineResponse20045 GetWebhookDeliveries(ctx, optional)
List webhook deliveries

Lists webhook_deliveries for the authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhooksApiGetWebhookDeliveriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiGetWebhookDeliveriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStatusEq** | **optional.String**| Return only webhook_deliveries matching the given &#x60;status&#x60; | 
 **filterWebhookContains** | **optional.String**| Return only webhook deliveries whose &#x60;webhook&#x60; component contains the given text | 
 **filterAttemptsContains** | **optional.String**| Return only webhook_deliveries whose &#x60;attempts&#x60; component contains the given text | 
 **filterStartedAtGte** | **optional.String**| Return only webhook_deliveries whose delivery started later than or at given ISO 8601 datetime | 
 **filterStartedAtLte** | **optional.String**| Return only webhook_deliveries whose delivery started earlier than or at given ISO 8601 datetime | 
 **filterFinishedAtGte** | **optional.String**| Return only webhook_deliveries whose delivery finished later than or at given ISO 8601 datetime | 
 **filterFinishedAtLte** | **optional.String**| Return only webhook_deliveries whose delivery finished earlier than or at given ISO 8601 datetime | 
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**InlineResponse20045**](inline_response_200_45.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhookDelivery**
> InlineResponse20046 GetWebhookDelivery(ctx, id)
Find webhook_delivery details by ID

Provides webhook_delivery debug data, such as timestamps, delivery status and attempts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Uniquely identifies the webhook_delivery. | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

