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

type SimCardNetworkPreference struct {
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at,omitempty"`
	MobileOperatorNetworksPreferences *[]MobileOperatorNetworkPreferencesResponse `json:"mobile_operator_networks_preferences,omitempty"`
	RecordType string `json:"record_type,omitempty"`
	SimCardId string `json:"sim_card_id,omitempty"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at,omitempty"`
}
