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

type CallerName struct {
	// The name of the requested phone number's owner as per the CNAM database
	CallerName string `json:"caller_name,omitempty"`
	// A caller-name lookup specific error code, expressed as a stringified 5-digit integer
	ErrorCode string `json:"error_code,omitempty"`
}