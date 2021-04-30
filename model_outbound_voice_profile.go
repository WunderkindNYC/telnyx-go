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

type OutboundVoiceProfile struct {
	// A user-supplied name to help with organization.
	Name string `json:"name,omitempty"`
	// Specifies the type of traffic allowed in this profile.
	TrafficType string `json:"traffic_type,omitempty"`
	ServicePlan string `json:"service_plan,omitempty"`
	// Must be no more than your global concurrent call limit. Null means no limit.
	ConcurrentCallLimit int32 `json:"concurrent_call_limit,omitempty"`
	// Specifies whether the outbound voice profile can be used. Disabled profiles will result in outbound calls being blocked for the associated Connections.
	Enabled bool `json:"enabled,omitempty"`
	Tags []string `json:"tags,omitempty"`
	UsagePaymentMethod string `json:"usage_payment_method,omitempty"`
	// The list of destinations you want to be able to call using this outbound voice profile formatted in alpha2.
	WhitelistedDestinations []string `json:"whitelisted_destinations,omitempty"`
	// Maximum rate (price per minute) for a Destination to be allowed when making outbound calls.
	MaxDestinationRate float64 `json:"max_destination_rate,omitempty"`
	// The maximum amount of usage charges, in USD, you want Telnyx to allow on this outbound voice profile in a day before disallowing new calls.
	DailySpendLimit string `json:"daily_spend_limit,omitempty"`
	// Specifies whether to enforce the daily_spend_limit on this outbound voice profile.
	DailySpendLimitEnabled bool `json:"daily_spend_limit_enabled,omitempty"`
	CallRecording *OutboundVoiceProfilesCallRecording `json:"call_recording,omitempty"`
	// The ID of the billing group associated with the outbound proflile. Defaults to null (for no group assigned).
	BillingGroupId string `json:"billing_group_id,omitempty"`
}
