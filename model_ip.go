/*
 * Telnyx API
 *
 * SIP trunking, SMS, MMS, Call Control and Telephony Data Services.
 *
 * API version: 2.0.0
 * Contact: support@telnyx.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package github.com/WunderkindNYC/telnyx-go

type Ip struct {
	// Identifies the type of resource.
	Id string `json:"id,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// ID of the IPConnection to which this IP should be attached.
	ConnectionId string `json:"connection_id,omitempty"`
	// IP adddress represented by this resource.
	IpAddress string `json:"ip_address,omitempty"`
	// Port to use when connecting to this IP.
	Port int32 `json:"port,omitempty"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at,omitempty"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at,omitempty"`
}
