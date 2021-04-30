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

type PortoutComment struct {
	// Comment body
	Body string `json:"body"`
	// Comment creation timestamp in ISO 8601 format
	CreatedAt string `json:"created_at"`
	Id string `json:"id"`
	// Identifies the associated port request
	PortoutId string `json:"portout_id,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// Identifies the user who created the comment. Will be null if created by Telnyx Admin
	UserId string `json:"user_id"`
}
