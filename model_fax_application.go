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

type FaxApplication struct {
	Active bool `json:"active,omitempty"`
	AnchorsiteOverride *AnchorsiteOverride `json:"anchorsite_override,omitempty"`
	ApplicationName string `json:"application_name,omitempty"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at,omitempty"`
	Id string `json:"id,omitempty"`
	Inbound *CreateFaxApplicationRequestInbound `json:"inbound,omitempty"`
	Outbound *CreateFaxApplicationRequestOutbound `json:"outbound,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at,omitempty"`
	WebhookEventFailoverUrl string `json:"webhook_event_failover_url,omitempty"`
	WebhookEventUrl string `json:"webhook_event_url,omitempty"`
	WebhookTimeoutSecs int32 `json:"webhook_timeout_secs,omitempty"`
}
