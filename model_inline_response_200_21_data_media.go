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

type InlineResponse20021DataMedia struct {
	// The url of the media requested to be sent.
	Url string `json:"url,omitempty"`
	// The MIME type of the requested media.
	ContentType string `json:"content_type,omitempty"`
	// The SHA256 hash of the requested media.
	Sha256 string `json:"sha256,omitempty"`
	// The size of the requested media.
	Size int32 `json:"size,omitempty"`
}
