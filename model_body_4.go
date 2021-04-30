/*
 * Telnyx API
 *
 * SIP trunking, SMS, MMS, Call Control and Telephony Data Services.
 *
 * API version: 2.0.0
 * Contact: support@telnyx.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package github.com/WunderkindNYC/telnyx-go

type Body4 struct {
	// Defaults to true
	Active bool `json:"active,omitempty"`
	// `Latency` directs Telnyx to route media through the site with the lowest round-trip time to the user's connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media.
	AnchorsiteOverride string `json:"anchorsite_override,omitempty"`
	ConnectionName string `json:"connection_name,omitempty"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileId string `json:"outbound_voice_profile_id,omitempty"`
	// One of UDP, TLS, or TCP. Applies only to connections with IP authentication or FQDN authentication.
	TransportProtocol string `json:"transport_protocol,omitempty"`
	// When enabled, Telnyx will generate comfort noise when you place the call on hold. If disabled, you will need to generate comfort noise or on hold music to avoid RTP timeout.
	DefaultOnHoldComfortNoiseEnabled bool `json:"default_on_hold_comfort_noise_enabled,omitempty"`
	// Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF digits sent to Telnyx will be accepted in all formats.
	DtmfType string `json:"dtmf_type,omitempty"`
	// Encode the SIP contact header sent by Telnyx to avoid issues for NAT or ALG scenarios.
	EncodeContactHeaderEnabled bool `json:"encode_contact_header_enabled,omitempty"`
	// Enable use of SRTP or ZRTP for encryption. Valid values are those listed or null. Cannot be set to non-null if the transport_portocol is TLS.
	EncryptedMedia string `json:"encrypted_media,omitempty"`
	// Enable on-net T38 if you prefer the sender and receiver negotiating T38 directly if both are on the Telnyx network. If this is disabled, Telnyx will be able to use T38 on just one leg of the call depending on each leg's settings.
	OnnetT38PassthroughEnabled bool `json:"onnet_t38_passthrough_enabled,omitempty"`
	RtcpSettings *FqdnConnectionsRtcpSettings `json:"rtcp_settings,omitempty"`
	Inbound *InboundConfiguration1 `json:"inbound,omitempty"`
}
