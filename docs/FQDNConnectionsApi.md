# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFQDNConnection**](FQDNConnectionsApi.md#CreateFQDNConnection) | **Post** /fqdn_connections | Creates an FQDN connection
[**DeleteFQDNConnection**](FQDNConnectionsApi.md#DeleteFQDNConnection) | **Delete** /fqdn_connections/{id} | Deletes an FQDN connection
[**FindFQDNConnections**](FQDNConnectionsApi.md#FindFQDNConnections) | **Get** /fqdn_connections | List all FQDN connections
[**GetFQDNConnection**](FQDNConnectionsApi.md#GetFQDNConnection) | **Get** /fqdn_connections/{id} | Retrieve a connection
[**UpdateFQDNConnection**](FQDNConnectionsApi.md#UpdateFQDNConnection) | **Patch** /fqdn_connections/{id} | Update a FQDN connection

# **CreateFQDNConnection**
> InlineResponse2012 CreateFQDNConnection(ctx, body)
Creates an FQDN connection

Creates a FQDN connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body4**](Body4.md)| Parameters that can be defined during FQDN connection creation | 

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFQDNConnection**
> InlineResponse2012 DeleteFQDNConnection(ctx, id)
Deletes an FQDN connection

Deletes an FQDN connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| FQDN Connection ID | 

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindFQDNConnections**
> InlineResponse20016 FindFQDNConnections(ctx, optional)
List all FQDN connections

Returns a list of your FQDN connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FQDNConnectionsApiFindFQDNConnectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FQDNConnectionsApiFindFQDNConnectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterConnectionNameContains** | **optional.String**| If present, connections with &lt;code&gt;connection_name&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | [default to null]
 **filterOutboundVoiceProfileId** | **optional.String**| Identifies the associated outbound voice profile. | 
 **sort** | **optional.String**| Specifies the sort order for results. By default sorting direction is ascending. To have the results sorterd in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to created_at]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFQDNConnection**
> InlineResponse2012 GetFQDNConnection(ctx, id)
Retrieve a connection

Retrieves the details of an existing FQDN connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| FQDN Connection ID | 

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFQDNConnection**
> InlineResponse2012 UpdateFQDNConnection(ctx, body, id)
Update a FQDN connection

Updates settings of an existing FQDN connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body5**](Body5.md)| Parameters that can be updated in a FQDN connection | 
  **id** | **string**| FQDN Connection ID | 

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

