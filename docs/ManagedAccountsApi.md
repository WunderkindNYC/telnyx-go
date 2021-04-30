# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManagedAccount**](ManagedAccountsApi.md#CreateManagedAccount) | **Post** /managed_accounts | Create a new managed account.
[**DisableManagedAccount**](ManagedAccountsApi.md#DisableManagedAccount) | **Post** /managed_accounts/{id}/actions/disable | Disables a managed account
[**EnableManagedAccount**](ManagedAccountsApi.md#EnableManagedAccount) | **Post** /managed_accounts/{id}/actions/enable | Enables a managed account
[**ListManagedAccounts**](ManagedAccountsApi.md#ListManagedAccounts) | **Get** /managed_accounts | Lists accounts managed by the current user.
[**RetrieveManagedAccount**](ManagedAccountsApi.md#RetrieveManagedAccount) | **Get** /managed_accounts/{id} | Retrieve a managed account

# **CreateManagedAccount**
> InlineResponse20013 CreateManagedAccount(ctx, body)
Create a new managed account.

Create a new managed account owned by the authenticated user. You need to be explictly approved by Telnyx in order to become a manager account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateManagedAccountRequest**](CreateManagedAccountRequest.md)| Parameters that define the managed account to be created | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableManagedAccount**
> InlineResponse20013 DisableManagedAccount(ctx, id)
Disables a managed account

Disables a managed account, forbidding it to use Telnyx services, including sending or receiving phone calls and SMS messages. Ongoing phone calls will not be affected. The managed account and its sub-users will still be able to log in via the mission control portal.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Managed Account User ID | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableManagedAccount**
> InlineResponse20013 EnableManagedAccount(ctx, id)
Enables a managed account

Enables a managed account and its sub-users to use Telnyx services.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Managed Account User ID | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListManagedAccounts**
> InlineResponse20012 ListManagedAccounts(ctx, optional)
Lists accounts managed by the current user.

Lists the accounts managed by the current user. Users need to be explictly approved by Telnyx in order to become manager accounts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagedAccountsApiListManagedAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagedAccountsApiListManagedAccountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterEmailContains** | **optional.String**| If present, email containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | [default to null]
 **filterEmailEq** | **optional.String**| If present, only returns results with the &lt;code&gt;email&lt;/code&gt; matching exactly the value given. | [default to null]
 **sort** | **optional.String**| Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;email&lt;/code&gt;: sorts the result by the     &lt;code&gt;email&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-email&lt;/code&gt;: sorts the result by the     &lt;code&gt;email&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to created_at]

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveManagedAccount**
> InlineResponse20013 RetrieveManagedAccount(ctx, id)
Retrieve a managed account

Retrieves the details of a single managed account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Managed Account User ID | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

