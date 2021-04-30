# Body20

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A user friendly name for the messaging profile. | [default to null]
**Enabled** | **bool** | Specifies whether the messaging profile is enabled or not. | [optional] [default to true]
**WebhookUrl** | **string** | The URL where webhooks related to this messaging profile will be sent. | [optional] 
**WebhookFailoverUrl** | **string** | The failover URL where webhooks related to this messaging profile will be sent if sending to the primary URL fails. | [optional] 
**WebhookApiVersion** | **string** | Determines which webhook format will be used, Telnyx API v1, v2, or a legacy 2010-04-01 format. | [optional] [default to WEBHOOK_API_VERSION.2_]
**NumberPoolSettings** | [***MessagingProfilesNumberPoolSettings**](messaging_profiles_number_pool_settings.md) |  | [optional] [default to null]
**UrlShortenerSettings** | [***MessagingProfilesUrlShortenerSettings**](messaging_profiles_url_shortener_settings.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

