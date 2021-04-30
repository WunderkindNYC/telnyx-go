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

type VerifyProfile struct {
	CreatedAt string `json:"created_at"`
	// For every request that is initiated via this Verify profile, this sets the number of seconds before a verification request code expires. Once the verification request expires, the user cannot use the code to verify their identity.
	DefaultTimeoutSecs int32 `json:"default_timeout_secs"`
	Id string `json:"id"`
	// Enables SMS text messaging for the Verify profile.
	MessagingEnabled bool `json:"messaging_enabled,omitempty"`
	// Optionally sets a messaging text template when sending the verification code. Uses `{code}` to template in the actual verification code.
	MessagingTemplate string `json:"messaging_template,omitempty"`
	// The name of the Verify profile.
	Name string `json:"name"`
	// Enables RCS messaging for the Verify profile.
	RcsEnabled bool `json:"rcs_enabled"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	UpdatedAt string `json:"updated_at"`
	// Enables VSMS for the Verify profile.
	VsmsEnabled bool `json:"vsms_enabled,omitempty"`
}
