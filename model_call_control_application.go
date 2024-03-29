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

type CallControlApplication struct {
	// Specifies whether the connection can be used.
	Active bool `json:"active,omitempty"`
	// `Latency` directs Telnyx to route media through the site with the lowest round-trip time to the user's connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media. 
	AnchorsiteOverride string `json:"anchorsite_override,omitempty"`
	// A user-assigned name to help manage the application.
	ApplicationName string `json:"application_name,omitempty"`
	// ISO 8601 formatted date of when the resource was created
	CreatedAt string `json:"created_at,omitempty"`
	// Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF digits sent to Telnyx will be accepted in all formats.
	DtmfType string `json:"dtmf_type,omitempty"`
	// Specifies whether calls to phone numbers associated with this connection should hangup after timing out.
	FirstCommandTimeout bool `json:"first_command_timeout,omitempty"`
	// Specifies how many seconds to wait before timing out a dial command.
	FirstCommandTimeoutSecs int32 `json:"first_command_timeout_secs,omitempty"`
	Id string `json:"id,omitempty"`
	Inbound *CallControlApplicationInbound `json:"inbound,omitempty"`
	Outbound *CallControlApplicationOutbound `json:"outbound,omitempty"`
	RecordType string `json:"record_type,omitempty"`
	// ISO 8601 formatted date of when the resource was last updated
	UpdatedAt string `json:"updated_at,omitempty"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	WebhookApiVersion string `json:"webhook_api_version,omitempty"`
	// The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as `https`.
	WebhookEventFailoverUrl string `json:"webhook_event_failover_url,omitempty"`
	// The URL where webhooks related to this connection will be sent. Must include a scheme, such as `https`.
	WebhookEventUrl string `json:"webhook_event_url,omitempty"`
	WebhookTimeoutSecs int32 `json:"webhook_timeout_secs,omitempty"`
}
