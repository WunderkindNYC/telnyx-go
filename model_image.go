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

// The media object containing an image
type Image struct {
	// Describes the specified media.
	Caption string `json:"caption,omitempty"`
	// The media object ID returned when the media is successfully uploaded to the media endpoint.
	Id string `json:"id,omitempty"`
	// The protocol and URL of the media to be sent. Use only with HTTP/HTTPS URLs. Either id or link must be provided, not both.
	Link string `json:"link,omitempty"`
}
