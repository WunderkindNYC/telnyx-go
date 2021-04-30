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

type AllOfinlineResponse200DataItems struct {
	// The current status of the SIM card. It will be one of the following: <br/> <ul>   <li><code>activating</code> - the card is being activated</li>   <li><code>active</code> - the card is active and ready for use</li>   <li><code>inactivating</code> - the card is being inactivated</li>   <li><code>inactive</code> - the card has been inactivated and cannot be used</li>   <li><code>data_limit_exceeded</code> - the card has exceeded its data consumption limit</li> </ul> Transitioning between the active and inactive states may take a period of time. 
	Status string `json:"status,omitempty"`
	// The ICCID is the identifier of the specific SIM card/chip. Each SIM is internationally identified by its integrated circuit card identifier (ICCID). ICCIDs are stored in the SIM card's memory and are also engraved or printed on the SIM card body during a process called personalization. 
	Iccid string `json:"iccid,omitempty"`
	// SIM cards are identified on their individual operator networks by a unique International Mobile Subscriber Identity (IMSI). <br/> Mobile network operators connect mobile phone calls and communicate with their market SIM cards using their IMSIs. The IMSI is stored in the Subscriber  Identity Module (SIM) inside the device and is sent by the device to the appropriate network. It is used to acquire the details of the device in the Home  Location Register (HLR) or the Visitor Location Register (VLR). 
	Imsi string `json:"imsi,omitempty"`
	// Mobile Station International Subscriber Directory Number (MSISDN) is a number used to identify a mobile phone number internationally. <br/> MSISDN is defined by the E.164 numbering plan. It includes a country code and a National Destination Code which identifies the subscriber's operator. 
	Msisdn string `json:"msisdn,omitempty"`
	// The group SIMCardGroup identification. This attribute can be <code>null</code> when it's present in an associated resource.
	SimCardGroupId string `json:"sim_card_group_id,omitempty"`
	// Searchable tags associated with the SIM card
	Tags []string `json:"tags,omitempty"`
}
