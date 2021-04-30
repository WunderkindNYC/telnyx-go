# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessagingProfile**](MessagingProfilesApi.md#CreateMessagingProfile) | **Post** /messaging_profiles | Create a messaging profile
[**DeleteMessagingProfile**](MessagingProfilesApi.md#DeleteMessagingProfile) | **Delete** /messaging_profiles/{id} | Delete a messaging profile
[**ListMessagingProfileMetrics**](MessagingProfilesApi.md#ListMessagingProfileMetrics) | **Get** /messaging_profile_metrics | List messaging profile metrics
[**ListMessagingProfilePhoneNumbers**](MessagingProfilesApi.md#ListMessagingProfilePhoneNumbers) | **Get** /messaging_profiles/{id}/phone_numbers | List phone numbers associated with a messaging profile
[**ListMessagingProfileShortCodes**](MessagingProfilesApi.md#ListMessagingProfileShortCodes) | **Get** /messaging_profiles/{id}/short_codes | List short codes associated with a messaging profile
[**ListMessagingProfiles**](MessagingProfilesApi.md#ListMessagingProfiles) | **Get** /messaging_profiles | List messaging profiles
[**RetrieveMessagingProfile**](MessagingProfilesApi.md#RetrieveMessagingProfile) | **Get** /messaging_profiles/{id} | Retrieve a messaging profile
[**RetrieveMessagingProfileDetailedMetrics**](MessagingProfilesApi.md#RetrieveMessagingProfileDetailedMetrics) | **Get** /messaging_profiles/{id}/metrics | Retrieve messaging profile metrics
[**UpdateMessagingProfile**](MessagingProfilesApi.md#UpdateMessagingProfile) | **Patch** /messaging_profiles/{id} | Update a messaging profile

# **CreateMessagingProfile**
> MessagingProfileResponse CreateMessagingProfile(ctx, body)
Create a messaging profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateMessagingProfileRequest**](CreateMessagingProfileRequest.md)| New Messaging Profile object | 

### Return type

[**MessagingProfileResponse**](Messaging Profile Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMessagingProfile**
> MessagingProfileResponse DeleteMessagingProfile(ctx, id)
Delete a messaging profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the messaging profile to retrieve | 

### Return type

[**MessagingProfileResponse**](Messaging Profile Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMessagingProfileMetrics**
> ListMessagingProfileMetricsResponse ListMessagingProfileMetrics(ctx, optional)
List messaging profile metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagingProfilesApiListMessagingProfileMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagingProfilesApiListMessagingProfileMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **id** | [**optional.Interface of string**](.md)| The id of the messaging profile(s) to retrieve | 
 **timeFrame** | **optional.String**| The timeframe for which you&#x27;d like to retrieve metrics. | [default to 24h]

### Return type

[**ListMessagingProfileMetricsResponse**](List Messaging Profile Metrics Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMessagingProfilePhoneNumbers**
> ListMessagingProfilePhoneNumbersResponse ListMessagingProfilePhoneNumbers(ctx, id, optional)
List phone numbers associated with a messaging profile

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

[**ListMessagingProfilePhoneNumbersResponse**](List Messaging Profile Phone Numbers Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMessagingProfileShortCodes**
> ListMessagingProfileShortCodesResponse ListMessagingProfileShortCodes(ctx, id, optional)
List short codes associated with a messaging profile

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

[**ListMessagingProfileShortCodesResponse**](List Messaging Profile Short Codes Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMessagingProfiles**
> ListMessagingProfilesResponse ListMessagingProfiles(ctx, optional)
List messaging profiles

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

[**ListMessagingProfilesResponse**](List Messaging Profiles Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveMessagingProfile**
> MessagingProfileResponse RetrieveMessagingProfile(ctx, id)
Retrieve a messaging profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the messaging profile to retrieve | 

### Return type

[**MessagingProfileResponse**](Messaging Profile Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveMessagingProfileDetailedMetrics**
> RetrieveMessagingProfileMetricsResponse RetrieveMessagingProfileDetailedMetrics(ctx, id, optional)
Retrieve messaging profile metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the messaging profile to retrieve | 
 **optional** | ***MessagingProfilesApiRetrieveMessagingProfileDetailedMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagingProfilesApiRetrieveMessagingProfileDetailedMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeFrame** | **optional.String**| The timeframe for which you&#x27;d like to retrieve metrics. | [default to 24h]

### Return type

[**RetrieveMessagingProfileMetricsResponse**](Retrieve Messaging Profile Metrics Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMessagingProfile**
> MessagingProfileResponse UpdateMessagingProfile(ctx, body, id)
Update a messaging profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateMessagingProfileRequest**](UpdateMessagingProfileRequest.md)| New Messaging Profile object | 
  **id** | [**string**](.md)| The id of the messaging profile to retrieve | 

### Return type

[**MessagingProfileResponse**](Messaging Profile Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

