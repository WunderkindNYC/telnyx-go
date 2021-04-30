# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessagingProfile**](MessagingProfilesApi.md#CreateMessagingProfile) | **Post** /messaging_profiles | Create a messaging profile
[**DeleteMessagingProfile**](MessagingProfilesApi.md#DeleteMessagingProfile) | **Delete** /messaging_profiles/{id} | Delete a messaging profile
[**ListMessagingProfilePhoneNumbers**](MessagingProfilesApi.md#ListMessagingProfilePhoneNumbers) | **Get** /messaging_profiles/{id}/phone_numbers | List all phone numbers associated with a messaging profile
[**ListMessagingProfileShortCodes**](MessagingProfilesApi.md#ListMessagingProfileShortCodes) | **Get** /messaging_profiles/{id}/short_codes | List all short codes associated with a messaging profile
[**ListMessagingProfiles**](MessagingProfilesApi.md#ListMessagingProfiles) | **Get** /messaging_profiles | List all messaging profiles
[**RetrieveMessagingProfile**](MessagingProfilesApi.md#RetrieveMessagingProfile) | **Get** /messaging_profiles/{id} | Retrieve a messaging profile
[**UpdateMessagingProfile**](MessagingProfilesApi.md#UpdateMessagingProfile) | **Patch** /messaging_profiles/{id} | Update a messaging profile

# **CreateMessagingProfile**
> InlineResponse20028 CreateMessagingProfile(ctx, body)
Create a messaging profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body20**](Body20.md)| New Messaging Profile object | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMessagingProfile**
> InlineResponse20028 DeleteMessagingProfile(ctx, id)
Delete a messaging profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the messaging profile to retrieve | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMessagingProfilePhoneNumbers**
> InlineResponse20025 ListMessagingProfilePhoneNumbers(ctx, id, optional)
List all phone numbers associated with a messaging profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the messaging profile to retrieve | 
 **optional** | ***MessagingProfilesApiListMessagingProfilePhoneNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagingProfilesApiListMessagingProfilePhoneNumbersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMessagingProfileShortCodes**
> InlineResponse20029 ListMessagingProfileShortCodes(ctx, id, optional)
List all short codes associated with a messaging profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the messaging profile to retrieve | 
 **optional** | ***MessagingProfilesApiListMessagingProfileShortCodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagingProfilesApiListMessagingProfileShortCodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMessagingProfiles**
> InlineResponse20027 ListMessagingProfiles(ctx, optional)
List all messaging profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagingProfilesApiListMessagingProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagingProfilesApiListMessagingProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveMessagingProfile**
> InlineResponse20028 RetrieveMessagingProfile(ctx, id)
Retrieve a messaging profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the messaging profile to retrieve | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMessagingProfile**
> InlineResponse20028 UpdateMessagingProfile(ctx, body, id)
Update a messaging profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body21**](Body21.md)| New Messaging Profile object | 
  **id** | [**string**](.md)| The id of the messaging profile to retrieve | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

