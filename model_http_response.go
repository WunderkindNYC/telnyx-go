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

// Response details, optional.
type HttpResponse struct {
	// Raw response body, limited to 10kB.
	Body string `json:"body,omitempty"`
	Headers *[]string `json:"headers,omitempty"`
	Status int32 `json:"status,omitempty"`
}
