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

type StartRecordingRequest struct {
	// When `dual`, final audio file will be stereo recorded with the first leg on channel A, and the rest on channel B.
	Channels string `json:"channels"`
	// Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.
	ClientState string `json:"client_state,omitempty"`
	// Use this field to avoid duplicate commands. Telnyx will ignore commands with the same `command_id`.
	CommandId string `json:"command_id,omitempty"`
	// The audio file format used when storing the call recording. Can be either `mp3` or `wav`.
	Format string `json:"format"`
	// If enabled, a beep sound will be played at the start of a recording.
	PlayBeep bool `json:"play_beep,omitempty"`
}
