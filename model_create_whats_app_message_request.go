/*
 * Telnyx API
 *
 * SIP trunking, SMS, MMS, Call Control and Telephony Data Services.
 *
 * API version: 2.0.0
 * Contact: support@telnyx.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package telnyx

type CreateWhatsAppMessageRequest struct {
	Audio *Audio `json:"audio,omitempty"`
	Contacts []Contact `json:"contacts,omitempty"`
	Document *Document `json:"document,omitempty"`
	Hsm *Hsm `json:"hsm,omitempty"`
	Image *Image `json:"image,omitempty"`
	Location *Location `json:"location,omitempty"`
	// Specifying preview_url in the request is optional when not including a URL in your message. To include a URL preview, set preview_url to true in the message body and make sure the URL begins with http:// or https://.
	PreviewUrl bool `json:"preview_url,omitempty"`
	Template *Template `json:"template,omitempty"`
	Text *Text `json:"text,omitempty"`
	// The WhatsApp ID (phone number) returned from contacts endpoint.
	To string `json:"to"`
	Type_ *MessageType `json:"type,omitempty"`
	Video *Video `json:"video,omitempty"`
	// The sender's WhatsApp ID.
	WhatsappUserId string `json:"whatsapp_user_id"`
}
