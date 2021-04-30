# CredentialConnectionCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Defaults to true | [optional] [default to null]
**UserName** | **string** | The user name to be used as part of the credentials. Must be 4-32 characters long and alphanumeric values only (no spaces or special characters). | [optional] [default to null]
**Password** | **string** | The password to be used as part of the credentials. Must be 8 to 128 characters long. | [optional] [default to null]
**AnchorsiteOverride** | **string** | &#x60;Latency&#x60; directs Telnyx to route media through the site with the lowest round-trip time to the user&#x27;s connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media. | [optional] [default to null]
**ConnectionName** | **string** |  | [optional] [default to null]
**SipUriCallingPreference** | **string** | This feature enables inbound SIP URI calls to your Credential Auth Connection. If enabled for all (unrestricted) then anyone who calls the SIP URI &lt;your-username&gt;@telnyx.com will be connected to your Connection. You can also choose to allow only calls that are originated on any Connections under your account (internal). | [optional] [default to SIP_URI_CALLING_PREFERENCE.DISABLED]
**DefaultOnHoldComfortNoiseEnabled** | **bool** | When enabled, Telnyx will generate comfort noise when you place the call on hold. If disabled, you will need to generate comfort noise or on hold music to avoid RTP timeout. | [optional] [default to false]
**DtmfType** | **string** | Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF digits sent to Telnyx will be accepted in all formats. | [optional] [default to DTMF_TYPE.RFC_2833]
**EncodeContactHeaderEnabled** | **bool** | Encode the SIP contact header sent by Telnyx to avoid issues for NAT or ALG scenarios. | [optional] [default to false]
**EncryptedMedia** | **string** | Enable use of SRTP or ZRTP for encryption. Valid values are those listed or null. Cannot be set to non-null if the transport_portocol is TLS. | [optional] [default to null]
**OnnetT38PassthroughEnabled** | **bool** | Enable on-net T38 if you prefer the sender and receiver negotiating T38 directly if both are on the Telnyx network. If this is disabled, Telnyx will be able to use T38 on just one leg of the call depending on each leg&#x27;s settings. | [optional] [default to false]
**OutboundVoiceProfileId** | **string** | Identifies the associated outbound voice profile. | [optional] [default to null]
**RtcpSettings** | [***CredentialConnectionsRtcpSettings**](credential_connections_rtcp_settings.md) |  | [optional] [default to null]
**Inbound** | [***InboundConfiguration**](Inbound configuration.md) |  | [optional] [default to null]
**Outbound** | [***OutboundConfiguration**](Outbound configuration.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

