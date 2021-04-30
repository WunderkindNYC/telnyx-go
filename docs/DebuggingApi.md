# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallControlDebuggingEventList**](DebuggingApi.md#CallControlDebuggingEventList) | **Get** /call_events | List call events

# **CallControlDebuggingEventList**
> InlineResponse2007 CallControlDebuggingEventList(ctx, optional)
List call events

Filters call events by given filter parameters. Events are ordered by `event_timestamp`. If filter for `call_leg_id` or `call_session_id` is not present, it only filters events from the last 24 hours.  **Note**: Only one `filter[event_timestamp]` can be passed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DebuggingApiCallControlDebuggingEventListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DebuggingApiCallControlDebuggingEventListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCallLegId** | [**optional.Interface of string**](.md)| The unique identifier of an individual call leg. | 
 **filterCallSessionId** | [**optional.Interface of string**](.md)| The unique identifier of the call control session. A session may include multiple call leg events. | 
 **filterStatus** | **optional.String**| Event status | 
 **filterType** | **optional.String**| Event type | 
 **filterEventTimestampGt** | **optional.String**| Event timestamp: greater than | 
 **filterEventTimestampGte** | **optional.String**| Event timestamp: greater than or equal | 
 **filterEventTimestampLt** | **optional.String**| Event timestamp: lower than | 
 **filterEventTimestampLte** | **optional.String**| Event timestamp: lower than or equal | 
 **filterEventTimestampEq** | **optional.String**| Event timestamp: equal | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

