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
import (
	"time"
)

// A Notification Channel
type NotificationChannel struct {
	// The destination associated with the channel type.
	ChannelDestination string `json:"channel_destination,omitempty"`
	// A Channel Type ID
	ChannelTypeId string `json:"channel_type_id,omitempty"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// A UUID.
	Id string `json:"id,omitempty"`
	// A UUID reference to the associated Notification Profile.
	NotificationProfileId string `json:"notification_profile_id,omitempty"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}