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

type InlineResponse2008Data struct {
	RecordType string `json:"record_type"`
	// ID that is unique to the call session and can be used to correlate webhook events
	CallSessionId string `json:"call_session_id"`
	// ID that is unique to the call and can be used to correlate webhook events
	CallLegId string `json:"call_leg_id"`
	// Unique identifier and token for controlling the call. For Dial command it will always be `null` (dialing is asynchronous).
	CallControlId string `json:"call_control_id"`
	// Indicates whether the call is alive or not. For Dial command it will always be `false` (dialing is asynchronous).
	IsAlive bool `json:"is_alive"`
}
