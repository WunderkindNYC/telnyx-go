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

type MessagingHostedNumberOrder struct {
	// Resource unique identifier.
	Id string `json:"id,omitempty"`
	// Automatically associate the number with this messaging profile ID when the order is complete.
	MessagingProfileId string `json:"messaging_profile_id,omitempty"`
	PhoneNumbers []HostedNumber `json:"phone_numbers,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	Status string `json:"status,omitempty"`
}
