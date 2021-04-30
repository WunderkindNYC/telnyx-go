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

type Address struct {
	// Uniquely identifies the address.
	Id string `json:"id,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// The first name associated with the address. An address must have either a first last name or a business name.
	FirstName string `json:"first_name,omitempty"`
	// The last name associated with the address. An address must have either a first last name or a business name.
	LastName string `json:"last_name,omitempty"`
	// The business name associated with the address. An address must have either a first last name or a business name.
	BusinessName string `json:"business_name,omitempty"`
	// The phone number associated with the address.
	PhoneNumber string `json:"phone_number,omitempty"`
	// The primary street address information about the address.
	StreetAddress string `json:"street_address,omitempty"`
	// Additional street address information about the address such as, but not limited to, unit number or apartment number.
	ExtendedAddress string `json:"extended_address,omitempty"`
	// The locality of the address. For US addresses, this corresponds to the city of the address.
	Locality string `json:"locality,omitempty"`
	// The locality of the address. For US addresses, this corresponds to the state of the address.
	AdministrativeArea string `json:"administrative_area,omitempty"`
	// The neighborhood of the address. This field is not used for addresses in the US but is used for some international addresses.
	Neighborhood string `json:"neighborhood,omitempty"`
	// The borough of the address. This field is not used for addresses in the US but is used for some international addresses.
	Borough string `json:"borough,omitempty"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,omitempty"`
	// The two-character (ISO 3166-1 alpha-2) country code of the address.
	CountryCode string `json:"country_code,omitempty"`
	// Indicates whether or not the address should be considered part of your list of addresses that appear for regular use.
	AddressBook bool `json:"address_book,omitempty"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at,omitempty"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at,omitempty"`
}
