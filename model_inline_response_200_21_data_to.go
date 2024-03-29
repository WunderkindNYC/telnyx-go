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

type InlineResponse20021DataTo struct {
	// Receiving address (+E.164 formatted phone number or short code).
	PhoneNumber string `json:"phone_number,omitempty"`
	Status string `json:"status,omitempty"`
	// The carrier of the receiver.
	Carrier string `json:"carrier,omitempty"`
	// The line-type of the receiver.
	LineType string `json:"line_type,omitempty"`
}
