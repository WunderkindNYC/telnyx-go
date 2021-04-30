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

// The request body when creating a verification.
type CreateVerificationRequest struct {
	// +E164 formatted phone number.
	PhoneNumber string `json:"phone_number"`
	// This is the number of seconds before the code of the request is expired. Once this request has expired, the code will no longer verify the user. Note: this will override the `default_timeout_secs` on the Verify profile.
	TimeoutSecs int32 `json:"timeout_secs,omitempty"`
	Type_ *VerificationType `json:"type"`
	// The identifier of the associated Verify profile.
	VerifyProfileId string `json:"verify_profile_id"`
}