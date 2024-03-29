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

type ConferenceMuteRequest struct {
	// Array of unique identifiers and tokens for controlling the call. When empty all participants will be muted.
	CallControlIds []string `json:"call_control_ids,omitempty"`
}
