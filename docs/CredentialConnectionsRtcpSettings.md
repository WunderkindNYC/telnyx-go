# CredentialConnectionsRtcpSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **string** | RTCP port by default is rtp+1, it can also be set to rtcp-mux | [optional] [default to rtp+1]
**ReportEnabled** | **bool** | DEPRECATED - RTCP reports are always sent to customers/vendors. For backwards compatibility, the value is always true. | [optional] [default to true]
**ReportFrequencySeconds** | **int32** | RTCP reports are sent to customers based on the frequency set. Frequency is in seconds and it can be set to values from 5 to 3000 seconds. | [optional] [default to 10]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

