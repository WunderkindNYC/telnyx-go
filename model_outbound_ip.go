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

type OutboundIp struct {
	// Set a phone number as the ani_override value to override caller id number on outbound calls.
	AniOverride string `json:"ani_override,omitempty"`
	// Specifies when we apply your ani_override setting. Only applies when ani_override is not blank.
	AniOverrideType string `json:"ani_override_type,omitempty"`
	// Forces all SIP calls originated on this connection to be \"parked\" instead of \"bridged\" to the destination specified on the URI. Parked calls will return ringback to the caller and will await for a Call Control command to define which action will be taken next.
	CallParkingEnabled bool `json:"call_parking_enabled,omitempty"`
	// When set, this will limit the total number of outbound calls to phone numbers associated with this connection.
	ChannelLimit int32 `json:"channel_limit,omitempty"`
	// Generate ringback tone through 183 session progress message with early media.
	GenerateRingbackTone bool `json:"generate_ringback_tone,omitempty"`
	// When set, ringback will not wait for indication before sending ringback tone to calling party.
	InstantRingbackEnabled bool `json:"instant_ringback_enabled,omitempty"`
	IpAuthenticationMethod string `json:"ip_authentication_method,omitempty"`
	IpAuthenticationToken string `json:"ip_authentication_token,omitempty"`
	// A 2-character country code specifying the country whose national dialing rules should be used. For example, if set to `US` then any US number can be dialed without preprending +1 to the number. When left blank, Telnyx will try US and GB dialing rules, in that order, by default.
	Localization string `json:"localization,omitempty"`
	OutboundVoiceProfileId string `json:"outbound_voice_profile_id,omitempty"`
	// This setting only affects connections with Fax-type Outbound Voice Profiles. The setting dictates whether or not Telnyx sends a t.38 reinvite.<br/><br/> By default, Telnyx will send the re-invite. If set to `customer`, the caller is expected to send the t.38 reinvite.
	T38ReinviteSource string `json:"t38_reinvite_source,omitempty"`
	// Numerical chars only, exactly 4 characters.
	TechPrefix string `json:"tech_prefix,omitempty"`
}