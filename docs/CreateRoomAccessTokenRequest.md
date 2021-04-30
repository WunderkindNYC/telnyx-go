# CreateRoomAccessTokenRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefreshTokenTtlSecs** | **int32** | The time to live in seconds of the Refresh Token, after that time the Refresh Token is invalid and can&#x27;t be used to refresh Access Token. | [optional] [default to 3600]
**TtlSecs** | **int32** | The time to live in seconds of the Access Token, after that time the Access Token is invalid and can&#x27;t be used to join a Room. | [optional] [default to 600]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

