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

type CallSpeakStartedPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlId string `json:"call_control_id,omitempty"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegId string `json:"call_leg_id,omitempty"`
	// ID that is unique to the call session and can be used to correlate webhook events.
	CallSessionId string `json:"call_session_id,omitempty"`
	// State received from a command.
	ClientState string `json:"client_state,omitempty"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionId string `json:"connection_id,omitempty"`
}
