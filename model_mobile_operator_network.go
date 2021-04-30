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

type MobileOperatorNetwork struct {
	// The mobile operator two-character (ISO 3166-1 alpha-2) origin country code.
	CountryCode string `json:"country_code,omitempty"`
	// Identifies the resource.
	Id string `json:"id,omitempty"`
	// MCC stands for Mobile Country Code. It's a three decimal digit that identifies a country.<br/><br/> This code is commonly seen joined with a Mobile Network Code (MNC) in a tuple that allows identifying a carrier known as PLMN (Public Land Mobile Network) code.
	Mcc string `json:"mcc,omitempty"`
	// MNC stands for Mobile Network Code. It's a two to three decimal digits that identify a network.<br/><br/>  This code is commonly seen joined with a Mobile Country Code (MCC) in a tuple that allows identifying a carrier known as PLMN (Public Land Mobile Network) code.
	Mnc string `json:"mnc,omitempty"`
	// The operator network name.
	Name string `json:"name,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// TADIG stands for Transferred Account Data Interchange Group. The TADIG code is a unique identifier for network operators in GSM mobile networks.
	Tadig string `json:"tadig,omitempty"`
}
