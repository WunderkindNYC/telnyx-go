# Fax

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | The connection ID to send the fax with. | [optional] [default to null]
**CreatedAt** | **string** | ISO 8601 timestamp when resource was created | [optional] [default to null]
**Direction** | **string** | The direction of the fax. | [optional] [default to null]
**From** | **string** | The phone number, in E.164 format, the fax will be sent from. | [optional] [default to null]
**Id** | **string** | Identifies the fax. | [optional] [default to null]
**MediaUrl** | **string** | The URL to the PDF used for the fax&#x27;s media. | [optional] [default to null]
**Quality** | **string** | The quality of the fax. Can be normal, high, very_high | [optional] [default to high]
**RecordType** | **string** | Identifies the type of the resource. | [optional] [default to null]
**Status** | **string** | Status of the fax | [optional] [default to null]
**StoreMedia** | **bool** | Should fax media be stored on temporary URL | [optional] [default to null]
**StoredMediaUrl** | **bool** | If store_media was set to true, this is a link to temporary location. Link expires after 7 days. | [optional] [default to null]
**To** | **string** | The phone number, in E.164 format, the fax will be sent to or SIP URI | [optional] [default to null]
**UpdatedAt** | **string** | ISO 8601 timestamp when resource was updated | [optional] [default to null]
**WebhookFailoverUrl** | **string** | Optional failover URL that will receive fax webhooks if &#x60;webhook_url&#x60; doesn&#x27;t return a 2XX response | [optional] [default to null]
**WebhookUrl** | **string** | URL that will receive fax webhooks | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

