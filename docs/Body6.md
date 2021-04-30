# Body6

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | ID of the FQDN connection to which this IP should be attached. | [optional] [default to null]
**Fqdn** | **string** | FQDN represented by this resource. | [default to null]
**Port** | **int32** | Port to use when connecting to this FQDN. | [optional] [default to 5060]
**DnsRecordType** | **string** | The DNS record type for the FQDN. For cases where a port is not set, the DNS record type must be &#x27;srv&#x27;. For cases where a port is set, the DNS record type must be &#x27;a&#x27;. If the DNS record type is &#x27;a&#x27; and a port is not specified, 5060 will be used. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

