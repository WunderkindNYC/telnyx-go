# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConference**](ConferenceCommandsApi.md#CreateConference) | **Post** /conferences | Create conference
[**GetConferences**](ConferenceCommandsApi.md#GetConferences) | **Get** /conferences | List conferences
[**HoldConference**](ConferenceCommandsApi.md#HoldConference) | **Post** /conferences/{id}/actions/hold | Hold conference participants
[**JoinConference**](ConferenceCommandsApi.md#JoinConference) | **Post** /conferences/{id}/actions/join | Join a conference
[**MuteConference**](ConferenceCommandsApi.md#MuteConference) | **Post** /conferences/{id}/actions/mute | Mute conference participants
[**UnholdConference**](ConferenceCommandsApi.md#UnholdConference) | **Post** /conferences/{id}/actions/unhold | Unhold conference participants
[**UnmuteConference**](ConferenceCommandsApi.md#UnmuteConference) | **Post** /conferences/{id}/actions/unmute | Unmute conference participants

# **CreateConference**
> InlineResponse20011 CreateConference(ctx, body)
Create conference

Create a conference from an existing call leg using a `call_control_id` and a conference name. Upon creating the conference, the call will be automatically bridged to the conference. Conferences will expire after all participants have left the conference or after 4 hours regardless of the number of active participants.  **Expected Webhooks:**  - `conference.created` - `conference.participant.joined` - `conference.participant.left` - `conference.ended` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateConferenceRequest**](CreateConferenceRequest.md)| Create a conference | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConferences**
> InlineResponse20010 GetConferences(ctx, optional)
List conferences

Lists conferences. Conferences are created on demand, and will expire after all participants have left the conference or after 4 hours regardless of the number of active participants. Conferences are listed in descending order by `expires_at`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConferenceCommandsApiGetConferencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConferenceCommandsApiGetConferencesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterName** | **optional.String**| If present, conferences will be filtered to those with a matching &#x60;name&#x60; attribute. Matching is case-sensitive | 
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HoldConference**
> InlineResponse20012 HoldConference(ctx, body, id)
Hold conference participants

Hold a list of participants in a conference call

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HoldConferenceParticipantsRequest**](HoldConferenceParticipantsRequest.md)|  | 
  **id** | **string**| Unique identifier of the conference | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JoinConference**
> InlineResponse20012 JoinConference(ctx, body, id)
Join a conference

Join an existing call leg to a conference. Issue the Join Conference command with the conference ID in the path and the `call_control_id` of the leg you wish to join to the conference as an attribute.  **Expected Webhooks:**  - `conference.participant.joined` - `conference.participant.left` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JoinAConferenceRequest**](JoinAConferenceRequest.md)| New Messaging Profile object | 
  **id** | **string**| Unique identifier of the conference | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MuteConference**
> InlineResponse20012 MuteConference(ctx, body, id)
Mute conference participants

Mute a list of participants in a conference call

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MuteConferenceParticipantsRequest**](MuteConferenceParticipantsRequest.md)|  | 
  **id** | **string**| Unique identifier of the conference | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnholdConference**
> InlineResponse20012 UnholdConference(ctx, body, id)
Unhold conference participants

Unhold a list of participants in a conference call

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UnholdConferenceParticipantsRequest**](UnholdConferenceParticipantsRequest.md)|  | 
  **id** | **string**| Unique identifier of the conference | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnmuteConference**
> InlineResponse20012 UnmuteConference(ctx, body, id)
Unmute conference participants

Unmute a list of participants in a conference call

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UnmuteConferenceParticipantsRequest**](UnmuteConferenceParticipantsRequest.md)|  | 
  **id** | **string**| Unique identifier of the conference | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

