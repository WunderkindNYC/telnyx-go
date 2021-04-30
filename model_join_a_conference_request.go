/*
 * Telnyx API
 *
 * SIP trunking, SMS, MMS, Call Control and Telephony Data Services.
 *
 * API version: 2.0.0
 * Contact: support@telnyx.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package github.com/WunderkindNYC/telnyx-go

type JoinAConferenceRequest struct {
	// Unique identifier and token for controlling the call
	CallControlId string `json:"call_control_id"`
	// Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.
	ClientState string `json:"client_state,omitempty"`
	// Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same `command_id` as one that has already been executed.
	CommandId string `json:"command_id,omitempty"`
	// Whether the conference should end and all remaining participant be hung up after the participant leaves the conference. Defaults to \"false\".
	EndConferenceOnExit bool `json:"end_conference_on_exit,omitempty"`
	// Whether the participant should be put on hold immediately after joining the conference. Defaults to \"false\".
	Hold bool `json:"hold,omitempty"`
	// The URL of an audio file to be played to the participant when they are put on hold after joining the conference. This property takes effect only if \"hold\" is set to \"true\".
	HoldAudioUrl string `json:"hold_audio_url,omitempty"`
	// Whether the participant should be muted immediately after joining the conference. Defaults to \"false\".
	Mute bool `json:"mute,omitempty"`
	// Whether the conference should be started after the participant joins the conference. Defaults to \"false\".
	StartConferenceOnEnter bool `json:"start_conference_on_enter,omitempty"`
}
