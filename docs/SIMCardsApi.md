# {{classname}}

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SimCardActivate**](SIMCardsApi.md#SimCardActivate) | **Post** /sim_cards/{id}/actions/activate | Request a SIM card activation
[**SimCardDeactivate**](SIMCardsApi.md#SimCardDeactivate) | **Post** /sim_cards/{id}/actions/deactivate | Request a SIM card deactivation
[**SimCardGet**](SIMCardsApi.md#SimCardGet) | **Get** /sim_cards/{id} | Get SIM card
[**SimCardRegister**](SIMCardsApi.md#SimCardRegister) | **Post** /actions/register/sim_cards | Register SIM cards
[**SimCardUpdate**](SIMCardsApi.md#SimCardUpdate) | **Patch** /sim_cards/{id} | Update a SIM card
[**SimCardsGet**](SIMCardsApi.md#SimCardsGet) | **Get** /sim_cards | Get all SIM cards

# **SimCardActivate**
> InlineResponse202 SimCardActivate(ctx, id)
Request a SIM card activation

The SIM card will be able to connect to the network once the activation is complete, thus making it possible to consume data.<br/> To activate a SIM card, it must be associated with a SIM card group that has a data plan.<br/> Transitioning to the active state may take a period of time. Until the transition is completed, the SIM card status will be <code>activating</code>. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse202**](inline_response_202.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardDeactivate**
> InlineResponse202 SimCardDeactivate(ctx, id)
Request a SIM card deactivation

The SIM card won't be able to connect to the network after the deactivation is completed, thus making it impossible to consume data.<br/> Transitioning to the inactive state may take a period of time.</br> Until the transition is completed, the SIM card status will be inactivating <code>inactivating</code>. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse202**](inline_response_202.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardGet**
> InlineResponse20057 SimCardGet(ctx, id, optional)
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

[**InlineResponse20057**](inline_response_200_57.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardRegister**
> InlineResponse200 SimCardRegister(ctx, body)
Register SIM cards

Register the SIM cards associated with the provided registration codes to the current user's account.<br/><br/> If <code>sim_card_group_id</code> is provided, the SIM cards will be associated with that group. Otherwise, the default group for the current user, which is associated with the default free data plan, will be used.<br/><br/> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body**](Body.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardUpdate**
> InlineResponse20057 SimCardUpdate(ctx, body, id)
Update a SIM card

Updates a SIM card's group and tags

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body36**](Body36.md)|  | 
  **id** | [**string**](.md)| Identifies the resource. | 

### Return type

[**InlineResponse20057**](inline_response_200_57.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimCardsGet**
> InlineResponse20056 SimCardsGet(ctx, optional)
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

[**InlineResponse20056**](inline_response_200_56.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

