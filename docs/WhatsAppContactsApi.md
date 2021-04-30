# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckContact**](WhatsAppContactsApi.md#CheckContact) | **Post** /whatsapp_contacts | Check Contact

# **CheckContact**
> CheckContactResponse CheckContact(ctx, body)
Check Contact

Verify that a phone number belongs to a valid WhatsApp account. You must ensure that the status is valid before you can message a user, and you'll get their WhatsApp ID to use for messaging.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CheckContactRequestBody**](CheckContactRequestBody.md)|  | 

### Return type

[**CheckContactResponse**](Check Contact Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

