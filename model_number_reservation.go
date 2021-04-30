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

type NumberReservation struct {
	// An ISO 8901 datetime string denoting when the numbers reservation was created.
	CreatedAt string `json:"created_at,omitempty"`
	// A customer reference string for customer look ups.
	CustomerReference string `json:"customer_reference,omitempty"`
	Id string `json:"id,omitempty"`
	PhoneNumbers []ReservedPhoneNumber `json:"phone_numbers,omitempty"`
	RecordType string `json:"record_type,omitempty"`
	// The status of the entire reservation.
	Status string `json:"status,omitempty"`
	// An ISO 8901 datetime string for when the number reservation was updated.
	UpdatedAt string `json:"updated_at,omitempty"`
}
