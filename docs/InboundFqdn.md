# InboundFqdn

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AniNumberFormat** | **string** | This setting allows you to set the format with which the caller&#x27;s number (ANI) is sent for inbound phone calls. | [optional] [default to ANI_NUMBER_FORMAT.E164_NATIONAL_3]
**ChannelLimit** | **int32** | When set, this will limit the total number of inbound calls to phone numbers associated with this connection. | [optional] [default to null]
**Codecs** | **[]string** | Defines the list of codecs that Telnyx will send for inbound calls to a specific number on your portal account, in priority order. This only works when the Connection the number is assigned to uses Media Handling mode: default. OPUS and H.264 codecs are available only when using TCP or TLS transport for SIP. | [optional] [default to ["G722","G711U","G711A","G729","OPUS","H.264"]]
**DefaultRoutingMethod** | **string** | Default routing method to be used when a number is associated with the connection. Must be one of the routing method types or null, other values are not allowed. | [optional] [default to null]
**DnisNumberFormat** | **string** |  | [optional] [default to DNIS_NUMBER_FORMAT.E164_1]
**GenerateRingbackTone** | **bool** | Generate ringback tone through 183 session progress message with early media. | [optional] [default to false]
**IsupHeadersEnabled** | **bool** | When set, inbound phone calls will receive ISUP parameters via SIP headers. (Only when available and only when using TCP or TLS transport.) | [optional] [default to false]
**PrackEnabled** | **bool** | Enable PRACK messages as defined in RFC3262. | [optional] [default to false]
**PrivacyZoneEnabled** | **bool** | By default, Telnyx does not send caller-id information when the caller has chosen to hide this information. When this option is enabled, Telnyx will send the SIP header Privacy:id plus the caller-id information so that the receiver side can choose when to hide it. | [optional] [default to false]
**SipCompactHeadersEnabled** | **bool** | Defaults to true. | [optional] [default to true]
**SipRegion** | **string** | Selects which &#x60;sip_region&#x60; to receive inbound calls from. If null, the default region (US) will be used. | [optional] [default to SIP_REGION.US]
**SipSubdomain** | **string** | Specifies a subdomain that can be used to receive Inbound calls to a Connection, in the same way a phone number is used, from a SIP endpoint. Example: the subdomain \&quot;example.sip.telnyx.com\&quot; can be called from any SIP endpoint by using the SIP URI \&quot;sip:@example.sip.telnyx.com\&quot; where the user part can be any alphanumeric value. Please note TLS encrypted calls are not allowed for subdomain calls. | [optional] [default to null]
**SipSubdomainReceiveSettings** | **string** | This option can be enabled to receive calls from: \&quot;Anyone\&quot; (any SIP endpoint in the public Internet) or \&quot;Only my connections\&quot; (any connection assigned to the same Telnyx user). | [optional] [default to SIP_SUBDOMAIN_RECEIVE_SETTINGS.FROM_ANYONE]
**Timeout1xxSecs** | **int32** | Time(sec) before aborting if connection is not made. | [optional] [default to 3]
**Timeout2xxSecs** | **int32** | Time(sec) before aborting if call is unanswered (min: 1, max: 600). | [optional] [default to 90]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

