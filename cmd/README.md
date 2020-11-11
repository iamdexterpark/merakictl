# merakictl cli

 
  
#### Commands
 
Command | Syntax | Description |
--- | --- | ---
| **claim** | merakictl claim [order/serial/licence] [TARGET] | Claim orders, licences, serial numbers into dashboard. |
| **create** | merakictl create [COMMAND] [SUBCOMMAND] [TARGET] [flags] | Creates (POST) new resources. |
| **get** | merakictl get [COMMAND] [SUBCOMMAND] [TARGET] [flags]| Operation for displaying (GET) api resources. |
| **update** | merakictl update [COMMAND] [SUBCOMMAND] [TARGET]  [flags]| Updates (PUT) targeted resources. |
| **remove** | merakictl remove [COMMAND] [SUBCOMMAND] [TARGET]  [flags]| Destructive (DELETE) API call for removing resources from the Dashboard. |
| **version** | merakictl version | Displays the version and associate release information of Merakictl. |


#### Subcommands
 
Subcommand | Abbreviated | Description |
--- | --- | --- |
| **organization** | org | Collection of Networks. |
| **network** | | Collection of Devices. | 
| **device** | | Meraki Products. | 
| **mx** | | UTM appliances with SD-WAN capability included. | 
| **ms**  | | Branch and campus access and aggregation switches with remote packet capture and cable testing. |
| **mr** | |  Industry-leading indoor and outdoor 802.11ax wireless access points with built-in location analytics. |
| **mv** | | Indoor and outdoor smart cameras, eliminating traditional external storage such as NVRs. |
| **mg** | |  Cellular gateways seamlessly transpose a wireless cellular signal to wired Ethernet, providing primary or failover connectivity. |
| **sm** | |  A complete enterprise solution for provisioning, monitoring, and securing of mobile devices. |
| **insight** | |  End-to-end visibility into end-users' experience on the network. |


#### Flags

Flag Type | Long | Short | Description |
--- | --- | --- | ---
| Global | **--input** | | Global flag for using a yaml file for passing config to the Dashboard API |
| Global | **--export** | -e | Global flag for extracting config from the Meraki API via get commands. |
| Global | **--organization** | -o | Global flag for Organization id. |
| Global | **--network** | -n | Global flag for Network id. |
| Global | **--device** | -d | Global flag for Serial Number. |
| Global | **--verbose** | -v | Global flag to display the http request & response for troubleshooting. |

 #### Common Filters

Filters are HTTP parameters that shape our API queries. The filters below, while not exhaustive, are most commonly used. 

Filter | Example | Description |
 --- | --- | --- 
t0 | |  The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
t1 | |  The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
timespan | |  The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
perPage | | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
startingAfter | | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
endingBefore | | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.


The filters "perPage", "startingAfter", and "endingBefore" are part of the Meraki Dashboard APIs pagination strategy. 
They should be considered overrides as Merakictl is designed to automatically handle pagination. 
To learn more see RFC5988 Web Linking.


##### Organization

All Organization level API calls. 

 HTTP | Operation | Syntax | Filters | Description |
----- | --------- | ------ | ----------- | ----------- |
 GET  | list | merakictl get organization list | | List the organizations that the user has privileges on.
 GET  | info | merakictl get organization info -o {organizationId} | | List a specific organization that the user has privileges on. 
 POST | {orgName} | merakictl create org {orgName} | | Create a new organization.
 PUT  | {orgName} | merakictl update org {orgName} | | Update an organization.
 DELETE | {orgName} | merakictl remove org {orgName} | | Delete an organization.
 GET  | actionbatches | merakictl get organization actionbatches -o {organizationId} | | Return The List Of Action Batches In The Organization.
 GET  | actionbatch | merakictl get organization actionbatch {actionBatchId} -o {organizationId} | | Return A Single Action Batch.
 GET  | admins | merakictl get organization admins -o {organizationId} | | List The Dashboard Administrators In This Organization.
 GET  | brandingpolicies | merakictl get organization brandingpolicies -o {organizationId} | | Return The Branding Policy IDs Of An Organization.
 GET  | brandingpolicy | merakictl get organization brandingpolicy {brandingPolicyId} -o {organizationId} | | Return The Branding Policies Of An Organization.
 GET  | configurationtemplates | merakictl get organization configurationtemplates -o {organizationId} | | List The Configuration Templates For This Organization.
 GET  | configurationtemplate | merakictl get organization configurationtemplate {configTemplateId} -o {organizationId} | | Return a Configuration Template For This Organization.
 GET  | devices | merakictl get organization devices -o {organizationId} | {{{perPage}}} {{{startingAfter}}} {{{endingBefore}}} |  List the devices in an organization.
 GET  | device | merakictl get organization device -d {deviceId} -o {organizationId} |  | Return A Single Device From The Inventory Of An Organization.
 GET  | inventory | merakictl get organization inventory |  {{{perPage}}} {{{startingAfter}}} {{{endingBefore}}} {{usedState}} {{search}} | Return The Device Inventory For An Organization.
 GET  | licences | merakictl get organization licences -o {organizationId} | {{{perPage}}} {{{startingAfter}}} {{{endingBefore}}} {{deviceSerial}} {{networkId}} {{state}} | List The Licenses For An Organization.
 GET  | licence | merakictl get organization licence {licenceId} -o {organizationId} | | List A Single License For An Organization.
 GET  | securitysettings | merakictl get organization securitysettings -o {organizationId} | | Returns The Login Security Settings For An Organization.
 GET  | networks | merakictl get organization networks -o {organizationId} | {{configTemplateId}} {{tags}} {{tagsFilterType}} {{{perPage}}} {{{startingAfter}}} {{{endingBefore}}} | List the networks that the user has privileges on in an organization.
 GET  | ldps | merakictl get organization ldps -o {organizationId} | | List the SAML IdPs in your organization.
 GET  | ldp | merakictl get organization ldp {ldpId} -o {organizationId} | | List a SAML IdP in your organization.
 GET  | saml | merakictl get organization saml -o {organizationId} | | Returns the SAML SSO enabled settings for an organization.
 GET  | samlroles | merakictl get organization samlroles -o {organizationId} | | List the SAML roles for this organization.
 GET  | samlrole | merakictl get organization samlrole {samalRoleId} -o {organizationId} | | List a single SAML role for this organization.
 GET  | snmp | merakictl get organization snmp -o {organizationId} | | Return the SNMP settings for an organization.
 GET  | apiresponsecount | merakictl get organization apiresponsecount -o {organizationId} | {{t0}} {{t1}} {{timespan}} | Return an aggregated overview of API requests data.
 GET  | apirequests | merakictl get organization apirequests -o {organizationId} | {{t0}} {{t1}} {{timespan}} {{{perPage}}} {{{startingAfter}}} {{{endingBefore}}} {{adminId}} {{path}} {{method}} {{responseCode}} {{sourceIp}} | List the API requests made by an organization
 GET  | configurationchanges | merakictl get organization configurationchanges -o {organizationId} | {{t0}} {{t1}} {{timespan}} {{{perPage}}} {{{startingAfter}}} {{{endingBefore}}} {{networkId}} {{adminId}} | View the Change Log for your organization.
 GET  | devicestatus | merakictl get organization devicestatus -o {organizationId}  | {{{perPage}}} {{{startingAfter}}} {{{endingBefore}}} | List the status of every Meraki device in the organization.
 GET  | losslatency | merakictl get organization losslatency -o {organizationId} | {{t0}} {{t1}} {{timespan}} {{uplink}} {{ip}} | Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago.
 GET  | licenseoverview | merakictl get organization licenseoverview -o {organizationId} | | Return an overview of the license state for an organization.
 GET  | openapi | merakictl get organization openapi -o {organizationId} | | Return the OpenAPI 2.0 Specification of the organization's API documentation in JSON.
 GET  | webhooklogs | merakictl get organization webhooklogs -o {organizationId} | {{t0}} {{t1}} {{timespan}} {{{perPage}}} {{{startingAfter}}} {{{endingBefore}}} {{url}} | Return the log of webhook POSTs sent.
 
 
##### Network

All Network level API calls. 

 HTTP | Operation | Syntax | Filters | Description |
----- | --------- | ------ | ----------- | ----------- |
 GET  |  | merakictl get network -n {networkId} | | Return a network.
 GET  | alertconfiguration | merakictl get network alertconfiguration -n {networkId} | | Return the alert configuration for this network.
 GET  | clientpolicy | merakictl get network clientpolicy {clientId} -n {networkId} | | Return the policy assigned to a client on the network. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | clientsplashauthorization | merakictl get network clientsplashauthorization {clientId} -n {networkId} | | Return the splash authorization for a client, for each SSID they've associated with through splash. Only enabled SSIDs with Click-through splash enabled will be included. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | devices | merakictl get network devices -n {networkId} | | List the devices in a network.
 GET  | firmwareupgrades | merakictl get network firmwareupgrades -n {networkId} | | Get current maintenance window for a network.
 GET  | floorplans | merakictl get network floorplans -n {networkId} | | List the floor plans that belong to your network.
 GET  | floorplan | merakictl get network floorplan {floorPlanId} -n {networkId} | | Find a floor plan by ID.
 GET  | grouppolicies | merakictl get network grouppolicies -n {networkId} | | List the group policies in a network.
 GET  | grouppolicy | merakictl get network grouppolicy {groupPolicyId} -n {networkId} | | Display a group policy.
 GET  | merakiauthusers | merakictl get network merakiauthusers -n {networkId} | | List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a wired network).
 GET  | merakiauthuser | merakictl get network merakiauthuser {merakiAuthUserId} -n {networkId} | | Return the Meraki Auth splash guest, RADIUS, or client VPN user.
 GET  | mqttbrokers | merakictl get network mqttbrokers -n {networkId} | | List the MQTT brokers for this network.
 GET  | mqttbroker | merakictl get network mqttbroker {mqttBrokerId} -n {networkId} | | Return an MQTT broker.
 GET  | netflow | merakictl get network netflow -n {networkId} | | Return the NetFlow traffic reporting settings for a network.
 GET  | channelutilization | merakictl get network channelutilization -n {networkId} | {{t0}} {{t1}} {{timespan}} {{resolution}} {{{perPage}}} {{{startingAfter}}} {{{endingBefore}}} | Get the channel utilization over each radio for all APs in a network.
 GET  | piikeys | merakictl get network piikeys -n {networkId} | {{username}} {{email}} {{mac}} {{serial}} {{imei}} {{bluetoothMac}} | List the keys required to access Personally Identifiable Information (PII) for a given identifier. Exactly one identifier will be accepted. If the organization contains org-wide Systems Manager users matching the key provided then there will be an entry with the key "0" containing the applicable keys.
 GET  | piirequests | merakictl get network piirequests -n {networkId} | | List the PII requests for this network or organization.
 GET  | piirequest | merakictl get network piirequest {requestId} -n {networkId} | | Return a PII request.
 GET  | smdevices | merakictl get network smdevices -n {networkId} | {{username}} {{email}} {{mac}} {{serial}} {{imei}} {{bluetoothMac}}  | Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier. These device IDs can be used with the Systems Manager API endpoints to retrieve device details. Exactly one identifier will be accepted.
 GET  | smowners | merakictl get network smowners -n {networkId} | {{username}} {{email}} {{mac}} {{serial}} {{imei}} {{bluetoothMac}}  | Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier. These owner IDs can be used with the Systems Manager API endpoints to retrieve owner details. Exactly one identifier will be accepted.
 GET  | settings | merakictl get network settings -n {networkId} | | Return the settings for a network.
 GET  | snmp | merakictl get network snmp -n {networkId} | | Return the SNMP settings for a network.
 GET  | syslog | merakictl get network syslog -n {networkId} | | List the syslog servers for a network.
 GET  | trafficanalysis | merakictl get network trafficanalysis -n {networkId} | | Return the traffic analysis settings for a network.
 GET  | trafficshaping | merakictl get network trafficshaping -n {networkId} | | Returns the application categories for traffic shaping rules.
 GET  | dscp | merakictl get network dscp -n {networkId} | | Returns the available DSCP tagging options for your traffic shaping rules.
 GET  | httpservers | merakictl get network httpservers -n {networkId} | | List the HTTP servers for a network.
 GET  | httpserver | merakictl get network httpserver {httpServerId} -n {networkId} | | Return an HTTP server for a network.
 GET  | webhooktest | merakictl get network webhooktest {weebhookTestId} -n {networkId} | | Return the status of a webhook test for a network.
 GET  | bluetoothclients | merakictl get network bluetoothclients -n {networkId} | {{t0}} {{t1}} {{timespan}} {{{perPage}}} {{{startingAfter}}} {{{endingBefore}}} {{includeConnectivityHistory}} | List the Bluetooth clients seen by APs in this network.
 GET  | bluetoothclient | merakictl get network bluetoothclient -n {networkId} | {{includeConnectivityHistory}} {{connectivityHistoryTimespan}} | Return a Bluetooth client. Bluetooth clients can be identified by their ID or their MAC.
 GET  | clienttraffichistory | merakictl get network clienttraffichistory {clientId} -n {networkId} | {{{perPage}}} {{{startingAfter}}} {{{endingBefore}}} | Return the client's network traffic data over time. Usage data is in kilobytes. This endpoint requires detailed traffic analysis to be enabled on the Network-wide > General page. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | clientusagehistory | merakictl get network clientusagehistory {clientId} -n {networkId} | | Return the client's daily usage history. Usage data is in kilobytes. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | networkclients | merakictl get network networkclients -n {networkId} | {{t0}} {{timespan}} {{{perPage}}} {{{startingAfter}}} {{{endingBefore}}} | List the clients that have used this network in the timespan.
 GET  | clientidentifier | merakictl get network clientidentifier {clientId} -n {networkId} | | Return the client associated with the given identifier. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | environmentalevents | merakictl get network environmentalevents -n {networkId} | {{includedEventTypes[]}} {{excludedEventTypes[]}} {{sensorSerial}} {{gatewaySerial}} {{{perPage}}} {{{startingAfter}}} {{{endingBefore}}} | List the environmental events for the network.
 GET  | events | merakictl get network events -n {networkId} | {{productType}} {{includedEventTypes[]}} {{excludedEventTypes[]}} {{deviceMac}} {{deviceSerial}} {{deviceName}} {{clientIp}} {{clientMac}} {{clientName}} {{smDeviceMac}} {{smDeviceName}} {{{perPage}}} {{{startingAfter}}} {{{endingBefore}}} | List the events for the network.
 GET  | splashloginattempts | merakictl get network splashloginattempts -n {networkId} | {{ssidNumber}} {{loginIdentifier}} {{timespan}} | List the splash login attempts for a network.
 GET  | traffic | merakictl get network traffic -n {networkId} | {{t0}} {{timespan}} {{deviceType}} | Return the traffic analysis data for this network. Traffic analysis with hostname visibility must be enabled on the network.


##### Device

All Device level API calls. 

 HTTP | Operation | Syntax | Filters | Description |
----- | --------- | ------ | ----------- | ----------- |
 GET  | mgmtinterface | merakictl get device mgmtinterface -d {deviceId} | | Return the management interface settings for a device.
 GET  | singledevice | merakictl get device singledevice -d {deviceId} |  | Return a single device.
 GET  | clients | merakictl get device clients -d {deviceId} | {{t0}} {{timespan}} | List the clients of a device, up to a maximum of a month ago. The usage of each client is returned in kilobytes. If the device is a switch, the switchport is returned; otherwise the switchport field is null.
 GET  | lldpcdp | merakictl get device lldpcdp -d {deviceId} | | List LLDP and CDP information for a device.
 GET  | uplinkloss | merakictl get device uplinkloss -d {deviceId} | {{t0}} {{t1}} {{timespan}} {{resolution}} {{uplink}} {{ip}} | Get the uplink loss percentage and latency in milliseconds for a wired network device.