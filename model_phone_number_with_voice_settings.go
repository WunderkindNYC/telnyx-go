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

type PhoneNumberWithVoiceSettings struct {
	CallForwarding *CallForwarding `json:"call_forwarding,omitempty"`
	CallRecording *CallRecording `json:"call_recording,omitempty"`
	CnamListing *CnamListing `json:"cnam_listing,omitempty"`
	// Identifies the connection associated with this phone number.
	ConnectionId string `json:"connection_id,omitempty"`
	// A customer reference string for customer look ups.
	CustomerReference string `json:"customer_reference,omitempty"`
	Emergency *EmergencySettings `json:"emergency,omitempty"`
	// Identifies the type of resource.
	Id string `json:"id,omitempty"`
	MediaFeatures *MediaFeatures `json:"media_features,omitempty"`
	// The phone number in +E164 format.
	PhoneNumber string `json:"phone_number,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// Controls whether a tech prefix is enabled for this phone number.
	TechPrefixEnabled bool `json:"tech_prefix_enabled,omitempty"`
	// This field allows you to rewrite the destination number of an inbound call before the call is routed to you. The value of this field may be any alphanumeric value, and the value will replace the number originally dialed.
	TranslatedNumber string `json:"translated_number,omitempty"`
	// Controls whether a number is billed per minute or uses your concurrent channels.
	UsagePaymentMethod string `json:"usage_payment_method,omitempty"`
}
