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

type ReservedPhoneNumber struct {
	// An ISO 8901 datetime string denoting when the individual number reservation was created.
	CreatedAt string `json:"created_at,omitempty"`
	// Errors the reservation could happen upon
	Errors string `json:"errors,omitempty"`
	// An ISO 8901 datetime string for when the individual number reservation is going to expire
	ExpiredAt string `json:"expired_at,omitempty"`
	Id string `json:"id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	RecordType string `json:"record_type,omitempty"`
	// The status of the phone number's reservation.
	Status string `json:"status,omitempty"`
	// An ISO 8901 datetime string for when the the individual number reservation was updated.
	UpdatedAt string `json:"updated_at,omitempty"`
}