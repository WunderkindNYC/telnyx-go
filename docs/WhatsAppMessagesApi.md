# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarkMessageAsRead**](WhatsAppMessagesApi.md#MarkMessageAsRead) | **Patch** /whatsapp_messages/{message_id} | Mark Message As Read
[**SendMessage**](WhatsAppMessagesApi.md#SendMessage) | **Post** /whatsapp_messages | Send Message

# **MarkMessageAsRead**
> MarkMessageAsRead(ctx, body, messageId)
Mark Message As Read

Mark a message as read.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MarkMessageAsReadRequestBody**](MarkMessageAsReadRequestBody.md)|  | 
  **messageId** | **string**| Message ID from Webhook | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendMessage**
> SuccessfulResponseWithDetailsAboutTheWhatsAppMessageSent_ SendMessage(ctx, body)
Send Message

Send text messages, media/documents, and message templates to your customers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateWhatsAppMessageRequest**](CreateWhatsAppMessageRequest.md)|  | 

### Return type

[**SuccessfulResponseWithDetailsAboutTheWhatsAppMessageSent_**](Successful response with details about the WhatsApp message sent..md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

