/*
 * Telnyx API
 *
 * SIP trunking, SMS, MMS, Call Control and Telephony Data Services.
 *
 * API version: 2.0.0
 * Contact: support@telnyx.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package telnyx

type CreateInboundIpRequest struct {
	// This setting allows you to set the format with which the caller's number (ANI) is sent for inbound phone calls.
	AniNumberFormat string `json:"ani_number_format,omitempty"`
	// When set, this will limit the total number of inbound calls to phone numbers associated with this connection.
	ChannelLimit int32 `json:"channel_limit,omitempty"`
	// Defines the list of codecs that Telnyx will send for inbound calls to a specific number on your portal account, in priority order. This only works when the Connection the number is assigned to uses Media Handling mode: default. OPUS and H.264 codecs are available only when using TCP or TLS transport for SIP.
	Codecs []string `json:"codecs,omitempty"`
	// Default routing method to be used when a number is associated with the connection. Must be one of the routing method types or left blank, other values are not allowed.
	DefaultRoutingMethod string `json:"default_routing_method,omitempty"`
	DnisNumberFormat string `json:"dnis_number_format,omitempty"`
	// Generate ringback tone through 183 session progress message with early media.
	GenerateRingbackTone bool `json:"generate_ringback_tone,omitempty"`
	// When set, inbound phone calls will receive ISUP parameters via SIP headers. (Only when available and only when using TCP or TLS transport.)
	IsupHeadersEnabled bool `json:"isup_headers_enabled,omitempty"`
	// Enable PRACK messages as defined in RFC3262.
	PrackEnabled bool `json:"prack_enabled,omitempty"`
	// By default, Telnyx does not send caller-id information when the caller has chosen to hide this information. When this option is enabled, Telnyx will send the SIP header Privacy:id plus the caller-id information so that the receiver side can choose when to hide it.
	PrivacyZoneEnabled bool `json:"privacy_zone_enabled,omitempty"`
	// Defaults to true.
	SipCompactHeadersEnabled bool `json:"sip_compact_headers_enabled,omitempty"`
	// Selects which `sip_region` to receive inbound calls from. If null, the default region (US) will be used.
	SipRegion string `json:"sip_region,omitempty"`
	// Specifies a subdomain that can be used to receive Inbound calls to a Connection, in the same way a phone number is used, from a SIP endpoint. Example: the subdomain \"example.sip.telnyx.com\" can be called from any SIP endpoint by using the SIP URI \"sip:@example.sip.telnyx.com\" where the user part can be any alphanumeric value. Please note TLS encrypted calls are not allowed for subdomain calls.
	SipSubdomain string `json:"sip_subdomain,omitempty"`
	// This option can be enabled to receive calls from: \"Anyone\" (any SIP endpoint in the public Internet) or \"Only my connections\" (any connection assigned to the same Telnyx user).
	SipSubdomainReceiveSettings string `json:"sip_subdomain_receive_settings,omitempty"`
	// Time(sec) before aborting if connection is not made.
	Timeout1xxSecs int32 `json:"timeout_1xx_secs,omitempty"`
	// Time(sec) before aborting if call is unanswered (min: 1, max: 600).
	Timeout2xxSecs int32 `json:"timeout_2xx_secs,omitempty"`
}
