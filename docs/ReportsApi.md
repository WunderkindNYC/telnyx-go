# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLedgerBillingGroupReport**](ReportsApi.md#CreateLedgerBillingGroupReport) | **Post** /ledger_billing_group_reports | Create a ledger billing group report
[**RetrieveLedgerBillingGroupReport**](ReportsApi.md#RetrieveLedgerBillingGroupReport) | **Get** /ledger_billing_group_reports/{id} | Retrieve a ledger billing group report

# **CreateLedgerBillingGroupReport**
> InlineResponse20020 CreateLedgerBillingGroupReport(ctx, body)
Create a ledger billing group report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body12**](Body12.md)| New ledger billing group report parameters | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveLedgerBillingGroupReport**
> InlineResponse20020 RetrieveLedgerBillingGroupReport(ctx, id)
Retrieve a ledger billing group report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the ledger billing group report | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

