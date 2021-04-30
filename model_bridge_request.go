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

type BridgeRequest struct {
	// The Call Control ID of the call you want to bridge with.
	CallControlId string `json:"call_control_id"`
	// Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.
	ClientState string `json:"client_state,omitempty"`
	// Use this field to avoid duplicate commands. Telnyx will ignore commands with the same `command_id`.
	CommandId string `json:"command_id,omitempty"`
	// Specifies behavior after the bridge ends (i.e. the opposite leg either hangs up or is transferred). If supplied with the value `self`, the current leg will be parked after unbridge. If not set, the default behavior is to hang up the leg.
	ParkAfterUnbridge string `json:"park_after_unbridge,omitempty"`
}
