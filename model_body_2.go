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

type Body2 struct {
	// The list of +E.164 formatted phone numbers to check for portability
	PhoneNumbers []string `json:"phone_numbers,omitempty"`
}
