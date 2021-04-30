# NotificationSetting

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedRecordType** | **string** |  | [optional] [default to null]
**AssociatedRecordTypeValue** | **string** |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | ISO 8601 formatted date indicating when the resource was created. | [optional] [default to null]
**Id** | **string** | A UUID. | [optional] [default to null]
**NotificationChannelId** | **string** | A UUID reference to the associated Notification Channel. | [optional] [default to null]
**NotificationEventConditionId** | **string** | A UUID reference to the associated Notification Event Condition. | [optional] [default to null]
**NotificationProfileId** | **string** | A UUID reference to the associated Notification Profile. | [optional] [default to null]
**Parameters** | [**[]NotificationSettingParameters**](NotificationSetting_parameters.md) |  | [optional] [default to null]
**Status** | **string** | Most preferences apply immediately; however, other may needs to propagate. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | ISO 8601 formatted date indicating when the resource was updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

