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

type TelephonyCredential struct {
	// ISO-8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at,omitempty"`
	// Defaults to false
	Expired bool `json:"expired,omitempty"`
	// ISO-8601 formatted date indicating when the resource will expire.
	ExpiresAt string `json:"expires_at,omitempty"`
	// Identifies the resource.
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// Identifies the resource this credential is associated with.
	ResourceId string `json:"resource_id,omitempty"`
	// The randomly generated SIP password for the credential.
	SipPassword string `json:"sip_password,omitempty"`
	// The randomly generated SIP username for the credential.
	SipUsername string `json:"sip_username,omitempty"`
	// ISO-8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at,omitempty"`
}
