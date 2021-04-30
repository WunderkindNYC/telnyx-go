# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NumberLookup**](NumberLookupApi.md#NumberLookup) | **Get** /number_lookup/{phone_number} | Lookup phone number data

# **NumberLookup**
> InlineResponse20023 NumberLookup(ctx, phoneNumber, optional)
Lookup phone number data

Returns information about the provided phone number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **phoneNumber** | **string**| The phone number to be looked up | 
 **optional** | ***NumberLookupApiNumberLookupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberLookupApiNumberLookupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Specifies the type of number lookup to be performed | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

