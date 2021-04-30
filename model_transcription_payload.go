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

type TranscriptionPayload struct {
	// Unique identifier and token for controlling the call.
	CallControlId string `json:"call_control_id,omitempty"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegId string `json:"call_leg_id,omitempty"`
	// ID that is unique to the call session and can be used to correlate webhook events.
	CallSessionId string `json:"call_session_id,omitempty"`
	// Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.
	ClientState string `json:"client_state,omitempty"`
	// Telnyx connection ID used in the call.
	ConnectionId string `json:"connection_id,omitempty"`
	TranscriptionData *TranscriptionPayloadTranscriptionData `json:"transcription_data,omitempty"`
}