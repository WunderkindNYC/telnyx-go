# UpdateMessagingProfileRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) | ISO 8601 formatted date indicating when the resource was created. | [optional] [default to null]
**Enabled** | **bool** | Specifies whether the messaging profile is enabled or not. | [optional] [default to null]
**Id** | **string** | Identifies the type of resource. | [optional] [default to null]
**Name** | **string** | A user friendly name for the messaging profile. | [optional] [default to null]
**NumberPoolSettings** | [***NumberPoolSettings**](NumberPoolSettings.md) |  | [optional] [default to null]
**RecordType** | **string** | Identifies the type of the resource. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | ISO 8601 formatted date indicating when the resource was updated. | [optional] [default to null]
**UrlShortenerSettings** | [***UrlShortenerSettings**](UrlShortenerSettings.md) |  | [optional] [default to null]
**V1Secret** | **string** | Secret used to authenticate with v1 endpoints. | [optional] [default to null]
**WebhookApiVersion** | **string** | Determines which webhook format will be used, Telnyx API v1, v2, or a legacy 2010-04-01 format. | [optional] [default to null]
**WebhookFailoverUrl** | **string** | The failover URL where webhooks related to this messaging profile will be sent if sending to the primary URL fails. | [optional] [default to null]
**WebhookUrl** | **string** | The URL where webhooks related to this messaging profile will be sent. | [optional] [default to null]
**WhitelistedDestinations** | **[]string** | Destinations to which the messaging profile is allowed to send. If set to &#x60;null&#x60;, all destinations will be allowed. Setting a value of &#x60;[\&quot;*\&quot;]&#x60; has the equivalent effect. The elements in the list must be valid ISO 3166-1 alpha-2 country codes. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

