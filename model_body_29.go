/*
 * Telnyx API
 *
 * SIP trunking, SMS, MMS, Call Control and Telephony Data Services.
 *
 * API version: 2.0.0
 * Contact: support@telnyx.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package github.com/WunderkindNYC/telnyx-go

type Body29 struct {
	// Indicates whether to enable emergency services on this number.
	EmergencyEnabled bool `json:"emergency_enabled"`
	// Identifies the address to be used with emergency services.
	EmergencyAddressId string `json:"emergency_address_id"`
}
