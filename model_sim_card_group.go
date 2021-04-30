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

type SimCardGroup struct {
	ConsumedData *ConsumedData `json:"consumed_data,omitempty"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at,omitempty"`
	// Upper limit on the amount of data the SIM cards, within the group, can use.
	DataLimit int32 `json:"data_limit,omitempty"`
	// Indicates whether the SIM card group is the users default group.<br/>The default group is created for the user and can not be removed.
	Default_ bool `json:"default,omitempty"`
	// Identifies the resource.
	Id string `json:"id,omitempty"`
	// A user friendly name for the SIM card group.
	Name string `json:"name,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at,omitempty"`
}