# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNumberOrderRegulatoryRequirements**](NumberOrderRegulatoryRequirementsApi.md#ListNumberOrderRegulatoryRequirements) | **Get** /regulatory_requirements | Get list of Number Order Regulatory Requirements
[**ListPhoneNumberRegulatoryRequirements**](NumberOrderRegulatoryRequirementsApi.md#ListPhoneNumberRegulatoryRequirements) | **Get** /phone_number_regulatory_requirements | Get Regulatory Requirements Per Number
[**RetrieveNumberOrderRegulatoryRequirement**](NumberOrderRegulatoryRequirementsApi.md#RetrieveNumberOrderRegulatoryRequirement) | **Get** /regulatory_requirements/{requirement_id} | Get Detailed Number Order Regulatory Requirement

# **ListNumberOrderRegulatoryRequirements**
> InlineResponse20052 ListNumberOrderRegulatoryRequirements(ctx, optional)
Get list of Number Order Regulatory Requirements

Gets a paginated list of Number Order Regulatory Requirements

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberOrderRegulatoryRequirementsApiListNumberOrderRegulatoryRequirementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberOrderRegulatoryRequirementsApiListNumberOrderRegulatoryRequirementsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterRequirementId** | **optional.String**| Filter number order regulatory requirements by requirement_id | 
 **filterFieldType** | **optional.String**| Filter number order regulatory requirements by field_type | 
 **filterRequirementType** | **optional.String**| Filter number order regulatory requirements by requirement_type | 

### Return type

[**InlineResponse20052**](inline_response_200_52.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPhoneNumberRegulatoryRequirements**
> InlineResponse20039 ListPhoneNumberRegulatoryRequirements(ctx, optional)
Get Regulatory Requirements Per Number

Gets a paginated list of Phone Number Regulatory Requirements

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberOrderRegulatoryRequirementsApiListPhoneNumberRegulatoryRequirementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberOrderRegulatoryRequirementsApiListPhoneNumberRegulatoryRequirementsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterPhoneNumber** | [**optional.Interface of []string**](string.md)| The list of phone numbers to retrieve regulatory requirements for | 

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveNumberOrderRegulatoryRequirement**
> InlineResponse20053 RetrieveNumberOrderRegulatoryRequirement(ctx, requirementId)
Get Detailed Number Order Regulatory Requirement

Gets a single Number Order Regulatory Requirement

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requirementId** | **string**| The number order regulatory requirement id | 

### Return type

[**InlineResponse20053**](inline_response_200_53.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

