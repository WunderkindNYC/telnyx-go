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

// The media features settings for a phone number.
type PhoneNumbersidvoiceMediaFeatures struct {
	// When RTP Auto-Adjust is enabled, the destination RTP address port will be automatically changed to match the source of the incoming RTP packets.
	RtpAutoAdjustEnabled bool `json:"rtp_auto_adjust_enabled,omitempty"`
	// Controls how media is handled for the phone number. default: media routed through Telnyx with transcode support. proxy: media routed through Telnyx with no transcode support.
	MediaHandlingMode string `json:"media_handling_mode,omitempty"`
	// When enabled, Telnyx will accept RTP packets from any customer-side IP address and port, not just those to which Telnyx is sending RTP.
	AcceptAnyRtpPacketsEnabled bool `json:"accept_any_rtp_packets_enabled,omitempty"`
	// Controls whether Telnyx will accept a T.38 re-INVITE for this phone number. Note that Telnyx will not send a T.38 re-INVITE; this option only controls whether one will be accepted.
	T38FaxGatewayEnabled bool `json:"t38_fax_gateway_enabled,omitempty"`
}