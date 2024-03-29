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

type Conference struct {
	// Identifies the connection (Call Control App) associated with the conference
	ConnectionId string `json:"connection_id,omitempty"`
	// ISO 8601 formatted date of when the conference was created
	CreatedAt string `json:"created_at"`
	// Reason why the conference ended
	EndReason string `json:"end_reason,omitempty"`
	EndedBy *ConferenceEndedBy `json:"ended_by,omitempty"`
	// ISO 8601 formatted date of when the conference will expire
	ExpiresAt string `json:"expires_at"`
	// Uniquely identifies the conference
	Id string `json:"id"`
	// Name of the conference
	Name string `json:"name"`
	RecordType string `json:"record_type"`
	// Region where the conference is hosted
	Region string `json:"region,omitempty"`
	// Status of the conference
	Status string `json:"status,omitempty"`
	// ISO 8601 formatted date of when the conference was last updated
	UpdatedAt string `json:"updated_at,omitempty"`
}
