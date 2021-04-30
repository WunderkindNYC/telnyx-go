# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrivateWirelessGateway**](PrivateWirelessGatewaysApi.md#CreatePrivateWirelessGateway) | **Post** /private_wireless_gateways | Create a Private Wireless Gateway
[**DeletePrivateWirelessGateway**](PrivateWirelessGatewaysApi.md#DeletePrivateWirelessGateway) | **Delete** /private_wireless_gateways/{id} | Delete a Private Wireless Gateway
[**GetPrivateWirelessGateway**](PrivateWirelessGatewaysApi.md#GetPrivateWirelessGateway) | **Get** /private_wireless_gateways/{id} | Get a Private Wireless Gateway
[**GetPrivateWirelessGateways**](PrivateWirelessGatewaysApi.md#GetPrivateWirelessGateways) | **Get** /private_wireless_gateways | Get all Private Wireless Gateways

# **CreatePrivateWirelessGateway**
> InlineResponse2023 CreatePrivateWirelessGateway(ctx, body)
Create a Private Wireless Gateway

Asynchronously create a Private Wireless Gateway for SIM cards for a previously created network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body4**](Body4.md)|  | 

### Return type

[**InlineResponse2023**](inline_response_202_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePrivateWirelessGateway**
> InlineResponse2023 DeletePrivateWirelessGateway(ctx, id)
Delete a Private Wireless Gateway

Deletes the Private Wireless Gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse2023**](inline_response_202_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPrivateWirelessGateway**
> InlineResponse2023 GetPrivateWirelessGateway(ctx, id)
Get a Private Wireless Gateway

Retrieve information about a Private Wireless Gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse2023**](inline_response_202_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPrivateWirelessGateways**
> InlineResponse20034 GetPrivateWirelessGateways(ctx, optional)
Get all Private Wireless Gateways

Get all Private Wireless Gateways belonging to the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PrivateWirelessGatewaysApiGetPrivateWirelessGatewaysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateWirelessGatewaysApiGetPrivateWirelessGatewaysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterName** | **optional.String**| The name of the Private Wireless Gateway. | 
 **filterIpRange** | **optional.String**| The IP address range of the Private Wireless Gateway. | 
 **filterRegionCode** | **optional.String**| The name of the region where the Private Wireless Gateway is deployed. | 
 **filterCreatedAt** | **optional.String**| Private Wireless Gateway resource creation date. | 
 **filterUpdatedAt** | **optional.String**| When the Private Wireless Gateway was last updated. | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

