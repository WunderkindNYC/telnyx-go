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

// A JSON object representation of the operation. The information present here will relate directly to the source of the OTA request.
type CompleteOtaUpdateSettings struct {
	MobileOperatorNetworksPreferences *[]MobileOperatorNetworkPreferencesResponse `json:"mobile_operator_networks_preferences,omitempty"`
}
