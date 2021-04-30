# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConferenceDialParticipantIn**](ConferenceCommandsApi.md#ConferenceDialParticipantIn) | **Post** /conferences/{id}/actions/dial_participant | Dial a new participant into a conference
[**ConferenceHoldParticipants**](ConferenceCommandsApi.md#ConferenceHoldParticipants) | **Post** /conferences/{id}/actions/hold | Hold conference participants
[**ConferenceJoin**](ConferenceCommandsApi.md#ConferenceJoin) | **Post** /conferences/{id}/actions/join | Join a conference
[**ConferenceMuteParticipants**](ConferenceCommandsApi.md#ConferenceMuteParticipants) | **Post** /conferences/{id}/actions/mute | Mute conference participants
[**ConferencePlayAudio**](ConferenceCommandsApi.md#ConferencePlayAudio) | **Post** /conferences/{id}/actions/play | Play audio to conference participants
[**ConferenceSpeakText**](ConferenceCommandsApi.md#ConferenceSpeakText) | **Post** /conferences/{id}/actions/speak | Speak text to conference participants
[**ConferenceStartRecording**](ConferenceCommandsApi.md#ConferenceStartRecording) | **Post** /conferences/{id}/actions/record_start | Conference recording start
[**ConferenceStopAudio**](ConferenceCommandsApi.md#ConferenceStopAudio) | **Post** /conferences/{id}/actions/stop | Stop audio being played on the conference
[**ConferenceStopRecording**](ConferenceCommandsApi.md#ConferenceStopRecording) | **Post** /conferences/{id}/actions/record_stop | Conference recording stop
[**ConferenceUnholdParticipants**](ConferenceCommandsApi.md#ConferenceUnholdParticipants) | **Post** /conferences/{id}/actions/unhold | Unhold conference participants
[**ConferenceUnmuteParticipants**](ConferenceCommandsApi.md#ConferenceUnmuteParticipants) | **Post** /conferences/{id}/actions/unmute | Unmute conference participants
[**ConferenceUpdate**](ConferenceCommandsApi.md#ConferenceUpdate) | **Post** /conferences/{id}/actions/update | Update conference participant
[**CreateConference**](ConferenceCommandsApi.md#CreateConference) | **Post** /conferences | Create conference
[**ListConferenceParticipants**](ConferenceCommandsApi.md#ListConferenceParticipants) | **Get** /conferences/{conference_id}/participants | List conference participants
[**ListConferences**](ConferenceCommandsApi.md#ListConferences) | **Get** /conferences | List conferences
[**RetrieveConference**](ConferenceCommandsApi.md#RetrieveConference) | **Get** /conferences/{id} | Retrieve a conference

# **ConferenceDialParticipantIn**
> ConferenceCommandResponse ConferenceDialParticipantIn(ctx, body, id)
Dial a new participant into a conference

Dials a phone number and, when the call is answered, automatically joins them into the specified conference.  **Expected Webhooks:**  - `call.hangup` - `call.answered` - `conference.participant.joined` - `conference.participant.left` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConferenceCallRequest**](ConferenceCallRequest.md)| Dial Into Conference request object | 
  **id** | **string**| Uniquely identifies the conference by id or name | 

### Return type

[**ConferenceCommandResponse**](Conference Command Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConferenceHoldParticipants**
> ConferenceCommandResponse ConferenceHoldParticipants(ctx, body, id)
Hold conference participants

Hold a list of participants in a conference call

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConferenceHoldRequest**](ConferenceHoldRequest.md)|  | 
  **id** | **string**| Uniquely identifies the conference by id or name | 

### Return type

[**ConferenceCommandResponse**](Conference Command Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConferenceJoin**
> ConferenceCommandResponse ConferenceJoin(ctx, body, id)
Join a conference

Join an existing call leg to a conference. Issue the Join Conference command with the conference ID in the path and the `call_control_id` of the leg you wish to join to the conference as an attribute.  **Expected Webhooks:**  - `conference.participant.joined` - `conference.participant.left` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JoinConferenceRequest**](JoinConferenceRequest.md)| Join Conference request object | 
  **id** | **string**| Uniquely identifies the conference by id or name | 

### Return type

[**ConferenceCommandResponse**](Conference Command Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConferenceMuteParticipants**
> ConferenceCommandResponse ConferenceMuteParticipants(ctx, body, id)
Mute conference participants

Mute a list of participants in a conference call

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConferenceMuteRequest**](ConferenceMuteRequest.md)|  | 
  **id** | **string**| Uniquely identifies the conference by id or name | 

### Return type

[**ConferenceCommandResponse**](Conference Command Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConferencePlayAudio**
> ConferenceCommandResponse ConferencePlayAudio(ctx, body, id)
Play audio to conference participants

Play audio to all or some participants on a conference call.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConferencePlayRequest**](ConferencePlayRequest.md)|  | 
  **id** | **string**| Uniquely identifies the conference by id or name | 

### Return type

[**ConferenceCommandResponse**](Conference Command Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConferenceSpeakText**
> ConferenceCommandResponse ConferenceSpeakText(ctx, body, id)
Speak text to conference participants

Convert text to speech and play it to all or some participants.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConferenceSpeakRequest**](ConferenceSpeakRequest.md)|  | 
  **id** | **string**| Specifies the conference by id or name | 

### Return type

[**ConferenceCommandResponse**](Conference Command Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConferenceStartRecording**
> ConferenceCommandResponse ConferenceStartRecording(ctx, body, id)
Conference recording start

Start recording the conference. Recording will stop on conference end, or via the Stop Recording command.  **Expected Webhooks:**  - `conference.recording.saved`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StartRecordingRequest**](StartRecordingRequest.md)|  | 
  **id** | **string**| Specifies the conference to record by id or name | 

### Return type

[**ConferenceCommandResponse**](Conference Command Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConferenceStopAudio**
> ConferenceCommandResponse ConferenceStopAudio(ctx, body, id)
Stop audio being played on the conference

Stop audio being played to all or some participants on a conference call.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConferenceStopRequest**](ConferenceStopRequest.md)|  | 
  **id** | **string**| Uniquely identifies the conference by id or name | 

### Return type

[**ConferenceCommandResponse**](Conference Command Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConferenceStopRecording**
> ConferenceCommandResponse ConferenceStopRecording(ctx, body, id)
Conference recording stop

Stop recording the conference.  **Expected Webhooks:**  - `conference.recording.saved` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StopRecordingRequest**](StopRecordingRequest.md)| Stop recording conference request | 
  **id** | **string**| Specifies the conference to stop the recording for by id or name | 

### Return type

[**ConferenceCommandResponse**](Conference Command Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConferenceUnholdParticipants**
> ConferenceCommandResponse ConferenceUnholdParticipants(ctx, body, id)
Unhold conference participants

Unhold a list of participants in a conference call

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConferenceUnholdRequest**](ConferenceUnholdRequest.md)|  | 
  **id** | **string**| Uniquely identifies the conference by id or name | 

### Return type

[**ConferenceCommandResponse**](Conference Command Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConferenceUnmuteParticipants**
> ConferenceCommandResponse ConferenceUnmuteParticipants(ctx, body, id)
Unmute conference participants

Unmute a list of participants in a conference call

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConferenceUnmuteRequest**](ConferenceUnmuteRequest.md)|  | 
  **id** | **string**| Uniquely identifies the conference by id or name | 

### Return type

[**ConferenceCommandResponse**](Conference Command Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConferenceUpdate**
> ConferenceCommandResponse ConferenceUpdate(ctx, body, id)
Update conference participant

Update conference participant supervisor_role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateConferenceRequest**](UpdateConferenceRequest.md)| Update Conference request object | 
  **id** | **string**| Uniquely identifies the conference by id or name | 

### Return type

[**ConferenceCommandResponse**](Conference Command Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConference**
> ConferenceResponse CreateConference(ctx, body)
Create conference

Create a conference from an existing call leg using a `call_control_id` and a conference name. Upon creating the conference, the call will be automatically bridged to the conference. Conferences will expire after all participants have left the conference or after 4 hours regardless of the number of active participants.  **Expected Webhooks:**  - `conference.created` - `conference.participant.joined` - `conference.participant.left` - `conference.ended` - `conference.recording.saved` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateConferenceRequest**](CreateConferenceRequest.md)| Create a conference | 

### Return type

[**ConferenceResponse**](Conference Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConferenceParticipants**
> ListParticipantsResponse ListConferenceParticipants(ctx, conferenceId, optional)
List conference participants

Lists conference participants

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **conferenceId** | **string**| Uniquely identifies the conference by id | 
 **optional** | ***ConferenceCommandsApiListConferenceParticipantsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConferenceCommandsApiListConferenceParticipantsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterMuted** | **optional.Bool**| If present, participants will be filtered to those who are/are not muted | 
 **filterOnHold** | **optional.Bool**| If present, participants will be filtered to those who are/are not put on hold | 
 **filterWhispering** | **optional.Bool**| If present, participants will be filtered to those who are whispering or are not | 
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**ListParticipantsResponse**](List Participants Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConferences**
> ListConferencesResponse ListConferences(ctx, optional)
List conferences

Lists conferences. Conferences are created on demand, and will expire after all participants have left the conference or after 4 hours regardless of the number of active participants. Conferences are listed in descending order by `expires_at`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConferenceCommandsApiListConferencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConferenceCommandsApiListConferencesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterName** | **optional.String**| If present, conferences will be filtered to those with a matching &#x60;name&#x60; attribute. Matching is case-sensitive | 
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**ListConferencesResponse**](List Conferences Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveConference**
> ConferenceResponse RetrieveConference(ctx, id)
Retrieve a conference

Retrieve an existing conference

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Uniquely identifies the conference by id | 

### Return type

[**ConferenceResponse**](Conference Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

