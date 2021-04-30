# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePhoneNumbersJobDeletePhoneNumbers**](NumberBackgroundJobsApi.md#CreatePhoneNumbersJobDeletePhoneNumbers) | **Post** /phone_numbers/jobs/delete_phone_numbers | Delete a batch of numbers
[**CreatePhoneNumbersJobUpdateEmergencySettings**](NumberBackgroundJobsApi.md#CreatePhoneNumbersJobUpdateEmergencySettings) | **Post** /phone_numbers/jobs/update_emergency_settings | Update the emergency settings from a batch of numbers
[**CreatePhoneNumbersJobUpdatePhoneNumber**](NumberBackgroundJobsApi.md#CreatePhoneNumbersJobUpdatePhoneNumber) | **Post** /phone_numbers/jobs/update_phone_numbers | Update a batch of numbers
[**ListPhoneNumbersJobs**](NumberBackgroundJobsApi.md#ListPhoneNumbersJobs) | **Get** /phone_numbers/jobs | Lists the phone numbers jobs
[**RetrievePhoneNumbersJob**](NumberBackgroundJobsApi.md#RetrievePhoneNumbersJob) | **Get** /phone_numbers/jobs/{id} | Retrieve a phone numbers job

# **CreatePhoneNumbersJobDeletePhoneNumbers**
> PhoneNumbersJobDeletePhoneNumbers CreatePhoneNumbersJobDeletePhoneNumbers(ctx, body)
Delete a batch of numbers

Creates a new background job to delete a batch of numbers. At most one thousand numbers can be updated per API call.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PhoneNumbersJobDeletePhoneNumbersRequest**](PhoneNumbersJobDeletePhoneNumbersRequest.md)|  | 

### Return type

[**PhoneNumbersJobDeletePhoneNumbers**](Phone Numbers Job Delete Phone Numbers.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePhoneNumbersJobUpdateEmergencySettings**
> PhoneNumbersEnableEmergency CreatePhoneNumbersJobUpdateEmergencySettings(ctx, body)
Update the emergency settings from a batch of numbers

Creates a background job to update the emergency settings of a collection of phone numbers. At most one thousand numbers can be updated per API call.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PhoneNumbersJobUpdateEmergencySettingsRequest**](PhoneNumbersJobUpdateEmergencySettingsRequest.md)|  | 

### Return type

[**PhoneNumbersEnableEmergency**](Phone Numbers Enable Emergency.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePhoneNumbersJobUpdatePhoneNumber**
> PhoneNumbersJobUpdatePhoneNumbers CreatePhoneNumbersJobUpdatePhoneNumber(ctx, body)
Update a batch of numbers

Creates a new background job to update a batch of numbers. At most one thousand numbers can be updated per API call. At least one of the updateable fields must be submitted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PhoneNumbersJobUpdatePhoneNumbersRequest**](PhoneNumbersJobUpdatePhoneNumbersRequest.md)|  | 

### Return type

[**PhoneNumbersJobUpdatePhoneNumbers**](Phone Numbers Job Update Phone Numbers.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPhoneNumbersJobs**
> ListPhoneNumbersBackgroundJobsResponse ListPhoneNumbersJobs(ctx, optional)
Lists the phone numbers jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberBackgroundJobsApiListPhoneNumbersJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberBackgroundJobsApiListPhoneNumbersJobsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterType** | **optional.String**| Filter the phone number jobs by type. | 
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **sort** | **optional.String**| Specifies the sort order for results. If not given, results are sorted by created_at in descending order. | 

### Return type

[**ListPhoneNumbersBackgroundJobsResponse**](List Phone Numbers Background Jobs Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrievePhoneNumbersJob**
> PhoneNumbersJob RetrievePhoneNumbersJob(ctx, id)
Retrieve a phone numbers job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the Phone Numbers Job. | 

### Return type

[**PhoneNumbersJob**](Phone Numbers Job.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

