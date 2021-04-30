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

// The set of features available for a specific messaging use case (SMS or MMS). Features can vary depending on the characteristics the phone number, as well as its current product configuration. 
type MessagingFeatureSet struct {
	// Send messages to and receive messages from numbers in the same country.
	DomesticTwoWay bool `json:"domestic_two_way"`
	// Receive messages from numbers in other countries.
	InternationalInbound bool `json:"international_inbound"`
	// Send messages to numbers in other countries.
	InternationalOutbound bool `json:"international_outbound"`
}