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
import (
	"time"
)

type ConferencePlaybackEndedPayload struct {
	// ID of the conference the text was spoken in.
	ConferenceId string `json:"conference_id,omitempty"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionId string `json:"connection_id,omitempty"`
	// ID that is unique to the call session that started the conference.
	CreatorCallSessionId string `json:"creator_call_session_id,omitempty"`
	// The URL to the audio file being played.
	MediaUrl string `json:"media_url,omitempty"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time `json:"occurred_at,omitempty"`
}
