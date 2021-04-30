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

type NotificationSetting struct {
	AssociatedRecordType string `json:"associated_record_type,omitempty"`
	AssociatedRecordTypeValue string `json:"associated_record_type_value,omitempty"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// A UUID.
	Id string `json:"id,omitempty"`
	// A UUID reference to the associated Notification Channel.
	NotificationChannelId string `json:"notification_channel_id,omitempty"`
	// A UUID reference to the associated Notification Event Condition.
	NotificationEventConditionId string `json:"notification_event_condition_id,omitempty"`
	// A UUID reference to the associated Notification Profile.
	NotificationProfileId string `json:"notification_profile_id,omitempty"`
	Parameters []NotificationSettingParameters `json:"parameters,omitempty"`
	// Most preferences apply immediately; however, other may needs to propagate.
	Status string `json:"status,omitempty"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
