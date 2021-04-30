# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostPortabilityCheck**](PhoneNumberPortingApi.md#PostPortabilityCheck) | **Post** /portability_checks | Run a portability check

# **PostPortabilityCheck**
> InlineResponse2012 PostPortabilityCheck(ctx, body)
Run a portability check

Runs a portability check, returning the results immediately.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body2**](Body2.md)|  | 

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

