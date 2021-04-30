# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMessagingPhoneNumbers**](PhoneNumbersApi.md#ListMessagingPhoneNumbers) | **Get** /messaging_phone_numbers | List all messaging capable phone numbers
[**RetrieveMessagingPhoneNumber**](PhoneNumbersApi.md#RetrieveMessagingPhoneNumber) | **Get** /messaging_phone_numbers/{id} | Retrieve a messaging phone number
[**UpdateMessagingPhoneNumber**](PhoneNumbersApi.md#UpdateMessagingPhoneNumber) | **Patch** /messaging_phone_numbers/{id} | Update a messaging phone number

# **ListMessagingPhoneNumbers**
> InlineResponse20025 ListMessagingPhoneNumbers(ctx, optional)
List all messaging capable phone numbers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PhoneNumbersApiListMessagingPhoneNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PhoneNumbersApiListMessagingPhoneNumbersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterMessagingProfileId** | **optional.String**| Filter by Messaging Profile ID. Use the string &#x60;null&#x60; for phone numbers without assigned profiles. A synonym for the &#x60;/messaging_profiles/{id}/phone_number&#x60; endpoint when querying about an extant profile. | 
 **filterPhoneNumber** | **optional.String**| A comma separated list of phone numbers to return | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveMessagingPhoneNumber**
> InlineResponse20026 RetrieveMessagingPhoneNumber(ctx, id)
Retrieve a messaging phone number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the phone number | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMessagingPhoneNumber**
> InlineResponse20026 UpdateMessagingPhoneNumber(ctx, body, id)
Update a messaging phone number

Update the settings for a specific number. To unbind a number from, a profile, set the `messaging_profile_id` to `null` or an empty string.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body19**](Body19.md)| Messaging Phone Number Update | 
  **id** | **string**| The id of the phone number | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

