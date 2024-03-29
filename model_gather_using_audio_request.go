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

type GatherUsingAudioRequest struct {
	// The URL of a file to be played back at the beginning of each prompt. The URL can point to either a WAV or MP3 file.
	AudioUrl string `json:"audio_url"`
	// Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.
	ClientState string `json:"client_state,omitempty"`
	// Use this field to avoid duplicate commands. Telnyx will ignore commands with the same `command_id`.
	CommandId string `json:"command_id,omitempty"`
	// The number of milliseconds to wait for input between digits.
	InterDigitTimeoutMillis int32 `json:"inter_digit_timeout_millis,omitempty"`
	// The URL of a file to play when digits don't match the `valid_digits` parameter or the number of digits is not between `min` and `max`. The URL can point to either a WAV or MP3 file.
	InvalidAudioUrl string `json:"invalid_audio_url,omitempty"`
	// The maximum number of digits to fetch. This parameter has a maximum value of 128.
	MaximumDigits int32 `json:"maximum_digits,omitempty"`
	// The maximum number of times the file should be played if there is no input from the user on the call.
	MaximumTries int32 `json:"maximum_tries,omitempty"`
	// The minimum number of digits to fetch. This parameter has a minimum value of 1.
	MinimumDigits int32 `json:"minimum_digits,omitempty"`
	// The digit used to terminate input if fewer than `maximum_digits` digits have been gathered.
	TerminatingDigit string `json:"terminating_digit,omitempty"`
	// The number of milliseconds to wait for a DTMF response after file playback ends before a replaying the sound file.
	TimeoutMillis int32 `json:"timeout_millis,omitempty"`
	// A list of all digits accepted as valid.
	ValidDigits string `json:"valid_digits,omitempty"`
}
