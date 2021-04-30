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

type ReplacedLinkClick struct {
	// The message ID associated with the clicked link.
	MessageId string `json:"message_id,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// ISO 8601 formatted date indicating when the message request was received.
	TimeClicked time.Time `json:"time_clicked,omitempty"`
	// Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short code).
	To string `json:"to,omitempty"`
	// The original link that was sent in the message.
	Url string `json:"url,omitempty"`
}