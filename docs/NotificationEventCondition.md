# NotificationEventCondition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMultipleChannels** | **bool** | Dictates whether a notification channel id needs to be provided when creating a notficiation setting. | [optional] [default to null]
**AssociatedRecordType** | **string** |  | [optional] [default to null]
**Asynchronous** | **bool** | Dictates whether a notification setting will take effect immediately. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | ISO 8601 formatted date indicating when the resource was created. | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**Id** | **string** | A UUID. | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**NotificationEventId** | **string** |  | [optional] [default to null]
**Parameters** | [**[]NotificationEventConditionParameters**](NotificationEventCondition_parameters.md) |  | [optional] [default to null]
**SupportedChannels** | **[]string** | Dictates the supported notification channel types that can be emitted. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | ISO 8601 formatted date indicating when the resource was updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

