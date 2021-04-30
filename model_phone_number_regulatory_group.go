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

type PhoneNumberRegulatoryGroup struct {
	PhoneNumber string `json:"phone_number,omitempty"`
	RecordType string `json:"record_type,omitempty"`
	RegulatoryGroupId string `json:"regulatory_group_id,omitempty"`
	RegulatoryRequirements []RegulatoryRequirement `json:"regulatory_requirements,omitempty"`
}
