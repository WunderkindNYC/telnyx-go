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

type CustomSipHeader struct {
	// The name of the header to add.
	Name string `json:"name"`
	// The value of the header.
	Value string `json:"value"`
}