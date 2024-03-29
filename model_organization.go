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

// Contact organization information
type Organization struct {
	// Name of the contact's company
	Company string `json:"company"`
	// Name of the contact's department
	Department string `json:"department,omitempty"`
	// Contact's business title
	Title string `json:"title,omitempty"`
}
