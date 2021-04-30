# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationChannels**](NotificationsApi.md#CreateNotificationChannels) | **Post** /notification_channels | Create a notification channel
[**CreateNotificationProfile**](NotificationsApi.md#CreateNotificationProfile) | **Post** /notification_profiles | Create a notification profile
[**CreateNotificationSetting**](NotificationsApi.md#CreateNotificationSetting) | **Post** /notification_settings | Add a Notification Setting
[**DeleteNotificationChannel**](NotificationsApi.md#DeleteNotificationChannel) | **Delete** /notification_channels/{id} | Delete a notification channel
[**DeleteNotificationProfile**](NotificationsApi.md#DeleteNotificationProfile) | **Delete** /notification_profiles/{id} | Delete a notification profile
[**DeleteNotificationSetting**](NotificationsApi.md#DeleteNotificationSetting) | **Delete** /notification_settings/{id} | Delete a notification setting
[**FindNotificationsEvents**](NotificationsApi.md#FindNotificationsEvents) | **Get** /notification_events | List all Notifications Events
[**FindNotificationsEventsConditions**](NotificationsApi.md#FindNotificationsEventsConditions) | **Get** /notification_event_conditions | List all Notifications Events Conditions
[**FindNotificationsProfiles**](NotificationsApi.md#FindNotificationsProfiles) | **Get** /notification_profiles | List all Notifications Profiles
[**ListNotificationChannels**](NotificationsApi.md#ListNotificationChannels) | **Get** /notification_channels | List notification channels
[**ListNotificationSettings**](NotificationsApi.md#ListNotificationSettings) | **Get** /notification_settings | List notification settings
[**RetrieveNotificationChannel**](NotificationsApi.md#RetrieveNotificationChannel) | **Get** /notification_channels/{id} | Retrieve a notification channel
[**RetrieveNotificationProfile**](NotificationsApi.md#RetrieveNotificationProfile) | **Get** /notification_profiles/{id} | Retrieve a notification profile
[**RetrieveNotificationSetting**](NotificationsApi.md#RetrieveNotificationSetting) | **Get** /notification_settings/{id} | Retrieve a notification setting
[**UpdateNotificationChannel**](NotificationsApi.md#UpdateNotificationChannel) | **Patch** /notification_channels/{id} | Update a notification channel
[**UpdateNotificationProfile**](NotificationsApi.md#UpdateNotificationProfile) | **Patch** /notification_profiles/{id} | Update a notification profile

# **CreateNotificationChannels**
> InlineResponse20016 CreateNotificationChannels(ctx, optional)
Create a notification channel

Create a notification channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationsApiCreateNotificationChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiCreateNotificationChannelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NotificationChannel**](NotificationChannel.md)| Add a Notification Channel | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNotificationProfile**
> InlineResponse20020 CreateNotificationProfile(ctx, optional)
Create a notification profile

Create a notification profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationsApiCreateNotificationProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiCreateNotificationProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NotificationProfile**](NotificationProfile.md)| Add a Notification Profile | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNotificationSetting**
> InlineResponse20022 CreateNotificationSetting(ctx, optional)
Add a Notification Setting

Add a notification setting.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationsApiCreateNotificationSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiCreateNotificationSettingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NotificationSetting**](NotificationSetting.md)|  | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNotificationChannel**
> InlineResponse20016 DeleteNotificationChannel(ctx, id)
Delete a notification channel

Delete a notification channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the resource. | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNotificationProfile**
> InlineResponse20020 DeleteNotificationProfile(ctx, id)
Delete a notification profile

Delete a notification profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the resource. | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNotificationSetting**
> InlineResponse20022 DeleteNotificationSetting(ctx, id)
Delete a notification setting

Delete a notification setting.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the resource. | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindNotificationsEvents**
> InlineResponse20018 FindNotificationsEvents(ctx, optional)
List all Notifications Events

Returns a list of your notifications events.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationsApiFindNotificationsEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiFindNotificationsEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindNotificationsEventsConditions**
> InlineResponse20017 FindNotificationsEventsConditions(ctx, optional)
List all Notifications Events Conditions

Returns a list of your notifications events conditions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationsApiFindNotificationsEventsConditionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiFindNotificationsEventsConditionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterAssociatedRecordTypeEq** | **optional.String**| Filter by the associated record type | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindNotificationsProfiles**
> InlineResponse20019 FindNotificationsProfiles(ctx, optional)
List all Notifications Profiles

Returns a list of your notifications profiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationsApiFindNotificationsProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiFindNotificationsProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNotificationChannels**
> InlineResponse20015 ListNotificationChannels(ctx, optional)
List notification channels

List notification channels.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationsApiListNotificationChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiListNotificationChannelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterChannelTypeIdEq** | **optional.String**| Filter by the id of a channel type | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNotificationSettings**
> InlineResponse20021 ListNotificationSettings(ctx, optional)
List notification settings

List notification settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationsApiListNotificationSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiListNotificationSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **filterNotificationProfileIdEq** | **optional.String**| Filter by the id of a notification profile | 
 **filterNotificationChannelEq** | **optional.String**| Filter by the id of a notification channel | 
 **filterNotificationEventConditionIdEq** | **optional.String**| Filter by the id of a notification channel | 
 **filterAssociatedRecordTypeEq** | **optional.String**| Filter by the associated record type | 
 **filterStatusEq** | **optional.String**| The status of a notification setting | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveNotificationChannel**
> InlineResponse20016 RetrieveNotificationChannel(ctx, id)
Retrieve a notification channel

Retrieve a notification channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the resource. | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveNotificationProfile**
> InlineResponse20020 RetrieveNotificationProfile(ctx, id)
Retrieve a notification profile

Retrieve a notification profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the resource. | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveNotificationSetting**
> InlineResponse20022 RetrieveNotificationSetting(ctx, id)
Retrieve a notification setting

Retrieve a notification setting.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The id of the resource. | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotificationChannel**
> InlineResponse20016 UpdateNotificationChannel(ctx, body, id)
Update a notification channel

Update a notification channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NotificationChannel**](NotificationChannel.md)| Update notification channel object | 
  **id** | [**string**](.md)| The id of the resource. | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotificationProfile**
> InlineResponse20020 UpdateNotificationProfile(ctx, body, id)
Update a notification profile

Update a notification profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NotificationProfile**](NotificationProfile.md)| Update notification profile object | 
  **id** | [**string**](.md)| The id of the resource. | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

