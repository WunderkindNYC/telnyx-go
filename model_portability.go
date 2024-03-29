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

type Portability struct {
	// Alternative SPID (Service Provider ID). Often used when a carrier is using a number from another carrier
	Altspid string `json:"altspid,omitempty"`
	// Alternative service provider name
	AltspidCarrierName string `json:"altspid_carrier_name,omitempty"`
	// Alternative service provider type
	AltspidCarrierType string `json:"altspid_carrier_type,omitempty"`
	// City name extracted from the locality in the Local Exchange Routing Guide (LERG) database
	City string `json:"city,omitempty"`
	// Type of number
	LineType string `json:"line_type,omitempty"`
	// Local Routing Number, if assigned to the requested phone number
	Lrn string `json:"lrn,omitempty"`
	// Operating Company Name (OCN) as per the Local Exchange Routing Guide (LERG) database
	Ocn string `json:"ocn,omitempty"`
	// ISO-formatted date when the requested phone number has been ported
	PortedDate string `json:"ported_date,omitempty"`
	// Indicates whether or not the requested phone number has been ported
	PortedStatus string `json:"ported_status,omitempty"`
	// SPID (Service Provider ID)
	Spid string `json:"spid,omitempty"`
	// Service provider name
	SpidCarrierName string `json:"spid_carrier_name,omitempty"`
	// Service provider type
	SpidCarrierType string `json:"spid_carrier_type,omitempty"`
	State string `json:"state,omitempty"`
}
