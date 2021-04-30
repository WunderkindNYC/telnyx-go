# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFax**](ProgrammableFaxCommandsApi.md#DeleteFax) | **Delete** /faxes/{id} | Delete a fax
[**ListFaxes**](ProgrammableFaxCommandsApi.md#ListFaxes) | **Get** /faxes | View a list of faxes
[**RefreshFax**](ProgrammableFaxCommandsApi.md#RefreshFax) | **Post** /faxes/{id}/actions/refresh | Refresh a fax
[**SendFax**](ProgrammableFaxCommandsApi.md#SendFax) | **Post** /faxes | Send a fax
[**ViewFax**](ProgrammableFaxCommandsApi.md#ViewFax) | **Get** /faxes/{id} | View a fax

# **DeleteFax**
> DeleteFax(ctx, id)
Delete a fax

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The unique identifier of a fax. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFaxes**
> InlineResponse2009 ListFaxes(ctx, optional)
View a list of faxes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProgrammableFaxCommandsApiListFaxesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgrammableFaxCommandsApiListFaxesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCreatedAtGte** | **optional.Time**| ISO 8601 date time for filtering faxes created after or on that date | 
 **filterCreatedAtGt** | **optional.Time**| ISO 8601 date time for filtering faxes created after that date | 
 **filterCreatedAtLte** | **optional.Time**| ISO 8601 formatted date time for filtering faxes created on or before that date | 
 **filterCreatedAtLt** | **optional.Time**| ISO 8601 formatted date time for filtering faxes created before that date | 
 **filterDirectionEq** | **optional.String**| The direction, inbound or outbound, for filtering faxes sent from this account | 
 **filterFromEq** | **optional.String**| The phone number, in E.164 format for filtering faxes sent from this number | 
 **pageSize** | **optional.Int32**| Number of fax resourcxes for the single page returned | 
 **pageNumber** | **optional.Int32**| Number of the page to be retrieved | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshFax**
> InlineResponse20010 RefreshFax(ctx, id)
Refresh a fax

This endpoint refreshes the media_url expiration for inbound faxes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The unique identifier of a fax. | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendFax**
> InlineResponse2022 SendFax(ctx, body)
Send a fax

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendFaxRequest**](SendFaxRequest.md)| Send fax request | 

### Return type

[**InlineResponse2022**](inline_response_202_2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewFax**
> InlineResponse2022 ViewFax(ctx, id)
View a fax

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The unique identifier of a fax. | 

### Return type

[**InlineResponse2022**](inline_response_202_2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

