# SimpleSimCard

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [default to null]
**Iccid** | **string** | The ICCID is the identifier of the specific SIM card/chip. Each SIM is internationally identified by its integrated circuit card identifier (ICCID). ICCIDs are stored in the SIM card&#x27;s memory and are also engraved or printed on the SIM card body during a process called personalization.  | [optional] [default to null]
**Id** | **string** | Identifies the resource. | [optional] [default to null]
**Imsi** | **string** | SIM cards are identified on their individual operator networks by a unique International Mobile Subscriber Identity (IMSI). &lt;br/&gt; Mobile network operators connect mobile phone calls and communicate with their market SIM cards using their IMSIs. The IMSI is stored in the Subscriber  Identity Module (SIM) inside the device and is sent by the device to the appropriate network. It is used to acquire the details of the device in the Home  Location Register (HLR) or the Visitor Location Register (VLR).  | [optional] [default to null]
**Msisdn** | **string** | Mobile Station International Subscriber Directory Number (MSISDN) is a number used to identify a mobile phone number internationally. &lt;br/&gt; MSISDN is defined by the E.164 numbering plan. It includes a country code and a National Destination Code which identifies the subscriber&#x27;s operator.  | [optional] [default to null]
**RecordType** | **string** |  | [optional] [default to null]
**SimCardGroupId** | **string** | The group SIMCardGroup identification. This attribute can be &lt;code&gt;null&lt;/code&gt; when it&#x27;s present in an associated resource. | [optional] [default to null]
**Status** | **string** | The current status of the SIM card. It will be one of the following: &lt;br/&gt; &lt;ul&gt;  &lt;li&gt;&lt;code&gt;registering&lt;/code&gt; - the card is being registered&lt;/li&gt;  &lt;li&gt;&lt;code&gt;enabling&lt;/code&gt; - the card is being enabled&lt;/li&gt;  &lt;li&gt;&lt;code&gt;enabled&lt;/code&gt; - the card is enabled and ready for use&lt;/li&gt;  &lt;li&gt;&lt;code&gt;disabling&lt;/code&gt; - the card is being disabled&lt;/li&gt;  &lt;li&gt;&lt;code&gt;disabled&lt;/code&gt; - the card has been disabled and cannot be used&lt;/li&gt;  &lt;li&gt;&lt;code&gt;data_limit_exceeded&lt;/code&gt; - the card has exceeded its data consumption limit&lt;/li&gt; &lt;/ul&gt; Transitioning between the enabled and disabled states may take a period of time.  | [optional] [default to null]
**Tags** | **[]string** | Searchable tags associated with the SIM card | [optional] [default to null]
**UpdatedAt** | **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

