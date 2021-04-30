# InlineResponse20044Data

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | **string** | Identifies the type of the resource. | [optional] [default to null]
**Id** | **string** | Identifies the type of resource. | [optional] [default to null]
**PhoneNumber** | **string** | +E.164 formatted phone number. | [optional] [default to null]
**MessagingProfileId** | **string** | Unique identifier for a messaging profile. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | ISO 8601 formatted date indicating when the resource was created. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | ISO 8601 formatted date indicating when the resource was updated. | [optional] [default to null]
**CountryCode** | **string** | ISO 3166-1 alpha-2 country code. | [optional] [default to null]
**Type_** | **string** | The type of the phone number | [optional] [default to null]
**Health** | [***InlineResponse20026DataHealth**](inline_response_200_26_data_health.md) |  | [optional] [default to null]
**EligibleMessagingProducts** | **[]string** | The messaging products that this number can be registered to use | [optional] [default to null]
**TrafficType** | **string** | The messaging traffic or use case for which the number is currently configured. | [optional] [default to null]
**MessagingProduct** | **string** | The messaging product that the number is registered to use | [optional] [default to null]
**Features** | [***InlineResponse20026DataFeatures**](inline_response_200_26_data_features.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

