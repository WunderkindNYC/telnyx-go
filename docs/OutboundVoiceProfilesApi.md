# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOutboundVoiceProfile**](OutboundVoiceProfilesApi.md#CreateOutboundVoiceProfile) | **Post** /outbound_voice_profiles | Create an outbound voice profile
[**DeleteOutboundVoiceProfile**](OutboundVoiceProfilesApi.md#DeleteOutboundVoiceProfile) | **Delete** /outbound_voice_profiles/{id} | Delete an outbound voice profile
[**ListOutboundVoiceProfiles**](OutboundVoiceProfilesApi.md#ListOutboundVoiceProfiles) | **Get** /outbound_voice_profiles | Get all outbound voice profiles
[**RetrieveOutboundVoiceProfile**](OutboundVoiceProfilesApi.md#RetrieveOutboundVoiceProfile) | **Get** /outbound_voice_profiles/{id} | Retrieve an outbound voice profile
[**UpdateOutboundVoiceProfile**](OutboundVoiceProfilesApi.md#UpdateOutboundVoiceProfile) | **Patch** /outbound_voice_profiles/{id} | Updates an existing outbound voice profile.

# **CreateOutboundVoiceProfile**
> OutboundVoiceProfileResponse CreateOutboundVoiceProfile(ctx, body)
Create an outbound voice profile

Create an outbound voice profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateOutboundVoiceProfileRequest**](CreateOutboundVoiceProfileRequest.md)| Parameters that can be defined when creating an outbound voice profile | 

### Return type

[**OutboundVoiceProfileResponse**](Outbound Voice Profile Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOutboundVoiceProfile**
> OutboundVoiceProfileResponse DeleteOutboundVoiceProfile(ctx, id)
Delete an outbound voice profile

Deletes an existing outbound voice profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**OutboundVoiceProfileResponse**](Outbound Voice Profile Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOutboundVoiceProfiles**
> ListOutboundVoiceProfilesResponse ListOutboundVoiceProfiles(ctx, optional)
Get all outbound voice profiles

Get all outbound voice profiles belonging to the user that match the given filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OutboundVoiceProfilesApiListOutboundVoiceProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OutboundVoiceProfilesApiListOutboundVoiceProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterNameContains** | **optional.String**| Optional filter on outbound voice profile name. | 
 **sort** | **optional.String**| Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt;-&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;name&lt;/code&gt;: sorts the result by the     &lt;code&gt;name&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-name&lt;/code&gt;: sorts the result by the     &lt;code&gt;name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; | [default to -created_at]

### Return type

[**ListOutboundVoiceProfilesResponse**](List Outbound Voice Profiles Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveOutboundVoiceProfile**
> OutboundVoiceProfileResponse RetrieveOutboundVoiceProfile(ctx, id)
Retrieve an outbound voice profile

Retrieves the details of an existing outbound voice profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**OutboundVoiceProfileResponse**](Outbound Voice Profile Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOutboundVoiceProfile**
> OutboundVoiceProfileResponse UpdateOutboundVoiceProfile(ctx, body, id)
Updates an existing outbound voice profile.

Updates an existing outbound voice profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateOutboundVoiceProfileRequest**](UpdateOutboundVoiceProfileRequest.md)| Parameters that can be updated on an outbound voice profile | 
  **id** | **string**| Identifies the resource. | 

### Return type

[**OutboundVoiceProfileResponse**](Outbound Voice Profile Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

