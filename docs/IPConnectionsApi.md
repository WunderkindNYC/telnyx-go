# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIPConnection**](IPConnectionsApi.md#CreateIPConnection) | **Post** /ip_connections | Creates an IP connection
[**DeleteIPConnection**](IPConnectionsApi.md#DeleteIPConnection) | **Delete** /ip_connections/{id} | Deletes an IP connection
[**FindConnections**](IPConnectionsApi.md#FindConnections) | **Get** /ip_connections | List all connections
[**GetIPConnection**](IPConnectionsApi.md#GetIPConnection) | **Get** /ip_connections/{id} | Retrieve a connection
[**UpdateIPConnection**](IPConnectionsApi.md#UpdateIPConnection) | **Patch** /ip_connections/{id} | Updates an IP connection

# **CreateIPConnection**
> InlineResponse2014 CreateIPConnection(ctx, body)
Creates an IP connection

Creates an IP connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body8**](Body8.md)| Parameters that can be defined during IP connection creation | 

### Return type

[**InlineResponse2014**](inline_response_201_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPConnection**
> InlineResponse2014 DeleteIPConnection(ctx, id)
Deletes an IP connection

Deletes an existing IP connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the type of resource. | 

### Return type

[**InlineResponse2014**](inline_response_201_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindConnections**
> InlineResponse20018 FindConnections(ctx, optional)
List all connections

Returns a list of your IP connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPConnectionsApiFindConnectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPConnectionsApiFindConnectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterConnectionNameContains** | **optional.String**| If present, connections with &lt;code&gt;connection_name&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | [default to null]
 **filterOutboundVoiceProfileId** | **optional.String**| Identifies the associated outbound voice profile. | 
 **sort** | **optional.String**| Specifies the sort order for results. By default sorting direction is ascending. To have the results sorterd in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to created_at]

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPConnection**
> InlineResponse2014 GetIPConnection(ctx, id)
Retrieve a connection

Retrieves the details of an existing ip connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| IP Connection ID | 

### Return type

[**InlineResponse2014**](inline_response_201_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIPConnection**
> InlineResponse2014 UpdateIPConnection(ctx, body, id)
Updates an IP connection

Updates settings of an existing IP connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body9**](Body9.md)| Parameters that can be updated in a IP connection | 
  **id** | **string**| Identifies the type of resource. | 

### Return type

[**InlineResponse2014**](inline_response_201_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

