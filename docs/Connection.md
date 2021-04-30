# Connection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifies the specific resource. | [optional] [default to null]
**RecordType** | **string** | Identifies the type of the resource. | [optional] [default to null]
**Active** | **bool** | Defaults to true | [optional] [default to null]
**AnchorsiteOverride** | **string** | &#x60;Latency&#x60; directs Telnyx to route media through the site with the lowest round-trip time to the user&#x27;s connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media. | [optional] [default to null]
**ConnectionName** | **string** |  | [optional] [default to null]
**CreatedAt** | **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] [default to null]
**OutboundVoiceProfileId** | **string** | Identifies the associated outbound voice profile. | [optional] [default to null]
**UpdatedAt** | **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] [default to null]
**WebhookEventUrl** | **string** | The URL where webhooks related to this connection will be sent. | [optional] 
**WebhookEventFailoverUrl** | **string** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. | [optional] 
**WebhookApiVersion** | **string** | Determines which webhook format will be used, Telnyx API v1 or v2. | [optional] [default to WEBHOOK_API_VERSION.1_]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

