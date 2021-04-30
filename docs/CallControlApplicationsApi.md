# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCallControlApplication**](CallControlApplicationsApi.md#CreateCallControlApplication) | **Post** /call_control_applications | Creates a Call Control application
[**DeleteCallControlApplication**](CallControlApplicationsApi.md#DeleteCallControlApplication) | **Delete** /call_control_applications/{id} | Deletes a Call Control application
[**FindCallControlApplications**](CallControlApplicationsApi.md#FindCallControlApplications) | **Get** /call_control_applications | List all Call Control applications
[**GetCallControlApplication**](CallControlApplicationsApi.md#GetCallControlApplication) | **Get** /call_control_applications/{id} | Retrieve a Call Control application
[**UpdateCallControlApplication**](CallControlApplicationsApi.md#UpdateCallControlApplication) | **Patch** /call_control_applications/{id} | Update a Call Control application

# **CreateCallControlApplication**
> InlineResponse201 CreateCallControlApplication(ctx, body)
Creates a Call Control application

Creates a Call Control application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatableAttributesForCallControlApplications**](UpdatableAttributesForCallControlApplications.md)| Parameters that can be set when creating a Call Control application | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCallControlApplication**
> InlineResponse201 DeleteCallControlApplication(ctx, id)
Deletes a Call Control application

Deletes a Call Control application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindCallControlApplications**
> InlineResponse2006 FindCallControlApplications(ctx, optional)
List all Call Control applications

Returns a list of your Call Control applications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CallControlApplicationsApiFindCallControlApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CallControlApplicationsApiFindCallControlApplicationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterConnectionNameContains** | **optional.String**| If present, connections with &lt;code&gt;connection_name&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | [default to null]
 **filterOutboundVoiceProfileId** | **optional.String**| Identifies the associated outbound voice profile. | 
 **sort** | **optional.String**| Specifies the sort order for results. By default sorting direction is ascending. To have the results sorterd in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to created_at]

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCallControlApplication**
> InlineResponse201 GetCallControlApplication(ctx, id)
Retrieve a Call Control application

Retrieves the details of an existing Call Control application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCallControlApplication**
> InlineResponse201 UpdateCallControlApplication(ctx, body, id)
Update a Call Control application

Updates settings of an existing Call Control application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatableAttributesForCallControlApplications1**](UpdatableAttributesForCallControlApplications1.md)| Parameters that can be updated in a Call Control application | 
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

