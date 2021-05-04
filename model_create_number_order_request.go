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

type CreateNumberOrderRequest struct {
	// Identifies the billing group associated with the phone number.
	BillingGroupId string `json:"billing_group_id,omitempty"`
	// Identifies the connection associated with this phone number.
	ConnectionId string `json:"connection_id,omitempty"`
	// An ISO 8901 datetime string denoting when the number order was created.
	CreatedAt string `json:"created_at,omitempty"`
	// A customer reference string for customer look ups.
	CustomerReference string `json:"customer_reference,omitempty"`
	Id                string `json:"id,omitempty"`
	// Identifies the messaging profile associated with the phone number.
	MessagingProfileId string        `json:"messaging_profile_id,omitempty"`
	PhoneNumbers       []PhoneNumber `json:"phone_numbers,omitempty"`
	// The count of phone numbers in the number order.
	PhoneNumbersCount int32  `json:"phone_numbers_count,omitempty"`
	RecordType        string `json:"record_type,omitempty"`
	// True if all requirements are met for every phone number, false otherwise.
	RequirementsMet bool `json:"requirements_met,omitempty"`
	// The status of the order.
	Status string `json:"status,omitempty"`
	// An ISO 8901 datetime string for when the number order was updated.
	UpdatedAt string `json:"updated_at,omitempty"`
}

type PhoneNumber struct {
	Id string `json:"id"`
	PhoneNumber string `json:"phone_number"`
	RecordType string `json:"record_type"`
	RegulatoryGroupId string `json:"regulatory_group_id"`
	RegulatoryRequirements []RegulatoryRequirement `json:"regulatory_requirements"`
	RequirementsMet bool `json:"requirements_met"`
	Status string `json:"status"`
}