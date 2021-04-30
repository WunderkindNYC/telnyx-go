# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListOutboundChannels**](InboundChannelsApi.md#ListOutboundChannels) | **Get** /phone_numbers/inbound_channels | Retrieve your inbound channels
[**UpdateOutboundChannels**](InboundChannelsApi.md#UpdateOutboundChannels) | **Patch** /phone_numbers/inbound_channels | Update inbound channels

# **ListOutboundChannels**
> InlineResponse20026 ListOutboundChannels(ctx, )
Retrieve your inbound channels

Returns the inbound channels for your account. Inbound channels allows you to use Channel Billing for calls to your Telnyx phone numbers. Please check the Telnyx Support Articles section for full information and examples of how to utilize Channel Billing.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOutboundChannels**
> InlineResponse20027 UpdateOutboundChannels(ctx, body)
Update inbound channels

Update the inbound channels for the account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body1**](Body1.md)| Inbound channels update | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

