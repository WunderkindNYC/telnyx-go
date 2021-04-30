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

type MessagingProfileDetailedMetric struct {
	Metrics []MessagingProfileMessageTypeMetrics `json:"metrics,omitempty"`
	// The timestamp of the aggregated data.
	Timestamp string `json:"timestamp,omitempty"`
}
