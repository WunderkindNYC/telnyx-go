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

type TexmlApplication struct {
	Active bool `json:"active,omitempty"`
	AnchorsiteOverride *AnchorsiteOverride `json:"anchorsite_override,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	DtmfType *DtmfType `json:"dtmf_type,omitempty"`
	FirstCommandTimeout bool `json:"first_command_timeout,omitempty"`
	FirstCommandTimeoutSecs int32 `json:"first_command_timeout_secs,omitempty"`
	FriendlyName string `json:"friendly_name,omitempty"`
	Id string `json:"id,omitempty"`
	Inbound *CreateTexmlApplicationRequestInbound `json:"inbound,omitempty"`
	Outbound *CreateTexmlApplicationRequestOutbound `json:"outbound,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// URL for Telnyx to send requests to containing information about call progress events.
	StatusCallback string `json:"status_callback,omitempty"`
	// HTTP request method Telnyx should use when requesting the status_callback URL.
	StatusCallbackMethod string `json:"status_callback_method,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	// URL to which Telnyx will deliver your XML Translator webhooks if we get an error response from your voice_url.
	VoiceFallbackUrl string `json:"voice_fallback_url,omitempty"`
	// HTTP request method Telnyx will use to interact with your XML Translator webhooks. Either 'get' or 'post'.
	VoiceMethod string `json:"voice_method,omitempty"`
	// URL to which Telnyx will deliver your XML Translator webhooks.
	VoiceUrl string `json:"voice_url,omitempty"`
}
