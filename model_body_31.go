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

type Body31 struct {
	// Controls whether a tech prefix is enabled for this phone number.
	TechPrefixEnabled bool `json:"tech_prefix_enabled,omitempty"`
	// This field allows you to rewrite the destination number of an inbound call before the call is routed to you. The value of this field may be any alphanumeric value, and the value will replace the number originally dialed.
	TranslatedNumber string `json:"translated_number,omitempty"`
	CallForwarding *PhoneNumbersidvoiceCallForwarding `json:"call_forwarding,omitempty"`
	CnamListing *PhoneNumbersidvoiceCnamListing `json:"cnam_listing,omitempty"`
	// Controls whether a number is billed per minute or uses your concurrent channels.
	UsagePaymentMethod string `json:"usage_payment_method,omitempty"`
	MediaFeatures *PhoneNumbersidvoiceMediaFeatures `json:"media_features,omitempty"`
	CallRecording *PhoneNumbersidvoiceCallRecording `json:"call_recording,omitempty"`
}