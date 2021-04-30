# ConnectionRtcpSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaptureEnabled** | **bool** | BETA - Enable the capture and storage of RTCP messages to create QoS reports on the Telnyx Mission Control Portal. | [optional] [default to false]
**Port** | **string** | RTCP port by default is rtp+1, it can also be set to rtcp-mux | [optional] [default to rtp+1]
**ReportFrequencySecs** | **int32** | RTCP reports are sent to customers based on the frequency set. Frequency is in seconds and it can be set to values from 5 to 3000 seconds. | [optional] [default to 5]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

