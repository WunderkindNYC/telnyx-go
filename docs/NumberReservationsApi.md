# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNumberReservations**](NumberReservationsApi.md#CreateNumberReservations) | **Post** /number_reservations | Create a Phone Number Reservation
[**ExtendNumberReservationExpiryTime**](NumberReservationsApi.md#ExtendNumberReservationExpiryTime) | **Post** /number_reservations/{number_reservation_id}/actions/extend | Extend a Phone Number Reservation
[**ListNumberReservations**](NumberReservationsApi.md#ListNumberReservations) | **Get** /number_reservations | Retrieve multiple Number Reservations
[**RetrieveNumberReservation**](NumberReservationsApi.md#RetrieveNumberReservation) | **Get** /number_reservations/{number_reservation_id} | Retrieve a Single Phone Number Reservation

# **CreateNumberReservations**
> InlineResponse20036 CreateNumberReservations(ctx, body)
Create a Phone Number Reservation

Creates a Phone Number Reservation for multiple numbers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body26**](Body26.md)|  | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtendNumberReservationExpiryTime**
> InlineResponse20036 ExtendNumberReservationExpiryTime(ctx, numberReservationId)
Extend a Phone Number Reservation

Extends reservation expiry time on all phone numbers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **numberReservationId** | **string**| The number reservation id | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNumberReservations**
> InlineResponse20035 ListNumberReservations(ctx, optional)
Retrieve multiple Number Reservations

Gets a paginated list of Phone Number Reservations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberReservationsApiListNumberReservationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberReservationsApiListNumberReservationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStatus** | **optional.String**| Filter number reservations by status | 
 **filterCreatedAtGt** | **optional.String**| Filter number reservations later than this value | 
 **filterCreatedAtLt** | **optional.String**| Filter number reservations earlier than this value | 
 **filterPhoneNumbersPhoneNumber** | **optional.String**| Filter number reservations having these phone numbers | 
 **filterCustomerReference** | **optional.String**| Filter number reservations via the customer reference set | 

### Return type

[**InlineResponse20035**](inline_response_200_35.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveNumberReservation**
> InlineResponse20036 RetrieveNumberReservation(ctx, numberReservationId)
Retrieve a Single Phone Number Reservation

Gets a single Phone Number Reservation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **numberReservationId** | **string**| The number reservation id | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

