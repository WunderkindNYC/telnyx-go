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

type ConferenceUnholdRequest struct {
	// List of unique identifiers and tokens for controlling the call. Enter each call control ID to be unheld.
	CallControlIds []string `json:"call_control_ids"`
}
