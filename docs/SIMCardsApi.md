# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkSIMCardNetworkPreferences**](SIMCardsApi.md#BulkSIMCardNetworkPreferences) | **Put** /actions/network_preferences/sim_cards | Bulk Network Preferences for SIM cards
[**SIMCardNetworkPreferencesDelete**](SIMCardsApi.md#SIMCardNetworkPreferencesDelete) | **Delete** /sim_cards/{sim_card_id}/network_preferences | DELETE network preferences
[**SIMCardNetworkPreferencesGet**](SIMCardsApi.md#SIMCardNetworkPreferencesGet) | **Get** /sim_cards/{sim_card_id}/network_preferences | Get network preferences
[**SIMCardNetworkPreferencesPut**](SIMCardsApi.md#SIMCardNetworkPreferencesPut) | **Put** /sim_cards/{sim_card_id}/network_preferences | Set network preferences
[**SIMCardPublicIPDelete**](SIMCardsApi.md#SIMCardPublicIPDelete) | **Delete** /sim_cards/{sim_card_id}/public_ip | Delete SIM card public IP
[**SIMCardPublicIPGet**](SIMCardsApi.md#SIMCardPublicIPGet) | **Get** /sim_cards/{sim_card_id}/public_ip | Get SIM card public IP definition
[**SIMCardPublicIPPost**](SIMCardsApi.md#SIMCardPublicIPPost) | **Post** /sim_cards/{sim_card_id}/public_ip | Set SIM card public IP
[**SimCardDelete**](SIMCardsApi.md#SimCardDelete) | **Delete** /sim_cards/{id} | Deletes a SIM card
[**SimCardDisable**](SIMCardsApi.md#SimCardDisable) | **Post** /sim_cards/{id}/actions/disable | Request a SIM card disable
[**SimCardEnable**](SIMCardsApi.md#SimCardEnable) | **Post** /sim_cards/{id}/actions/enable | Request a SIM card enable
[**SimCardGet**](SIMCardsApi.md#SimCardGet) | **Get** /sim_cards/{id} | Get SIM card
[**SimCardRegister**](SIMCardsApi.md#SimCardRegister) | **Post** /actions/register/sim_cards | Register SIM cards
[**SimCardSetStandby**](SIMCardsApi.md#SimCardSetStandby) | **Post** /sim_cards/{id}/actions/set_standby | Request setting a SIM card to standby
[**SimCardUpdate**](SIMCardsApi.md#SimCardUpdate) | **Patch** /sim_cards/{id} | Update a SIM card
[**SimCardsGet**](SIMCardsApi.md#SimCardsGet) | **Get** /sim_cards | Get all SIM cards

# **BulkSIMCardNetworkPreferences**
> InlineResponse202 BulkSIMCardNetworkPreferences(ctx, optional)
Bulk Network Preferences for SIM cards

This API allows dispatching the same operation described for the PUT sim_cards/:sim_card_id/network_preferences API for multiple SIM cards at once.<br/><br/> Although, a SIM card network preference may fail individually under any validation triggered as a consequence of its state. For example, a SIM can't have an in-progress OTA update for applying a Network Preference, so they'll fail when requested in this API. In that scenario, the specific error will be present in the response along with the successful definitions in the \"errors\" response node. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SIMCardsApiBulkSIMCardNetworkPreferencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SIMCardsApiBulkSIMCardNetworkPreferencesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body**](Body.md)|  | 

### Return type

[**InlineResponse202**](inline_response_202.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SIMCardNetworkPreferencesDelete**
> InlineResponse20040 SIMCardNetworkPreferencesDelete(ctx, simCardId)
DELETE network preferences

This API asynchronously removes the custom-defined network preferences settings. After this operation is done the Telnyx default settings, the same applied for an unaltered SIM card, will be in place. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **simCardId** | [**string**](.md)| Identifies a SIM card. | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SIMCardNetworkPreferencesGet**
> InlineResponse20040 SIMCardNetworkPreferencesGet(ctx, simCardId, optional)
Get network preferences

It returns the network preferences currently applied in the SIM card. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **simCardId** | [**string**](.md)| Identifies a SIM card. | 
 **optional** | ***SIMCardsApiSIMCardNetworkPreferencesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SIMCardsApiSIMCardNetworkPreferencesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeOtaUpdates** | **optional.Bool**| It includes the associated OTA update objects in the response when present. | [default to false]

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SIMCardNetworkPreferencesPut**
> InlineResponse20040 SIMCardNetworkPreferencesPut(ctx, simCardId, optional)
Set network preferences

This API allows setting or updating a SIM card network preference. <br/><br/> Every SIM card has default network preferences defined on Telnyx. These preferences will determine how a SIMCard will connect to the network by considering a list of preferable operators.<br/><br/> There can be multiple scenarios where an operator can be preferred over another, for example, when a specific mobile operator can provide better network latency or better pricing. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **simCardId** | [**string**](.md)| Identifies a SIM card. | 
 **optional** | ***SIMCardsApiSIMCardNetworkPreferencesPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SIMCardsApiSIMCardNetworkPreferencesPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body5**](Body5.md)|  | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SIMCardPublicIPDelete**
> InlineResponse20041 SIMCardPublicIPDelete(ctx, simCardId)
Delete SIM card public IP

This API asynchronously removes an existing public IP from a SIM card. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **simCardId** | [**string**](.md)| Identifies a SIM card. | 

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SIMCardPublicIPGet**
> InlineResponse20041 SIMCardPublicIPGet(ctx, simCardId)
Get SIM card public IP definition

It returns the public IP requested for a SIM card. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **simCardId** | [**string**](.md)| Identifies a SIM card. | 

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SIMCardPublicIPPost**
> InlineResponse20041 SIMCardPublicIPPost(ctx, simCardId)
Set SIM card public IP

This API will asynchronously map a random public IP to a SIM card, making it reachable on the public internet. <br/><br/>  The request will return the resource as \"provisioning\" right away, and it'll eventually get \"provisioned\", meaning it completed the setup. <br/><br/>  Setting up a public IP will generate charges, and we won't be able to provide the IP if the account doesn't have a balance. In that case, this operation will succeed, but its status will change to failed eventually, and the resource will be updated with the associated status. The IP resource will need to be deleted and re-created with our DELETE and POST APIs in this scenario.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **simCardId** | [**string**](.md)| Identifies a SIM card. | 

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardDelete**
> InlineResponse20039 SimCardDelete(ctx, id)
Deletes a SIM card

The SIM card will be decommissioned, removed from your account and you will stop being charged.<br />The SIM card won't be able to connect to the network after the deletion is completed, thus making it impossible to consume data.<br/> Transitioning to the disabled state may take a period of time.</br> Until the transition is completed, the SIM card status will be disabling <code>disabling</code>.<br />In order to re-enable the SIM card, you will need to re-register it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardDisable**
> InlineResponse2024 SimCardDisable(ctx, id)
Request a SIM card disable

The SIM card won't be able to connect to the network after the disabling is completed, thus making it impossible to consume data.<br/> Transitioning to the disabled state may take a period of time.</br> Until the transition is completed, the SIM card status will be <code>disabling</code>. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse2024**](inline_response_202_4.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardEnable**
> InlineResponse2024 SimCardEnable(ctx, id)
Request a SIM card enable

The SIM card will be able to connect to the network once the enabling is complete, thus making it possible to consume data.<br/> To enable a SIM card, it must be associated with SIM card group.<br/> Transitioning to the enabled state may take a period of time. Until the transition is completed, the SIM card status will be <code>enabling</code>. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse2024**](inline_response_202_4.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardGet**
> InlineResponse20039 SimCardGet(ctx, id, optional)
Get SIM card

Returns the details regarding a specific SIM card.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 
 **optional** | ***SIMCardsApiSimCardGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SIMCardsApiSimCardGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSimCardGroup** | **optional.Bool**| It includes the associated SIM card group object in the response when present. | [default to false]

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardRegister**
> InlineResponse2021 SimCardRegister(ctx, body)
Register SIM cards

Register the SIM cards associated with the provided registration codes to the current user's account.<br/><br/> If <code>sim_card_group_id</code> is provided, the SIM cards will be associated with that group. Otherwise, the default group for the current user will be used.<br/><br/> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SimCardRegistration**](SimCardRegistration.md)|  | 

### Return type

[**InlineResponse2021**](inline_response_202_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardSetStandby**
> InlineResponse2024 SimCardSetStandby(ctx, id)
Request setting a SIM card to standby

The SIM card will be able to connect to the network once the process to set it to standby has been completed, thus making it possible to consume data.<br/> To set a SIM card to standby, it must be associated with SIM card group.<br/> Transitioning to the standby state may take a period of time. Until the transition is completed, the SIM card status will be <code>setting_standby</code>. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse2024**](inline_response_202_4.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardUpdate**
> InlineResponse20039 SimCardUpdate(ctx, body, id)
Update a SIM card

Updates a SIM card's group and tags

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SimCard**](SimCard.md)|  | 
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardsGet**
> InlineResponse20038 SimCardsGet(ctx, optional)
Get all SIM cards

Get all SIM cards belonging to the user that match the given filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SIMCardsApiSimCardsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SIMCardsApiSimCardsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| The page number to load | [default to 1]
 **pageSize** | **optional.Int32**| The size of the page | [default to 20]
 **includeSimCardGroup** | **optional.Bool**| It includes the associated SIM card group object in the response when present. | [default to false]
 **filterSimCardGroupId** | [**optional.Interface of string**](.md)| A valid SIM card group ID. | 
 **filterTags** | [**optional.Interface of []string**](string.md)| A list of SIM card tags to filter on.&lt;br/&gt;&lt;br/&gt; If the SIM card contains &lt;b&gt;&lt;i&gt;all&lt;/i&gt;&lt;/b&gt; of the given &lt;code&gt;tags&lt;/code&gt; they will be found.&lt;br/&gt;&lt;br/&gt; For example, if the SIM cards have the following tags: &lt;ul&gt;   &lt;li&gt;&lt;code&gt;[&#x27;customers&#x27;, &#x27;staff&#x27;, &#x27;test&#x27;]&lt;/code&gt;   &lt;li&gt;&lt;code&gt;[&#x27;test&#x27;]&lt;/code&gt;&lt;/li&gt;   &lt;li&gt;&lt;code&gt;[&#x27;customers&#x27;]&lt;/code&gt;&lt;/li&gt; &lt;/ul&gt; Searching for &lt;code&gt;[&#x27;customers&#x27;, &#x27;test&#x27;]&lt;/code&gt; returns only the first because it&#x27;s the only one with both tags.&lt;br/&gt; Searching for &lt;code&gt;test&lt;/code&gt; returns the first two SIMs, because both of them have such tag.&lt;br/&gt; Searching for &lt;code&gt;customers&lt;/code&gt; returns the first and last SIMs.&lt;br/&gt;  | 
 **filterIccid** | **optional.String**| A search string to partially match for the SIM card&#x27;s ICCID. | 

### Return type

[**InlineResponse20038**](inline_response_200_38.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

