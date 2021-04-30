# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFaxApplication**](ProgrammableFaxApplicationsApi.md#CreateFaxApplication) | **Post** /fax_applications | Creates a Fax Application
[**DeleteFaxApplication**](ProgrammableFaxApplicationsApi.md#DeleteFaxApplication) | **Delete** /fax_applications/{id} | Deletes a Fax Application
[**GetFaxApplication**](ProgrammableFaxApplicationsApi.md#GetFaxApplication) | **Get** /fax_applications/{id} | Retrieve a Fax Application
[**ListFaxApplications**](ProgrammableFaxApplicationsApi.md#ListFaxApplications) | **Get** /fax_applications | List all Fax Applications
[**UpdateFaxApplication**](ProgrammableFaxApplicationsApi.md#UpdateFaxApplication) | **Patch** /fax_applications/{id} | Update a Fax Application

# **CreateFaxApplication**
> InlineResponse2011 CreateFaxApplication(ctx, body)
Creates a Fax Application

Creates a new Fax Application based on the parameters sent in the request. The application name and webhook URL are required. Once created, you can assign phone numbers to your application using the `/phone_numbers` endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateFaxApplicationRequest**](CreateFaxApplicationRequest.md)| Parameters that can be set when creating a Fax Application | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFaxApplication**
> InlineResponse2011 DeleteFaxApplication(ctx, id)
Deletes a Fax Application

Permanently deletes a Fax Application. Deletion may be prevented if the application is in use by phone numbers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFaxApplication**
> InlineResponse2011 GetFaxApplication(ctx, id)
Retrieve a Fax Application

Return the details of an existing Fax Application inside the 'data' attribute of the response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFaxApplications**
> InlineResponse2008 ListFaxApplications(ctx, optional)
List all Fax Applications

This endpoint returns a list of your Fax Applications inside the 'data' attribute of the response. You can adjust which applications are listed by using filters. Fax Applications are used to configure how you send and receive faxes using the Programmable Fax API with Telnyx.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProgrammableFaxApplicationsApiListFaxApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgrammableFaxApplicationsApiListFaxApplicationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterApplicationNameContains** | **optional.String**| If present, applications with &lt;code&gt;application_name&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | [default to null]
 **filterOutboundVoiceProfileId** | **optional.String**| Identifies the associated outbound voice profile. | 
 **sort** | **optional.String**| Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-connection_name&lt;/code&gt;: sorts the result by the     &lt;code&gt;connection_name&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to created_at]

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFaxApplication**
> InlineResponse2011 UpdateFaxApplication(ctx, body, id)
Update a Fax Application

Updates settings of an existing Fax Application based on the parameters of the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateFaxApplicationRequest**](UpdateFaxApplicationRequest.md)| Parameters to be updated for the Fax Application | 
  **id** | **string**| Identifies the resource. | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

