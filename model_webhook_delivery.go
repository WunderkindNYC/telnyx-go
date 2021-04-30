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
import (
	"time"
)

// Record of all attempts to deliver a webhook.
type WebhookDelivery struct {
	// Detailed delivery attempts, ordered by most recent.
	Attempts []Attempt `json:"attempts,omitempty"`
	// ISO 8601 timestamp indicating when the last webhook response has been received.
	FinishedAt time.Time `json:"finished_at,omitempty"`
	// Uniquely identifies the webhook_delivery record.
	Id string `json:"id,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// ISO 8601 timestamp indicating when the first request attempt was initiated.
	StartedAt time.Time `json:"started_at,omitempty"`
	// Delivery status: 'delivered' when successfuly delivered or 'failed' if all attempts have failed.
	Status string `json:"status,omitempty"`
	// Uniquely identifies the user that owns the webhook_delivery record.
	UserId string `json:"user_id,omitempty"`
	Webhook *WebhookDeliveryWebhook `json:"webhook,omitempty"`
}
