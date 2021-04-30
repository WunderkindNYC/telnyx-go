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

type CreateNumberPoolMessageRequest struct {
	// Automatically detect if an SMS message is unusually long and exceeds a recommended limit of message parts.
	AutoDetect bool `json:"auto_detect,omitempty"`
	// A list of media URLs. The total media size must be less than 1 MB.  **Required for MMS**
	MediaUrls []string `json:"media_urls,omitempty"`
	// Unique identifier for a messaging profile.
	MessagingProfileId string `json:"messaging_profile_id"`
	// Subject of multimedia message
	Subject string `json:"subject,omitempty"`
	// Message body (i.e., content) as a non-empty string.  **Required for SMS**
	Text string `json:"text,omitempty"`
	To string `json:"to"`
	// The protocol for sending the message, either SMS or MMS.
	Type_ string `json:"type,omitempty"`
	// If the profile this number is associated with has webhooks, use them for delivery notifications. If webhooks are also specified on the message itself, they will be attempted first, then those on the profile.
	UseProfileWebhooks bool `json:"use_profile_webhooks,omitempty"`
	// The failover URL where webhooks related to this message will be sent if sending to the primary URL fails.
	WebhookFailoverUrl string `json:"webhook_failover_url,omitempty"`
	// The URL where webhooks related to this message will be sent.
	WebhookUrl string `json:"webhook_url,omitempty"`
}
