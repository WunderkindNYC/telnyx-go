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

type InlineResponse20041Data struct {
	// Identifies the resource.
	Id string `json:"id,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// The URL at which the CSV file can be retrieved.
	Url string `json:"url,omitempty"`
	// Indicates the completion level of the CSV report. Only complete CSV download requests will be able to be retrieved.
	Status string `json:"status,omitempty"`
}
