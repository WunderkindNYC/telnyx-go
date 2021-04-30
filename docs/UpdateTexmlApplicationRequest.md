# UpdateTexmlApplicationRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | [optional] [default to null]
**AnchorsiteOverride** | [***AnchorsiteOverride**](AnchorsiteOverride.md) |  | [optional] [default to null]
**DtmfType** | [***DtmfType**](DtmfType.md) |  | [optional] [default to null]
**FirstCommandTimeout** | **bool** |  | [optional] [default to null]
**FirstCommandTimeoutSecs** | **int32** |  | [optional] [default to null]
**FriendlyName** | **string** |  | [default to null]
**Inbound** | [***CreateTexmlApplicationRequestInbound**](CreateTexmlApplicationRequest_inbound.md) |  | [optional] [default to null]
**Outbound** | [***CreateTexmlApplicationRequestOutbound**](CreateTexmlApplicationRequest_outbound.md) |  | [optional] [default to null]
**StatusCallback** | **string** | URL for Telnyx to send requests to containing information about call progress events. | [optional] [default to null]
**StatusCallbackMethod** | **string** | HTTP request method Telnyx should use when requesting the status_callback URL. | [optional] [default to STATUS_CALLBACK_METHOD.POST]
**VoiceFallbackUrl** | **string** | URL to which Telnyx will deliver your XML Translator webhooks if we get an error response from your voice_url. | [optional] [default to null]
**VoiceMethod** | **string** | HTTP request method Telnyx will use to interact with your XML Translator webhooks. Either &#x27;get&#x27; or &#x27;post&#x27;. | [optional] [default to VOICE_METHOD.POST]
**VoiceUrl** | **string** | URL to which Telnyx will deliver your XML Translator webhooks. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

