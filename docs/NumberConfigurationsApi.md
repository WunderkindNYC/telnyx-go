# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePhoneNumber**](NumberConfigurationsApi.md#DeletePhoneNumber) | **Delete** /phone_numbers/{id} | Delete a phone number
[**EnableEmergencyPhoneNumber**](NumberConfigurationsApi.md#EnableEmergencyPhoneNumber) | **Post** /phone_numbers/{id}/actions/enable_emergency | Enable emergency for a phone number
[**FindPhoneNumberVoices**](NumberConfigurationsApi.md#FindPhoneNumberVoices) | **Get** /phone_numbers/voice | List voice settings for multiple phone numbers
[**FindPhoneNumbers**](NumberConfigurationsApi.md#FindPhoneNumbers) | **Get** /phone_numbers | List all phone numbers
[**GetPhoneNumber**](NumberConfigurationsApi.md#GetPhoneNumber) | **Get** /phone_numbers/{id} | Retrieve the settings for a phone number
[**ListPhoneNumberMessagingSettings**](NumberConfigurationsApi.md#ListPhoneNumberMessagingSettings) | **Get** /phone_numbers/messaging | List all phone numbers&#x27; messaging settings
[**RetrievePhoneNumberMessagingSettings**](NumberConfigurationsApi.md#RetrievePhoneNumberMessagingSettings) | **Get** /phone_numbers/{id}/messaging | Retrieve the messaging settings for a phone number
[**RetrievePhoneNumberVoice**](NumberConfigurationsApi.md#RetrievePhoneNumberVoice) | **Get** /phone_numbers/{id}/voice | Retrieve the voice settings for a phone number
[**UpdatePhoneNumber**](NumberConfigurationsApi.md#UpdatePhoneNumber) | **Patch** /phone_numbers/{id} | Update the settings for a phone number
[**UpdatePhoneNumberMessagingSettings**](NumberConfigurationsApi.md#UpdatePhoneNumberMessagingSettings) | **Patch** /phone_numbers/{id}/messaging | Update the messaging settings for a phone number
[**UpdatePhoneNumberVoice**](NumberConfigurationsApi.md#UpdatePhoneNumberVoice) | **Patch** /phone_numbers/{id}/voice | Update the voice settings for a phone number

# **DeletePhoneNumber**
> InlineResponse20046 DeletePhoneNumber(ctx, id)
Delete a phone number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableEmergencyPhoneNumber**
> InlineResponse20047 EnableEmergencyPhoneNumber(ctx, body, id)
Enable emergency for a phone number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body29**](Body29.md)|  | 
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindPhoneNumberVoices**
> InlineResponse20045 FindPhoneNumberVoices(ctx, optional)
List voice settings for multiple phone numbers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberConfigurationsApiFindPhoneNumberVoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberConfigurationsApiFindPhoneNumberVoicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterPhoneNumber** | **optional.String**| Filter by phone number. Requires at least three digits.              Non-numerical characters will result in no values being returned. | 
 **filterConnectionNameContains** | **optional.String**| Filter contains connection name. Requires at least three characters | 
 **filterUsagePaymentMethod** | **optional.String**| Filter by usage_payment_method. | 
 **sort** | **optional.String**| Specifies the sort order for results. If not given, results are sorted by created_at in descending order | 

### Return type

[**InlineResponse20045**](inline_response_200_45.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindPhoneNumbers**
> InlineResponse20040 FindPhoneNumbers(ctx, optional)
List all phone numbers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberConfigurationsApiFindPhoneNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberConfigurationsApiFindPhoneNumbersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterTag** | **optional.String**| Filter by phone number tags | 
 **filterPhoneNumber** | **optional.String**| Filter by phone number. Requires at least three digits.              Non-numerical characters will result in no values being returned. | 
 **filterStatus** | **optional.String**| Filter by phone number status | 
 **filterVoiceConnectionNameContains** | **optional.String**| Filter contains connection name. Requires at least three characters | 
 **filterVoiceConnectionNameStartsWith** | **optional.String**| Filter starts with connection name. Requires at least three characters | 
 **filterVoiceConnectionNameEndsWith** | **optional.String**| Filter ends with connection name. Requires at least three characters | 
 **filterVoiceConnectionNameEq** | **optional.String**| Filter by connection name | 
 **filterUsagePaymentMethod** | **optional.String**| Filter by usage_payment_method. | 
 **filterMessagingMessagingProfileNameContains** | **optional.String**| Filter contains messaging profile name. Requires at least three characters | 
 **filterMessagingMessagingProfileNameStartsWith** | **optional.String**| Filter starts with messaging profile name. Requires at least three characters | 
 **filterMessagingMessagingProfileNameEndsWith** | **optional.String**| Filter ends with messaging profile name. Requires at least three characters | 
 **filterMessagingMessagingProfileNameEq** | **optional.String**| Filter by messaging profile name | 
 **sort** | **optional.String**| Specifies the sort order for results. If not given, results are sorted by created_at in descending order | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPhoneNumber**
> InlineResponse20046 GetPhoneNumber(ctx, id)
Retrieve the settings for a phone number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPhoneNumberMessagingSettings**
> InlineResponse20044 ListPhoneNumberMessagingSettings(ctx, )
List all phone numbers' messaging settings

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20044**](inline_response_200_44.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrievePhoneNumberMessagingSettings**
> InlineResponse20048 RetrievePhoneNumberMessagingSettings(ctx, id)
Retrieve the messaging settings for a phone number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the type of resource. | 

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrievePhoneNumberVoice**
> InlineResponse20047 RetrievePhoneNumberVoice(ctx, id)
Retrieve the voice settings for a phone number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePhoneNumber**
> InlineResponse20046 UpdatePhoneNumber(ctx, body, id)
Update the settings for a phone number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body28**](Body28.md)| Updated settings for the phone number | 
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePhoneNumberMessagingSettings**
> InlineResponse20048 UpdatePhoneNumberMessagingSettings(ctx, body, id)
Update the messaging settings for a phone number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body30**](Body30.md)| Updated messaging settings for the phone number | 
  **id** | **string**| Identifies the type of resource. | 

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePhoneNumberVoice**
> InlineResponse20047 UpdatePhoneNumberVoice(ctx, body, id)
Update the voice settings for a phone number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body31**](Body31.md)| Updated voice settings for the phone number | 
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

