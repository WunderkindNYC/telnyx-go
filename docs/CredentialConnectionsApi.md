# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredentialConnection**](CredentialConnectionsApi.md#CreateCredentialConnection) | **Post** /credential_connections | Create a credential connection
[**DeleteCredentialConnection**](CredentialConnectionsApi.md#DeleteCredentialConnection) | **Delete** /credential_connections/{id} | Delete a credential connection
[**ListCredentialConnections**](CredentialConnectionsApi.md#ListCredentialConnections) | **Get** /credential_connections | List credential connections
[**RetrieveCredentialConnection**](CredentialConnectionsApi.md#RetrieveCredentialConnection) | **Get** /credential_connections/{id} | Retrieve a credential connection
[**UpdateCredentialConnection**](CredentialConnectionsApi.md#UpdateCredentialConnection) | **Patch** /credential_connections/{id} | Update a credential connection

# **CreateCredentialConnection**
> CredentialConnectionResponse CreateCredentialConnection(ctx, body)
Create a credential connection

Creates a credential connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCredentialConnectionRequest**](CreateCredentialConnectionRequest.md)| Parameters that can be defined during credential connection creation | 

### Return type

[**CredentialConnectionResponse**](Credential Connection Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCredentialConnection**
> CredentialConnectionResponse DeleteCredentialConnection(ctx, id)
Delete a credential connection

Deletes an existing credential connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**CredentialConnectionResponse**](Credential Connection Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCredentialConnections**
> ListCredentialConnectionsResponse ListCredentialConnections(ctx, optional)
List credential connections

Returns a list of your credential connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CredentialConnectionsApiListCredentialConnectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialConnectionsApiListCredentialConnectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterConnectionNameContains** | **optional.String**| If present, connections with &lt;code&gt;connection_name&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | [default to null]
 **filterOutboundOutboundVoiceProfileId** | **optional.String**| Identifies the associated outbound voice profile. | 
 **sort** | **optional.String**| Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to created_at]

### Return type

[**ListCredentialConnectionsResponse**](List Credential Connections Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCredentialConnection**
> CredentialConnectionResponse RetrieveCredentialConnection(ctx, id)
Retrieve a credential connection

Retrieves the details of an existing credential connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**CredentialConnectionResponse**](Credential Connection Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCredentialConnection**
> CredentialConnectionResponse UpdateCredentialConnection(ctx, body, id)
Update a credential connection

Updates settings of an existing credential connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateCredentialConnectionRequest**](UpdateCredentialConnectionRequest.md)| Parameters that can be updated in a credential connection | 
  **id** | **string**| Identifies the resource. | 

### Return type

[**CredentialConnectionResponse**](Credential Connection Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

