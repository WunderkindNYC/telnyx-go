# CreateWhatsAppMessageRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audio** | [***Audio**](Audio.md) |  | [optional] [default to null]
**Contacts** | [**[]Contact**](Contact.md) |  | [optional] [default to null]
**Document** | [***Document**](Document.md) |  | [optional] [default to null]
**Hsm** | [***Hsm**](Hsm.md) |  | [optional] [default to null]
**Image** | [***Image**](Image.md) |  | [optional] [default to null]
**Location** | [***Location**](Location.md) |  | [optional] [default to null]
**PreviewUrl** | **bool** | Specifying preview_url in the request is optional when not including a URL in your message. To include a URL preview, set preview_url to true in the message body and make sure the URL begins with http:// or https://. | [optional] [default to null]
**Template** | [***Template**](Template.md) |  | [optional] [default to null]
**Text** | [***Text**](Text.md) |  | [optional] [default to null]
**To** | **string** | The WhatsApp ID (phone number) returned from contacts endpoint. | [default to null]
**Type_** | [***MessageType**](MessageType.md) |  | [optional] [default to null]
**Video** | [***Video**](Video.md) |  | [optional] [default to null]
**WhatsappUserId** | **string** | The sender&#x27;s WhatsApp ID. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

