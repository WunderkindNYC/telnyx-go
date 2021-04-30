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
// AnchorsiteOverride : `Latency` directs Telnyx to route media through the site with the lowest round-trip time to the user's connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media.
type AnchorsiteOverride string

// List of AnchorsiteOverride
const (
	LATENCY_AnchorsiteOverride AnchorsiteOverride = "Latency"
	CHICAGO_IL_AnchorsiteOverride AnchorsiteOverride = "Chicago, IL"
	ASHBURN_VA_AnchorsiteOverride AnchorsiteOverride = "Ashburn, VA"
	SAN_JOSE_CA_AnchorsiteOverride AnchorsiteOverride = "San Jose, CA"
	SYDNEY_AUSTRALIA_AnchorsiteOverride AnchorsiteOverride = "Sydney, Australia"
	AMSTERDAM_NETHERLANDS_AnchorsiteOverride AnchorsiteOverride = "Amsterdam, Netherlands"
	LONDON_UK_AnchorsiteOverride AnchorsiteOverride = "London, UK"
	TORONTO_CANADA_AnchorsiteOverride AnchorsiteOverride = "Toronto, Canada"
	VANCOUVER_CANADA_AnchorsiteOverride AnchorsiteOverride = "Vancouver, Canada"
	FRANKFURT_GERMANY_AnchorsiteOverride AnchorsiteOverride = "Frankfurt, Germany"
)
