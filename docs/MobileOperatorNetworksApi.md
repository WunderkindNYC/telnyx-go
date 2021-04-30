# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MobileOperatorNetworksGet**](MobileOperatorNetworksApi.md#MobileOperatorNetworksGet) | **Get** /mobile_operator_networks | List mobile operator networks

# **MobileOperatorNetworksGet**
> InlineResponse20014 MobileOperatorNetworksGet(ctx, optional)
List mobile operator networks

Telnyx has a set of GSM mobile operators partners that are available through our mobile network roaming. This resource is entirely managed by Telnyx and may change over time. That means that this resource won't allow any write operations for it. Still, it's available so it can be used as a support resource that can be related to other resources or become a configuration option.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MobileOperatorNetworksApiMobileOperatorNetworksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MobileOperatorNetworksApiMobileOperatorNetworksGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterNameStartsWith** | **optional.String**| Filter by name starting with. | 
 **filterNameContains** | **optional.String**| Filter by name containing match. | 
 **filterNameEndsWith** | **optional.String**| Filter by name ending with. | 
 **filterCountryCode** | **optional.String**| Filter by exact country_code. | 
 **filterMcc** | **optional.String**| Filter by exact MCC. | 
 **filterMnc** | **optional.String**| Filter by exact MNC. | 
 **filterTadig** | **optional.String**| Filter by exact TADIG. | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

