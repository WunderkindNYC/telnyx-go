# Body13

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short code).  **Required if sending with a phone number, short code, or alphanumeric sender ID.**  | [optional] [default to null]
**MessagingProfileId** | **string** | Unique identifier for a messaging profile.  **Required if sending via number pool or with an alphanumeric sender ID.**  | [optional] [default to null]
**To** | [***OneOfbody13To**](OneOfbody13To.md) |  | [default to null]
**Text** | **string** | Message body (i.e., content) as a non-empty string.  **Required for SMS** | [optional] [default to null]
**Subject** | **string** | Subject of multimedia message | [optional] [default to null]
**MediaUrls** | **[]string** | A list of media URLs. The total media size must be less than 1 MB.  **Required for MMS** | [optional] [default to null]
**WebhookUrl** | **string** | The URL where webhooks related to this message will be sent. | [optional] [default to null]
**WebhookFailoverUrl** | **string** | The failover URL where webhooks related to this message will be sent if sending to the primary URL fails. | [optional] [default to null]
**UseProfileWebhooks** | **bool** | If the profile this number is associated with has webhooks, use them for delivery notifications. If webhooks are also specified on the message itself, they will be attempted first, then those on the profile. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

