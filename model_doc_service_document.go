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

type DocServiceDocument struct {
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at,omitempty"`
	// Identifies the resource.
	Id string `json:"id,omitempty"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type,omitempty"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at,omitempty"`
	// The document's content_type.
	ContentType string `json:"content_type,omitempty"`
	// The filename of the document.
	Filename string `json:"filename,omitempty"`
	// The document's SHA256 hash provided for optional verification purposes.
	Sha256 string `json:"sha256,omitempty"`
	Size *DocServiceDocumentSize `json:"size,omitempty"`
}
