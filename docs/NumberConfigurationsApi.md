# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePhoneNumber**](NumberConfigurationsApi.md#DeletePhoneNumber) | **Delete** /phone_numbers/{id} | Delete a phone number
[**EnableEmergencyPhoneNumber**](NumberConfigurationsApi.md#EnableEmergencyPhoneNumber) | **Post** /phone_numbers/{id}/actions/enable_emergency | Enable emergency for a phone number
[**ListPhoneNumbers**](NumberConfigurationsApi.md#ListPhoneNumbers) | **Get** /phone_numbers | List phone numbers
[**ListPhoneNumbersWithMessagingSettings**](NumberConfigurationsApi.md#ListPhoneNumbersWithMessagingSettings) | **Get** /phone_numbers/messaging | List phone numbers with messaging settings
[**ListPhoneNumbersWithVoiceSettings**](NumberConfigurationsApi.md#ListPhoneNumbersWithVoiceSettings) | **Get** /phone_numbers/voice | List phone numbers with voice settings
[**RetrievePhoneNumber**](NumberConfigurationsApi.md#RetrievePhoneNumber) | **Get** /phone_numbers/{id} | Retrieve a phone number
[**RetrievePhoneNumberWithMessagingSettings**](NumberConfigurationsApi.md#RetrievePhoneNumberWithMessagingSettings) | **Get** /phone_numbers/{id}/messaging | Retrieve a phone number with messaging settings
[**RetrievePhoneNumberWithVoiceSettings**](NumberConfigurationsApi.md#RetrievePhoneNumberWithVoiceSettings) | **Get** /phone_numbers/{id}/voice | Retrieve a phone number with voice settings
[**UpdatePhoneNumber**](NumberConfigurationsApi.md#UpdatePhoneNumber) | **Patch** /phone_numbers/{id} | Update a phone number
[**UpdatePhoneNumberWithMessagingSettings**](NumberConfigurationsApi.md#UpdatePhoneNumberWithMessagingSettings) | **Patch** /phone_numbers/{id}/messaging | Update a phone number with messaging settings
[**UpdatePhoneNumberWithVoiceSettings**](NumberConfigurationsApi.md#UpdatePhoneNumberWithVoiceSettings) | **Patch** /phone_numbers/{id}/voice | Update a phone number with voice settings

# **DeletePhoneNumber**
> PhoneNumberResponse1 DeletePhoneNumber(ctx, id)
Delete a phone number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**PhoneNumberResponse1**](Phone Number Response_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableEmergencyPhoneNumber**
> PhoneNumberEnableEmergency EnableEmergencyPhoneNumber(ctx, body, id)
Enable emergency for a phone number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PhoneNumberEnableEmergencyRequest**](PhoneNumberEnableEmergencyRequest.md)|  | 
  **id** | **string**| Identifies the resource. | 

### Return type

[**PhoneNumberEnableEmergency**](Phone Number Enable Emergency.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPhoneNumbers**
> ListPhoneNumbersResponse ListPhoneNumbers(ctx, optional)
List phone numbers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberConfigurationsApiListPhoneNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberConfigurationsApiListPhoneNumbersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterTag** | **optional.String**| Filter by phone number tags. | 
 **filterPhoneNumber** | **optional.String**| Filter by phone number. Requires at least three digits.              Non-numerical characters will result in no values being returned. | 
 **filterStatus** | **optional.String**| Filter by phone number status. | 
 **filterVoiceConnectionNameContains** | **optional.String**| Filter contains connection name. Requires at least three characters. | 
 **filterVoiceConnectionNameStartsWith** | **optional.String**| Filter starts with connection name. Requires at least three characters. | 
 **filterVoiceConnectionNameEndsWith** | **optional.String**| Filter ends with connection name. Requires at least three characters. | 
 **filterVoiceConnectionNameEq** | **optional.String**| Filter by connection name. | 
 **filterUsagePaymentMethod** | **optional.String**| Filter by usage_payment_method. | 
 **filterBillingGroupId** | **optional.String**| Filter by the billing_group_id associated with phone numbers. To filter to only phone numbers that have no billing group associated them, set the value of this filter to the string &#x27;null&#x27;. | 
 **filterEmergencyAddressId** | **optional.String**| Filter by the emergency_address_id associated with phone numbers. To filter only phone numbers that have no emergency address associated with them, set the value of this filter to the string &#x27;null&#x27;. | 
 **filterCustomerReference** | **optional.String**| Filter numbers via the customer_reference set. | 
 **sort** | **optional.String**| Specifies the sort order for results. If not given, results are sorted by created_at in descending order. | 

### Return type

[**ListPhoneNumbersResponse**](List Phone Numbers Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPhoneNumbersWithMessagingSettings**
> ListMessagingSettingsResponse ListPhoneNumbersWithMessagingSettings(ctx, optional)
List phone numbers with messaging settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberConfigurationsApiListPhoneNumbersWithMessagingSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberConfigurationsApiListPhoneNumbersWithMessagingSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**ListMessagingSettingsResponse**](List Messaging Settings Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPhoneNumbersWithVoiceSettings**
> ListPhoneNumbersWithVoiceSettingsResponse ListPhoneNumbersWithVoiceSettings(ctx, optional)
List phone numbers with voice settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberConfigurationsApiListPhoneNumbersWithVoiceSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberConfigurationsApiListPhoneNumbersWithVoiceSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterPhoneNumber** | **optional.String**| Filter by phone number. Requires at least three digits.              Non-numerical characters will result in no values being returned. | 
 **filterConnectionNameContains** | **optional.String**| Filter contains connection name. Requires at least three characters. | 
 **filterCustomerReference** | **optional.String**| Filter numbers via the customer_reference set. | 
 **filterUsagePaymentMethod** | **optional.String**| Filter by usage_payment_method. | 
 **sort** | **optional.String**| Specifies the sort order for results. If not given, results are sorted by created_at in descending order. | 

### Return type

[**ListPhoneNumbersWithVoiceSettingsResponse**](List Phone Numbers With Voice Settings Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrievePhoneNumber**
> PhoneNumberResponse RetrievePhoneNumber(ctx, id)
Retrieve a phone number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**PhoneNumberResponse**](Phone Number Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrievePhoneNumberWithMessagingSettings**
> RetrieveMessagingSettingsResponse RetrievePhoneNumberWithMessagingSettings(ctx, id)
Retrieve a phone number with messaging settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the type of resource. | 

### Return type

[**RetrieveMessagingSettingsResponse**](Retrieve Messaging Settings Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrievePhoneNumberWithVoiceSettings**
> RetrievePhoneNumberVoiceResponse RetrievePhoneNumberWithVoiceSettings(ctx, id)
Retrieve a phone number with voice settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**RetrievePhoneNumberVoiceResponse**](Retrieve Phone Number Voice Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePhoneNumber**
> PhoneNumberResponse UpdatePhoneNumber(ctx, body, id)
Update a phone number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePhoneNumberRequest**](UpdatePhoneNumberRequest.md)| Updated settings for the phone number. | 
  **id** | **string**| Identifies the resource. | 

### Return type

[**PhoneNumberResponse**](Phone Number Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePhoneNumberWithMessagingSettings**
> RetrieveMessagingSettingsResponse UpdatePhoneNumberWithMessagingSettings(ctx, body, id)
Update a phone number with messaging settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePhoneNumberMessagingSettingsRequest**](UpdatePhoneNumberMessagingSettingsRequest.md)| Updated messaging settings for the phone number | 
  **id** | **string**| Identifies the type of resource. | 

### Return type

[**RetrieveMessagingSettingsResponse**](Retrieve Messaging Settings Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePhoneNumberWithVoiceSettings**
> RetrievePhoneNumberVoiceResponse UpdatePhoneNumberWithVoiceSettings(ctx, body, id)
Update a phone number with voice settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePhoneNumberVoiceSettingsRequest**](UpdatePhoneNumberVoiceSettingsRequest.md)| Updated voice settings for the phone number. | 
  **id** | **string**| Identifies the resource. | 

### Return type

[**RetrievePhoneNumberVoiceResponse**](Retrieve Phone Number Voice Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

