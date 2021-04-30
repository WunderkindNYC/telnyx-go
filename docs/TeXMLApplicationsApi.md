# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTexmlApplication**](TeXMLApplicationsApi.md#CreateTexmlApplication) | **Post** /texml_applications | Creates a TeXML Application
[**DeleteTexmlApplication**](TeXMLApplicationsApi.md#DeleteTexmlApplication) | **Delete** /texml_applications/{id} | Deletes a TeXML Application
[**FindTexmlApplications**](TeXMLApplicationsApi.md#FindTexmlApplications) | **Get** /texml_applications | List all TeXML Applications
[**GetTexmlApplication**](TeXMLApplicationsApi.md#GetTexmlApplication) | **Get** /texml_applications/{id} | Retrieve a TeXML Application
[**UpdateTexmlApplication**](TeXMLApplicationsApi.md#UpdateTexmlApplication) | **Patch** /texml_applications/{id} | Update a TeXML Application

# **CreateTexmlApplication**
> InlineResponse2019 CreateTexmlApplication(ctx, body)
Creates a TeXML Application

Creates a TeXML Application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateTexmlApplicationRequest**](CreateTexmlApplicationRequest.md)| Parameters that can be set when creating a TeXML Application | 

### Return type

[**InlineResponse2019**](inline_response_201_9.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTexmlApplication**
> InlineResponse2019 DeleteTexmlApplication(ctx, id)
Deletes a TeXML Application

Deletes a TeXML Application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse2019**](inline_response_201_9.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTexmlApplications**
> InlineResponse20044 FindTexmlApplications(ctx, optional)
List all TeXML Applications

Returns a list of your TeXML Applications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeXMLApplicationsApiFindTexmlApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeXMLApplicationsApiFindTexmlApplicationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterFriendlyNameContains** | **optional.String**| If present, applications with &lt;code&gt;friendly_name&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | [default to null]
 **filterOutboundVoiceProfileId** | **optional.String**| Identifies the associated outbound voice profile. | 
 **sort** | **optional.String**| Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to created_at]

### Return type

[**InlineResponse20044**](inline_response_200_44.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTexmlApplication**
> InlineResponse2019 GetTexmlApplication(ctx, id)
Retrieve a TeXML Application

Retrieves the details of an existing TeXML Application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse2019**](inline_response_201_9.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTexmlApplication**
> InlineResponse2019 UpdateTexmlApplication(ctx, body, id)
Update a TeXML Application

Updates settings of an existing TeXML Application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateTexmlApplicationRequest**](UpdateTexmlApplicationRequest.md)| Parameters that can be updated in a TeXML Application | 
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse2019**](inline_response_201_9.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

