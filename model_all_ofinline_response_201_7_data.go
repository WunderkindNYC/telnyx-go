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

type AllOfinlineResponse2017Data struct {
	// ISO 8601 formatted date-time indicating the start time.
	StartTime string `json:"start_time,omitempty"`
	// ISO 8601 formatted date-time indicating the end time.
	EndTime string `json:"end_time,omitempty"`
	// Indicates the status of the report, which is updated asynchronously.
	Status string `json:"status,omitempty"`
	// The URL where the report content, when generated, will be published to.
	ReportUrl string `json:"report_url,omitempty"`
}
