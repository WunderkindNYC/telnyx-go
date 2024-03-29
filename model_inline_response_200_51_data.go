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

type InlineResponse20051Data struct {
	Id string `json:"id"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// Comment body
	Body string `json:"body"`
	// Identifies the associated port request
	PortoutId string `json:"portout_id,omitempty"`
	// Identifies the user who created the comment. Will be null if created by Telnyx Admin
	UserId string `json:"user_id"`
	// Comment creation timestamp in ISO 8601 format
	CreatedAt string `json:"created_at"`
}
