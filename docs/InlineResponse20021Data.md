# InlineResponse20021Data

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | **string** | Identifies the type of the resource. | [optional] [default to null]
**Direction** | **string** | The direction of the message. Inbound messages are sent to you whereas outbound messages are sent from you. | [optional] [default to null]
**Id** | **string** | Identifies the type of resource. | [optional] [default to null]
**Type_** | **string** | The type of message. | [optional] [default to null]
**MessagingProfileId** | **string** | Unique identifier for a messaging profile. | [optional] [default to null]
**From** | **string** | Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short code). | [optional] [default to null]
**To** | [**[]InlineResponse20021DataTo**](inline_response_200_21_data_to.md) |  | [optional] [default to null]
**Text** | **string** | Message body (i.e., content) as a non-empty string.  **Required for SMS** | [optional] [default to null]
**Subject** | **string** | Subject of multimedia message | [optional] [default to null]
**Media** | [**[]InlineResponse20021DataMedia**](inline_response_200_21_data_media.md) |  | [optional] [default to null]
**WebhookUrl** | **string** | The URL where webhooks related to this message will be sent. | [optional] [default to null]
**WebhookFailoverUrl** | **string** | The failover URL where webhooks related to this message will be sent if sending to the primary URL fails. | [optional] [default to null]
**Encoding** | **string** | Encoding scheme used for the message body. | [optional] [default to null]
**Parts** | **int32** | Number of parts into which the message&#x27;s body must be split. | [optional] [default to null]
**Tags** | **[]string** | Tags associated with the resource. | [optional] [default to null]
**Cost** | [***InlineResponse20021DataCost**](inline_response_200_21_data_cost.md) |  | [optional] [default to null]
**ReceivedAt** | [**time.Time**](time.Time.md) | ISO 8601 formatted date indicating when the message request was received. | [optional] [default to null]
**SentAt** | [**time.Time**](time.Time.md) | ISO 8601 formatted date indicating when the message was sent. | [optional] [default to null]
**CompletedAt** | [**time.Time**](time.Time.md) | ISO 8601 formatted date indicating when the message was finalized. | [optional] [default to null]
**ValidUntil** | [**time.Time**](time.Time.md) | Message must be out of the queue by this time or else it will be discarded and marked as &#x27;sending_failed&#x27;. Once the message moves out of the queue, this field will be nulled | [optional] [default to null]
**Errors** | [**[]InlineResponse20021DataErrors**](inline_response_200_21_data_errors.md) | These errors may point at addressees when referring to unsuccessful/unconfirmed delivery statuses. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

