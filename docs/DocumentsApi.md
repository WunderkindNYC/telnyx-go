# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDocument**](DocumentsApi.md#CreateDocument) | **Post** /documents | Upload a document
[**DeleteDocument**](DocumentsApi.md#DeleteDocument) | **Delete** /documents/{id} | Delete a document
[**DownloadDocServiceDocument**](DocumentsApi.md#DownloadDocServiceDocument) | **Get** /documents/{id}/download | Download a document
[**ListDocumentLinks**](DocumentsApi.md#ListDocumentLinks) | **Get** /document_links | List all document links
[**ListDocuments**](DocumentsApi.md#ListDocuments) | **Get** /documents | List all documents
[**RetrieveDocument**](DocumentsApi.md#RetrieveDocument) | **Get** /documents/{id} | Retrieve a document
[**UpdateDocument**](DocumentsApi.md#UpdateDocument) | **Patch** /documents/{id} | Update a document

# **CreateDocument**
> InlineResponse2007 CreateDocument(ctx, body, file)
Upload a document

Upload a document.<br /><br />Uploaded files must be linked to a service within 30 minutes or they will be automatically deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateDocServiceDocumentRequest**](CreateDocServiceDocumentRequest.md)|  | 
  **file** | ***os.File*****os.File**|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDocument**
> InlineResponse2007 DeleteDocument(ctx, id)
Delete a document

Delete a document.<br /><br />A document can only be deleted if it's not linked to a service. If it is linked to a service, it must be unlinked prior to deleting.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the resource. | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadDocServiceDocument**
> *os.File DownloadDocServiceDocument(ctx, id)
Download a document

Download a document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the resource. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDocumentLinks**
> InlineResponse2005 ListDocumentLinks(ctx, optional)
List all document links

List all documents links ordered by created_at descending.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiListDocumentLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiListDocumentLinksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterDocumentId** | [**optional.Interface of string**](.md)| Identifies the associated document to filter on. | 
 **filterLinkedRecordType** | **optional.String**| The &#x60;linked_record_type&#x60; of the document to filter on. | 
 **filterLinkedResourceId** | [**optional.Interface of string**](.md)| The &#x60;linked_resource_id&#x60; of the document to filter on. | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDocuments**
> InlineResponse2006 ListDocuments(ctx, optional)
List all documents

List all documents ordered by created_at descending.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DocumentsApiListDocumentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DocumentsApiListDocumentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveDocument**
> InlineResponse2007 RetrieveDocument(ctx, id)
Retrieve a document

Retrieve a document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the resource. | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDocument**
> InlineResponse2007 UpdateDocument(ctx, body, id)
Update a document

Update a document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DocServiceDocument**](DocServiceDocument.md)|  | 
  **id** | [**string**](.md)| The id of the resource. | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

