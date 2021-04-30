# Participant

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | **string** | Call Control ID associated with the partiipant of the conference | [default to null]
**CallLegId** | **string** | Uniquely identifies the call leg associated with the participant | [default to null]
**Conference** | [***ParticipantConference**](Participant_conference.md) |  | [default to null]
**CreatedAt** | **string** | ISO 8601 formatted date of when the participant was created | [default to null]
**EndConferenceOnExit** | **bool** | Whether the conference will end and all remaining participants be hung up after the participant leaves the conference. | [default to null]
**Id** | **string** | Uniquely identifies the participant | [default to null]
**Muted** | **bool** | Whether the participant is muted. | [default to null]
**OnHold** | **bool** | Whether the participant is put on_hold. | [default to null]
**RecordType** | **string** |  | [default to null]
**SoftEndConferenceOnExit** | **bool** | Whether the conference will end after the participant leaves the conference. | [default to null]
**Status** | **string** | The status of the participant with respect to the lifecycle within the conference | [default to null]
**UpdatedAt** | **string** | ISO 8601 formatted date of when the participant was last updated | [default to null]
**WhisperCallControlIds** | **[]string** | Array of unique call_control_ids the participant can whisper to.. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

