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

type InlineResponse2005Data struct {
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// Identifies the type of resource.
	Id string `json:"id,omitempty"`
	// Identifies the organization that owns the resource.
	OrganizationId string `json:"organization_id,omitempty"`
	// A user-specified name for the billing group
	Name string `json:"name,omitempty"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ISO 8601 formatted date indicating when the resource was removed.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
