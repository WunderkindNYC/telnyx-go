# StartForkingRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientState** | **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] [default to null]
**CommandId** | **string** | Use this field to avoid duplicate commands. Telnyx will ignore commands with the same &#x60;command_id&#x60;. | [optional] [default to null]
**Rx** | **string** | The network target, &lt;udp:ip_address:port&gt;, where the call&#x27;s incoming RTP media packets should be forwarded. | [optional] [default to null]
**StreamType** | **string** | Optionally specify a media type to stream. If &#x60;decrpyted&#x60; selected, Telnyx will decrypt incoming SIP media before forking to the target. &#x60;rx&#x60; and &#x60;tx&#x60; are required fields if &#x60;decrypted&#x60; selected. | [optional] [default to STREAM_TYPE.RAW]
**Target** | **string** | The network target, &lt;udp:ip_address:port&gt;, where the call&#x27;s RTP media packets should be forwarded. Both incoming and outgoing media packets will be delivered to the specified target, and information about the stream will be included in the encapsulation protocol header, including the direction (0 &#x3D; inbound; 1 &#x3D; outbound), leg (0 &#x3D; A-leg; 1 &#x3D; B-leg), and call_leg_id. | [optional] [default to null]
**Tx** | **string** | The network target, &lt;udp:ip_address:port&gt;, where the call&#x27;s outgoing RTP media packets should be forwarded. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

