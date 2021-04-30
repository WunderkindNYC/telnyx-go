# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNumberOrderDocument**](NumberOrderDocumentsApi.md#CreateNumberOrderDocument) | **Post** /number_order_documents | Create a number order document
[**ListNumberOrderDocuments**](NumberOrderDocumentsApi.md#ListNumberOrderDocuments) | **Get** /number_order_documents | List number order documents
[**RetrieveNumberOrderDocument**](NumberOrderDocumentsApi.md#RetrieveNumberOrderDocument) | **Get** /number_order_documents/{number_order_document_id} | Retrieve a number order document
[**UpdateNumberOrderDocument**](NumberOrderDocumentsApi.md#UpdateNumberOrderDocument) | **Patch** /number_order_documents/{number_order_document_id} | Update a number order document

# **CreateNumberOrderDocument**
> NumberOrderDocumentResponse CreateNumberOrderDocument(ctx, body)
Create a number order document

Upload a phone number order document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateNumberOrderDocumentRequest**](CreateNumberOrderDocumentRequest.md)|  | 

### Return type

[**NumberOrderDocumentResponse**](Number Order Document Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNumberOrderDocuments**
> ListNumberOrderDocumentsResponse ListNumberOrderDocuments(ctx, optional)
List number order documents

Gets a paginated list of number order documents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberOrderDocumentsApiListNumberOrderDocumentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberOrderDocumentsApiListNumberOrderDocumentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterRequirementId** | **optional.String**| Filter number order documents by &#x60;requirement_id&#x60;. | 
 **filterCreatedAtGt** | **optional.String**| Filter number order documents after this datetime. | 
 **filterCreatedAtLt** | **optional.String**| Filter number order documents from before this datetime. | 
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**ListNumberOrderDocumentsResponse**](List Number Order Documents Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveNumberOrderDocument**
> NumberOrderDocumentResponse RetrieveNumberOrderDocument(ctx, numberOrderDocumentId)
Retrieve a number order document

Gets a single number order document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **numberOrderDocumentId** | **string**| The number order document ID. | 

### Return type

[**NumberOrderDocumentResponse**](Number Order Document Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNumberOrderDocument**
> NumberOrderDocumentResponse UpdateNumberOrderDocument(ctx, body, numberOrderDocumentId)
Update a number order document

Updates a number order document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateNumberOrderDocumentRequest**](UpdateNumberOrderDocumentRequest.md)|  | 
  **numberOrderDocumentId** | **string**| The number order document ID. | 

### Return type

[**NumberOrderDocumentResponse**](Number Order Document Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

