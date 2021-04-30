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

type SimCardGroupCreate struct {
	// Upper limit on the amount of data the SIM cards, within the group, can use.
	DataLimit int32 `json:"data_limit,omitempty"`
	// A user friendly name for the SIM card group.
	Name string `json:"name"`
}