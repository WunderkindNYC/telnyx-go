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

type OutboundCallRecording struct {
	// When call_recording_type is 'by_caller_phone_number', only outbound calls using one of these numbers will be recorded. Numbers must be specified in E164 format.
	CallRecordingCallerPhoneNumbers []string `json:"call_recording_caller_phone_numbers,omitempty"`
	// When using 'dual' channels, the final audio file will be a stereo recording with the first leg on channel A, and the rest on channel B.
	CallRecordingChannels string `json:"call_recording_channels,omitempty"`
	// The audio file format for calls being recorded.
	CallRecordingFormat string `json:"call_recording_format,omitempty"`
	// Specifies which calls are recorded.
	CallRecordingType string `json:"call_recording_type,omitempty"`
}
