# MessagingProfilesUrlShortenerSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | One of the domains provided by the Telnyx URL shortener service.  | [default to null]
**Prefix** | **string** | Optional prefix that can be used to identify your brand, and will appear in the Telnyx generated URLs after the domain name.  | [optional] [default to null]
**ReplaceBlacklistOnly** | **bool** | Use the link replacement tool only for links that are specifically blacklisted by Telnyx.  | [optional] [default to null]
**SendWebhooks** | **bool** | Receive webhooks for when your replaced links are clicked. Webhooks are sent to the webhooks on the messaging profile.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

