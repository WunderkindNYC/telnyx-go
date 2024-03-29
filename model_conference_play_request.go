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

type ConferencePlayRequest struct {
	// The URL of the file to be played back in the conference. The URL can point to either a WAV or MP3 file.
	AudioUrl string `json:"audio_url"`
	// List of call control ids identifying participants the audio file should be played to. If not given, the audio file will be played to the entire conference.
	CallControlIds []string `json:"call_control_ids,omitempty"`
	Loop *Loopcount `json:"loop,omitempty"`
}
