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

type InlineResponse2007Data struct {
	RecordType string `json:"record_type"`
	// Uniquely identifies an individual call leg.
	CallLegId string `json:"call_leg_id"`
	// Uniquely identifies the call control session. A session may include multiple call leg events.
	CallSessionId string `json:"call_session_id"`
	// Event timestamp
	EventTimestamp string `json:"event_timestamp"`
	// Event name
	Name string `json:"name"`
	// Event type
	Type_ string `json:"type"`
	// Event metadata, which includes raw event, and extra information based on event type
	Metadata *interface{} `json:"metadata"`
}
