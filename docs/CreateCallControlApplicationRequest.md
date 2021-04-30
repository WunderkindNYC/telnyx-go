# CreateCallControlApplicationRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Specifies whether the connection can be used. | [optional] [default to true]
**AnchorsiteOverride** | **string** | &lt;code&gt;Latency&lt;/code&gt; directs Telnyx to route media through the site with the lowest round-trip time to the user&#x27;s connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media.  | [optional] [default to ANCHORSITE_OVERRIDE.LATENCY]
**ApplicationName** | **string** | A user-assigned name to help manage the application. | [default to null]
**DtmfType** | **string** | Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF digits sent to Telnyx will be accepted in all formats. | [optional] [default to DTMF_TYPE.RFC_2833]
**FirstCommandTimeout** | **bool** | Specifies whether calls to phone numbers associated with this connection should hangup after timing out. | [optional] [default to false]
**FirstCommandTimeoutSecs** | **int32** | Specifies how many seconds to wait before timing out a dial command. | [optional] [default to 30]
**Inbound** | [***CallControlApplicationInbound**](CallControlApplicationInbound.md) |  | [optional] [default to null]
**Outbound** | [***CallControlApplicationOutbound**](CallControlApplicationOutbound.md) |  | [optional] [default to null]
**WebhookApiVersion** | **string** | Determines which webhook format will be used, Telnyx API v1 or v2. | [optional] [default to WEBHOOK_API_VERSION.1_]
**WebhookEventFailoverUrl** | **string** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as &#x27;https&#x27;. | [optional] 
**WebhookEventUrl** | **string** | The URL where webhooks related to this connection will be sent. Must include a scheme, such as &#x27;https&#x27;. | [default to null]
**WebhookTimeoutSecs** | **int32** | Specifies how many seconds to wait before timing out a webhook. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

