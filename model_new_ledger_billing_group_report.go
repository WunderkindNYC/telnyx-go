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

type NewLedgerBillingGroupReport struct {
	// Month of the ledger billing group report
	Month int32 `json:"month,omitempty"`
	// Year of the ledger billing group report
	Year int32 `json:"year,omitempty"`
}
