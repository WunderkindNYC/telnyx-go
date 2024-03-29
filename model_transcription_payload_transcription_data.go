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

type TranscriptionPayloadTranscriptionData struct {
	// Speech recognition confidence level.
	Confidence float64 `json:"confidence,omitempty"`
	// Recognized text.
	Transcript string `json:"transcript,omitempty"`
}
