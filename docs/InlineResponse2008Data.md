# InlineResponse2008Data

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | **string** |  | [default to null]
**CallSessionId** | **string** | ID that is unique to the call session and can be used to correlate webhook events | [default to null]
**CallLegId** | **string** | ID that is unique to the call and can be used to correlate webhook events | [default to null]
**CallControlId** | **string** | Unique identifier and token for controlling the call. For Dial command it will always be &#x60;null&#x60; (dialing is asynchronous). | [default to null]
**IsAlive** | **bool** | Indicates whether the call is alive or not. For Dial command it will always be &#x60;false&#x60; (dialing is asynchronous). | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

