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

type ValidateAddressRequest struct {
	AdministrativeArea string `json:"administrative_area"`
	CountryCode string `json:"country_code"`
	ExtendedAddress string `json:"extended_address,omitempty"`
	Locality string `json:"locality"`
	PostalCode string `json:"postal_code"`
	StreetAddress string `json:"street_address"`
}
