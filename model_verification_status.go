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
// VerificationStatus : The possible statuses of the verification request.
type VerificationStatus string

// List of VerificationStatus
const (
	PENDING_VerificationStatus VerificationStatus = "pending"
	SMS_DELIVERY_FAILED_VerificationStatus VerificationStatus = "sms_delivery_failed"
	ACCEPTED_VerificationStatus VerificationStatus = "accepted"
	EXPIRED_VerificationStatus VerificationStatus = "expired"
	NOT_ENOUGH_CREDIT_VerificationStatus VerificationStatus = "not_enough_credit"
	NETWORK_ERROR_VerificationStatus VerificationStatus = "network_error"
	NUMBER_UNREACHABLE_VerificationStatus VerificationStatus = "number_unreachable"
	INTERNAL_ERROR_VerificationStatus VerificationStatus = "internal_error"
	INVALID_DESTINATION_VerificationStatus VerificationStatus = "invalid_destination"
	TIMED_OUT_VerificationStatus VerificationStatus = "timed_out"
)
