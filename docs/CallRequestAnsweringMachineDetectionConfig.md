# CallRequestAnsweringMachineDetectionConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfterGreetingSilenceMillis** | **int32** | Silence duration threshold after a greeting message or voice for it be considered human. | [optional] [default to 800]
**BetweenWordsSilenceMillis** | **int32** | Maximum threshold for silence between words. | [optional] [default to 50]
**GreetingDurationMillis** | **int32** | Maximum threshold of a human greeting. If greeting longer than this value, considered machine. | [optional] [default to 3500]
**GreetingSilenceDurationMillis** | **int32** | If machine already detected, maximum threshold for silence between words. If exceeded, the greeting is considered ended. | [optional] [default to 1500]
**GreetingTotalAnalysisTimeMillis** | **int32** | If machine already detected, maximum timeout threshold to determine the end of the machine greeting. | [optional] [default to 5000]
**InitialSilenceMillis** | **int32** | If initial silence duration is greater than this value, consider it a machine. | [optional] [default to 3500]
**MaximumNumberOfWords** | **int32** | If number of detected words is greater than this value, consder it a machine. | [optional] [default to 5]
**MaximumWordLengthMillis** | **int32** | If a single word lasts longer than this threshold, consider it a machine. | [optional] [default to 3500]
**SilenceThreshold** | **int32** | Minimum noise threshold for any analysis. | [optional] [default to 256]
**TotalAnalysisTimeMillis** | **int32** | Maximum timeout threshold for overall detection. | [optional] [default to 3500]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

