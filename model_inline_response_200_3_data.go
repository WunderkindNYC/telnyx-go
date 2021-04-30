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

type InlineResponse2003Data struct {
	RecordType string `json:"record_type,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	VanityFormat string `json:"vanity_format,omitempty"`
	// Specifies whether the phone number is an exact match based on the search criteria or not.
	BestEffort bool `json:"best_effort,omitempty"`
	// Specifies whether the phone number can receive calls immediately after purchase or not.
	Quickship bool `json:"quickship,omitempty"`
	// Specifies whether the phone number can be reserved before purchase or not.
	Reservable bool `json:"reservable,omitempty"`
	RegionInformation []InlineResponse2003RegionInformation `json:"region_information,omitempty"`
	CostInformation *InlineResponse2003CostInformation `json:"cost_information,omitempty"`
	RegulatoryRequirements []InlineResponse20039RegulatoryRequirements `json:"regulatory_requirements,omitempty"`
}
