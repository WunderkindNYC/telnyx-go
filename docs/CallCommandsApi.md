# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallControlAnswer**](CallCommandsApi.md#CallControlAnswer) | **Post** /calls/{call_control_id}/actions/answer | Answer call
[**CallControlBridge**](CallCommandsApi.md#CallControlBridge) | **Post** /calls/{call_control_id}/actions/bridge | Bridge calls
[**CallControlDial**](CallCommandsApi.md#CallControlDial) | **Post** /calls | Dial
[**CallControlForkStart**](CallCommandsApi.md#CallControlForkStart) | **Post** /calls/{call_control_id}/actions/fork_start | Forking start
[**CallControlForkStop**](CallCommandsApi.md#CallControlForkStop) | **Post** /calls/{call_control_id}/actions/fork_stop | Forking stop
[**CallControlGatherUsingAudio**](CallCommandsApi.md#CallControlGatherUsingAudio) | **Post** /calls/{call_control_id}/actions/gather_using_audio | Gather using audio
[**CallControlGatherUsingSpeak**](CallCommandsApi.md#CallControlGatherUsingSpeak) | **Post** /calls/{call_control_id}/actions/gather_using_speak | Gather using speak
[**CallControlHangup**](CallCommandsApi.md#CallControlHangup) | **Post** /calls/{call_control_id}/actions/hangup | Hangup call
[**CallControlPlaybackStart**](CallCommandsApi.md#CallControlPlaybackStart) | **Post** /calls/{call_control_id}/actions/playback_start | Play audio URL
[**CallControlPlaybackStop**](CallCommandsApi.md#CallControlPlaybackStop) | **Post** /calls/{call_control_id}/actions/playback_stop | Stop audio playback
[**CallControlRecordStart**](CallCommandsApi.md#CallControlRecordStart) | **Post** /calls/{call_control_id}/actions/record_start | Recording start
[**CallControlRecordStop**](CallCommandsApi.md#CallControlRecordStop) | **Post** /calls/{call_control_id}/actions/record_stop | Recording stop
[**CallControlReject**](CallCommandsApi.md#CallControlReject) | **Post** /calls/{call_control_id}/actions/reject | Reject call
[**CallControlSendDTMF**](CallCommandsApi.md#CallControlSendDTMF) | **Post** /calls/{call_control_id}/actions/send_dtmf | Send DTMF
[**CallControlSpeak**](CallCommandsApi.md#CallControlSpeak) | **Post** /calls/{call_control_id}/actions/speak | Speak text
[**CallControlTransfer**](CallCommandsApi.md#CallControlTransfer) | **Post** /calls/{call_control_id}/actions/transfer | Transfer call

# **CallControlAnswer**
> InlineResponse2009 CallControlAnswer(ctx, body, callControlId)
Answer call

Answer an incoming call. You must issue this command before executing subsequent commands on an incoming call.  **Expected Webhooks:**  - `call.answered` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AnswerRequest**](AnswerRequest.md)| Answer call request | 
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallControlBridge**
> InlineResponse2009 CallControlBridge(ctx, body, callControlId)
Bridge calls

Bridge two call control calls.  **Expected Webhooks:**  - `call.bridged` for Leg A - `call.bridged` for Leg B 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BridgeRequest**](BridgeRequest.md)| Bridge call request | 
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallControlDial**
> InlineResponse2008 CallControlDial(ctx, body)
Dial

Dial a number from a given connection. A successful response will include a `call_leg_id` which can be used to correlate the command with subsequent webhooks.  **Expected Webhooks:**  - `call.initiated` - `call.answered` or `call.hangup` - `call.machine.detection.ended` if `answering_machine_detection` was requested - `call.machine.greeting.ended` if `answering_machine_detection` was requested to detect the end of machine greeting 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DialRequest**](DialRequest.md)| Dial request | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallControlForkStart**
> InlineResponse2009 CallControlForkStart(ctx, body, callControlId)
Forking start

Call forking allows you to stream the media from a call to a specific target in realtime.  This stream can be used to enable realtime audio analysis to support a  variety of use cases, including fraud detection, or the creation of AI-generated audio responses.  Requests must specify either the `target` attribute or the `rx` and `tx` attributes.  **Expected Webhooks:**  - `call.fork.started` - `call.fork.stopped`  **Simple Telnyx RTP Encapsulation Protocol (STREP)**  *Note: This header/encapsulation is not used when the `rx` and `tx` parameters have been specified; it only applies when media is forked using the `target` attribute.*  If the destination for forked media is specified using the \"target\" attribute, the RTP will be encapsulated in an extra Telnyx protocol, which adds a 24 byte header to the RTP payload in each packet. The STREP header includes the Call Control `call_leg_id` for stream identification, along with bits that represent the direction (inbound or outbound) of the media. This 24-byte header sits between the UDP header and the RTP header.  The STREP header makes it possible to fork RTP for multiple calls (or two RTP streams for the same call) to the same IP:port, where the streams can be demultiplexed by your application using the information in the header. Of course, it's still possible to ignore this header completely, for example, if sending forked media for different calls to different ports or IP addresses. In this case, simply strip 24 bytes (or use the second byte to find the header length) from the received UDP payload to get the RTP (RTP header and payload).  ``` STREP Specification    0                   1                   2                   3   0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1  +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+  |1 1|Version|L|D|    HeaderLen  |  reserved (2 bytes)           |  +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+  |       reserved (4 bytes, for UDP ports or anything else)      |  +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+  |               The call_leg_id                                 |  |                   from Call Control                           |  |                       (128 bits / 16 bytes)                   |  |                           (this is binary data)               |  +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+   11    Static bits 11, always set to 11 to easily distinguish forked media    from RTP (10) and T.38 media (usually 00) and SIP (which begins    with a capital letter, so begins with bits 01). This is a magic number.   Version    Four bits to indicate the version number of the protocol, starting at 0001.   L    One bit to represent the leg of the call (A or B).    0 represents the A (first) leg of the call.    1 represents the B (second) leg of the call.   D    One bit to represent the direction of this RTP stream.    0 represents media received by Telnyx.    1 represents media transmitted by Telnyx.   HeaderLen (1 byte)    The length of the header in bytes.    Note that this value does not include the length of the payload. The total    size of the RTP can be calculated by subtracting the HeaderLen from the UDP    length (minus 8 for the UDP header).    In version 1, this value will always be 24.   Reserved (6 bytes)    Reserved for future use and to make sure that the header is a multiple of 32 bits   Call Leg ID    A 128-bit identifier for the call leg.    This is the call_leg_id from Call Control. ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StartForkingRequest**](StartForkingRequest.md)| Fork media request | 
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallControlForkStop**
> InlineResponse2009 CallControlForkStop(ctx, body, callControlId)
Forking stop

Stop forking a call.  **Expected Webhooks:**  - `call.fork.stopped` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StopForkingRequest**](StopForkingRequest.md)| Stop forking media request | 
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallControlGatherUsingAudio**
> InlineResponse2009 CallControlGatherUsingAudio(ctx, body, callControlId)
Gather using audio

Play an audio file on the call until the required DTMF signals are gathered to build interactive menus.  You can pass a list of valid digits along with an 'invalid_audio_url', which will be played back at the beginning of each prompt. Playback will be interrupted when a DTMF signal is received. The [Answer](/docs/api/v2/call-control/Call-Commands#CallControlAnswer) command must be issued before the `gather_using_audio` command.  **Expected Webhooks:**  - `call.playback.started` - `call.playback.ended` - `call.dtmf.received` (you may receive many of these webhooks) - `call.gather.ended` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GatherUsingAudioWithAllParametersRequest**](GatherUsingAudioWithAllParametersRequest.md)| Gather using audio request | 
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallControlGatherUsingSpeak**
> InlineResponse2009 CallControlGatherUsingSpeak(ctx, body, callControlId)
Gather using speak

Convert text to speech and play it on the call until the required DTMF signals are gathered to build interactive menus.  You can pass a list of valid digits along with an 'invalid_payload', which will be played back at the beginning of each prompt. Speech will be interrupted when a DTMF signal is received. The [Answer](/docs/api/v2/call-control/Call-Commands#CallControlAnswer) command must be issued before the `gather_using_speak` command.  **Expected Webhooks:**  - `call.dtmf.received` (you may receive many of these webhooks) - `call.gather.ended` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GatherUsingSpeakWithAllParametersRequest**](GatherUsingSpeakWithAllParametersRequest.md)| Gather using speak request | 
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallControlHangup**
> InlineResponse2009 CallControlHangup(ctx, body, callControlId)
Hangup call

Hang up the call.  **Expected Webhooks:**  - `call.hangup` - `call.recording.saved` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HangupRequest**](HangupRequest.md)| Hangup request | 
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallControlPlaybackStart**
> InlineResponse2009 CallControlPlaybackStart(ctx, body, callControlId)
Play audio URL

Play an audio file on the call. If multiple play audio commands are issued consecutively, the audio files will be placed in a queue awaiting playback.  *Notes:*  - When `overlay` is enabled, `loop` is limited to 1, and `target_legs` is limited to `self`. - A customer cannot Play Audio with `overlay=true` unless there is a Play Audio with `overlay=false` actively playing.  **Expected Webhooks:**  - `call.playback.started` - `call.playback.ended` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PlayAudioUrlRequest**](PlayAudioUrlRequest.md)| Play audio URL request | 
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallControlPlaybackStop**
> InlineResponse2009 CallControlPlaybackStop(ctx, body, callControlId)
Stop audio playback

Stop audio being played on the call.  **Expected Webhooks:**  - `call.playback.ended` or `call.speak.ended` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PlaybackStopRequest**](PlaybackStopRequest.md)| Stop audio playback request | 
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallControlRecordStart**
> InlineResponse2009 CallControlRecordStart(ctx, body, callControlId)
Recording start

Start recording the call. Recording will stop on call hang-up, or can be initiated via the Stop Recording command.  **Expected Webhooks:**  There are no webhooks associated with this Command. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StartRecordingRequest**](StartRecordingRequest.md)| Start recording audio request | 
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallControlRecordStop**
> InlineResponse2009 CallControlRecordStop(ctx, body, callControlId)
Recording stop

Stop recording the call.  **Expected Webhooks:**  - `call.recording.saved` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StropRecordingRequest**](StropRecordingRequest.md)| Answer call request | 
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallControlReject**
> InlineResponse2009 CallControlReject(ctx, body, callControlId)
Reject call

Reject an incoming call.  **Expected Webhooks:**  - `call.hangup` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RejectRequest**](RejectRequest.md)| Reject request | 
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallControlSendDTMF**
> InlineResponse2009 CallControlSendDTMF(ctx, body, callControlId)
Send DTMF

Sends DTMF tones from this leg. DTMF tones will be heard by the other end of the call.  **Expected Webhooks:**  There are no webhooks associated with this command. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendDtmfRequest**](SendDtmfRequest.md)| Send DTMF request | 
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallControlSpeak**
> InlineResponse2009 CallControlSpeak(ctx, body, callControlId)
Speak text

Convert text to speech and play it back on the call. If multiple speak text commands are issued consecutively, the audio files will be placed in a queue awaiting playback.  **Expected Webhooks:**  - `call.speak.started` - `call.speak.ended` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SpeakRequest**](SpeakRequest.md)| Speak request | 
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallControlTransfer**
> InlineResponse2009 CallControlTransfer(ctx, body, callControlId)
Transfer call

Transfer a call to a new destination. If the transfer is unsuccessful, a `call.hangup` webhook will be sent indicating that the transfer could not be completed. The original call will remain active and may be issued additional commands, potentially transfering the call to an alternate destination.  **Expected Webhooks:**  - `call.initiated` - `call.bridged` to Leg B - `call.answered` or `call.hangup` - `call.machine.detection.ended` if `answering_machine_detection` was requested - `call.machine.greeting.ended` if `answering_machine_detection` was requested to detect the end of machine greeting 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransferCallRequest**](TransferCallRequest.md)| Transfer call request | 
  **callControlId** | **string**| Unique identifier and token for controlling the call | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

