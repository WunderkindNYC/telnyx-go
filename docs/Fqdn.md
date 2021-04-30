# Fqdn

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifies the resource. | [optional] [default to null]
**RecordType** | **string** | Identifies the type of the resource. | [optional] [default to null]
**ConnectionId** | **string** | ID of the FQDN connection to which this FQDN is attached. | [optional] [default to null]
**Fqdn** | **string** | FQDN represented by this resource. | [optional] [default to null]
**Port** | **int32** | Port to use when connecting to this FQDN. | [optional] [default to 5060]
**DnsRecordType** | **string** | The DNS record type for the FQDN. For cases where a port is not set, the DNS record type must be &#x27;srv&#x27;. For cases where a port is set, the DNS record type must be &#x27;a&#x27;. If the DNS record type is &#x27;a&#x27; and a port is not specified, 5060 will be used. | [optional] [default to null]
**CreatedAt** | **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] [default to null]
**UpdatedAt** | **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

