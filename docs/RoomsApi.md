# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoom**](RoomsApi.md#CreateRoom) | **Post** /rooms | Create a room
[**DeleteRoom**](RoomsApi.md#DeleteRoom) | **Delete** /rooms/{room_id} | Delete a room
[**ListRooms**](RoomsApi.md#ListRooms) | **Get** /rooms | View a list of rooms.
[**ViewRoom**](RoomsApi.md#ViewRoom) | **Get** /rooms/{room_id} | View a room

# **CreateRoom**
> InlineResponse2015 CreateRoom(ctx, body)
Create a room

Synchronously create a Room.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRoomRequest**](CreateRoomRequest.md)| Parameters that can be defined during room creation. | 

### Return type

[**InlineResponse2015**](inline_response_201_5.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRoom**
> DeleteRoom(ctx, id)
Delete a room

Synchronously delete a Room. Participants from that room will be kicked out, they won't be able to join that room anymore, and you won't be charged anymore for that room.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The unique identifier of a room. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRooms**
> InlineResponse20035 ListRooms(ctx, optional)
View a list of rooms.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RoomsApiListRoomsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoomsApiListRoomsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterDateCreatedAtEq** | **optional.String**| ISO 8601 date for filtering rooms created on that date. | 
 **filterDateCreatedAtGte** | **optional.String**| ISO 8601 date for filtering rooms created after that date. | 
 **filterDateCreatedAtLte** | **optional.String**| ISO 8601 date for filtering rooms created before that date. | 
 **filterDateUpdatedAtEq** | **optional.String**| ISO 8601 date for filtering rooms updated on that date. | 
 **filterDateUpdatedAtGte** | **optional.String**| ISO 8601 date for filtering rooms updated after that date. | 
 **filterDateUpdatedAtLte** | **optional.String**| ISO 8601 date for filtering rooms updated before that date. | 
 **filterUniqueName** | **optional.String**| Unique_name for filtering rooms. | 
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]

### Return type

[**InlineResponse20035**](inline_response_200_35.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewRoom**
> InlineResponse2015 ViewRoom(ctx, roomId)
View a room

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roomId** | [**string**](.md)| The unique identifier of a room. | 

### Return type

[**InlineResponse2015**](inline_response_201_5.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

