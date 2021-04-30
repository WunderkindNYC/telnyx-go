# PrivateWirelessGatewayStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** | This attribute is an &lt;a href&#x3D;\&quot;https://developers.telnyx.com/docs/api/v2/overview#errors\&quot;&gt;error code&lt;/a&gt; related to the failure reason. | [optional] [default to null]
**ErrorDescription** | **string** | This attribute provides a human-readable explanation of why a failure happened. | [optional] [default to null]
**Value** | **string** | The current status or failure details of the Private Wireless Gateway. &lt;ul&gt;  &lt;li&gt;&lt;code&gt;provisioning&lt;/code&gt; - the Private Wireless Gateway is being provisioned.&lt;/li&gt;  &lt;li&gt;&lt;code&gt;provisioned&lt;/code&gt; - the Private Wireless Gateway was provisioned and able to receive connections.&lt;/li&gt;  &lt;li&gt;&lt;code&gt;failed&lt;/code&gt; - the provisioning had failed for a reason and it requires an intervention.&lt;/li&gt;  &lt;li&gt;&lt;code&gt;decommissioning&lt;/code&gt; - the Private Wireless Gateway is being removed from the network.&lt;/li&gt;  &lt;/ul&gt;  Transitioning between the provisioning and provisioned states may take some time. | [optional] [default to VALUE.PROVISIONING]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

