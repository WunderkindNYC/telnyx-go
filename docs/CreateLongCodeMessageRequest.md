# CreateLongCodeMessageRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDetect** | **bool** | Automatically detect if an SMS message is unusually long and exceeds a recommended limit of message parts. | [optional] [default to false]
**From** | **string** | Phone number, in +E.164 format, used to send the message. | [default to null]
**MediaUrls** | **[]string** | A list of media URLs. The total media size must be less than 1 MB.  **Required for MMS** | [optional] [default to null]
**Subject** | **string** | Subject of multimedia message | [optional] [default to null]
**Text** | **string** | Message body (i.e., content) as a non-empty string.  **Required for SMS** | [optional] [default to null]
**To** | **string** |  | [default to null]
**Type_** | **string** | The protocol for sending the message, either SMS or MMS. | [optional] [default to null]
**UseProfileWebhooks** | **bool** | If the profile this number is associated with has webhooks, use them for delivery notifications. If webhooks are also specified on the message itself, they will be attempted first, then those on the profile. | [optional] [default to true]
**WebhookFailoverUrl** | **string** | The failover URL where webhooks related to this message will be sent if sending to the primary URL fails. | [optional] [default to null]
**WebhookUrl** | **string** | The URL where webhooks related to this message will be sent. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

