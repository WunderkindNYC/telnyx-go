# Go API client for telnyx

SIP trunking, SMS, MMS, Call Control and Telephony Data Services.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./telnyx"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.telnyx.com/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AddressesApi* | [**CreateAddress**](docs/AddressesApi.md#createaddress) | **Post** /addresses | Creates an address
*AddressesApi* | [**DeleteAddress**](docs/AddressesApi.md#deleteaddress) | **Delete** /addresses/{id} | Deletes an address
*AddressesApi* | [**FindAddresss**](docs/AddressesApi.md#findaddresss) | **Get** /addresses | List all addresses
*AddressesApi* | [**GetAddress**](docs/AddressesApi.md#getaddress) | **Get** /addresses/{id} | Retrieve an address
*BillingGroupsApi* | [**CreateBillingGroup**](docs/BillingGroupsApi.md#createbillinggroup) | **Post** /billing_groups | Create a billing group
*BillingGroupsApi* | [**DeleteBillingGroup**](docs/BillingGroupsApi.md#deletebillinggroup) | **Delete** /billing_groups/{id} | Delete a billing group
*BillingGroupsApi* | [**ListBillingGroups**](docs/BillingGroupsApi.md#listbillinggroups) | **Get** /billing_groups | List all billing groups
*BillingGroupsApi* | [**RetrieveBillingGroup**](docs/BillingGroupsApi.md#retrievebillinggroup) | **Get** /billing_groups/{id} | Retrieve a billing group
*BillingGroupsApi* | [**UpdateBillingGroup**](docs/BillingGroupsApi.md#updatebillinggroup) | **Patch** /billing_groups/{id} | Update a billing group
*CSVDownloadsApi* | [**CreateCsvDownload**](docs/CSVDownloadsApi.md#createcsvdownload) | **Post** /phone_numbers/csv_downloads | create a new CSV download request
*CSVDownloadsApi* | [**FindCsvDownloads**](docs/CSVDownloadsApi.md#findcsvdownloads) | **Get** /phone_numbers/csv_downloads | List your submitted CSV download requests
*CSVDownloadsApi* | [**RetrieveCsvDownload**](docs/CSVDownloadsApi.md#retrievecsvdownload) | **Get** /phone_numbers/csv_downloads/{id} | Get a single submitted CSV download request.
*CallCommandsApi* | [**CallControlAnswer**](docs/CallCommandsApi.md#callcontrolanswer) | **Post** /calls/{call_control_id}/actions/answer | Answer call
*CallCommandsApi* | [**CallControlBridge**](docs/CallCommandsApi.md#callcontrolbridge) | **Post** /calls/{call_control_id}/actions/bridge | Bridge calls
*CallCommandsApi* | [**CallControlDial**](docs/CallCommandsApi.md#callcontroldial) | **Post** /calls | Dial
*CallCommandsApi* | [**CallControlForkStart**](docs/CallCommandsApi.md#callcontrolforkstart) | **Post** /calls/{call_control_id}/actions/fork_start | Forking start
*CallCommandsApi* | [**CallControlForkStop**](docs/CallCommandsApi.md#callcontrolforkstop) | **Post** /calls/{call_control_id}/actions/fork_stop | Forking stop
*CallCommandsApi* | [**CallControlGatherUsingAudio**](docs/CallCommandsApi.md#callcontrolgatherusingaudio) | **Post** /calls/{call_control_id}/actions/gather_using_audio | Gather using audio
*CallCommandsApi* | [**CallControlGatherUsingSpeak**](docs/CallCommandsApi.md#callcontrolgatherusingspeak) | **Post** /calls/{call_control_id}/actions/gather_using_speak | Gather using speak
*CallCommandsApi* | [**CallControlHangup**](docs/CallCommandsApi.md#callcontrolhangup) | **Post** /calls/{call_control_id}/actions/hangup | Hangup call
*CallCommandsApi* | [**CallControlPlaybackStart**](docs/CallCommandsApi.md#callcontrolplaybackstart) | **Post** /calls/{call_control_id}/actions/playback_start | Play audio URL
*CallCommandsApi* | [**CallControlPlaybackStop**](docs/CallCommandsApi.md#callcontrolplaybackstop) | **Post** /calls/{call_control_id}/actions/playback_stop | Stop audio playback
*CallCommandsApi* | [**CallControlRecordStart**](docs/CallCommandsApi.md#callcontrolrecordstart) | **Post** /calls/{call_control_id}/actions/record_start | Recording start
*CallCommandsApi* | [**CallControlRecordStop**](docs/CallCommandsApi.md#callcontrolrecordstop) | **Post** /calls/{call_control_id}/actions/record_stop | Recording stop
*CallCommandsApi* | [**CallControlReject**](docs/CallCommandsApi.md#callcontrolreject) | **Post** /calls/{call_control_id}/actions/reject | Reject call
*CallCommandsApi* | [**CallControlSendDTMF**](docs/CallCommandsApi.md#callcontrolsenddtmf) | **Post** /calls/{call_control_id}/actions/send_dtmf | Send DTMF
*CallCommandsApi* | [**CallControlSpeak**](docs/CallCommandsApi.md#callcontrolspeak) | **Post** /calls/{call_control_id}/actions/speak | Speak text
*CallCommandsApi* | [**CallControlTransfer**](docs/CallCommandsApi.md#callcontroltransfer) | **Post** /calls/{call_control_id}/actions/transfer | Transfer call
*CallControlApplicationsApi* | [**CreateCallControlApplication**](docs/CallControlApplicationsApi.md#createcallcontrolapplication) | **Post** /call_control_applications | Creates a Call Control application
*CallControlApplicationsApi* | [**DeleteCallControlApplication**](docs/CallControlApplicationsApi.md#deletecallcontrolapplication) | **Delete** /call_control_applications/{id} | Deletes a Call Control application
*CallControlApplicationsApi* | [**FindCallControlApplications**](docs/CallControlApplicationsApi.md#findcallcontrolapplications) | **Get** /call_control_applications | List all Call Control applications
*CallControlApplicationsApi* | [**GetCallControlApplication**](docs/CallControlApplicationsApi.md#getcallcontrolapplication) | **Get** /call_control_applications/{id} | Retrieve a Call Control application
*CallControlApplicationsApi* | [**UpdateCallControlApplication**](docs/CallControlApplicationsApi.md#updatecallcontrolapplication) | **Patch** /call_control_applications/{id} | Update a Call Control application
*CallInformationApi* | [**RetrieveCallStatus**](docs/CallInformationApi.md#retrievecallstatus) | **Get** /calls/{call_control_id} | Retrieve call status
*ConferenceCommandsApi* | [**CreateConference**](docs/ConferenceCommandsApi.md#createconference) | **Post** /conferences | Create conference
*ConferenceCommandsApi* | [**GetConferences**](docs/ConferenceCommandsApi.md#getconferences) | **Get** /conferences | List conferences
*ConferenceCommandsApi* | [**HoldConference**](docs/ConferenceCommandsApi.md#holdconference) | **Post** /conferences/{id}/actions/hold | Hold conference participants
*ConferenceCommandsApi* | [**JoinConference**](docs/ConferenceCommandsApi.md#joinconference) | **Post** /conferences/{id}/actions/join | Join a conference
*ConferenceCommandsApi* | [**MuteConference**](docs/ConferenceCommandsApi.md#muteconference) | **Post** /conferences/{id}/actions/mute | Mute conference participants
*ConferenceCommandsApi* | [**UnholdConference**](docs/ConferenceCommandsApi.md#unholdconference) | **Post** /conferences/{id}/actions/unhold | Unhold conference participants
*ConferenceCommandsApi* | [**UnmuteConference**](docs/ConferenceCommandsApi.md#unmuteconference) | **Post** /conferences/{id}/actions/unmute | Unmute conference participants
*ConnectionsApi* | [**FindAllConnections**](docs/ConnectionsApi.md#findallconnections) | **Get** /connections | List all connections
*ConnectionsApi* | [**GetConnection**](docs/ConnectionsApi.md#getconnection) | **Get** /connections/{id} | Retrieve a connection
*CredentialConnectionsApi* | [**CreateCredentialConnection**](docs/CredentialConnectionsApi.md#createcredentialconnection) | **Post** /credential_connections | Creates a credential connection
*CredentialConnectionsApi* | [**DeleteCredentialConnection**](docs/CredentialConnectionsApi.md#deletecredentialconnection) | **Delete** /credential_connections/{id} | Deletes a credential connection
*CredentialConnectionsApi* | [**FindCredentialConnections**](docs/CredentialConnectionsApi.md#findcredentialconnections) | **Get** /credential_connections | List all credential connections
*CredentialConnectionsApi* | [**GetCredentialConnection**](docs/CredentialConnectionsApi.md#getcredentialconnection) | **Get** /credential_connections/{id} | Retrieve a connection
*CredentialConnectionsApi* | [**UpdateCredentialConnection**](docs/CredentialConnectionsApi.md#updatecredentialconnection) | **Patch** /credential_connections/{id} | Updates a credential connection
*DebuggingApi* | [**CallControlDebuggingEventList**](docs/DebuggingApi.md#callcontroldebuggingeventlist) | **Get** /call_events | List call events
*FQDNConnectionsApi* | [**CreateFQDNConnection**](docs/FQDNConnectionsApi.md#createfqdnconnection) | **Post** /fqdn_connections | Creates an FQDN connection
*FQDNConnectionsApi* | [**DeleteFQDNConnection**](docs/FQDNConnectionsApi.md#deletefqdnconnection) | **Delete** /fqdn_connections/{id} | Deletes an FQDN connection
*FQDNConnectionsApi* | [**FindFQDNConnections**](docs/FQDNConnectionsApi.md#findfqdnconnections) | **Get** /fqdn_connections | List all FQDN connections
*FQDNConnectionsApi* | [**GetFQDNConnection**](docs/FQDNConnectionsApi.md#getfqdnconnection) | **Get** /fqdn_connections/{id} | Retrieve a connection
*FQDNConnectionsApi* | [**UpdateFQDNConnection**](docs/FQDNConnectionsApi.md#updatefqdnconnection) | **Patch** /fqdn_connections/{id} | Update a FQDN connection
*FQDNsApi* | [**AddFQDN**](docs/FQDNsApi.md#addfqdn) | **Post** /fqdns | Create an FQDN
*FQDNsApi* | [**DeleteFQDN**](docs/FQDNsApi.md#deletefqdn) | **Delete** /fqdns/{id} | Delete an FQDN
*FQDNsApi* | [**FQDNsGet**](docs/FQDNsApi.md#fqdnsget) | **Get** /fqdns | Get all FQDNs
*FQDNsApi* | [**GetFQDNDetails**](docs/FQDNsApi.md#getfqdndetails) | **Get** /fqdns/{id} | Get FQDN
*FQDNsApi* | [**UpdateFQDN**](docs/FQDNsApi.md#updatefqdn) | **Patch** /fqdns/{id} | Update FQDN
*IPConnectionsApi* | [**CreateIPConnection**](docs/IPConnectionsApi.md#createipconnection) | **Post** /ip_connections | Creates an IP connection
*IPConnectionsApi* | [**DeleteIPConnection**](docs/IPConnectionsApi.md#deleteipconnection) | **Delete** /ip_connections/{id} | Deletes an IP connection
*IPConnectionsApi* | [**FindConnections**](docs/IPConnectionsApi.md#findconnections) | **Get** /ip_connections | List all connections
*IPConnectionsApi* | [**GetIPConnection**](docs/IPConnectionsApi.md#getipconnection) | **Get** /ip_connections/{id} | Retrieve a connection
*IPConnectionsApi* | [**UpdateIPConnection**](docs/IPConnectionsApi.md#updateipconnection) | **Patch** /ip_connections/{id} | Updates an IP connection
*IPsApi* | [**AddIP**](docs/IPsApi.md#addip) | **Post** /ips | Create an IP
*IPsApi* | [**DeleteIP**](docs/IPsApi.md#deleteip) | **Delete** /ips/{id} | Delete an IP
*IPsApi* | [**GetIPDetails**](docs/IPsApi.md#getipdetails) | **Get** /ips/{id} | Get IP
*IPsApi* | [**IPsGet**](docs/IPsApi.md#ipsget) | **Get** /ips | Get all IPs
*IPsApi* | [**UpdateIP**](docs/IPsApi.md#updateip) | **Patch** /ips/{id} | Update IP
*InboundChannelsApi* | [**ListOutboundChannels**](docs/InboundChannelsApi.md#listoutboundchannels) | **Get** /phone_numbers/inbound_channels | Retrieve your inbound channels
*InboundChannelsApi* | [**UpdateOutboundChannels**](docs/InboundChannelsApi.md#updateoutboundchannels) | **Patch** /phone_numbers/inbound_channels | Update inbound channels
*MessagesApi* | [**CreateLongCodeMessage**](docs/MessagesApi.md#createlongcodemessage) | **Post** /messages/long_code | Send a Long Code message
*MessagesApi* | [**CreateMessage**](docs/MessagesApi.md#createmessage) | **Post** /messages | Send a message
*MessagesApi* | [**CreateNumberPoolMessage**](docs/MessagesApi.md#createnumberpoolmessage) | **Post** /messages/number_pool | Send a message using Number Pool
*MessagesApi* | [**CreateShortCodeMessage**](docs/MessagesApi.md#createshortcodemessage) | **Post** /messages/short_code | Send a Short Code message
*MessagesApi* | [**RetrieveMessage**](docs/MessagesApi.md#retrievemessage) | **Get** /messages/{id} | Retrieve a message
*MessagingHostedNumbersApi* | [**DeleteMessagingHostedNumber**](docs/MessagingHostedNumbersApi.md#deletemessaginghostednumber) | **Delete** /messaging_hosted_numbers/{id} | Delete Messaging Hosted Number
*MessagingHostedNumbersApi* | [**GetMessagingHostedNumberOrder**](docs/MessagingHostedNumbersApi.md#getmessaginghostednumberorder) | **Get** /messaging_hosted_number_orders/{id} | Get Messaging Hosted Numbers Order Information
*MessagingHostedNumbersApi* | [**ListMessagingHostedNumberOrder**](docs/MessagingHostedNumbersApi.md#listmessaginghostednumberorder) | **Get** /messaging_hosted_number_orders | List All Messaging Hosted Number Orders
*MessagingHostedNumbersApi* | [**NewMessagingHostedNumberOrder**](docs/MessagingHostedNumbersApi.md#newmessaginghostednumberorder) | **Post** /messaging_hosted_number_orders | New Messaging Hosted Numbers Order
*MessagingHostedNumbersApi* | [**UploadFilesMessagingHostedNumberOrder**](docs/MessagingHostedNumbersApi.md#uploadfilesmessaginghostednumberorder) | **Post** /messaging_hosted_number_orders/{id}/actions/file_upload | Upload LOA and Bill required for a Messaging Hosted Number Order
*MessagingProfilesApi* | [**CreateMessagingProfile**](docs/MessagingProfilesApi.md#createmessagingprofile) | **Post** /messaging_profiles | Create a messaging profile
*MessagingProfilesApi* | [**DeleteMessagingProfile**](docs/MessagingProfilesApi.md#deletemessagingprofile) | **Delete** /messaging_profiles/{id} | Delete a messaging profile
*MessagingProfilesApi* | [**ListMessagingProfilePhoneNumbers**](docs/MessagingProfilesApi.md#listmessagingprofilephonenumbers) | **Get** /messaging_profiles/{id}/phone_numbers | List all phone numbers associated with a messaging profile
*MessagingProfilesApi* | [**ListMessagingProfileShortCodes**](docs/MessagingProfilesApi.md#listmessagingprofileshortcodes) | **Get** /messaging_profiles/{id}/short_codes | List all short codes associated with a messaging profile
*MessagingProfilesApi* | [**ListMessagingProfiles**](docs/MessagingProfilesApi.md#listmessagingprofiles) | **Get** /messaging_profiles | List all messaging profiles
*MessagingProfilesApi* | [**RetrieveMessagingProfile**](docs/MessagingProfilesApi.md#retrievemessagingprofile) | **Get** /messaging_profiles/{id} | Retrieve a messaging profile
*MessagingProfilesApi* | [**UpdateMessagingProfile**](docs/MessagingProfilesApi.md#updatemessagingprofile) | **Patch** /messaging_profiles/{id} | Update a messaging profile
*MessagingURLDomainsApi* | [**GetAllMessagingUrlDomains**](docs/MessagingURLDomainsApi.md#getallmessagingurldomains) | **Get** /messaging_url_domains | List all available messaging URL domains
*NumberConfigurationsApi* | [**DeletePhoneNumber**](docs/NumberConfigurationsApi.md#deletephonenumber) | **Delete** /phone_numbers/{id} | Delete a phone number
*NumberConfigurationsApi* | [**EnableEmergencyPhoneNumber**](docs/NumberConfigurationsApi.md#enableemergencyphonenumber) | **Post** /phone_numbers/{id}/actions/enable_emergency | Enable emergency for a phone number
*NumberConfigurationsApi* | [**FindPhoneNumberVoices**](docs/NumberConfigurationsApi.md#findphonenumbervoices) | **Get** /phone_numbers/voice | List voice settings for multiple phone numbers
*NumberConfigurationsApi* | [**FindPhoneNumbers**](docs/NumberConfigurationsApi.md#findphonenumbers) | **Get** /phone_numbers | List all phone numbers
*NumberConfigurationsApi* | [**GetPhoneNumber**](docs/NumberConfigurationsApi.md#getphonenumber) | **Get** /phone_numbers/{id} | Retrieve the settings for a phone number
*NumberConfigurationsApi* | [**ListPhoneNumberMessagingSettings**](docs/NumberConfigurationsApi.md#listphonenumbermessagingsettings) | **Get** /phone_numbers/messaging | List all phone numbers&#x27; messaging settings
*NumberConfigurationsApi* | [**RetrievePhoneNumberMessagingSettings**](docs/NumberConfigurationsApi.md#retrievephonenumbermessagingsettings) | **Get** /phone_numbers/{id}/messaging | Retrieve the messaging settings for a phone number
*NumberConfigurationsApi* | [**RetrievePhoneNumberVoice**](docs/NumberConfigurationsApi.md#retrievephonenumbervoice) | **Get** /phone_numbers/{id}/voice | Retrieve the voice settings for a phone number
*NumberConfigurationsApi* | [**UpdatePhoneNumber**](docs/NumberConfigurationsApi.md#updatephonenumber) | **Patch** /phone_numbers/{id} | Update the settings for a phone number
*NumberConfigurationsApi* | [**UpdatePhoneNumberMessagingSettings**](docs/NumberConfigurationsApi.md#updatephonenumbermessagingsettings) | **Patch** /phone_numbers/{id}/messaging | Update the messaging settings for a phone number
*NumberConfigurationsApi* | [**UpdatePhoneNumberVoice**](docs/NumberConfigurationsApi.md#updatephonenumbervoice) | **Patch** /phone_numbers/{id}/voice | Update the voice settings for a phone number
*NumberOrderDocumentsApi* | [**CreateNumberOrderDocument**](docs/NumberOrderDocumentsApi.md#createnumberorderdocument) | **Post** /number_order_documents | Upload Number Order Document
*NumberOrderDocumentsApi* | [**ListNumberOrderDocuments**](docs/NumberOrderDocumentsApi.md#listnumberorderdocuments) | **Get** /number_order_documents | Get Uploaded Number Order Documents
*NumberOrderDocumentsApi* | [**RetrieveNumberOrderDocument**](docs/NumberOrderDocumentsApi.md#retrievenumberorderdocument) | **Get** /number_order_documents/{number_order_document_id} | Retrieve a Single Number Order Document
*NumberOrderDocumentsApi* | [**UpdateNumberOrderDocument**](docs/NumberOrderDocumentsApi.md#updatenumberorderdocument) | **Patch** /number_order_documents/{number_order_document_id} | Update Number Order Document
*NumberOrderRegulatoryRequirementsApi* | [**ListNumberOrderRegulatoryRequirements**](docs/NumberOrderRegulatoryRequirementsApi.md#listnumberorderregulatoryrequirements) | **Get** /regulatory_requirements | Get list of Number Order Regulatory Requirements
*NumberOrderRegulatoryRequirementsApi* | [**ListPhoneNumberRegulatoryRequirements**](docs/NumberOrderRegulatoryRequirementsApi.md#listphonenumberregulatoryrequirements) | **Get** /phone_number_regulatory_requirements | Get Regulatory Requirements Per Number
*NumberOrderRegulatoryRequirementsApi* | [**RetrieveNumberOrderRegulatoryRequirement**](docs/NumberOrderRegulatoryRequirementsApi.md#retrievenumberorderregulatoryrequirement) | **Get** /regulatory_requirements/{requirement_id} | Get Detailed Number Order Regulatory Requirement
*NumberOrdersApi* | [**CreateNumberOrder**](docs/NumberOrdersApi.md#createnumberorder) | **Post** /number_orders | Create Phone Number Order
*NumberOrdersApi* | [**ListNumberOrders**](docs/NumberOrdersApi.md#listnumberorders) | **Get** /number_orders | Retrieve multiple Number Orders
*NumberOrdersApi* | [**RetrieveNumberOrder**](docs/NumberOrdersApi.md#retrievenumberorder) | **Get** /number_orders/{number_order_id} | Retrieve a single phone number order
*NumberOrdersApi* | [**UpdateNumberOrder**](docs/NumberOrdersApi.md#updatenumberorder) | **Patch** /number_orders/{number_order_id} | Update phone number order
*NumberPortoutApi* | [**FindPortoutComments**](docs/NumberPortoutApi.md#findportoutcomments) | **Get** /portouts/{id}/comments | List all comments for a portout request
*NumberPortoutApi* | [**FindPortoutRequest**](docs/NumberPortoutApi.md#findportoutrequest) | **Get** /portouts/{id} | Retrieve a portout request
*NumberPortoutApi* | [**ListPortoutRequest**](docs/NumberPortoutApi.md#listportoutrequest) | **Get** /portouts | Retrieve a list of portout requests
*NumberPortoutApi* | [**PostPortRequestComment**](docs/NumberPortoutApi.md#postportrequestcomment) | **Post** /portouts/{id}/comments | Create a comment on a portout request
*NumberPortoutApi* | [**UpdatePortoutRequest**](docs/NumberPortoutApi.md#updateportoutrequest) | **Patch** /portouts/{id}/{status} | Update Status
*NumberReservationsApi* | [**CreateNumberReservations**](docs/NumberReservationsApi.md#createnumberreservations) | **Post** /number_reservations | Create a Phone Number Reservation
*NumberReservationsApi* | [**ExtendNumberReservationExpiryTime**](docs/NumberReservationsApi.md#extendnumberreservationexpirytime) | **Post** /number_reservations/{number_reservation_id}/actions/extend | Extend a Phone Number Reservation
*NumberReservationsApi* | [**ListNumberReservations**](docs/NumberReservationsApi.md#listnumberreservations) | **Get** /number_reservations | Retrieve multiple Number Reservations
*NumberReservationsApi* | [**RetrieveNumberReservation**](docs/NumberReservationsApi.md#retrievenumberreservation) | **Get** /number_reservations/{number_reservation_id} | Retrieve a Single Phone Number Reservation
*NumberSearchApi* | [**ListAvailablePhoneNumbers**](docs/NumberSearchApi.md#listavailablephonenumbers) | **Get** /available_phone_numbers | List available phone numbers
*OutboundVoiceProfilesApi* | [**CreateOutboundVoiceProfile**](docs/OutboundVoiceProfilesApi.md#createoutboundvoiceprofile) | **Post** /outbound_voice_profiles | Create an outbound voice profile
*OutboundVoiceProfilesApi* | [**DeleteOutboundVoiceProfile**](docs/OutboundVoiceProfilesApi.md#deleteoutboundvoiceprofile) | **Delete** /outbound_voice_profiles/{id} | Delete an outbound voice profile
*OutboundVoiceProfilesApi* | [**GetOutboundVoiceProfile**](docs/OutboundVoiceProfilesApi.md#getoutboundvoiceprofile) | **Get** /outbound_voice_profiles/{id} | Retrieve an outbound voice profile
*OutboundVoiceProfilesApi* | [**GetOutboundVoiceProfiles**](docs/OutboundVoiceProfilesApi.md#getoutboundvoiceprofiles) | **Get** /outbound_voice_profiles | Get all outbound voice profiles
*OutboundVoiceProfilesApi* | [**UpdateOutboundVoiceProfile**](docs/OutboundVoiceProfilesApi.md#updateoutboundvoiceprofile) | **Patch** /outbound_voice_profiles/{id} | Updates an existing outbound voice profile.
*PhoneNumbersApi* | [**ListMessagingPhoneNumbers**](docs/PhoneNumbersApi.md#listmessagingphonenumbers) | **Get** /messaging_phone_numbers | List all messaging capable phone numbers
*PhoneNumbersApi* | [**RetrieveMessagingPhoneNumber**](docs/PhoneNumbersApi.md#retrievemessagingphonenumber) | **Get** /messaging_phone_numbers/{id} | Retrieve a messaging phone number
*PhoneNumbersApi* | [**UpdateMessagingPhoneNumber**](docs/PhoneNumbersApi.md#updatemessagingphonenumber) | **Patch** /messaging_phone_numbers/{id} | Update a messaging phone number
*ReportingApi* | [**CreateWdrReport**](docs/ReportingApi.md#createwdrreport) | **Post** /wireless/detail_records_reports | Create a Wireless Detail Records (WDRs) Report
*ReportingApi* | [**DeleteWdrReport**](docs/ReportingApi.md#deletewdrreport) | **Delete** /wireless/detail_records_reports/{id} | Delete a Wireless Detail Record (WDR) Report
*ReportingApi* | [**GetWdrReport**](docs/ReportingApi.md#getwdrreport) | **Get** /wireless/detail_records_reports/{id} | Get a Wireless Detail Record (WDR) Report
*ReportingApi* | [**GetWdrReports**](docs/ReportingApi.md#getwdrreports) | **Get** /wireless/detail_records_reports | Get all Wireless Detail Records (WDRs) Reports
*ReportsApi* | [**CreateLedgerBillingGroupReport**](docs/ReportsApi.md#createledgerbillinggroupreport) | **Post** /ledger_billing_group_reports | Create a ledger billing group report
*ReportsApi* | [**RetrieveLedgerBillingGroupReport**](docs/ReportsApi.md#retrieveledgerbillinggroupreport) | **Get** /ledger_billing_group_reports/{id} | Retrieve a ledger billing group report
*SIMCardGroupsApi* | [**SimCardGroupDelete**](docs/SIMCardGroupsApi.md#simcardgroupdelete) | **Delete** /sim_card_groups/{id} | Delete a SIM card group
*SIMCardGroupsApi* | [**SimCardGroupUpdate**](docs/SIMCardGroupsApi.md#simcardgroupupdate) | **Patch** /sim_card_groups/{id} | Update a SIM card group
*SIMCardGroupsApi* | [**SimCardGroupsGet**](docs/SIMCardGroupsApi.md#simcardgroupsget) | **Get** /sim_card_groups/{id} | Get SIM card gruop
*SIMCardGroupsApi* | [**SimCardGroupsGetAll**](docs/SIMCardGroupsApi.md#simcardgroupsgetall) | **Get** /sim_card_groups | Get all SIM card groups
*SIMCardGroupsApi* | [**SimCardGroupsPost**](docs/SIMCardGroupsApi.md#simcardgroupspost) | **Post** /sim_card_groups | Create a SIM card group
*SIMCardsApi* | [**SimCardActivate**](docs/SIMCardsApi.md#simcardactivate) | **Post** /sim_cards/{id}/actions/activate | Request a SIM card activation
*SIMCardsApi* | [**SimCardDeactivate**](docs/SIMCardsApi.md#simcarddeactivate) | **Post** /sim_cards/{id}/actions/deactivate | Request a SIM card deactivation
*SIMCardsApi* | [**SimCardGet**](docs/SIMCardsApi.md#simcardget) | **Get** /sim_cards/{id} | Get SIM card
*SIMCardsApi* | [**SimCardRegister**](docs/SIMCardsApi.md#simcardregister) | **Post** /actions/register/sim_cards | Register SIM cards
*SIMCardsApi* | [**SimCardUpdate**](docs/SIMCardsApi.md#simcardupdate) | **Patch** /sim_cards/{id} | Update a SIM card
*SIMCardsApi* | [**SimCardsGet**](docs/SIMCardsApi.md#simcardsget) | **Get** /sim_cards | Get all SIM cards
*ShortCodesApi* | [**ListShortCodes**](docs/ShortCodesApi.md#listshortcodes) | **Get** /short_codes | List all short codes
*ShortCodesApi* | [**RetrieveShortCode**](docs/ShortCodesApi.md#retrieveshortcode) | **Get** /short_codes/{id} | Retrieve a short code
*ShortCodesApi* | [**UpdateShortCode**](docs/ShortCodesApi.md#updateshortcode) | **Patch** /short_codes/{id} | Update short code

## Documentation For Models

 - [Address](docs/Address.md)
 - [AllOfinlineResponse20055DataItems](docs/AllOfinlineResponse20055DataItems.md)
 - [AllOfinlineResponse20056DataItems](docs/AllOfinlineResponse20056DataItems.md)
 - [AllOfinlineResponse20057Data](docs/AllOfinlineResponse20057Data.md)
 - [AllOfinlineResponse20058DataItems](docs/AllOfinlineResponse20058DataItems.md)
 - [AllOfinlineResponse200DataItems](docs/AllOfinlineResponse200DataItems.md)
 - [AllOfinlineResponse2017Data](docs/AllOfinlineResponse2017Data.md)
 - [AllOfinlineResponse202Data](docs/AllOfinlineResponse202Data.md)
 - [AnswerRequest](docs/AnswerRequest.md)
 - [Body](docs/Body.md)
 - [Body1](docs/Body1.md)
 - [Body10](docs/Body10.md)
 - [Body11](docs/Body11.md)
 - [Body12](docs/Body12.md)
 - [Body13](docs/Body13.md)
 - [Body14](docs/Body14.md)
 - [Body15](docs/Body15.md)
 - [Body16](docs/Body16.md)
 - [Body17](docs/Body17.md)
 - [Body18](docs/Body18.md)
 - [Body19](docs/Body19.md)
 - [Body2](docs/Body2.md)
 - [Body20](docs/Body20.md)
 - [Body21](docs/Body21.md)
 - [Body22](docs/Body22.md)
 - [Body23](docs/Body23.md)
 - [Body24](docs/Body24.md)
 - [Body25](docs/Body25.md)
 - [Body26](docs/Body26.md)
 - [Body27](docs/Body27.md)
 - [Body28](docs/Body28.md)
 - [Body29](docs/Body29.md)
 - [Body3](docs/Body3.md)
 - [Body30](docs/Body30.md)
 - [Body31](docs/Body31.md)
 - [Body32](docs/Body32.md)
 - [Body33](docs/Body33.md)
 - [Body34](docs/Body34.md)
 - [Body35](docs/Body35.md)
 - [Body36](docs/Body36.md)
 - [Body37](docs/Body37.md)
 - [Body4](docs/Body4.md)
 - [Body5](docs/Body5.md)
 - [Body6](docs/Body6.md)
 - [Body7](docs/Body7.md)
 - [Body8](docs/Body8.md)
 - [Body9](docs/Body9.md)
 - [BridgeRequest](docs/BridgeRequest.md)
 - [CallControlApplication](docs/CallControlApplication.md)
 - [CallControlApplicationsInbound](docs/CallControlApplicationsInbound.md)
 - [CallControlApplicationsOutbound](docs/CallControlApplicationsOutbound.md)
 - [CallsAnsweringMachineDetectionConfig](docs/CallsAnsweringMachineDetectionConfig.md)
 - [CommandResult](docs/CommandResult.md)
 - [Connection](docs/Connection.md)
 - [CreateConferenceRequest](docs/CreateConferenceRequest.md)
 - [CredentialConnection](docs/CredentialConnection.md)
 - [CredentialConnectionCreate](docs/CredentialConnectionCreate.md)
 - [CredentialConnectionUpdate](docs/CredentialConnectionUpdate.md)
 - [CredentialConnectionsRtcpSettings](docs/CredentialConnectionsRtcpSettings.md)
 - [CustomSipHeader](docs/CustomSipHeader.md)
 - [DialRequest](docs/DialRequest.md)
 - [Fqdn](docs/Fqdn.md)
 - [FqdnConnection](docs/FqdnConnection.md)
 - [FqdnConnectionsRtcpSettings](docs/FqdnConnectionsRtcpSettings.md)
 - [GatherUsingAudioWithAllParametersRequest](docs/GatherUsingAudioWithAllParametersRequest.md)
 - [GatherUsingSpeakWithAllParametersRequest](docs/GatherUsingSpeakWithAllParametersRequest.md)
 - [HangupRequest](docs/HangupRequest.md)
 - [HoldConferenceParticipantsRequest](docs/HoldConferenceParticipantsRequest.md)
 - [InboundConfiguration](docs/InboundConfiguration.md)
 - [InboundConfiguration1](docs/InboundConfiguration1.md)
 - [InboundConfiguration2](docs/InboundConfiguration2.md)
 - [InboundConfiguration3](docs/InboundConfiguration3.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse20010](docs/InlineResponse20010.md)
 - [InlineResponse20010Meta](docs/InlineResponse20010Meta.md)
 - [InlineResponse20011](docs/InlineResponse20011.md)
 - [InlineResponse20011Data](docs/InlineResponse20011Data.md)
 - [InlineResponse20012](docs/InlineResponse20012.md)
 - [InlineResponse20013](docs/InlineResponse20013.md)
 - [InlineResponse20014](docs/InlineResponse20014.md)
 - [InlineResponse20015](docs/InlineResponse20015.md)
 - [InlineResponse20016](docs/InlineResponse20016.md)
 - [InlineResponse20017](docs/InlineResponse20017.md)
 - [InlineResponse20018](docs/InlineResponse20018.md)
 - [InlineResponse20019](docs/InlineResponse20019.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse20020](docs/InlineResponse20020.md)
 - [InlineResponse20020Data](docs/InlineResponse20020Data.md)
 - [InlineResponse20021](docs/InlineResponse20021.md)
 - [InlineResponse20021Data](docs/InlineResponse20021Data.md)
 - [InlineResponse20021DataCost](docs/InlineResponse20021DataCost.md)
 - [InlineResponse20021DataErrors](docs/InlineResponse20021DataErrors.md)
 - [InlineResponse20021DataMedia](docs/InlineResponse20021DataMedia.md)
 - [InlineResponse20021DataSource](docs/InlineResponse20021DataSource.md)
 - [InlineResponse20021DataTo](docs/InlineResponse20021DataTo.md)
 - [InlineResponse20022](docs/InlineResponse20022.md)
 - [InlineResponse20022Data](docs/InlineResponse20022Data.md)
 - [InlineResponse20022PhoneNumbers](docs/InlineResponse20022PhoneNumbers.md)
 - [InlineResponse20023](docs/InlineResponse20023.md)
 - [InlineResponse20024](docs/InlineResponse20024.md)
 - [InlineResponse20025](docs/InlineResponse20025.md)
 - [InlineResponse20026](docs/InlineResponse20026.md)
 - [InlineResponse20026Data](docs/InlineResponse20026Data.md)
 - [InlineResponse20026DataFeatures](docs/InlineResponse20026DataFeatures.md)
 - [InlineResponse20026DataFeaturesSms](docs/InlineResponse20026DataFeaturesSms.md)
 - [InlineResponse20026DataHealth](docs/InlineResponse20026DataHealth.md)
 - [InlineResponse20027](docs/InlineResponse20027.md)
 - [InlineResponse20028](docs/InlineResponse20028.md)
 - [InlineResponse20028Data](docs/InlineResponse20028Data.md)
 - [InlineResponse20029](docs/InlineResponse20029.md)
 - [InlineResponse20029Data](docs/InlineResponse20029Data.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse20030](docs/InlineResponse20030.md)
 - [InlineResponse20030Data](docs/InlineResponse20030Data.md)
 - [InlineResponse20031](docs/InlineResponse20031.md)
 - [InlineResponse20032](docs/InlineResponse20032.md)
 - [InlineResponse20032Data](docs/InlineResponse20032Data.md)
 - [InlineResponse20033](docs/InlineResponse20033.md)
 - [InlineResponse20033Data](docs/InlineResponse20033Data.md)
 - [InlineResponse20033Meta](docs/InlineResponse20033Meta.md)
 - [InlineResponse20034](docs/InlineResponse20034.md)
 - [InlineResponse20034Data](docs/InlineResponse20034Data.md)
 - [InlineResponse20035](docs/InlineResponse20035.md)
 - [InlineResponse20035Data](docs/InlineResponse20035Data.md)
 - [InlineResponse20036](docs/InlineResponse20036.md)
 - [InlineResponse20037](docs/InlineResponse20037.md)
 - [InlineResponse20037CallRecording](docs/InlineResponse20037CallRecording.md)
 - [InlineResponse20038](docs/InlineResponse20038.md)
 - [InlineResponse20039](docs/InlineResponse20039.md)
 - [InlineResponse20039Data](docs/InlineResponse20039Data.md)
 - [InlineResponse20039RegulatoryRequirements](docs/InlineResponse20039RegulatoryRequirements.md)
 - [InlineResponse2003CostInformation](docs/InlineResponse2003CostInformation.md)
 - [InlineResponse2003Data](docs/InlineResponse2003Data.md)
 - [InlineResponse2003Meta](docs/InlineResponse2003Meta.md)
 - [InlineResponse2003RegionInformation](docs/InlineResponse2003RegionInformation.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse20040](docs/InlineResponse20040.md)
 - [InlineResponse20041](docs/InlineResponse20041.md)
 - [InlineResponse20041Data](docs/InlineResponse20041Data.md)
 - [InlineResponse20042](docs/InlineResponse20042.md)
 - [InlineResponse20042Data](docs/InlineResponse20042Data.md)
 - [InlineResponse20043](docs/InlineResponse20043.md)
 - [InlineResponse20043Data](docs/InlineResponse20043Data.md)
 - [InlineResponse20044](docs/InlineResponse20044.md)
 - [InlineResponse20044Data](docs/InlineResponse20044Data.md)
 - [InlineResponse20045](docs/InlineResponse20045.md)
 - [InlineResponse20045Data](docs/InlineResponse20045Data.md)
 - [InlineResponse20045Emergency](docs/InlineResponse20045Emergency.md)
 - [InlineResponse20046](docs/InlineResponse20046.md)
 - [InlineResponse20046Data](docs/InlineResponse20046Data.md)
 - [InlineResponse20047](docs/InlineResponse20047.md)
 - [InlineResponse20048](docs/InlineResponse20048.md)
 - [InlineResponse20049](docs/InlineResponse20049.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse20050](docs/InlineResponse20050.md)
 - [InlineResponse20050Data](docs/InlineResponse20050Data.md)
 - [InlineResponse20051](docs/InlineResponse20051.md)
 - [InlineResponse20051Data](docs/InlineResponse20051Data.md)
 - [InlineResponse20052](docs/InlineResponse20052.md)
 - [InlineResponse20053](docs/InlineResponse20053.md)
 - [InlineResponse20054](docs/InlineResponse20054.md)
 - [InlineResponse20055](docs/InlineResponse20055.md)
 - [InlineResponse20056](docs/InlineResponse20056.md)
 - [InlineResponse20057](docs/InlineResponse20057.md)
 - [InlineResponse20058](docs/InlineResponse20058.md)
 - [InlineResponse2005Data](docs/InlineResponse2005Data.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse2007](docs/InlineResponse2007.md)
 - [InlineResponse2007Data](docs/InlineResponse2007Data.md)
 - [InlineResponse2008](docs/InlineResponse2008.md)
 - [InlineResponse2008Data](docs/InlineResponse2008Data.md)
 - [InlineResponse2009](docs/InlineResponse2009.md)
 - [InlineResponse2009Data](docs/InlineResponse2009Data.md)
 - [InlineResponse201](docs/InlineResponse201.md)
 - [InlineResponse2011](docs/InlineResponse2011.md)
 - [InlineResponse2012](docs/InlineResponse2012.md)
 - [InlineResponse2013](docs/InlineResponse2013.md)
 - [InlineResponse2014](docs/InlineResponse2014.md)
 - [InlineResponse2015](docs/InlineResponse2015.md)
 - [InlineResponse2016](docs/InlineResponse2016.md)
 - [InlineResponse2017](docs/InlineResponse2017.md)
 - [InlineResponse202](docs/InlineResponse202.md)
 - [InlineResponse404](docs/InlineResponse404.md)
 - [InlineResponseDefault](docs/InlineResponseDefault.md)
 - [InlineResponseDefault1](docs/InlineResponseDefault1.md)
 - [InlineResponseDefaultErrors](docs/InlineResponseDefaultErrors.md)
 - [InlineResponseDefaultMeta](docs/InlineResponseDefaultMeta.md)
 - [InlineResponseDefaultSource](docs/InlineResponseDefaultSource.md)
 - [Ip](docs/Ip.md)
 - [IpConnection](docs/IpConnection.md)
 - [JoinAConferenceRequest](docs/JoinAConferenceRequest.md)
 - [MessagingProfilesNumberPoolSettings](docs/MessagingProfilesNumberPoolSettings.md)
 - [MessagingProfilesUrlShortenerSettings](docs/MessagingProfilesUrlShortenerSettings.md)
 - [Metadata](docs/Metadata.md)
 - [MuteConferenceParticipantsRequest](docs/MuteConferenceParticipantsRequest.md)
 - [NumberOrdersPhoneNumbers](docs/NumberOrdersPhoneNumbers.md)
 - [NumberOrdersRegulatoryRequirements](docs/NumberOrdersRegulatoryRequirements.md)
 - [NumberReservationsPhoneNumbers](docs/NumberReservationsPhoneNumbers.md)
 - [OneOfbody13To](docs/OneOfbody13To.md)
 - [OneOfbody14To](docs/OneOfbody14To.md)
 - [OneOfbody15To](docs/OneOfbody15To.md)
 - [OneOfbody16To](docs/OneOfbody16To.md)
 - [OutboundConfiguration](docs/OutboundConfiguration.md)
 - [OutboundConfiguration1](docs/OutboundConfiguration1.md)
 - [OutboundVoiceProfile](docs/OutboundVoiceProfile.md)
 - [OutboundVoiceProfile1](docs/OutboundVoiceProfile1.md)
 - [OutboundVoiceProfile2](docs/OutboundVoiceProfile2.md)
 - [OutboundVoiceProfilesCallRecording](docs/OutboundVoiceProfilesCallRecording.md)
 - [PhoneNumbersidvoiceCallForwarding](docs/PhoneNumbersidvoiceCallForwarding.md)
 - [PhoneNumbersidvoiceCallRecording](docs/PhoneNumbersidvoiceCallRecording.md)
 - [PhoneNumbersidvoiceCnamListing](docs/PhoneNumbersidvoiceCnamListing.md)
 - [PhoneNumbersidvoiceMediaFeatures](docs/PhoneNumbersidvoiceMediaFeatures.md)
 - [PlayAudioUrlRequest](docs/PlayAudioUrlRequest.md)
 - [PlaybackStopRequest](docs/PlaybackStopRequest.md)
 - [RejectRequest](docs/RejectRequest.md)
 - [SendDtmfRequest](docs/SendDtmfRequest.md)
 - [SpeakRequest](docs/SpeakRequest.md)
 - [StartForkingRequest](docs/StartForkingRequest.md)
 - [StartRecordingRequest](docs/StartRecordingRequest.md)
 - [StopForkingRequest](docs/StopForkingRequest.md)
 - [StropRecordingRequest](docs/StropRecordingRequest.md)
 - [TransferCallRequest](docs/TransferCallRequest.md)
 - [UnholdConferenceParticipantsRequest](docs/UnholdConferenceParticipantsRequest.md)
 - [UnmuteConferenceParticipantsRequest](docs/UnmuteConferenceParticipantsRequest.md)
 - [UpdatableAttributesForCallControlApplications](docs/UpdatableAttributesForCallControlApplications.md)
 - [UpdatableAttributesForCallControlApplications1](docs/UpdatableAttributesForCallControlApplications1.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author

support@telnyx.com