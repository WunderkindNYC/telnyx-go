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

type Fax struct {
	// The connection ID to send the fax with.
	ConnectionId string `json:"connection_id,omitempty"`
	// ISO 8601 timestamp when resource was created
	CreatedAt string `json:"created_at,omitempty"`
	// The direction of the fax.
	Direction string `json:"direction,omitempty"`
	// The phone number, in E.164 format, the fax will be sent from.
	From string `json:"from,omitempty"`
	// Identifies the fax.
	Id string `json:"id,omitempty"`
	// The URL to the PDF used for the fax's media.
	MediaUrl string `json:"media_url,omitempty"`
	// The quality of the fax. Can be normal, high, very_high
	Quality string `json:"quality,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// Status of the fax
	Status string `json:"status,omitempty"`
	// Should fax media be stored on temporary URL
	StoreMedia bool `json:"store_media,omitempty"`
	// If store_media was set to true, this is a link to temporary location. Link expires after 7 days.
	StoredMediaUrl bool `json:"stored_media_url,omitempty"`
	// The phone number, in E.164 format, the fax will be sent to or SIP URI
	To string `json:"to,omitempty"`
	// ISO 8601 timestamp when resource was updated
	UpdatedAt string `json:"updated_at,omitempty"`
	// Optional failover URL that will receive fax webhooks if `webhook_url` doesn't return a 2XX response
	WebhookFailoverUrl string `json:"webhook_failover_url,omitempty"`
	// URL that will receive fax webhooks
	WebhookUrl string `json:"webhook_url,omitempty"`
}
