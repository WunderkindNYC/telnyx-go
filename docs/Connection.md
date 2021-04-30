# Connection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Defaults to true | [optional] [default to null]
**AnchorsiteOverride** | [***AnchorsiteOverride**](AnchorsiteOverride.md) |  | [optional] [default to null]
**ConnectionName** | **string** |  | [optional] [default to null]
**CreatedAt** | **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] [default to null]
**Id** | **string** | Identifies the specific resource. | [optional] [default to null]
**OutboundVoiceProfileId** | **string** |  | [optional] [default to null]
**RecordType** | **string** | Identifies the type of the resource. | [optional] [default to null]
**UpdatedAt** | **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] [default to null]
**WebhookApiVersion** | **string** | Determines which webhook format will be used, Telnyx API v1 or v2. | [optional] [default to WEBHOOK_API_VERSION.1_]
**WebhookEventFailoverUrl** | **string** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. | [optional] 
**WebhookEventUrl** | **string** | The URL where webhooks related to this connection will be sent. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

