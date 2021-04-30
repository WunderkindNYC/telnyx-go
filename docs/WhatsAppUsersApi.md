# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUser**](WhatsAppUsersApi.md#GetUser) | **Get** /whatsapp_users/{whatsapp_user_id} | Get User
[**UpdateWhatsAppWebhook**](WhatsAppUsersApi.md#UpdateWhatsAppWebhook) | **Patch** /whatsapp_users/{whatsapp_user_id} | Update WhatsApp User

# **GetUser**
> SuccessfulResponseWithDetailsAboutTheWhatsAppMessageSent1 GetUser(ctx, whatsappUserId)
Get User

Retrieve WhatsApp user information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **whatsappUserId** | **string**| User&#x27;s WhatsApp ID | 

### Return type

[**SuccessfulResponseWithDetailsAboutTheWhatsAppMessageSent1**](Successful response with details about the WhatsApp message sent._1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWhatsAppWebhook**
> UpdateWhatsAppWebhook(ctx, body, whatsappUserId)
Update WhatsApp User

Update webhook URL used for a given WhatsApp user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateUserRequest**](UpdateUserRequest.md)|  | 
  **whatsappUserId** | **string**| User&#x27;s WhatsApp ID | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

