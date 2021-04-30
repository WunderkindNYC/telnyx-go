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

type Connection struct {
	// Defaults to true
	Active bool `json:"active,omitempty"`
	AnchorsiteOverride *AnchorsiteOverride `json:"anchorsite_override,omitempty"`
	ConnectionName string `json:"connection_name,omitempty"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at,omitempty"`
	// Identifies the specific resource.
	Id string `json:"id,omitempty"`
	OutboundVoiceProfileId string `json:"outbound_voice_profile_id,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at,omitempty"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	WebhookApiVersion string `json:"webhook_api_version,omitempty"`
	// The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails.
	WebhookEventFailoverUrl string `json:"webhook_event_failover_url,omitempty"`
	// The URL where webhooks related to this connection will be sent.
	WebhookEventUrl string `json:"webhook_event_url,omitempty"`
}
