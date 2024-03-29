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

type NotificationEventConditionParameters struct {
	DataType string `json:"data_type,omitempty"`
	Name string `json:"name,omitempty"`
	Optional bool `json:"optional,omitempty"`
}
