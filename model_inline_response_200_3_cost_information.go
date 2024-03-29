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

type InlineResponse2003CostInformation struct {
	UpfrontCost string `json:"upfront_cost,omitempty"`
	MonthlyCost string `json:"monthly_cost,omitempty"`
	// The ISO 4217 code for the currency
	Currency string `json:"currency,omitempty"`
}
