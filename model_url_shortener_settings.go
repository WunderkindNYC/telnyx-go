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

// The URL shortener feature allows automatic replacement of URLs that were generated using a public URL shortener service. Some examples include bit.do, bit.ly, goo.gl, ht.ly, is.gd, ow.ly, rebrand.ly, t.co, tiny.cc, and tinyurl.com. Such URLs are replaced with with links generated by Telnyx. The use of custom links can improve branding and message deliverability.  To disable this feature, set the object field to `null`. 
type UrlShortenerSettings struct {
	// One of the domains provided by the Telnyx URL shortener service. 
	Domain string `json:"domain"`
	// Optional prefix that can be used to identify your brand, and will appear in the Telnyx generated URLs after the domain name. 
	Prefix string `json:"prefix,omitempty"`
	// Use the link replacement tool only for links that are specifically blacklisted by Telnyx. 
	ReplaceBlacklistOnly bool `json:"replace_blacklist_only,omitempty"`
	// Receive webhooks for when your replaced links are clicked. Webhooks are sent to the webhooks on the messaging profile. 
	SendWebhooks bool `json:"send_webhooks,omitempty"`
}
