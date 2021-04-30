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

type InlineResponse20022PhoneNumbers struct {
	RecordType string `json:"record_type,omitempty"`
	// Identifies the type of resource.
	Id string `json:"id,omitempty"`
	// The messaging hosted phone number (+E.164 format)
	PhoneNumber string `json:"phone_number,omitempty"`
	Status string `json:"status,omitempty"`
}
