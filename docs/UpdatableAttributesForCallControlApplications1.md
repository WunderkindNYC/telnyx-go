# UpdatableAttributesForCallControlApplications1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionName** | **string** | A user-assigned name to help manage the connection. | [default to null]
**Active** | **bool** | Specifies whether the connection can be used. | [optional] [default to true]
**AnchorsiteOverride** | **string** | &#x60;Latency&#x60; directs Telnyx to route media through the site with the lowest round-trip time to the user&#x27;s connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media. | [optional] [default to ANCHORSITE_OVERRIDE.LATENCY]
**DtmfType** | **string** | Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF digits sent to Telnyx will be accepted in all formats. | [optional] [default to DTMF_TYPE.RFC_2833]
**WebhookEventUrl** | **string** | The URL where webhooks related to this connection will be sent. Must include a scheme, such as &#x27;https&#x27;. | 
**WebhookEventFailoverUrl** | **string** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as &#x27;https&#x27;. | [optional] 
**WebhookApiVersion** | **string** | Determines which webhook format will be used, Telnyx API v1 or v2. | [optional] [default to WEBHOOK_API_VERSION.1_]
**HangupOnTimeout** | **bool** | Specifies whether calls to phone numbers associated with this connection should hangup after when timing out. | [optional] [default to false]
**CallControlTimeout** | **int32** | Specifies how many seconds to wait before timing out a dial command. | [optional] [default to 30]
**OutboundVoiceProfileId** | **string** | Identifies the associated outbound voice profile. | [optional] [default to null]
**Inbound** | [***CallControlApplicationsInbound**](call_control_applications_inbound.md) |  | [optional] [default to null]
**Outbound** | [***CallControlApplicationsOutbound**](call_control_applications_outbound.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

