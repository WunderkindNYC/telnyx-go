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

type Body struct {
	MobileOperatorNetworksPreferences *[]MobileOperatorNetworkPreferencesRequest `json:"mobile_operator_networks_preferences,omitempty"`
	SimCardIds []string `json:"sim_card_ids,omitempty"`
}
