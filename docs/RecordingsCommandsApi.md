# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRecordings**](RecordingsCommandsApi.md#ListRecordings) | **Get** /recordings | List recordings
[**RetrieveRecording**](RecordingsCommandsApi.md#RetrieveRecording) | **Get** /recordings/{id} | Retrieve a recording

# **ListRecordings**
> ListRecordingsResponse ListRecordings(ctx, optional)
List recordings

Lists recordings for the authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecordingsCommandsApiListRecordingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecordingsCommandsApiListRecordingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterConferenceId** | **optional.String**| Return only recordings associated with a given conference | 
 **filterCreatedAtGte** | **optional.String**| Return only recordings created later than or at given ISO 8601 datetime | 
 **filterCreatedAtLte** | **optional.String**| Return only recordings created earlier than or at given ISO 8601 datetime | 
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**ListRecordingsResponse**](List Recordings Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveRecording**
> GetRecordingResponse RetrieveRecording(ctx, id)
Retrieve a recording

Retrieve a recording from the authenticated user's recordings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Uniquely identifies the recording | 

### Return type

[**GetRecordingResponse**](Get Recording Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

