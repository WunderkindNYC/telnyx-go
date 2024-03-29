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

// Webhook delivery attempt details.
type Attempt struct {
	// Webhook delivery errors.
	Errors []ModelError `json:"errors,omitempty"`
	// ISO 8601 timestamp indicating when the attempt has finished.
	FinishedAt time.Time `json:"finished_at,omitempty"`
	Http *Http `json:"http,omitempty"`
	// ISO 8601 timestamp indicating when the attempt was initiated.
	StartedAt time.Time `json:"started_at,omitempty"`
	Status string `json:"status,omitempty"`
}
