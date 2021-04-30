# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SimCardGroupDelete**](SIMCardGroupsApi.md#SimCardGroupDelete) | **Delete** /sim_card_groups/{id} | Delete a SIM card group
[**SimCardGroupUpdate**](SIMCardGroupsApi.md#SimCardGroupUpdate) | **Patch** /sim_card_groups/{id} | Update a SIM card group
[**SimCardGroupsGet**](SIMCardGroupsApi.md#SimCardGroupsGet) | **Get** /sim_card_groups/{id} | Get SIM card gruop
[**SimCardGroupsGetAll**](SIMCardGroupsApi.md#SimCardGroupsGetAll) | **Get** /sim_card_groups | Get all SIM card groups
[**SimCardGroupsPost**](SIMCardGroupsApi.md#SimCardGroupsPost) | **Post** /sim_card_groups | Create a SIM card group

# **SimCardGroupDelete**
> Body34 SimCardGroupDelete(ctx, id)
Delete a SIM card group

Permanently deletes a SIM card group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**Body34**](body_34.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardGroupUpdate**
> Body35 SimCardGroupUpdate(ctx, body, id)
Update a SIM card group

Updates a SIM card group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body35**](Body35.md)|  | 
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**Body35**](body_35.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardGroupsGet**
> Body34 SimCardGroupsGet(ctx, id, optional)
Get SIM card gruop

Returns the details regarding a specific SIM card group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 
 **optional** | ***SIMCardGroupsApiSimCardGroupsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SIMCardGroupsApiSimCardGroupsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSimCards** | **optional.Bool**| It includes the associated SIM card objects in the response when present. | [default to false]
 **includeDataPlans** | **optional.Bool**| It includes the associated data plan objects in the response when present. | [default to false]

### Return type

[**Body34**](body_34.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardGroupsGetAll**
> InlineResponse20055 SimCardGroupsGetAll(ctx, optional)
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
 **includeSimCards** | **optional.Bool**| It includes the associated SIM card objects in the response when present. | [default to false]
 **includeDataPlans** | **optional.Bool**| It includes the associated data plan objects in the response when present. | [default to false]
 **filterName** | [**optional.Interface of string**](.md)| A valid SIM card group name. | 

### Return type

[**InlineResponse20055**](inline_response_200_55.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardGroupsPost**
> Body34 SimCardGroupsPost(ctx, body)
Create a SIM card group

Creates a new SIM card group object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body34**](Body34.md)|  | 

### Return type

[**Body34**](body_34.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

