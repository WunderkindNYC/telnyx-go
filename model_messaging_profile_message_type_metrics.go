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

type MessagingProfileMessageTypeMetrics struct {
	// The number of outbound messages successfully delivered.
	Delivered float64 `json:"delivered,omitempty"`
	// The metric type.
	Label string `json:"label,omitempty"`
	// The ratio of outbound messages sent that resulted in errors.
	OutboundErrorRatio float64 `json:"outbound_error_ratio,omitempty"`
	// The number of inbound messages received.
	Received float64 `json:"received,omitempty"`
	// The number of outbound messages sent.
	Sent float64 `json:"sent,omitempty"`
}
