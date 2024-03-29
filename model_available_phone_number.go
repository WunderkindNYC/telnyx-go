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

type AvailablePhoneNumber struct {
	// Specifies whether the phone number is an exact match based on the search criteria or not.
	BestEffort bool `json:"best_effort,omitempty"`
	CostInformation *CostInformation `json:"cost_information,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	// Specifies whether the phone number can receive calls immediately after purchase or not.
	Quickship bool `json:"quickship,omitempty"`
	RecordType string `json:"record_type,omitempty"`
	RegionInformation []RegionInformation `json:"region_information,omitempty"`
	RegulatoryRequirements []RegulatoryRequirement `json:"regulatory_requirements,omitempty"`
	// Specifies whether the phone number can be reserved before purchase or not.
	Reservable bool `json:"reservable,omitempty"`
	VanityFormat string `json:"vanity_format,omitempty"`
}
