# AnswerRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingGroupId** | **string** | Use this field to set the Billing Group ID for the call. Must be a valid and existing Billing Group ID. | [optional] [default to null]
**ClientState** | **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] [default to null]
**CommandId** | **string** | Use this field to avoid duplicate commands. Telnyx will ignore commands with the same &#x60;command_id&#x60;. | [optional] [default to null]
**WebhookUrl** | **string** | Use this field to override the URL for which Telnyx will send subsequent webhooks to for this call. | [optional] [default to null]
**WebhookUrlMethod** | **string** | HTTP request type used for &#x60;webhook_url&#x60;. | [optional] [default to WEBHOOK_URL_METHOD.POST]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

