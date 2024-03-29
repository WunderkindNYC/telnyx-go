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

// The CNAM listing settings for a phone number.
type PhoneNumbersidvoiceCnamListing struct {
	// Enables CNAM listings for this number. Requires cnam_listing_details to also be set.
	CnamListingEnabled bool `json:"cnam_listing_enabled,omitempty"`
	// The CNAM listing details for this number. Must be alphanumeric characters or spaces with a maximum length of 15. Requires cnam_listing_enabled to also be set to true.
	CnamListingDetails string `json:"cnam_listing_details,omitempty"`
}
