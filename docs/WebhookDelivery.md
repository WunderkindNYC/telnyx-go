# WebhookDelivery

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | [**[]Attempt**](attempt.md) | Detailed delivery attempts, ordered by most recent. | [optional] [default to null]
**FinishedAt** | [**time.Time**](time.Time.md) | ISO 8601 timestamp indicating when the last webhook response has been received. | [optional] [default to null]
**Id** | **string** | Uniquely identifies the webhook_delivery record. | [optional] [default to null]
**RecordType** | **string** | Identifies the type of the resource. | [optional] [default to null]
**StartedAt** | [**time.Time**](time.Time.md) | ISO 8601 timestamp indicating when the first request attempt was initiated. | [optional] [default to null]
**Status** | **string** | Delivery status: &#x27;delivered&#x27; when successfuly delivered or &#x27;failed&#x27; if all attempts have failed. | [optional] [default to null]
**UserId** | **string** | Uniquely identifies the user that owns the webhook_delivery record. | [optional] [default to null]
**Webhook** | [***WebhookDeliveryWebhook**](webhook_delivery_webhook.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

