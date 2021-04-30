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

type PortingOrder struct {
	ActivationSettings *PortingOrderActivationSettings `json:"activation_settings,omitempty"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	Documents *PortingOrderDocuments `json:"documents,omitempty"`
	EndUser *PortingEndUser `json:"end_user,omitempty"`
	// Uniquely identifies this porting order
	Id string `json:"id,omitempty"`
	Misc *PortingOrderMisc `json:"misc,omitempty"`
	// Identifies the old service provider
	OldServiceProviderOcn string `json:"old_service_provider_ocn,omitempty"`
	PhoneNumberConfiguration *PortingOrderPhoneNumberConfiguration `json:"phone_number_configuration,omitempty"`
	// Count of phone numbers associated with this porting order
	PortingPhoneNumbersCount int32 `json:"porting_phone_numbers_count,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	Status *PortingOrderStatus `json:"status,omitempty"`
	// A key to reference this porting order when contacting Telnyx customer support
	SupportKey string `json:"support_key,omitempty"`
	// ISO 8601 formatted date indicating when the resource was created.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	UserFeedback *PortingOrderUserFeedback `json:"user_feedback,omitempty"`
	// A customer-specified reference number for customer bookkeeping purposes
	UserReference string `json:"user_reference,omitempty"`
	WebhookUrl string `json:"webhook_url,omitempty"`
}
