# ReferRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientState** | **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] [default to null]
**CommandId** | **string** | Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same &#x60;command_id&#x60; as one that has already been executed. | [optional] [default to null]
**CustomHeaders** | [**[]CustomSipHeader**](CustomSipHeader.md) | Custom headers to be added to the SIP INVITE. | [optional] [default to null]
**SipAddress** | **string** | The SIP URI to which the call will be referred to. | [default to null]
**SipAuthPassword** | **string** | SIP Authentication password used for SIP challenges. | [optional] [default to null]
**SipAuthUsername** | **string** | SIP Authentication username used for SIP challenges. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

