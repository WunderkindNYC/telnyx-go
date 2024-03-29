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

type NumberReservationsPhoneNumbers struct {
	Id string `json:"id,omitempty"`
	RecordType string `json:"record_type,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	// The status of the phone number's reservation
	Status string `json:"status,omitempty"`
	// An ISO 8901 datetime string denoting when the individual number reservation was created
	CreatedAt string `json:"created_at,omitempty"`
	// An ISO 8901 datetime string for when the the individual number reservation was updated
	UpdatedAt string `json:"updated_at,omitempty"`
	// An ISO 8901 datetime string for when the individual number reservation is going to expire
	ExpiredAt string `json:"expired_at,omitempty"`
}
