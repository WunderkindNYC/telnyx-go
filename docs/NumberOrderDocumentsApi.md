# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNumberOrderDocument**](NumberOrderDocumentsApi.md#CreateNumberOrderDocument) | **Post** /number_order_documents | Upload Number Order Document
[**ListNumberOrderDocuments**](NumberOrderDocumentsApi.md#ListNumberOrderDocuments) | **Get** /number_order_documents | Get Uploaded Number Order Documents
[**RetrieveNumberOrderDocument**](NumberOrderDocumentsApi.md#RetrieveNumberOrderDocument) | **Get** /number_order_documents/{number_order_document_id} | Retrieve a Single Number Order Document
[**UpdateNumberOrderDocument**](NumberOrderDocumentsApi.md#UpdateNumberOrderDocument) | **Patch** /number_order_documents/{number_order_document_id} | Update Number Order Document

# **CreateNumberOrderDocument**
> InlineResponse20032 CreateNumberOrderDocument(ctx, body)
Upload Number Order Document

Upload a Phone Number Order Document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body22**](Body22.md)|  | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNumberOrderDocuments**
> InlineResponse20031 ListNumberOrderDocuments(ctx, optional)
Get Uploaded Number Order Documents

Gets a paginated list of Number Order Documents

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberOrderDocumentsApiListNumberOrderDocumentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberOrderDocumentsApiListNumberOrderDocumentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterRequirementId** | **optional.String**| Filter number order documents by requirement_id | 
 **filterCreatedAtGt** | **optional.String**| Filter number order documents after this datetime | 
 **filterCreatedAtLt** | **optional.String**| Filter number order documents from before this datetime | 

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveNumberOrderDocument**
> InlineResponse20032 RetrieveNumberOrderDocument(ctx, numberOrderDocumentId)
Retrieve a Single Number Order Document

Gets a single Number Order Document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **numberOrderDocumentId** | **string**| The number order document id | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNumberOrderDocument**
> InlineResponse20032 UpdateNumberOrderDocument(ctx, body, numberOrderDocumentId)
Update Number Order Document

Updates a Number Order Document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body23**](Body23.md)|  | 
  **numberOrderDocumentId** | **string**| The number order document id | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

