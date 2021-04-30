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

type PrivateWirelessGatewayRequest struct {
	// The private wireless gateway name.
	Name string `json:"name,omitempty"`
	// The identification of the related network resource.
	NetworkId string `json:"network_id,omitempty"`
}