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

// Optional configuration parameters to modify 'answering_machine_detection' performance.
type CallRequestAnsweringMachineDetectionConfig struct {
	// Silence duration threshold after a greeting message or voice for it be considered human.
	AfterGreetingSilenceMillis int32 `json:"after_greeting_silence_millis,omitempty"`
	// Maximum threshold for silence between words.
	BetweenWordsSilenceMillis int32 `json:"between_words_silence_millis,omitempty"`
	// Maximum threshold of a human greeting. If greeting longer than this value, considered machine.
	GreetingDurationMillis int32 `json:"greeting_duration_millis,omitempty"`
	// If machine already detected, maximum threshold for silence between words. If exceeded, the greeting is considered ended.
	GreetingSilenceDurationMillis int32 `json:"greeting_silence_duration_millis,omitempty"`
	// If machine already detected, maximum timeout threshold to determine the end of the machine greeting.
	GreetingTotalAnalysisTimeMillis int32 `json:"greeting_total_analysis_time_millis,omitempty"`
	// If initial silence duration is greater than this value, consider it a machine.
	InitialSilenceMillis int32 `json:"initial_silence_millis,omitempty"`
	// If number of detected words is greater than this value, consder it a machine.
	MaximumNumberOfWords int32 `json:"maximum_number_of_words,omitempty"`
	// If a single word lasts longer than this threshold, consider it a machine.
	MaximumWordLengthMillis int32 `json:"maximum_word_length_millis,omitempty"`
	// Minimum noise threshold for any analysis.
	SilenceThreshold int32 `json:"silence_threshold,omitempty"`
	// Maximum timeout threshold for overall detection.
	TotalAnalysisTimeMillis int32 `json:"total_analysis_time_millis,omitempty"`
}
