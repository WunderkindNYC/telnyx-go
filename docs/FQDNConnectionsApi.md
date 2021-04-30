# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFqdnConnection**](FQDNConnectionsApi.md#CreateFqdnConnection) | **Post** /fqdn_connections | Create an Fqdn connection
[**DeleteFqdnConnection**](FQDNConnectionsApi.md#DeleteFqdnConnection) | **Delete** /fqdn_connections/{id} | Delete an Fqdn connection
[**ListFqdnConnections**](FQDNConnectionsApi.md#ListFqdnConnections) | **Get** /fqdn_connections | List Fqdn connections
[**RetrieveFqdnConnection**](FQDNConnectionsApi.md#RetrieveFqdnConnection) | **Get** /fqdn_connections/{id} | Retrieve an Fqdn connection
[**UpdateFqdnConnection**](FQDNConnectionsApi.md#UpdateFqdnConnection) | **Patch** /fqdn_connections/{id} | Update an Fqdn connection

# **CreateFqdnConnection**
> FqdnConnectionResponse CreateFqdnConnection(ctx, body)
Create an Fqdn connection

Creates a FQDN connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateFqdnConnectionRequest**](CreateFqdnConnectionRequest.md)| Parameters that can be defined during FQDN connection creation | 

### Return type

[**FqdnConnectionResponse**](Fqdn Connection Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFqdnConnection**
> FqdnConnectionResponse DeleteFqdnConnection(ctx, id)
Delete an Fqdn connection

Deletes an FQDN connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| FQDN Connection ID | 

### Return type

[**FqdnConnectionResponse**](Fqdn Connection Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFqdnConnections**
> ListFqdnConnectionsResponse ListFqdnConnections(ctx, optional)
List Fqdn connections

Returns a list of your FQDN connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FQDNConnectionsApiListFqdnConnectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FQDNConnectionsApiListFqdnConnectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterConnectionNameContains** | **optional.String**| If present, connections with &lt;code&gt;connection_name&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | [default to null]
 **sort** | **optional.String**| Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to created_at]

### Return type

[**ListFqdnConnectionsResponse**](List Fqdn Connections Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveFqdnConnection**
> FqdnConnectionResponse RetrieveFqdnConnection(ctx, id)
Retrieve an Fqdn connection

Retrieves the details of an existing FQDN connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| FQDN Connection ID | 

### Return type

[**FqdnConnectionResponse**](Fqdn Connection Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFqdnConnection**
> FqdnConnectionResponse UpdateFqdnConnection(ctx, body, id)
Update an Fqdn connection

Updates settings of an existing FQDN connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateFqdnConnectionRequest**](UpdateFqdnConnectionRequest.md)| Parameters that can be updated in a FQDN connection | 
  **id** | **string**| FQDN Connection ID | 

### Return type

[**FqdnConnectionResponse**](Fqdn Connection Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

