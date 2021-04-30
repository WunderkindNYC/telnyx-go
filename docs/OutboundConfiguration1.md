# OutboundConfiguration1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallParkingEnabled** | **bool** | Forces all SIP calls originated on this connection to be \&quot;parked\&quot; instead of \&quot;bridged\&quot; to the destination specified on the URI. Parked calls will return ringback to the caller and will await for a Call Control command to define which action will be taken next. | [optional] [default to false]
**AniOverride** | **string** | Set a phone number as the ani_override value to override caller id number on outbound calls. | [optional] 
**AniOverrideType** | **string** | Specifies when we apply your ani_override setting. Only applies when ani_override is not blank. | [optional] [default to ANI_OVERRIDE_TYPE.ALWAYS]
**ChannelLimit** | **int32** | When set, this will limit the total number of outbound calls to phone numbers associated with this connection. | [optional] [default to null]
**InstantRingbackEnabled** | **bool** | When set, ringback will not wait for indication before sending ringback tone to calling party. | [optional] [default to true]
**GenerateRingbackTone** | **bool** | Generate ringback tone through 183 session progress message with early media. | [optional] [default to false]
**Localization** | **string** | A 2-character country code specifying the country whose national dialing rules should be used. For example, if set to &#x60;US&#x60; then any US number can be dialed without preprending +1 to the number. When left blank, Telnyx will try US and GB dialing rules, in that order, by default. | [optional] [default to null]
**T38ReinviteSource** | **string** | This setting only affects connections with Fax-type Outbound Voice Profiles. The setting dictates whether or not Telnyx sends a t.38 reinvite.&lt;br/&gt;&lt;br/&gt; By default, Telnyx will send the re-invite. If set to &#x60;customer&#x60;, the caller is expected to send the t.38 reinvite. | [optional] [default to T38_REINVITE_SOURCE.TELNYX]
**TechPrefix** | **string** | Numerical chars only, exactly 4 characters. | [optional] 
**IpAuthenticationMethod** | **string** |  | [optional] [default to tech-prefix]
**IpAuthenticationToken** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

