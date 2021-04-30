# Address

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Uniquely identifies the address. | [optional] [default to null]
**RecordType** | **string** | Identifies the type of the resource. | [optional] [default to null]
**FirstName** | **string** | The first name associated with the address. An address must have either a first last name or a business name. | [optional] [default to null]
**LastName** | **string** | The last name associated with the address. An address must have either a first last name or a business name. | [optional] [default to null]
**BusinessName** | **string** | The business name associated with the address. An address must have either a first last name or a business name. | [optional] [default to null]
**PhoneNumber** | **string** | The phone number associated with the address. | [optional] [default to null]
**StreetAddress** | **string** | The primary street address information about the address. | [optional] [default to null]
**ExtendedAddress** | **string** | Additional street address information about the address such as, but not limited to, unit number or apartment number. | [optional] [default to null]
**Locality** | **string** | The locality of the address. For US addresses, this corresponds to the city of the address. | [optional] [default to null]
**AdministrativeArea** | **string** | The locality of the address. For US addresses, this corresponds to the state of the address. | [optional] [default to null]
**Neighborhood** | **string** | The neighborhood of the address. This field is not used for addresses in the US but is used for some international addresses. | [optional] [default to null]
**Borough** | **string** | The borough of the address. This field is not used for addresses in the US but is used for some international addresses. | [optional] [default to null]
**PostalCode** | **string** | The postal code of the address. | [optional] [default to null]
**CountryCode** | **string** | The two-character (ISO 3166-1 alpha-2) country code of the address. | [optional] [default to null]
**AddressBook** | **bool** | Indicates whether or not the address should be considered part of your list of addresses that appear for regular use. | [optional] [default to true]
**CreatedAt** | **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] [default to null]
**UpdatedAt** | **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

