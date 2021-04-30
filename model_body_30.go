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

type Body30 struct {
	// Unique identifier for a messaging profile.
	MessagingProfileId string `json:"messaging_profile_id,omitempty"`
	// The requested messaging product the number should be on
	MessagingProduct string `json:"messaging_product,omitempty"`
}
