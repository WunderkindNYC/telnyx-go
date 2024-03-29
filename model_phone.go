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

type Phone struct {
	// WhatsApp ID
	Id string `json:"id,omitempty"`
	Phone string `json:"phone,omitempty"`
	// Standard Values: CELL, MAIN, IPHONE, HOME, WORK
	Type_ string `json:"type,omitempty"`
}
