# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoomAccessToken**](RoomsAccessTokensApi.md#CreateRoomAccessToken) | **Post** /rooms/{room_id}/actions/create_access_token | Create Access Token to join a room
[**RefreshRoomAccessToken**](RoomsAccessTokensApi.md#RefreshRoomAccessToken) | **Post** /rooms/{room_id}/actions/refresh_access_token | Refresh Access Token to join a room

# **CreateRoomAccessToken**
> InlineResponse2016 CreateRoomAccessToken(ctx, body, roomId)
Create Access Token to join a room

Synchronously create an Access Token to join a Room. Access Token is necessary to join a Telnyx Room. Access Token will expire after `ttl_secs`, a Refresh Token is also provided to refresh an Access Token, the Refresh Token expires after `refresh_token_ttl_secs`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRoomAccessTokenRequest**](CreateRoomAccessTokenRequest.md)| Parameters that can be defined during Room Access Token creation. | 
  **roomId** | [**string**](.md)| The unique identifier of a room. | 

### Return type

[**InlineResponse2016**](inline_response_201_6.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshRoomAccessToken**
> InlineResponse2017 RefreshRoomAccessToken(ctx, body, roomId)
Refresh Access Token to join a room

Synchronously refresh an Access Token to join a Room. Access Token is necessary to join a Telnyx Room. Access Token will expire after `ttl_secs`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RefreshRoomAccessTokenRequest**](RefreshRoomAccessTokenRequest.md)| Parameters that can be defined during Room Access Token refresh. | 
  **roomId** | [**string**](.md)| The unique identifier of a room. | 

### Return type

[**InlineResponse2017**](inline_response_201_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

