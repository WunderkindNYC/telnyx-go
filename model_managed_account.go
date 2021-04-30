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

type ManagedAccount struct {
	// The managed account's V2 API access key
	ApiKey string `json:"api_key"`
	// The managed account's V1 API token
	ApiToken string `json:"api_token"`
	// The manager account's email, which serves as the V1 API user identifier
	ApiUser string `json:"api_user"`
	Balance *ManagedAccountBalance `json:"balance,omitempty"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// The managed account's email.
	Email string `json:"email"`
	// Uniquely identifies the managed account.
	Id string `json:"id"`
	// The ID of the manager account associated with the managed account.
	ManagerAccountId string `json:"manager_account_id"`
	// The organization the managed account is associated with.
	OrganizationName string `json:"organization_name,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
}