# CallControlApplicationInbound

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelLimit** | **int32** | When set, this will limit the total number of inbound calls to phone numbers associated with this connection. | [optional] [default to null]
**SipSubdomain** | **string** | Specifies a subdomain that can be used to receive Inbound calls to a Connection, in the same way a phone number is used, from a SIP endpoint. Example: the subdomain \&quot;example.sip.telnyx.com\&quot; can be called from any SIP endpoint by using the SIP URI \&quot;sip:@example.sip.telnyx.com\&quot; where the user part can be any alphanumeric value. Please note TLS encrypted calls are not allowed for subdomain calls. | [optional] [default to null]
**SipSubdomainReceiveSettings** | **string** | This option can be enabled to receive calls from: \&quot;Anyone\&quot; (any SIP endpoint in the public Internet) or \&quot;Only my connections\&quot; (any connection assigned to the same Telnyx user). | [optional] [default to SIP_SUBDOMAIN_RECEIVE_SETTINGS.FROM_ANYONE]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

