# UpdateCredentialConnectionRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Defaults to true | [optional] [default to null]
**AnchorsiteOverride** | [***AnchorsiteOverride**](AnchorsiteOverride.md) |  | [optional] [default to null]
**ConnectionName** | **string** |  | [optional] [default to null]
**DefaultOnHoldComfortNoiseEnabled** | **bool** | When enabled, Telnyx will generate comfort noise when you place the call on hold. If disabled, you will need to generate comfort noise or on hold music to avoid RTP timeout. | [optional] [default to false]
**DtmfType** | [***DtmfType**](DtmfType.md) |  | [optional] [default to null]
**EncodeContactHeaderEnabled** | **bool** | Encode the SIP contact header sent by Telnyx to avoid issues for NAT or ALG scenarios. | [optional] [default to false]
**EncryptedMedia** | [***EncryptedMedia**](EncryptedMedia.md) |  | [optional] [default to null]
**Inbound** | [***CredentialInbound**](CredentialInbound.md) |  | [optional] [default to null]
**OnnetT38PassthroughEnabled** | **bool** | Enable on-net T38 if you prefer the sender and receiver negotiating T38 directly if both are on the Telnyx network. If this is disabled, Telnyx will be able to use T38 on just one leg of the call depending on each leg&#x27;s settings. | [optional] [default to false]
**Outbound** | [***CredentialOutbound**](CredentialOutbound.md) |  | [optional] [default to null]
**Password** | **string** | The password to be used as part of the credentials. Must be 8 to 128 characters long. | [optional] [default to null]
**RtcpSettings** | [***ConnectionRtcpSettings**](ConnectionRtcpSettings.md) |  | [optional] [default to null]
**SipUriCallingPreference** | **string** | This feature enables inbound SIP URI calls to your Credential Auth Connection. If enabled for all (unrestricted) then anyone who calls the SIP URI &lt;your-username&gt;@telnyx.com will be connected to your Connection. You can also choose to allow only calls that are originated on any Connections under your account (internal). | [optional] [default to SIP_URI_CALLING_PREFERENCE.DISABLED]
**UserName** | **string** | The user name to be used as part of the credentials. Must be 4-32 characters long and alphanumeric values only (no spaces or special characters). | [optional] [default to null]
**WebhookApiVersion** | **string** | Determines which webhook format will be used, Telnyx API v1 or v2. | [optional] [default to WEBHOOK_API_VERSION.1_]
**WebhookEventFailoverUrl** | **string** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as &#x27;https&#x27;. | [optional] 
**WebhookEventUrl** | **string** | The URL where webhooks related to this connection will be sent. Must include a scheme, such as &#x27;https&#x27;. | [optional] [default to null]
**WebhookTimeoutSecs** | **int32** | Specifies how many seconds to wait before timing out a webhook. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

