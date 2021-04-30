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

type SendDtmfRequest struct {
	// DTMF digits to send. Valid digits are 0-9, A-D, *, and #. Pauses can be added using w (0.5s) and W (1s).
	Digits string `json:"digits"`
	// Specifies for how many milliseconds each digit will be played in the audio stream. Ranges from 100 to 500ms
	DurationMillis int32 `json:"duration_millis,omitempty"`
	// Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.
	ClientState string `json:"client_state,omitempty"`
	// Use this field to avoid duplicate commands. Telnyx will ignore commands with the same `command_id`.
	CommandId string `json:"command_id,omitempty"`
}