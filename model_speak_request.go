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

type SpeakRequest struct {
	// The text or SSML to be converted into speech. There is a 5,000 character limit.
	Payload string `json:"payload"`
	// The type of the provided payload. The payload can either be plain text, or Speech Synthesis Markup Language (SSML).
	PayloadType string `json:"payload_type,omitempty"`
	// This parameter impacts speech quality, language options and payload types. When using `basic`, only the `en-US` language and payload type `text` are allowed.
	ServiceLevel string `json:"service_level,omitempty"`
	// When specified, it stops the current audio being played.  Specify `current` to stop the current audio being played, and to play the next file in the queue. Specify `all` to stop the current audio file being played and to also clear all audio files from the queue.
	Stop string `json:"stop,omitempty"`
	// The gender of the voice used to speak back the text.
	Voice string `json:"voice"`
	// The language you want spoken.
	Language string `json:"language"`
	// Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.
	ClientState string `json:"client_state,omitempty"`
	// Use this field to avoid duplicate commands. Telnyx will ignore commands with the same `command_id`.
	CommandId string `json:"command_id,omitempty"`
}
