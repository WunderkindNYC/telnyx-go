# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCallControlApplication**](CallControlApplicationsApi.md#CreateCallControlApplication) | **Post** /call_control_applications | Create a call control application
[**DeleteCallControlApplication**](CallControlApplicationsApi.md#DeleteCallControlApplication) | **Delete** /call_control_applications/{id} | Delete a call control application
[**ListCallControlApplications**](CallControlApplicationsApi.md#ListCallControlApplications) | **Get** /call_control_applications | List call control applications
[**RetrieveCallControlApplication**](CallControlApplicationsApi.md#RetrieveCallControlApplication) | **Get** /call_control_applications/{id} | Retrieve a call control application
[**UpdateCallControlApplication**](CallControlApplicationsApi.md#UpdateCallControlApplication) | **Patch** /call_control_applications/{id} | Update a call control application

# **CreateCallControlApplication**
> CallControlApplicationResponse CreateCallControlApplication(ctx, body)
Create a call control application

Create a call control application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCallControlApplicationRequest**](CreateCallControlApplicationRequest.md)| Create call control application request. | 

### Return type

[**CallControlApplicationResponse**](Call Control Application Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCallControlApplication**
> CallControlApplicationResponse DeleteCallControlApplication(ctx, id)
Delete a call control application

Deletes a call control application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**CallControlApplicationResponse**](Call Control Application Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCallControlApplications**
> ListCallControlApplicationsResponse ListCallControlApplications(ctx, optional)
List call control applications

Return a list of call control applications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CallControlApplicationsApiListCallControlApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CallControlApplicationsApiListCallControlApplicationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterApplicationNameContains** | **optional.String**| If present, applications with &lt;code&gt;application_name&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | [default to null]
 **filterOutboundVoiceProfileId** | **optional.String**| Identifies the associated outbound voice profile. | 
 **sort** | **optional.String**| Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to created_at]

### Return type

[**ListCallControlApplicationsResponse**](List Call Control Applications Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCallControlApplication**
> CallControlApplicationResponse RetrieveCallControlApplication(ctx, id)
Retrieve a call control application

Retrieves the details of an existing call control application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**CallControlApplicationResponse**](Call Control Application Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCallControlApplication**
> CallControlApplicationResponse UpdateCallControlApplication(ctx, body, id)
Update a call control application

Updates settings of an existing call control application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateCallControlApplicationRequest**](UpdateCallControlApplicationRequest.md)| Update call control application request. | 
  **id** | **string**| Identifies the resource. | 

### Return type

[**CallControlApplicationResponse**](Call Control Application Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

