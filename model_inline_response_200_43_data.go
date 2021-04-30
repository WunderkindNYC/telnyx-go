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

type InlineResponse20043Data struct {
	// The number of channels set for the account
	Channels int32 `json:"channels,omitempty"`
	// Identifies the type of the response
	RecordType string `json:"record_type,omitempty"`
}
