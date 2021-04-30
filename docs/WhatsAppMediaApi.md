# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMedia**](WhatsAppMediaApi.md#DeleteMedia) | **Delete** /whatsapp_media/{whatsapp_user_id}/{media_id} | Delete Media
[**DownloadMedia**](WhatsAppMediaApi.md#DownloadMedia) | **Get** /whatsapp_media/{whatsapp_user_id}/{media_id} | Download Media
[**UploadMedia**](WhatsAppMediaApi.md#UploadMedia) | **Post** /whatsapp_media | Upload Media

# **DeleteMedia**
> DeleteMedia(ctx, whatsappUserId, mediaId)
Delete Media

Delete uploaded media.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **whatsappUserId** | **string**| User&#x27;s WhatsApp ID | 
  **mediaId** | **string**| WhatsApp&#x27;s Media ID | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadMedia**
> *os.File DownloadMedia(ctx, whatsappUserId, mediaId)
Download Media

Retrieve uploaded media. Media is typically available for 30 days after uploading.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **whatsappUserId** | **string**| User&#x27;s WhatsApp ID | 
  **mediaId** | **string**| WhatsApp&#x27;s Media ID | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadMedia**
> UploadMedia_ UploadMedia(ctx, mediaContentType, uploadFile, whatsappUserId)
Upload Media

Upload media. Accepted media types include image/jpeg, image/png, video/mp4, video/3gpp, audio/aac, audio/ogg.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaContentType** | **string**|  | 
  **uploadFile** | ***os.File*****os.File**|  | 
  **whatsappUserId** | **string**|  | 

### Return type

[**UploadMedia_**](Upload Media..md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

