# CallControlApplication

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Specifies whether the connection can be used. | [optional] [default to true]
**AnchorsiteOverride** | **string** | &#x60;Latency&#x60; directs Telnyx to route media through the site with the lowest round-trip time to the user&#x27;s connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media.  | [optional] [default to ANCHORSITE_OVERRIDE.LATENCY]
**ApplicationName** | **string** | A user-assigned name to help manage the application. | [optional] [default to null]
**CreatedAt** | **string** | ISO 8601 formatted date of when the resource was created | [optional] [default to null]
**DtmfType** | **string** | Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF digits sent to Telnyx will be accepted in all formats. | [optional] [default to DTMF_TYPE.RFC_2833]
**FirstCommandTimeout** | **bool** | Specifies whether calls to phone numbers associated with this connection should hangup after timing out. | [optional] [default to false]
**FirstCommandTimeoutSecs** | **int32** | Specifies how many seconds to wait before timing out a dial command. | [optional] [default to 30]
**Id** | **string** |  | [optional] [default to null]
**Inbound** | [***CallControlApplicationInbound**](CallControlApplicationInbound.md) |  | [optional] [default to null]
**Outbound** | [***CallControlApplicationOutbound**](CallControlApplicationOutbound.md) |  | [optional] [default to null]
**RecordType** | **string** |  | [optional] [default to RECORD_TYPE.APPLICATION]
**UpdatedAt** | **string** | ISO 8601 formatted date of when the resource was last updated | [optional] [default to null]
**WebhookApiVersion** | **string** | Determines which webhook format will be used, Telnyx API v1 or v2. | [optional] [default to WEBHOOK_API_VERSION.1_]
**WebhookEventFailoverUrl** | **string** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as &#x60;https&#x60;. | [optional] 
**WebhookEventUrl** | **string** | The URL where webhooks related to this connection will be sent. Must include a scheme, such as &#x60;https&#x60;. | [optional] [default to null]
**WebhookTimeoutSecs** | **int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

