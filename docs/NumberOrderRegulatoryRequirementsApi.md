# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNumberOrderRegulatoryRequirements**](NumberOrderRegulatoryRequirementsApi.md#ListNumberOrderRegulatoryRequirements) | **Get** /regulatory_requirements | List number order regulatory requirements
[**ListPhoneNumberRegulatoryRequirements**](NumberOrderRegulatoryRequirementsApi.md#ListPhoneNumberRegulatoryRequirements) | **Get** /phone_numbers_regulatory_requirements | List regulatory requirements per number
[**RetrieveNumberOrderRegulatoryRequirement**](NumberOrderRegulatoryRequirementsApi.md#RetrieveNumberOrderRegulatoryRequirement) | **Get** /regulatory_requirements/{requirement_id} | Retrieve a number order regulatory requirement

# **ListNumberOrderRegulatoryRequirements**
> ListNumberOrderRegulatoryRequirementsResponse ListNumberOrderRegulatoryRequirements(ctx, optional)
List number order regulatory requirements

Gets a paginated list of number order regulatory requirements.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberOrderRegulatoryRequirementsApiListNumberOrderRegulatoryRequirementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberOrderRegulatoryRequirementsApiListNumberOrderRegulatoryRequirementsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterRequirementId** | **optional.String**| Filter number order regulatory requirements by &#x60;requirement_id&#x60;. | 
 **filterFieldType** | **optional.String**| Filter number order regulatory requirements by &#x60;field_type&#x60;. | 
 **filterRequirementType** | **optional.String**| Filter number order regulatory requirements by &#x60;requirement_type&#x60;. | 
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**ListNumberOrderRegulatoryRequirementsResponse**](List Number Order Regulatory Requirements Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPhoneNumberRegulatoryRequirements**
> ListPhoneNumberRegulatoryRequirementsResponse ListPhoneNumberRegulatoryRequirements(ctx, optional)
List regulatory requirements per number

Gets a paginated list of phone number regulatory requirements.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberOrderRegulatoryRequirementsApiListPhoneNumberRegulatoryRequirementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberOrderRegulatoryRequirementsApiListPhoneNumberRegulatoryRequirementsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterPhoneNumber** | [**optional.Interface of []string**](string.md)| The list of phone numbers to retrieve regulatory requirements for. | 
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**ListPhoneNumberRegulatoryRequirementsResponse**](List Phone Number Regulatory Requirements Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveNumberOrderRegulatoryRequirement**
> NumberOrderRegulatoryRequirementResponse RetrieveNumberOrderRegulatoryRequirement(ctx, requirementId)
Retrieve a number order regulatory requirement

Gets a single number order regulatory requirement.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requirementId** | **string**| The number order regulatory requirement ID. | 

### Return type

[**NumberOrderRegulatoryRequirementResponse**](Number Order Regulatory Requirement Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

