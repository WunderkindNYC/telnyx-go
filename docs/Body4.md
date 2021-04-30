# Body4

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Defaults to true | [optional] [default to null]
**AnchorsiteOverride** | **string** | &#x60;Latency&#x60; directs Telnyx to route media through the site with the lowest round-trip time to the user&#x27;s connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media. | [optional] [default to null]
**ConnectionName** | **string** |  | [optional] [default to null]
**OutboundVoiceProfileId** | **string** | Identifies the associated outbound voice profile. | [optional] [default to null]
**TransportProtocol** | **string** | One of UDP, TLS, or TCP. Applies only to connections with IP authentication or FQDN authentication. | [optional] [default to TRANSPORT_PROTOCOL.UDP]
**DefaultOnHoldComfortNoiseEnabled** | **bool** | When enabled, Telnyx will generate comfort noise when you place the call on hold. If disabled, you will need to generate comfort noise or on hold music to avoid RTP timeout. | [optional] [default to true]
**DtmfType** | **string** | Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF digits sent to Telnyx will be accepted in all formats. | [optional] [default to DTMF_TYPE.RFC_2833]
**EncodeContactHeaderEnabled** | **bool** | Encode the SIP contact header sent by Telnyx to avoid issues for NAT or ALG scenarios. | [optional] [default to false]
**EncryptedMedia** | **string** | Enable use of SRTP or ZRTP for encryption. Valid values are those listed or null. Cannot be set to non-null if the transport_portocol is TLS. | [optional] [default to null]
**OnnetT38PassthroughEnabled** | **bool** | Enable on-net T38 if you prefer the sender and receiver negotiating T38 directly if both are on the Telnyx network. If this is disabled, Telnyx will be able to use T38 on just one leg of the call depending on each leg&#x27;s settings. | [optional] [default to false]
**RtcpSettings** | [***FqdnConnectionsRtcpSettings**](fqdn_connections_rtcp_settings.md) |  | [optional] [default to null]
**Inbound** | [***InboundConfiguration1**](Inbound configuration_1.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

