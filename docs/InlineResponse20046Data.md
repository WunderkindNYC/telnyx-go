# InlineResponse20046Data

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifies the resource. | [optional] [default to null]
**RecordType** | **string** | Identifies the type of the resource. | [optional] [default to null]
**PhoneNumber** | **string** | The +E.164-formatted phone number associated with this record | [optional] [default to null]
**Status** | **string** | The phone number&#x27;s current status | [optional] [default to null]
**Tags** | **[]string** | A list of user-assigned tags to help manage the phone number. | [optional] [default to null]
**ExternalPin** | **string** | If someone attempts to port your phone number away from Telnyx and your phone number has an external PIN set, Telnyx will attempt to verify that you provided the correct external PIN to the winning carrier. Note that not all carriers cooperate with this security mechanism. | [optional] [default to null]
**ConnectionName** | **string** | The user-assigned name of the connection to be associated with this phone number. | [optional] [default to null]
**ConnectionId** | **string** | Identifies the connection associated with the phone number. | [optional] [default to null]
**MessagingEnabled** | **bool** | Indicates whether messaging services are enabled for this number. | [optional] [default to null]
**MessagingProfileId** | **string** | Identifies the messaging profile associated with the phone number. | [optional] [default to null]
**BillingGroupId** | **string** | Identifies the billing group associated with the phone number. | [optional] [default to null]
**EmergencyEnabled** | **bool** | Indicates whether emergency services are enabled for this number. | [optional] [default to null]
**EmergencyAddressId** | **string** | Identifies the emergency address associated with the phone number. | [optional] [default to null]
**CallForwardingEnabled** | **bool** | Indicates if call forwarding will be enabled for this number if forwards_to and forwarding_type are filled in. Defaults to true for backwards compatibility with APIV1 use of numbers endpoints. | [optional] [default to true]
**CnamListingEnabled** | **bool** | Indicates whether a CNAM listing is enabled for this number. | [optional] [default to null]
**CallerIdNameEnabled** | **bool** | Indicates whether caller ID is enabled for this number. | [optional] [default to null]
**CallRecordingEnabled** | **bool** | Indicates whether call recording is enabled for this number. | [optional] [default to null]
**T38FaxGatewayEnabled** | **bool** | Indicates whether T38 Fax Gateway for inbound calls to this number. | [optional] [default to null]
**PurchasedAt** | **string** | ISO 8601 formatted date indicating when the resource was purchased. | [optional] [default to null]
**CreatedAt** | **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] [default to null]
**UpdatedAt** | **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

