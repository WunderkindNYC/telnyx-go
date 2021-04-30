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

type RegulatoryRequirement struct {
	Description string `json:"description,omitempty"`
	FieldType string `json:"field_type,omitempty"`
	Label string `json:"label,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	RequirementType string `json:"requirement_type,omitempty"`
}
