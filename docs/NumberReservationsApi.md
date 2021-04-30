# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNumberReservation**](NumberReservationsApi.md#CreateNumberReservation) | **Post** /number_reservations | Create a number reservation
[**ExtendNumberReservationExpiryTime**](NumberReservationsApi.md#ExtendNumberReservationExpiryTime) | **Post** /number_reservations/{number_reservation_id}/actions/extend | Extend a number reservation
[**ListNumberReservations**](NumberReservationsApi.md#ListNumberReservations) | **Get** /number_reservations | List number reservations
[**RetrieveNumberReservation**](NumberReservationsApi.md#RetrieveNumberReservation) | **Get** /number_reservations/{number_reservation_id} | Retrieve a number reservation

# **CreateNumberReservation**
> NumberReservationResponse CreateNumberReservation(ctx, body)
Create a number reservation

Creates a Phone Number Reservation for multiple numbers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateNumberReservationRequest**](CreateNumberReservationRequest.md)|  | 

### Return type

[**NumberReservationResponse**](Number Reservation Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtendNumberReservationExpiryTime**
> NumberReservationResponse ExtendNumberReservationExpiryTime(ctx, numberReservationId)
Extend a number reservation

Extends reservation expiry time on all phone numbers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **numberReservationId** | **string**| The number reservation ID. | 

### Return type

[**NumberReservationResponse**](Number Reservation Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNumberReservations**
> ListNumberReservationsResponse ListNumberReservations(ctx, optional)
List number reservations

Gets a paginated list of phone number reservations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumberReservationsApiListNumberReservationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumberReservationsApiListNumberReservationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterStatus** | **optional.String**| Filter number reservations by status. | 
 **filterCreatedAtGt** | **optional.String**| Filter number reservations later than this value. | 
 **filterCreatedAtLt** | **optional.String**| Filter number reservations earlier than this value. | 
 **filterPhoneNumbersPhoneNumber** | **optional.String**| Filter number reservations having these phone numbers. | 
 **filterCustomerReference** | **optional.String**| Filter number reservations via the customer reference set. | 
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**ListNumberReservationsResponse**](List Number Reservations Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveNumberReservation**
> NumberReservationResponse RetrieveNumberReservation(ctx, numberReservationId)
Retrieve a number reservation

Gets a single phone number reservation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **numberReservationId** | **string**| The number reservation ID. | 

### Return type

[**NumberReservationResponse**](Number Reservation Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

