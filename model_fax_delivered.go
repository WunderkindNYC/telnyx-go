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

type FaxDelivered struct {
	// The type of event being delivered.
	EventType string `json:"event_type,omitempty"`
	// Identifies the type of resource.
	Id string `json:"id,omitempty"`
	Payload *FaxDeliveredPayload `json:"payload,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
}
