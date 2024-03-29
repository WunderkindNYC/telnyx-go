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

type HoldConferenceParticipantsRequest struct {
	// List of unique identifiers and tokens for controlling the call. When empty all participants will be placed on hold.
	CallControlIds []string `json:"call_control_ids,omitempty"`
	// The URL of a file to be played back at the beginning of each prompt. The URL can point to either a WAV or MP3 file.
	AudioUrl string `json:"audio_url,omitempty"`
}
