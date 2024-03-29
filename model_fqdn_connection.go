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

type FqdnConnection struct {
	// Defaults to true
	Active bool `json:"active,omitempty"`
	AnchorsiteOverride *AnchorsiteOverride `json:"anchorsite_override,omitempty"`
	ConnectionName string `json:"connection_name"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at,omitempty"`
	// When enabled, Telnyx will generate comfort noise when you place the call on hold. If disabled, you will need to generate comfort noise or on hold music to avoid RTP timeout.
	DefaultOnHoldComfortNoiseEnabled bool `json:"default_on_hold_comfort_noise_enabled,omitempty"`
	DtmfType *DtmfType `json:"dtmf_type,omitempty"`
	// Encode the SIP contact header sent by Telnyx to avoid issues for NAT or ALG scenarios.
	EncodeContactHeaderEnabled bool `json:"encode_contact_header_enabled,omitempty"`
	EncryptedMedia *EncryptedMedia `json:"encrypted_media,omitempty"`
	// Identifies the resource.
	Id string `json:"id,omitempty"`
	Inbound *InboundFqdn `json:"inbound,omitempty"`
	// Enable on-net T38 if you prefer that the sender and receiver negotiate T38 directly when both are on the Telnyx network. If this is disabled, Telnyx will be able to use T38 on just one leg of the call according to each leg's settings.
	OnnetT38PassthroughEnabled bool `json:"onnet_t38_passthrough_enabled,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	RtcpSettings *ConnectionRtcpSettings `json:"rtcp_settings,omitempty"`
	TransportProtocol *FqdnConnectionTransportProtocol `json:"transport_protocol,omitempty"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at,omitempty"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	WebhookApiVersion string `json:"webhook_api_version,omitempty"`
	// The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverUrl string `json:"webhook_event_failover_url,omitempty"`
	// The URL where webhooks related to this connection will be sent. Must include a scheme, such as 'https'.
	WebhookEventUrl string `json:"webhook_event_url,omitempty"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs int32 `json:"webhook_timeout_secs,omitempty"`
}
