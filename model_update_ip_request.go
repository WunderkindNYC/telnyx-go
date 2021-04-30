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

type UpdateIpRequest struct {
	// ID of the IP Connection to which this IP should be attached.
	ConnectionId string `json:"connection_id,omitempty"`
	// IP adddress represented by this resource.
	IpAddress string `json:"ip_address"`
	// Port to use when connecting to this IP.
	Port int32 `json:"port,omitempty"`
}
