# InboundConfiguration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AniNumberFormat** | **string** | This setting allows you to set the format with which the caller&#x27;s number (ANI) is sent for inbound phone calls. | [optional] [default to ANI_NUMBER_FORMAT.E164_NATIONAL_3]
**DnisNumberFormat** | **string** |  | [optional] [default to DNIS_NUMBER_FORMAT.E164_1]
**Codecs** | **[]string** | Defines the list of codecs that Telnyx will send for inbound calls to a specific number on your portal account, in priority order. This only works when the Connection the number is assigned to uses Media Handling mode: default. OPUS and H.264 codecs are available only when using TCP or TLS transport for SIP. | [optional] [default to ["G722","G711U","G711A","G729","OPUS","H.264"]]
**ChannelLimit** | **int32** | When set, this will limit the total number of inbound calls to phone numbers associated with this connection. | [optional] [default to null]
**GenerateRingbackTone** | **bool** | Generate ringback tone through 183 session progress message with early media. | [optional] [default to false]
**IsupHeadersEnabled** | **bool** | When set, inbound phone calls will receive ISUP parameters via SIP headers. (Only when available and only when using TCP or TLS transport.) | [optional] [default to false]
**PrackEnabled** | **bool** | Enable PRACK messages as defined in RFC3262. | [optional] [default to false]
**PrivacyZoneEnabled** | **bool** | By default, Telnyx does not send caller-id information when the caller has chosen to hide this information. When this option is enabled, Telnyx will send the SIP header Privacy:id plus the caller-id information so that the receiver side can choose when to hide it. | [optional] [default to false]
**SipCompactHeadersEnabled** | **bool** | Defaults to true. | [optional] [default to true]
**Timeout1xxSecs** | **string** | Time(sec) before aborting if connection is not made (min: 1, max: 20). | [optional] [default to 3]
**Timeout2xxSecs** | **string** | Time(sec) before aborting if call is unanswered (min: 1, max: 600). | [optional] [default to 90]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

