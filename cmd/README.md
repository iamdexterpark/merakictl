# Merakictl CLI Reference 

 
  
## Root Commands
 
Command | Syntax | Description |
--- | --- | ---
| **claim** | merakictl claim [order/serial/licence] | Claim orders, licences, serial numbers into dashboard. |
| **create** | merakictl create [COMMAND] [SUBCOMMAND] [FLAGS] | Creates (POST) new resources. |
| **get** | merakictl get [COMMAND] [SUBCOMMAND] [FLAGS]| Operation for displaying (GET) api resources. |
| **update** | merakictl update [COMMAND] [SUBCOMMAND] [FLAGS]| Updates (PUT) targeted resources. |
| **remove** | merakictl remove [COMMAND] [SUBCOMMAND] [FLAGS]| Destructive (DELETE) API call for removing resources from the Dashboard. |
| **version** | merakictl version | Displays the version and associate release information of Merakictl. |


## Subcommands
 
Subcommand | Abbreviated | Description |
--- | --- | --- |
| **organization** | org | Collection of Networks. |
| **network** | net | Collection of Devices. | 
| **device** | dev | Meraki Products. | 
| **mx** | | UTM appliances with SD-WAN capability included. | 
| **ms**  | | Branch and campus access and aggregation switches with remote packet capture and cable testing. |
| **mr** | |  Industry-leading indoor and outdoor 802.11ax wireless access points with built-in location analytics. |
| **mv** | | Indoor and outdoor smart cameras, eliminating traditional external storage such as NVRs. |
| **mg** | |  Cellular gateways seamlessly transpose a wireless cellular signal to wired Ethernet, providing primary or failover connectivity. |
| **sm** | |  A complete enterprise solution for provisioning, monitoring, and securing of mobile devices. |
| **insight** | in |  End-to-end visibility into end-users' experience on the network. |


## Flags

Flag Type | Long | Short | Description |
--- | --- | --- | ---
| Global | **--input** | | Global flag for using a yaml file for passing config to the Dashboard API |
| Global | **--export** | -e | Global flag for extracting config from the Meraki API via get commands. |
| Global | **--diff** | | Global flag for diffing a yaml/json file with config in the Dashboard API |
| Global | **--verbose** | -v | Global flag to display the http request & response for troubleshooting. |
| Global | **--organization** | -o | Global flag for matching a string against Organization ids. |
| Global | **--network** | -n | Global flag for matching a string against Network ids. |
| Global | **--hostname** | -h | Global flag for matching a string against device hostnames. |

## Common Filters

Filters are HTTP parameters that shape our API queries. The list of filters below is not exhaustive but provide details around those most commonly use.

Filter | Example | Description |
 --- | --- | --- 
t0 | |  The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
t1 | |  The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
timespan | |  The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
resolution | | The time resolution in seconds for returned data. The valid resolutions are: 600, 1200, 3600, 14400, 86400. The default is 86400.
autoResolution | | Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false. 
deviceSerial | | The serial of the Meraki device which the list of events will be filtered with
uplink | | Optional filter for a specific WAN uplink. Valid uplinks are wan1, wan2, cellular. Default will return all uplinks.
ip | | Optional filter for a specific destination IP. Default will return all destination IPs. 
search | |  
configTemplateId | | An optional parameter that is the ID of a config template. Will return all networks bound to that template. 
tags | | An optional parameter to filter networks by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). 
usedState | |  
tagsFilterType | | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. 
username | | The username of a Systems Manager user
email | | The email of a network user account or a Systems Manager device 
mac | | The MAC of a network client device or a Systems Manager device 
serial | | The serial of a Systems Manager device 
imei | | The IMEI of a Systems Manager device 
bluetoothMac | | The MAC of a Bluetooth client 
includeConnectivityHistory | |  Include the connectivity history for this client
connectivityHistoryTimespan | | The timespan, in seconds, for the connectivityHistory data. By default 1 day, 86400, will be used.
includedEventTypes[] | |  A list of event types. The returned events will be filtered to only include events with these types.
excludedEventTypes[] | |  A list of event types. The returned events will be filtered to exclude events with these types.
sensorSerial | | The serial of the sensor device which the list of events will be filtered with
gatewaySerial | | The serial of the environmental gateway device which the list of events will be filtered with
loginIdentifier | | The username, email, or phone number used during login 
deviceType | |  
sortOrder | | Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order.
networkIds | |  
objectType | | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. 
clientId | | Filter results by network client to return per-device, per-band AP channel utilization metrics inner joined by the queried client's connection history.
apTag | | Filter results by AP tag to return AP channel utilization metrics for devices labeled with the given tag; either :clientId or :deviceSerial must be jointly specified.
band | | Filter results by band (either '2.4' or '5'). 
ssid | | Filter results by SSID
ssidNumber | | An SSID number to include. If not specified, events for all SSIDs will be returned.
vlan | | Filter results by VLAN
fields | | Partial selection: If present, this call will return only the selected fields of ["rawDistribution", "avg"]. All fields will be returned by default. Selected fields must be entered as a comma separated string. 
wifiMacs | |
serials | |  
ids | | Filter users by id(s). 
scope | | Specifiy a scope (one of all, none, withAny, withAll, withoutAny, withoutAll) and a set of tags. 
withDetails | | Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response 
perPage | | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
startingAfter | | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
endingBefore | | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.


The filters `perPage`, `startingAfter`, and `endingBefore` are part of the Meraki Dashboard APIs pagination strategy. 
They should be considered overrides as Merakictl is designed to automatically handle pagination. 
To learn more see [RFC5988 Web Linking](https://tools.ietf.org/html/rfc5988).


## Organization

All Organization level API calls. 

 HTTP | Operation | Syntax | Filters | Description |
----- | --------- | ------ | ----------- | ----------- |
 GET  | list | merakictl get organization list | | List the organizations that the user has privileges on.
 GET  | info | merakictl get organization info {organizationId} | | List a specific organization that the user has privileges on. 
 POST | {orgName} | merakictl create org {orgName} | | Create a new organization.
 PUT  | {orgName} | merakictl update org {orgName} | | Update an organization.
 DELETE | {orgName} | merakictl remove org {orgName} | | Delete an organization.
 GET  | actionbatches | merakictl get organization actionbatches {organizationId} | | Return The List Of Action Batches In The Organization.
 GET  | actionbatch | merakictl get organization actionbatch {actionBatchId} {organizationId} | | Return A Single Action Batch.
 GET  | admins | merakictl get organization admins {organizationId} | | List The Dashboard Administrators In This Organization.
 GET  | brandingpolicies | merakictl get organization brandingpolicies {organizationId} | | Return The Branding Policy IDs Of An Organization.
 GET  | brandingpolicy | merakictl get organization brandingpolicy {brandingPolicyId} {organizationId} | | Return The Branding Policies Of An Organization.
 GET  | configurationtemplates | merakictl get organization configurationtemplates {organizationId} | | List The Configuration Templates For This Organization.
 GET  | configurationtemplate | merakictl get organization configurationtemplate {configTemplateId} {organizationId} | | Return a Configuration Template For This Organization.
 GET  | devices | merakictl get organization devices {organizationId} | {perPage} {startingAfter} {endingBefore} |  List the devices in an organization.
 GET  | device | merakictl get organization device {serial} {organizationId} |  | Return A Single Device From The Inventory Of An Organization.
 GET  | inventory | merakictl get organization inventory |  {perPage} {startingAfter} {endingBefore} {usedState} {search} | Return The Device Inventory For An Organization.
 GET  | licences | merakictl get organization licences {organizationId} | {perPage} {startingAfter} {endingBefore} {deviceSerial} {networkId} {state} | List The Licenses For An Organization.
 GET  | licence | merakictl get organization licence {licenceId} {organizationId} | | List A Single License For An Organization.
 GET  | securitysettings | merakictl get organization securitysettings {organizationId} | | Returns The Login Security Settings For An Organization.
 GET  | networks | merakictl get organization networks {organizationId} | {configTemplateId} {tags} {tagsFilterType} {perPage} {startingAfter} {endingBefore} | List the networks that the user has privileges on in an organization.
 GET  | ldps | merakictl get organization ldps {organizationId} | | List the SAML IdPs in your organization.
 GET  | ldp | merakictl get organization ldp {ldpId} {organizationId} | | List a SAML IdP in your organization.
 GET  | saml | merakictl get organization saml {organizationId} | | Returns the SAML SSO enabled settings for an organization.
 GET  | samlroles | merakictl get organization samlroles {organizationId} | | List the SAML roles for this organization.
 GET  | samlrole | merakictl get organization samlrole {samalRoleId} {organizationId} | | List a single SAML role for this organization.
 GET  | snmp | merakictl get organization snmp {organizationId} | | Return the SNMP settings for an organization.
 GET  | apiresponsecount | merakictl get organization apiresponsecount {organizationId} | {t0} {t1} {timespan} | Return an aggregated overview of API requests data.
 GET  | apirequests | merakictl get organization apirequests {organizationId} | {t0} {t1} {timespan} {perPage} {startingAfter} {endingBefore} {adminId} {path} {method} {responseCode} {sourceIp} | List the API requests made by an organization
 GET  | configurationchanges | merakictl get organization configurationchanges {organizationId} | {t0} {t1} {timespan} {perPage} {startingAfter} {endingBefore} {networkId} {adminId} | View the Change Log for your organization.
 GET  | devicestatus | merakictl get organization devicestatus {organizationId}  | {perPage} {startingAfter} {endingBefore} | List the status of every Meraki device in the organization.
 GET  | losslatency | merakictl get organization losslatency {organizationId} | {t0} {t1} {timespan} {uplink} {ip} | Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago.
 GET  | licenseoverview | merakictl get organization licenseoverview {organizationId} | | Return an overview of the license state for an organization.
 GET  | openapi | merakictl get organization openapi {organizationId} | | Return the OpenAPI 2.0 Specification of the organization's API documentation in JSON.
 GET  | webhooklogs | merakictl get organization webhooklogs {organizationId} | {t0} {t1} {timespan} {perPage} {startingAfter} {endingBefore} {url} | Return the log of webhook POSTs sent.
 
 
## Network

All Network level API calls. 

 HTTP | Operation | Syntax | Filters | Description |
----- | --------- | ------ | ----------- | ----------- |
 GET  |  | merakictl get network {networkId} | | Return a network.
 GET  | alertconfiguration | merakictl get network alertconfiguration {networkId} | | Return the alert configuration for this network.
 GET  | clientpolicy | merakictl get network clientpolicy {clientId} {networkId} | | Return the policy assigned to a client on the network. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | clientsplashauthorization | merakictl get network clientsplashauthorization {clientId} {networkId} | | Return the splash authorization for a client, for each SSID they've associated with through splash. Only enabled SSIDs with Click-through splash enabled will be included. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | devices | merakictl get network devices {networkId} | | List the devices in a network.
 GET  | firmwareupgrades | merakictl get network firmwareupgrades {networkId} | | Get current maintenance window for a network.
 GET  | floorplans | merakictl get network floorplans {networkId} | | List the floor plans that belong to your network.
 GET  | floorplan | merakictl get network floorplan {floorPlanId} {networkId} | | Find a floor plan by ID.
 GET  | grouppolicies | merakictl get network grouppolicies {networkId} | | List the group policies in a network.
 GET  | grouppolicy | merakictl get network grouppolicy {groupPolicyId} {networkId} | | Display a group policy.
 GET  | merakiauthusers | merakictl get network merakiauthusers {networkId} | | List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a wired network).
 GET  | merakiauthuser | merakictl get network merakiauthuser {merakiAuthUserId} {networkId} | | Return the Meraki Auth splash guest, RADIUS, or client VPN user.
 GET  | mqttbrokers | merakictl get network mqttbrokers {networkId} | | List the MQTT brokers for this network.
 GET  | mqttbroker | merakictl get network mqttbroker {mqttBrokerId} {networkId} | | Return an MQTT broker.
 GET  | netflow | merakictl get network netflow {networkId} | | Return the NetFlow traffic reporting settings for a network.
 GET  | channelutilization | merakictl get network channelutilization {networkId} | {t0} {t1} {timespan} {resolution} {perPage} {startingAfter} {endingBefore} | Get the channel utilization over each radio for all APs in a network.
 GET  | piikeys | merakictl get network piikeys {networkId} | {username} {email} {mac} {serial} {imei} {bluetoothMac} | List the keys required to access Personally Identifiable Information (PII) for a given identifier. Exactly one identifier will be accepted. If the organization contains org-wide Systems Manager users matching the key provided then there will be an entry with the key "0" containing the applicable keys.
 GET  | piirequests | merakictl get network piirequests {networkId} | | List the PII requests for this network or organization.
 GET  | piirequest | merakictl get network piirequest {requestId} {networkId} | | Return a PII request.
 GET  | smdevices | merakictl get network smdevices {networkId} | {username} {email} {mac} {serial} {imei} {bluetoothMac}  | Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier. These device IDs can be used with the Systems Manager API endpoints to retrieve device details. Exactly one identifier will be accepted.
 GET  | smowners | merakictl get network smowners {networkId} | {username} {email} {mac} {serial} {imei} {bluetoothMac}  | Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier. These owner IDs can be used with the Systems Manager API endpoints to retrieve owner details. Exactly one identifier will be accepted.
 GET  | settings | merakictl get network settings {networkId} | | Return the settings for a network.
 GET  | snmp | merakictl get network snmp {networkId} | | Return the SNMP settings for a network.
 GET  | syslog | merakictl get network syslog {networkId} | | List the syslog servers for a network.
 GET  | trafficanalysis | merakictl get network trafficanalysis {networkId} | | Return the traffic analysis settings for a network.
 GET  | trafficshaping | merakictl get network trafficshaping {networkId} | | Returns the application categories for traffic shaping rules.
 GET  | dscp | merakictl get network dscp {networkId} | | Returns the available DSCP tagging options for your traffic shaping rules.
 GET  | httpservers | merakictl get network httpservers {networkId} | | List the HTTP servers for a network.
 GET  | httpserver | merakictl get network httpserver {httpServerId} {networkId} | | Return an HTTP server for a network.
 GET  | webhooktest | merakictl get network webhooktest {weebhookTestId} {networkId} | | Return the status of a webhook test for a network.
 GET  | bluetoothclients | merakictl get network bluetoothclients {networkId} | {t0} {t1} {timespan} {perPage} {startingAfter} {endingBefore} {includeConnectivityHistory} | List the Bluetooth clients seen by APs in this network.
 GET  | bluetoothclient | merakictl get network bluetoothclient {networkId} | {includeConnectivityHistory} {connectivityHistoryTimespan} | Return a Bluetooth client. Bluetooth clients can be identified by their ID or their MAC.
 GET  | clienttraffichistory | merakictl get network clienttraffichistory {clientId} {networkId} | {perPage} {startingAfter} {endingBefore} | Return the client's network traffic data over time. Usage data is in kilobytes. This endpoint requires detailed traffic analysis to be enabled on the Network-wide > General page. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | clientusagehistory | merakictl get network clientusagehistory {clientId} {networkId} | | Return the client's daily usage history. Usage data is in kilobytes. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | networkclients | merakictl get network networkclients {networkId} | {t0} {timespan} {perPage} {startingAfter} {endingBefore} | List the clients that have used this network in the timespan.
 GET  | clientidentifier | merakictl get network clientidentifier {clientId} {networkId} | | Return the client associated with the given identifier. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | environmentalevents | merakictl get network environmentalevents {networkId} | {includedEventTypes[]} {excludedEventTypes[]} {sensorSerial} {gatewaySerial} {perPage} {startingAfter} {endingBefore} | List the environmental events for the network.
 GET  | events | merakictl get network events {networkId} | {productType} {includedEventTypes[]} {excludedEventTypes[]} {deviceMac} {deviceSerial} {deviceName} {clientIp} {clientMac} {clientName} {smDeviceMac} {smDeviceName} {perPage} {startingAfter} {endingBefore} | List the events for the network.
 GET  | splashloginattempts | merakictl get network splashloginattempts {networkId} | {ssidNumber} {loginIdentifier} {timespan} | List the splash login attempts for a network.
 GET  | traffic | merakictl get network traffic {networkId} | {t0} {timespan} {deviceType} | Return the traffic analysis data for this network. Traffic analysis with hostname visibility must be enabled on the network.


## Device

All Device level API calls. 

 HTTP | Operation | Syntax | Filters | Description |
----- | --------- | ------ | ----------- | ----------- |
 GET  | mgmtinterface | merakictl get device mgmtinterface {serial} | | Return the management interface settings for a device.
 GET  | singledevice | merakictl get device singledevice {serial} |  | Return a single device.
 GET  | clients | merakictl get device clients {serial} | {t0} {timespan} | List the clients of a device, up to a maximum of a month ago. The usage of each client is returned in kilobytes. If the device is a switch, the switchport is returned; otherwise the switchport field is null.
 GET  | lldpcdp | merakictl get device lldpcdp {serial} | | List LLDP and CDP information for a device.
 GET  | uplinkloss | merakictl get device uplinkloss {serial} | {t0} {t1} {timespan} {resolution} {uplink} {ip} | Get the uplink loss percentage and latency in milliseconds for a wired network device.
 
 
 
## MX Appliance 
 
 All MX level API calls. 
 
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
 GET  | connectivitymonitoringdestinations | merakictl get mx connectivitymonitoringdestinations {networkId} | | Return the connectivity testing destinations for an MX network.
 GET  | contentfilteringcategories | merakictl get mx contentfilteringcategories {networkId} | | List all available content filtering categories for an MX network.
 GET  | contentfiltering | merakictl get mx contentfiltering {networkId} | | Return the content filtering settings for an MX network.
 GET  | cellularfirewallrules | merakictl get mx cellularfirewallrules {networkId} | | Return the cellular firewall rules for an MX network.
 GET  | firewalledservices | merakictl get mx firewalledservices {networkId} | | List the appliance services and their accessibility rules.
 GET  | firewalledservice | merakictl get mx firewalledservice {service} {networkId} | | Return the accessibility settings of the given service ('ICMP', 'web', or 'SNMP').
 GET  | inboundfirewallrules | merakictl get mx inboundfirewallrules {networkId} | | Return the inbound firewall rules for an MX network.
 GET  | l3firewallrules | merakictl get mx l3firewallrules {networkId} | | Return the L3 firewall rules for an MX network.
 GET  | l7applicationcategories | merakictl get mx l7applicationcategories {networkId} | | Return the L7 firewall application categories and their associated applications for an MX network.
 GET  | l7firewallrules | merakictl get mx l7firewallrules {networkId} | | List the MX L7 firewall rules for an MX network.
 GET  | onetomanynatrules | merakictl get mx onetomanynatrules {networkId} | | Return the 1:Many NAT mapping rules for an MX network.
 GET  | onetoonenatrules | merakictl get mx onetoonenatrules {networkId} | | Return the 1:1 NAT mapping rules for an MX network.
 GET  | portforwardingrules | merakictl get mx portforwardingrules {networkId} | | Return the port forwarding rules for an MX network.
 GET  | ports | merakictl get mx ports {networkId} | | List per-port VLAN settings for all ports of a MX.
 GET  | port | merakictl get mx port {portId} {networkId} | | Return per-port VLAN settings for a single MX port.
 GET  | networkintrusion | merakictl get mx networkintrusion {networkId} | | Returns all supported intrusion settings for an MX network.
 GET  | organizationintrusion | merakictl get mx organizationintrusion {organizationId} | | Returns all supported intrusion settings for an organization.
 GET  | malware | merakictl get mx malware {networkId} | | Returns all supported malware settings for an MX network.
 GET  | settings | merakictl get mx settings {networkId} | | Return the appliance settings for a network.
 GET  | singlelan | merakictl get mx singlelan {networkId} | | Return single LAN configuration.
 GET  | staticroutes | merakictl get mx staticroutes {networkId} | | List the static routes for an MX or teleworker network.
 GET  | staticroute | merakictl get mx staticroute {staticRouteId} {networkId} | | Return a static route for an MX or teleworker network.
 GET  | customperformanceclasses | merakictl get mx customperformanceclasses {networkId} | | List all custom performance classes for an MX network.
 GET  | customperformanceclass | merakictl get mx customperformanceclass {customPerformanceClassId} {networkId} | | Return a custom performance class for an MX network.
 GET  | trafficshapingrules | merakictl get mx trafficshapingrules {networkId} | | Display the traffic shaping settings rules for an MX network.
 GET  | uplinkbandwidth | merakictl get mx uplinkbandwidth {networkId} | | Returns the uplink bandwidth settings for your MX network..
 GET  | uplinkselection | merakictl get mx uplinkselection {networkId} | | Show uplink selection settings for an MX network.
 GET  | trafficshaping | merakictl get mx trafficshaping {networkId} | | Display the traffic shaping settings for an MX network.
 GET  | vlanssettings | merakictl get mx vlanssettings {networkId} | | Returns the enabled status of VLANs for the network.
 GET  | vlans | merakictl get mx vlans {networkId} | | List the VLANs for an MX network.
 GET  | vlan | merakictl get mx vlan {vlanId} {networkId} | | Return a VLAN.
 GET  | sitetositevpn | merakictl get mx sitetositevpn {networkId} | | Return the site-to-site VPN settings of a network. Only valid for MX networks.
 GET  | thirdpartyvpnpeers | merakictl get mx thirdpartyvpnpeers {organizationId} | | Return the third party VPN peers for an organization.
 GET  | vpnfirewallrules | merakictl get mx vpnfirewallrules {organizationId} | | Return the firewall rules for an organization's site-to-site VPN.
 GET  | warmspare | merakictl get mx warmspare {networkId} | | Return MX warm spare settings.
 GET  | securityevents | merakictl get mx securityevents {clientId} {networkId} | {t0} {t1} {timespan} {perPage} {startingAfter} {endingBefore} {sortOrder} | List the security events for a client. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | dhcpsubnets | merakictl get mx dhcpsubnets {serial} | | Return the DHCP subnet information for an appliance.
 GET  | performance | merakictl get mx performance {serial} | | Return the performance score for a single MX. Only primary MX devices supported. If no data is available, a 204 error code is returned.
 GET  | networksecurityevents | merakictl get mx networksecurityevents {networkId} | {t0} {t1} {timespan} {perPage} {startingAfter} {endingBefore} {sortOrder} | List the security events for a network.
 GET  | organizationsecurityevents | merakictl get mx organizationsecurityevents {organizationId} | {t0} {t1} {timespan} {perPage} {startingAfter} {endingBefore} {sortOrder} | List the security events for an organization.
 GET  | uplinkstatuses | merakictl get mx uplinkstatuses {organizationId} | {perPage} {startingAfter} {endingBefore} | List the uplink status of every Meraki MX and Z series appliances in the organization.
 GET  | vpnstats | merakictl get mx vpnstats {organizationId} | {perPage} {startingAfter} {endingBefore} {networkIds} {t0} {t1} {timespan} | Show VPN history stat for networks in an organization.
 GET  | vpnstatuses | merakictl get mx vpnstatuses {organizationId} | {perPage} {startingAfter} {endingBefore} {networkIds} | Show VPN status for networks in an organization.


## Camera MV  
 
 All MV level API calls. 
 
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
 GET  | qualityandretention | merakictl get mv qualityandretention {serial} | | Returns quality and retention settings for the given camera.
 GET  | qualityretentionprofiles | merakictl get mv qualityretentionprofiles {networkId} | | List the quality retention profiles for this network.
 GET  | qualityretentionprofile | merakictl get mv qualityretentionprofile {qualityRetentionProfileId} {networkId} | | Retrieve a single quality retention profile.
 GET  | schedules | merakictl get mvschedules  {networkId} | | Returns a list of all camera recording schedules.
 GET  | objectdetectionmodels | merakictl get mv objectdetectionmodels {serial} | | Returns the MV Sense object detection model list for the given camera.
 GET  | sense | merakictl get mv  {serial} | | Returns sense settings for a given camera.
 GET  | settings | merakictl get mv settings {serial} | | Returns video settings for the given camera.
 GET  | videolink | merakictl get mv videolink {serial} | {timestamp} | Returns video link to the specified camera. If a timestamp is supplied, it links to that timestamp.
 GET  | live | merakictl get mv live {serial} | | Returns live state from camera of analytics zones.
 GET  | overview | merakictl get mv overview {serial} | {t0} {t1} {timespan} {objectType} | Returns an overview of aggregate analytics data for a timespan.
 GET  | recent | merakictl get mv recent {serial} | {objectType} | Returns most recent record for analytics zones.
 GET  | history | merakictl get mv history {serial} {zoneId} | {t0} {t1} {timespan} {resolution} {objectType} | Return historical records for analytic zones.
 GET  | zones | merakictl get mv zones {serial} | | Returns all configured analytic zones for this camera.
 
 
## Cellular Gateway mg  
  
  All MG level API calls.
  
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
 GET  | connectivitymonitor | merakictl get mg connectivitymonitor {networkId} | | Return the connectivity testing destinations for an MG network.
 GET  | dhcp | merakictl get mg dhcp {networkId} | | List common DHCP settings of MGs.
 GET  | lan | merakictl get mg lan {serial} | | Show the LAN Settings of a MG.
 GET  | portforwardingrules | merakictl get mg portforwardingrules {serial} | | Returns the port forwarding rules for a single MG.
 GET  | subnetpool | merakictl get mg subnetpool {networkId} | | Return the subnet pool and mask configured for MGs in the network.
 GET  | uplink | merakictl get mg uplink {networkId} | | Returns the uplink settings for your MG network.
 


## Switch MS  
 
 All MS level API calls. 
 
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
  GET  | accesscontrollists | merakictl get ms accesscontrollists {networkId} | | Return the access control lists for a MS network.
  GET  | accesspolicies | merakictl get ms accesspolicies {networkId} | | List the access policies for a switch network. Only returns access policies with 'my RADIUS server' as authentication method.
  GET  | accesspolicy | merakictl get ms {networkId} {accessPolicyNumber} | | Return a specific access policy for a switch network.
  GET  | switchportsprofile | merakictl get ms switchportsprofile {organizationId} {configTemplateId} {profileId} {portId} | | Return a switch profile port.
  GET  | switchportprofile | merakictl get ms switchportprofile {organizationId} {configTemplateId} {profileId} | | Return all the ports of a switch profile.
  GET  | switchprofiles | merakictl get ms switchprofiles {organizationId} {configTemplateId} | | List the switch profiles for your switch template configuration.
  GET  | dhcpserverpolicy | merakictl get ms dhcpserverpolicy {networkId} | | Return the DHCP server policy.
  GET  | dscp | merakictl get ms dscp {networkId} | | Return the DSCP to CoS mappings.
  GET  | linkaggregations | merakictl get ms linkaggregations {networkId} | | List link aggregation groups.
  GET  | mtu | merakictl get ms mtu {networkId} | | Return the MTU configuration.
  GET  | portschedules | merakictl get ms portschedules {networkId} | | List switch port schedules.
  GET  | ports | merakictl get ms ports {serial} | | List the switch ports for a switch.
  GET  | port | merakictl get ms port {serial} {portId} | | Return a switch port.
  GET  | qosruleids | merakictl get ms qosruleids {networkId} | | Return the quality of service rule IDs by order in which they will be processed by the switch.
  GET  | qosrules | merakictl get ms qosrules {networkId} | | List quality of service rules.
  GET  | qosrule | merakictl get ms qosrule {networkId} {qosRuleId} | | Return a quality of service rule.
  GET  | dhcp | merakictl get ms dhcp {serial} {interfaceId} | | Return a layer 3 interface DHCP configuration for a switch.
  GET  | l3interfaces | merakictl get ms l3interfaces {serial} | | List layer 3 interfaces for a switch.
  GET  | l3interface | merakictl get ms l3interface {serial} {interfaceId} | | Return a layer 3 interface for a switch.
  GET  | rendezvouspoints | merakictl get ms rendezvouspoints {networkId} | | List multicast rendezvous points.
  GET  | rendezvouspoint | merakictl get ms rendezvouspoint {networkId} {rendezvousPointId} | | Return a multicast rendezvous point.
  GET  | multicast | merakictl get ms multicast {networkId} | | Return multicast settings for a network.
  GET  | ospf | merakictl get ms ospf {networkId} | | Return layer 3 OSPF routing configuration.
  GET  | staticroutes | merakictl get ms staticroutes {serial} | | List layer 3 static routes for a switch.
  GET  | staticroute | merakictl get ms staticroute {serial} {staticRouteId} | | Return a layer 3 static route for a switch.
  GET  | settings | merakictl get ms settings {networkId} | | Returns the switch network settings.
  GET  | stackdhcp | merakictl get ms stackdhcp {networkId} {switchStackId} {interfaceId} | | Return a layer 3 interface DHCP configuration for a switch stack.
  GET  | stackl3interfaces | merakictl get ms stackl3interfaces {networkId} {switchStackId} | | List layer 3 interfaces for a switch stack.
  GET  | stackl3interface | merakictl get ms stackl3interface {networkId} {switchStackId} {interfaceId} | | Return a layer 3 interface from a switch stack.
  GET  | stackstaticroutes | merakictl get ms stackstaticroutes {networkId} {switchStackId} | | List layer 3 static routes for a switch stack.
  GET  | stackstaticroute | merakictl get ms stackstaticroute {networkId} {switchStackId} {staticRouteId} | | Return a layer 3 static route for a switch stack.
  GET  | switchstacks | merakictl get ms switchstacks {networkId} | | List the switch stacks in a network.
  GET  | switchstack | merakictl get ms switchstack {networkId} {switchStackId} | | Show a switch stack.
  GET  | stormcontrol | merakictl get ms stormcontrol {networkId} | | Return the storm control configuration for a switch network.
  GET  | stp | merakictl get ms stp {networkId} | | storm control.
  GET  | warmspare | merakictl get ms warmspare {serial} | | Return warm spare configuration for a switch.
  GET  | packets | merakictl get ms packets {serial} | {t0} {timespan} | Return the packet counters for all the ports of a switch.
  GET  | portsstatuses | merakictl get ms portsstatuses {serial} | {t0} {timespan}  | Return the status for all the ports of a switch.




## Wireless MR  
 
 All MR level API calls. 
 
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
  GET  | alternatemgmtinterface | merakictl get mr alternatemgmtinterface {networkId} | | Return alternate management interface and devices with IP assigned.
  GET  | bluetoothsettings | merakictl get mr bluetoothsettings {networkId} | | Return the Bluetooth settings for a network. Bluetooth settings must be enabled on the network.
  GET  | bluetoothdevicesettings | merakictl get mr bluetoothdevicesettings {serial} | | Return the bluetooth settings for a wireless device.
  GET  | radiosettings | merakictl get mr radiosettings {serial} | | Return the radio settings of a device.
  GET  | rfprofiles | merakictl get mr rfprofiles {networkId} | {includeTemplateProfiles} | List the non-basic RF profiles for this network.
  GET  | rfprofile | merakictl get mr rfprofile {networkId} {rfProfileId} | | Return a RF profile.
  GET  | wirelesssettings | merakictl get mr wirelesssettings {networkId} | | Return the wireless settings for a network.
  GET  | l3FirewallRules | merakictl get mr l3FirewallRules {networkId} {number} | | Return the L3 firewall rules for an SSID on an MR network.
  GET  | l7FirewallRules | merakictl get mr l7FirewallRules {networkId} {number} | | Return the L7 firewall rules for an SSID on an MR network.
  GET  | identitypsks | merakictl get mr identitypsks {networkId} {number} | | List all Identity PSKs in a wireless network.
  GET  | identitypsk | merakictl get mr identitypsk {networkId} {number} {identityPskId} | | Return an Identity PSK.
  GET  | splashsettings | merakictl get mr splashsettings {networkId} {number} | | Display the splash page settings for the given SSID.
  GET  | trafficshapingrules | merakictl get mr trafficshapingrules {networkId} {number} | | Display the traffic shaping settings for a SSID on an MR network.
  GET  | ssids | merakictl get mr ssids {networkId} | | List the MR SSIDs in a network.
  GET  | ssid | merakictl get mr ssid {networkId} {number} | | Return a single MR SSID.
  GET  | airmarshal | merakictl get mr airmarshal {networkId} | {t0} {timespan} | List Air Marshal scan results from a network.
  GET  | channelutilizationhistory | merakictl get mr channelutilizationhistory {networkId} | {t0} {t1} {timespan} {resolution} {autoResolution} {clientId} {deviceSerial} {apTag} {band} | Return AP channel utilization over time for a device or network client.
  GET  | clientcounthistory | merakictl get mr clientcounthistory {networkId} | {t0} {t1} {timespan} {resolution} {autoResolution} {clientId} {deviceSerial} {apTag} {band} {ssid} | Return wireless client counts over time for a network, device, or network client.
  GET  | connectionstat | merakictl get mr connectionstat {networkId} | {t0} {t1} {timespan} {band} {ssid} {vlan} {apTag} | Aggregated connectivity info for a given client on this network. Clients are identified by their MAC.
  GET  | connectionstats  | merakictl get mr connectionstats  {networkId} | {t0} {t1} {timespan} {band} {ssid} {vlan} {apTag}  |  Aggregated connectivity info for this network, grouped by clients.
  GET  | connectivityevents | merakictl get mr connectivityevents {networkId} {clientId} | {perPage} {startingAfter} {endingBefore} {t0} {t1} {timespan} {types} {includedSeverities} {band} {ssidNumber} {deviceSerial} | List the wireless connectivity events for a client within a network in the timespan.
  GET  | latencyhistory | merakictl get mr latencyhistory {networkId} {clientId} | {t0} {t1} {timespan} {resolution} | Return the latency history for a client. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP. The latency data is from a sample of 2% of packets and is grouped into 4 traffic categories: background, best effort, video, voice. Within these categories the sampled packet counters are bucketed by latency in milliseconds..
  GET  | latencystat | merakictl get mr latencystat {networkId} {clientId}| {t0} {t1} {timespan} {band} {ssid} {vlan} {apTag} {fields} | Aggregated latency info for a given client on this network. Clients are identified by their MAC.
  GET  | latencystats | merakictl get mr latencystats {networkId} | {t0} {t1} {timespan} {band} {ssid} {vlan} {apTag} {fields} | Aggregated latency info for this network, grouped by clients.
  GET  | deviceconnectionstats | merakictl get mr deviceconnectionstats {serial} | {t0} {t1} {timespan} {band} {ssid} {vlan} {apTag} | Aggregated connectivity info for a given AP on this network.
  GET  | networkconnectionstats | merakictl get mr networkconnectionstats {networkId} | {t0} {t1} {timespan} {band} {ssid} {vlan} {apTag} | Aggregated connectivity info for this network.
  GET  | dataratehistory | merakictl get mr dataratehistory {networkId} | {t0} {t1} {timespan} {resolution} {autoResolution} {clientId} {deviceSerial} {apTag} {band} {ssid} | Return PHY data rates over time for a network, device, or network client.
  GET  | connectionstats | merakictl get mr connectionstats {networkId} |{t0} {t1} {timespan} {band} {ssid} {vlan} {apTag} | Aggregated connectivity info for this network, grouped by node.
  GET  | latencystats | merakictl get mr latencystats {networkId} | {t0} {t1} {timespan} {band} {ssid} {vlan} {apTag} {fields} | Aggregated latency info for this network, grouped by node.
  GET  | failedconnections | merakictl get mr failedconnections {networkId} | {t0} {t1} {timespan} {band} {ssid} {vlan} {apTag} {serial} {clientId} | List of all failed client connection events on this network in a given time range.
  GET  | latencyhistory | merakictl get mr latencyhistory {networkId} | {t0} {t1} {timespan} {resolution} {autoResolution} {clientId} {deviceSerial} {apTag} {band} {ssid} {accessCategory} | Return average wireless latency over time for a network, device, or network client.
  GET  | devicelatencystats | merakictl get mr latencystats {serial} | {t0} {t1} {timespan} {band} {ssid} {vlan} {apTag} {fields} | Aggregated latency info for a given AP on this network.
  GET  | networklatencystats | merakictl get mr networklatencystats {networkId} | {t0} {t1} {timespan} {band} {ssid} {vlan} {apTag} {fields} | Aggregated latency info for this network.
  GET  | meshstatuses | merakictl get mr meshstatuses {networkId} | {perPage} {startingAfter} {endingBefore} | List wireless mesh statuses for repeaters.
  GET  | signalqualityhistory | merakictl get mr signalqualityhistory {networkId} | {t0} {t1} {timespan} {resolution} {autoResolution} {clientId} {deviceSerial} {apTag} {band} {ssid} | Return signal quality (SNR/RSSI) over time for a device or network client.
  GET  | status | merakictl get mr status {serial} | | Return the SSID statuses of an access point.
  GET  | usagehistory | merakictl get mr usagehistory {networkId} |  {t0} {t1} {timespan} {resolution} {autoResolution} {clientId} {deviceSerial} {apTag} {band} {ssid} | Return AP usage over time for a device or network client.



## SM  
 
 All SM level API calls. 
 
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
  GET  | apnscert | merakictl get sm apnscert {organizationId} | | Get the organization's APNS certificate.
  GET  | bypassactivationlockattempts | merakictl get sm bypassactivationlockattempts {networkId} {attemptId} | | Bypass activation lock attempt status.
  GET  | certs | merakictl get sm certs {networkId} {deviceId | | List the certs on a device.
  GET  | deviceprofiles | merakictl get sm deviceprofiles {networkId} {deviceId} | | Get the profiles associated with a device.
  GET  | networkadapters | merakictl get sm networkadapters {networkId} {deviceId} | | List the network adapters of a device.
  GET  | restrictions | merakictl get sm restrictions {networkId} {deviceId} | | List the restrictions on a device.
  GET  | securitycenters | merakictl get sm securitycenters {networkId} {deviceId} | | List the security centers on a device.
  GET  | softwares | merakictl get sm softwares {networkId} {deviceId} | | Get a list of softwares associated with a device.
  GET  | wlanlists | merakictl get sm wlanlists {networkId} {deviceId} | | List the saved SSID names on a device.
  GET  | devices | merakictl get sm devices {networkId} | {fields} {wifiMacs} {serials} {ids} {scope} {perPage} {startingAfter} {endingBefore} | List the devices enrolled in an SM network with various specified fields and filters.
  GET  | profiles | merakictl get sm profiles {networkId} | | List all profiles in a network.
  GET  | targetgroups | merakictl get sm targetgroups {networkId} | {withDetails} | List the target groups in this network.
  GET  | targetgroup | merakictl get sm targetgroup {networkId} {targetGroupId} | {withDetails} | Return a target group.
  GET  | deviceprofiles | merakictl get sm deviceprofiles {networkId} {clientId} | | Get the profiles associated with a user.
  GET  | softwares | merakictl get sm softwares {networkId} {clientId} | | Get a list of softwares associated with a user.
  GET  | users | merakictl get sm users {networkId} | {ids} {usernames} {emails} {scope} | List the owners in an SM network with various specified fields and filters.
  GET  | vppaccount | merakictl get sm vppaccount {organizationId} {vppAccountId} | | Get a hash containing the unparsed token of the VPP account with the given ID.
  GET  | vppaccounts | merakictl get sm vppaccounts {organizationId} | | List the VPP accounts in the organization.
  GET  | cellularusagehistory | merakictl get sm cellularusagehistory {networkId} {deviceId} | | Return the client's daily cellular data usage history. Usage data is in kilobytes.
  GET  | connectivity | merakictl get sm connectivity {networkId} {deviceId} | {perPage} {startingAfter} {endingBefore} | Returns historical connectivity data (whether a device is regularly checking in to Dashboard).
  GET  | desktoplogs | merakictl get sm desktoplogs {networkId} {deviceId} | {perPage} {startingAfter} {endingBefore} | Return historical records of various Systems Manager network connection details for desktop devices.
  GET  | devicecommandlogs | merakictl get sm devicecommandlogs {networkId} {deviceId} | {perPage} {startingAfter} {endingBefore} | Return historical records of commands sent to Systems Manager devices. Note that this will include the name of the Dashboard user who initiated the command if it was generated by a Dashboard admin rather than the automatic behavior of the system; you may wish to filter this out of any reports.
  GET  | performancehistory | merakictl get sm performancehistory {networkId} {deviceId} | {perPage} {startingAfter} {endingBefore} | Return historical records of various Systems Manager client metrics for desktop devices.


## Insight  
 
 All Insight level API calls. 
 
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
  GET  | monitoredmediaservers | merakictl get insight monitoredmediaservers {organizationId} | | List the monitored media servers for this organization. Only valid for organizations with Meraki Insight.
  GET  | monitoredmediaserver | merakictl get insight monitoredmediaserver {organizationId} {monitoredMediaServerId} | | Return a monitored media server for this organization. Only valid for organizations with Meraki Insight.