# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SimCardGroupDelete**](SIMCardGroupsApi.md#SimCardGroupDelete) | **Delete** /sim_card_groups/{id} | Delete a SIM card group
[**SimCardGroupUpdate**](SIMCardGroupsApi.md#SimCardGroupUpdate) | **Patch** /sim_card_groups/{id} | Update a SIM card group
[**SimCardGroupsGet**](SIMCardGroupsApi.md#SimCardGroupsGet) | **Get** /sim_card_groups/{id} | Get SIM card group
[**SimCardGroupsGetAll**](SIMCardGroupsApi.md#SimCardGroupsGetAll) | **Get** /sim_card_groups | Get all SIM card groups
[**SimCardGroupsPost**](SIMCardGroupsApi.md#SimCardGroupsPost) | **Post** /sim_card_groups | Create a SIM card group

# **SimCardGroupDelete**
> InlineResponse20037 SimCardGroupDelete(ctx, id)
Delete a SIM card group

Permanently deletes a SIM card group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardGroupUpdate**
> InlineResponse20037 SimCardGroupUpdate(ctx, body, id)
Update a SIM card group

Updates a SIM card group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SimCardGroupPatch**](SimCardGroupPatch.md)|  | 
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardGroupsGet**
> InlineResponse20037 SimCardGroupsGet(ctx, id)
Get SIM card group

Returns the details regarding a specific SIM card group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardGroupsGetAll**
> InlineResponse20036 SimCardGroupsGetAll(ctx, optional)
Get all SIM card groups

Get all SIM card groups belonging to the user that match the given filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SIMCardGroupsApiSimCardGroupsGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SIMCardGroupsApiSimCardGroupsGetAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterName** | [**optional.Interface of string**](.md)| A valid SIM card group name. | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardGroupsPost**
> InlineResponse20037 SimCardGroupsPost(ctx, body)
Create a SIM card group

Creates a new SIM card group object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SimCardGroupCreate**](SimCardGroupCreate.md)|  | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

