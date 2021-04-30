# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAvailablePhoneNumbers**](NumberSearchApi.md#ListAvailablePhoneNumbers) | **Get** /available_phone_numbers | List available phone numbers

# **ListAvailablePhoneNumbers**
> InlineResponse2003 ListAvailablePhoneNumbers(ctx, optional)
List available phone numbers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberSearchApiListAvailablePhoneNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberSearchApiListAvailablePhoneNumbersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterPhoneNumberStartsWith** | **optional.String**| Filter numbers starting with a pattern | 
 **filterPhoneNumberEndsWith** | **optional.String**| Filter numbers ending with a pattern | 
 **filterPhoneNumberContains** | **optional.String**| Filter numbers containing a pattern | 
 **filterLocality** | **optional.String**| Filter phone numbers by city | 
 **filterAdministrativeArea** | **optional.String**| Filter phone numbers by US state/CA province | 
 **filterCountryCode** | **optional.String**| Filter phone numbers by ISO 2 country code | 
 **filterNationalDestinationCode** | **optional.String**| Filter by the national destination code of the number. This filter is only applicable to North American numbers | 
 **filterRateCenter** | **optional.String**| Filter phone numbers by NANP rate center. This filter is only applicable to North American numbers | 
 **filterNumberType** | **optional.String**| Filter phone numbers by number type | 
 **filterFeatures** | [**optional.Interface of []string**](string.md)| Filter if the phone number should be used for voice, fax, mms, sms, emergency | 
 **filterLimit** | **optional.Int32**| Limits the number of results | 
 **filterBestEffort** | **optional.Bool**| Filter to determine if best effort results should be included. | 
 **filterQuickship** | **optional.Bool**| Filter to exclude phone numbers that need additional time after to purchase to receive phone calls. | 
 **filterReservable** | **optional.Bool**| Filter to exclude phone numbers that cannot be reserved before purchase. | 
 **filterExcludeRegulatoryRequirements** | **optional.Bool**| Filter to determine if the phone number has regulatory requirements | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

