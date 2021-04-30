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

type PortingOrderEndUserLocation struct {
	// State, province, or similar of billing address
	AdministrativeArea string `json:"administrative_area,omitempty"`
	// ISO3166-1 alpha-2 country code of billing address
	CountryCode string `json:"country_code,omitempty"`
	// Second line of billing address
	ExtendedAddress string `json:"extended_address,omitempty"`
	// City or municipality of billing address
	Locality string `json:"locality,omitempty"`
	// Postal Code of billing address
	PostalCode string `json:"postal_code,omitempty"`
	// First line of billing address
	StreetAddress string `json:"street_address,omitempty"`
}
