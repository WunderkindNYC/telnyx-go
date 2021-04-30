# MediaFeatures

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptAnyRtpPacketsEnabled** | **bool** | When enabled, Telnyx will accept RTP packets from any customer-side IP address and port, not just those to which Telnyx is sending RTP. | [optional] [default to false]
**MediaHandlingMode** | **string** | Controls how media is handled for the phone number. default: media routed through Telnyx with transcode support. proxy: media routed through Telnyx with no transcode support. | [optional] [default to MEDIA_HANDLING_MODE.DEFAULT_]
**RtpAutoAdjustEnabled** | **bool** | When RTP Auto-Adjust is enabled, the destination RTP address port will be automatically changed to match the source of the incoming RTP packets. | [optional] [default to true]
**T38FaxGatewayEnabled** | **bool** | Controls whether Telnyx will accept a T.38 re-INVITE for this phone number. Note that Telnyx will not send a T.38 re-INVITE; this option only controls whether one will be accepted. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

