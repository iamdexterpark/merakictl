# Merakictl CLI Reference 

## Organization

All Organization level API calls. 

 HTTP | Operation | Syntax | Filters | Description |
----- | --------- | ------ | ----------- | ----------- |
 GET  | list | merakictl get organization list | | List the organizations that the user has privileges on.
 GET  | details | merakictl get organization info {organizationId} | | List a specific organization that the user has privileges on. 
 GET  | actionBatches | merakictl get organization actionBatches {organizationId} | | Return The List Of Action Batches In The Organization.
 GET  | actionBatch | merakictl get organization actionBatch {actionBatchId} {organizationId} | | Return A Single Action Batch.
 GET  | admins | merakictl get organization admins {organizationId} | | List The Dashboard Administrators In This Organization.
 GET  | brandingPolicies | merakictl get organization brandingpolicies {organizationId} | | Return The Branding Policy IDs Of An Organization.
 GET  | brandingPolicy | merakictl get organization brandingPolicy {brandingPolicyId} {organizationId} | | Return The Branding Policies Of An Organization.
 GET  | configurationTemplates | merakictl get organization configurationTemplates {organizationId} | | List The Configuration Templates For This Organization.
 GET  | configurationTemplate | merakictl get organization configurationTemplate {configTemplateId} {organizationId} | | Return a Configuration Template For This Organization.
 GET  | devices | merakictl get organization devices {organizationId} | --perPage --startingAfter --endingBefore --configurationUpdatedAfter |  List the devices in an organization.
 GET  | inventoryDevices | merakictl get organization inventoryDevices {organizationId}  |  --perPage --startingAfter --endingBefore --usedState --search | Return The Device Inventory For An Organization.
 GET  | inventoryDevice | merakictl get organization inventoryDevice {deviceId} {organizationId} |  | Return a single device from the inventory of an organizatio.
 GET  | licences | merakictl get organization licences {organizationId} | --perPage --startingAfter --endingBefore --deviceSerial --networkId --state | List The Licenses For An Organization.
 GET  | licence | merakictl get organization licence {licenceId} {organizationId} | | List A Single License For An Organization.
 GET  | loginSecurity | merakictl get organization loginSecurity {organizationId} | | Returns The Login Security Settings For An Organization.
 GET  | networks | merakictl get organization networks {organizationId} | --configTemplateId --tags --tagsFilterType --perPage --startingAfter --endingBefore | List the networks that the user has privileges on in an organization.
 GET  | idps | merakictl get organization idps {organizationId} | | List the SAML IdPs in your organization.
 GET  | idp | merakictl get organization idp {ldpId} {organizationId} | | List a SAML IdP in your organization.
 GET  | saml | merakictl get organization saml {organizationId} | | Returns the SAML SSO enabled settings for an organization.
 GET  | samlRoles | merakictl get organization samlRoles {organizationId} | | List the SAML roles for this organization.
 GET  | samlRole | merakictl get organization samlRole {samalRoleId} {organizationId} | | List a single SAML role for this organization.
 GET  | snmp | merakictl get organization snmp {organizationId} | | Return the SNMP settings for an organization.
 GET  | apiRequestsOverview | merakictl get organization apiRequestsOverview {organizationId} | --t0 --t1 --timespan | Return an aggregated overview of API requests data.
 GET  | apiRequests | merakictl get organization apiRequests {organizationId} | --t0 --t1 --timespan --perPage --startingAfter --endingBefore --adminId --path --method --responseCode --sourceIp | List the API requests made by an organization
 GET  | configurationChanges | merakictl get organization configurationChanges {organizationId} | --t0 --t1 --timespan --perPage --startingAfter --endingBefore --networkId --adminId | View the Change Log for your organization.
 GET  | deviceStatuses | merakictl get organization deviceStatuses {organizationId}  | --perPage --startingAfter --endingBefore | List the status of every Meraki device in the organization.
 GET  | uplinksLossAndLatency | merakictl get organization uplinksLossAndLatency {organizationId} | --t0 --t1 --timespan --uplink --ip | Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago.
 GET  | licenseOverview | merakictl get organization licenseOverview {organizationId} | | Return an overview of the license state for an organization.
 GET  | openApiSpec | merakictl get organization openapiSpec {organizationId} | | Return the OpenAPI 2.0 Specification of the organization's API documentation in JSON.
 GET  | alertTypes | merakictl get organization alertTypes {organizationId} | | Return a list of alert types to be used with managing webhook alerts.
 GET  | webhookLogs | merakictl get organization webhookLogs {organizationId} | --t0 --t1 --timespan --perPage --startingAfter --endingBefore --url | Return the log of webhook POSTs sent.
 
 
## Network

All Network level API calls. 

 HTTP | Operation | Syntax | Filters | Description |
----- | --------- | ------ | ----------- | ----------- |
 GET  | details | merakictl get network details {networkId} | | Return a network.
 GET  | alertSettings | merakictl get network alertSettings {networkId} | | Return the alert configuration for this network.
 GET  | clientPolicy | merakictl get network clientPolicy {clientId} {networkId} | | Return the policy assigned to a client on the network. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | splashAuthorizationStatus | merakictl get network splashAuthorizationStatus {clientId} {networkId} | | Return the splash authorization for a client, for each SSID they've associated with through splash. Only enabled SSIDs with Click-through splash enabled will be included. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | devices | merakictl get network devices {networkId} | | List the devices in a network.
 GET  | firmwareUpgrades | merakictl get network firmwareUpgrades {networkId} | | Get current maintenance window for a network.
 GET  | floorPlans | merakictl get network floorPlans {networkId} | | List the floor plans that belong to your network.
 GET  | floorPlan | merakictl get network floorPlan {floorPlanId} {networkId} | | Find a floor plan by ID.
 GET  | groupPolicies | merakictl get network groupPolicies {networkId} | | List the group policies in a network.
 GET  | groupPolicy | merakictl get network groupPolicy {groupPolicyId} {networkId} | | Display a group policy.
 GET  | merakiAuthUsers | merakictl get network merakiAuthUsers {networkId} | | List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a wired network).
 GET  | merakiAuthUser | merakictl get network merakiAuthUser {merakiAuthUserId} {networkId} | | Return the Meraki Auth splash guest, RADIUS, or client VPN user.
 GET  | mqttBrokers | merakictl get network mqttBrokers {networkId} | | List the MQTT brokers for this network.
 GET  | mqttBroker | merakictl get network mqttBroker {mqttBrokerId} {networkId} | | Return an MQTT broker.
 GET  | netflow | merakictl get network netflow {networkId} | | Return the NetFlow traffic reporting settings for a network.
 GET  | channelUtilization | merakictl get network channelUtilization {networkId} | --t0 --t1 --timespan --resolution --perPage --startingAfter --endingBefore | Get the channel utilization over each radio for all APs in a network.
 GET  | piiKeys | merakictl get network piiKeys {networkId} | --username --email --mac --serial --imei --bluetoothMac | List the keys required to access Personally Identifiable Information (PII) for a given identifier. Exactly one identifier will be accepted. If the organization contains org-wide Systems Manager users matching the key provided then there will be an entry with the key "0" containing the applicable keys.
 GET  | piiRequests | merakictl get network piiRequests {networkId} | | List the PII requests for this network or organization.
 GET  | piiRequest | merakictl get network piiRequest {requestId} {networkId} | | Return a PII request.
 GET  | smDevices | merakictl get network smDevices {networkId} | --username --email --mac --serial --imei --bluetoothMac  | Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier. These device IDs can be used with the Systems Manager API endpoints to retrieve device details. Exactly one identifier will be accepted.
 GET  | smOwners | merakictl get network smOwners {networkId} | --username --email --mac --serial --imei --bluetoothMac  | Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier. These owner IDs can be used with the Systems Manager API endpoints to retrieve owner details. Exactly one identifier will be accepted.
 GET  | settings | merakictl get network settings {networkId} | | Return the settings for a network.
 GET  | snmp | merakictl get network snmp {networkId} | | Return the SNMP settings for a network.
 GET  | syslog | merakictl get network syslog {networkId} | | List the syslog servers for a network.
 GET  | trafficAnalysis | merakictl get network trafficAnalysis {networkId} | | Return the traffic analysis settings for a network.
 GET  | trafficShaping | merakictl get network trafficShaping {networkId} | | Returns the application categories for traffic shaping rules.
 GET  | dscpTaggingOptions | merakictl get network dscpTaggingOptions {networkId} | | Returns the available DSCP tagging options for your traffic shaping rules.
 GET  | httpServers | merakictl get network httpServers {networkId} | | List the HTTP servers for a network.
 GET  | httpServer | merakictl get network httpServer {httpServerId} {networkId} | | Return an HTTP server for a network.
 GET  | webhookTest | merakictl get network webhookTest {weebhookTestId} {networkId} | | Return the status of a webhook test for a network.
 GET  | bluetoothclients | merakictl get network bluetoothClients {networkId} | --t0 --t1 --timespan --perPage --startingAfter --endingBefore --includeConnectivityHistory | List the Bluetooth clients seen by APs in this network.
 GET  | bluetoothClient | merakictl get network bluetoothClient {bluetoothClientId} {networkId} | --includeConnectivityHistory --connectivityHistoryTimespan | Return a Bluetooth client. Bluetooth clients can be identified by their ID or their MAC.
 GET  | clientTrafficHistory | merakictl get network clientTrafficHistory {clientId} {networkId} | --perPage --startingAfter --endingBefore | Return the client's network traffic data over time. Usage data is in kilobytes. This endpoint requires detailed traffic analysis to be enabled on the Network-wide > General page. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | clientUsageHistory | merakictl get network clientUsageHistory {clientId} {networkId} | | Return the client's daily usage history. Usage data is in kilobytes. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | networkClients | merakictl get network networkClients {networkId} | --t0 --timespan --perPage --startingAfter --endingBefore | List the clients that have used this network in the timespan.
 GET  | clientIdentifier | merakictl get network clientIdentifier {clientId} {networkId} | | Return the client associated with the given identifier. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | environmentalEvents | merakictl get network environmentalEvents {networkId} | --includedEventTypes[] --excludedEventTypes[] --sensorSerial --gatewaySerial --perPage --startingAfter --endingBefore | List the environmental events for the network.
 GET  | events | merakictl get network events {networkId} | --productType --includedEventTypes[] --excludedEventTypes[] --deviceMac --deviceSerial --deviceName --clientIp --clientMac --clientName --smDeviceMac --smDeviceName --perPage --startingAfter --endingBefore | List the events for the network.
 GET  | splashLoginAttempts | merakictl get network splashLoginAttempts {networkId} | --splashLoginAttempts --ssidNumber --loginIdentifier --timespan | List the splash login attempts for a network.
 GET  | traffic | merakictl get network traffic {networkId} | --t0 --timespan --deviceType | Return the traffic analysis data for this network. Traffic analysis with hostname visibility must be enabled on the network.


## Device

All Device level API calls. 

 HTTP | Operation | Syntax | Filters | Description |
----- | --------- | ------ | ----------- | ----------- |
 GET  | managementInterface | merakictl get device managementInterface {serial} | | Return the management interface settings for a device.
 GET  | device | merakictl get device details {serial} |  | Return a single device.
 GET  | clients | merakictl get device clients {serial} | --t0 --timespan | List the clients of a device, up to a maximum of a month ago. The usage of each client is returned in kilobytes. If the device is a switch, the switchport is returned; otherwise the switchport field is null.
 GET  | lldpcdp | merakictl get device lldpcdp {serial} | | List LLDP and CDP information for a device.
 GET  | uplinkloss | merakictl get device uplinkloss {serial} | --t0 --t1 --timespan --resolution --uplink --ip | Get the uplink loss percentage and latency in milliseconds for a wired network device.
 
 
 
## MX Appliance 
 
 All MX level API calls. 
 
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
 GET  | connectivityMonitoringDestinations | merakictl get mx connectivityMonitoringDestinations {networkId} | | Return the connectivity testing destinations for an MX network.
 GET  | contentFilteringCategories | merakictl get mx contentFilteringCategories {networkId} | | List all available content filtering categories for an MX network.
 GET  | contentFiltering | merakictl get mx contentFiltering {networkId} | | Return the content filtering settings for an MX network.
 GET  | cellularFirewallRules | merakictl get mx cellularFirewallRules {networkId} | | Return the cellular firewall rules for an MX network.
 GET  | firewalledServices | merakictl get mx firewalledServices {networkId} | | List the appliance services and their accessibility rules.
 GET  | firewalledService | merakictl get mx firewalledService {service} {networkId} | | Return the accessibility settings of the given service ('ICMP', 'web', or 'SNMP').
 GET  | inboundFirewallRules | merakictl get mx inboundFirewallRules {networkId} | | Return the inbound firewall rules for an MX network.
 GET  | l3FirewallRules | merakictl get mx l3FirewallRules {networkId} | | Return the L3 firewall rules for an MX network.
 GET  | l7ApplicationCategories | merakictl get mx l7ApplicationCategories {networkId} | | Return the L7 firewall application categories and their associated applications for an MX network.
 GET  | l7FirewallRule | merakictl get mx l7FirewallRule {networkId} | | List the MX L7 firewall rules for an MX network.
 GET  | oneToManyNatRules | merakictl get mx oneToManyNatRules {networkId} | | Return the 1:Many NAT mapping rules for an MX network.
 GET  | oneToOneNatRules | merakictl get mx oneToOneNatRules {networkId} | | Return the 1:1 NAT mapping rules for an MX network.
 GET  | portForwardingRules | merakictl get mx portForwardingRules {networkId} | | Return the port forwarding rules for an MX network.
 GET  | ports | merakictl get mx ports {networkId} | | List per-port VLAN settings for all ports of a MX.
 GET  | port | merakictl get mx port {portId} {networkId} | | Return per-port VLAN settings for a single MX port.
 GET  | networkIntrusion | merakictl get mx networkIntrusion {networkId} | | Returns all supported intrusion settings for an MX network.
 GET  | organizationIntrusion | merakictl get mx organizationIntrusion {organizationId} | | Returns all supported intrusion settings for an organization.
 GET  | malware | merakictl get mx malware {networkId} | | Returns all supported malware settings for an MX network.
 GET  | settings | merakictl get mx settings {networkId} | | Return the appliance settings for a network.
 GET  | lan | merakictl get mx lan {networkId} | | Return single LAN configuration.
 GET  | staticRoutes | merakictl get mx staticRoutes {networkId} | | List the static routes for an MX or teleworker network.
 GET  | staticRoute | merakictl get mx staticRoute {staticRouteId} {networkId} | | Return a static route for an MX or teleworker network.
 GET  | customPerformanceClasses | merakictl get mx customPerformanceClasses {networkId} | | List all custom performance classes for an MX network.
 GET  | customPerformanceClass | merakictl get mx customPerformanceClass {customPerformanceClassId} {networkId} | | Return a custom performance class for an MX network.
 GET  | trafficShapingRules | merakictl get mx trafficShapingRules {networkId} | | Display the traffic shaping settings rules for an MX network.
 GET  | uplinkBandwidth | merakictl get mx uplinkBandwidth {networkId} | | Returns the uplink bandwidth settings for your MX network..
 GET  | uplinkSelection | merakictl get mx uplinkSelection {networkId} | | Show uplink selection settings for an MX network.
 GET  | trafficShaping | merakictl get mx trafficShaping {networkId} | | Display the traffic shaping settings for an MX network.
 GET  | vlanSettings | merakictl get mx vlanSettings {networkId} | | Returns the enabled status of VLANs for the network.
 GET  | vlans | merakictl get mx vlans {networkId} | | List the VLANs for an MX network.
 GET  | vlan | merakictl get mx vlan {vlanId} {networkId} | | Return a VLAN.
 GET  | bgp | merakictl get mx bgp {networkId} | | Return a Hub BGP Configuration.
 GET  | siteToSiteVPN | merakictl get mx siteToSiteVPN {networkId} | | Return the site-to-site VPN settings of a network. Only valid for MX networks.
 GET  | thirdPartyVPNPeers | merakictl get mx thirdPartyVPNPeers {organizationId} | | Return the third party VPN peers for an organization.
 GET  | vpnFirewallRules | merakictl get mx vpnFirewallRules {organizationId} | | Return the firewall rules for an organization's site-to-site VPN.
 GET  | warmspare | merakictl get mx warmspare {networkId} | | Return MX warm spare settings.
 GET  | securityEvents | merakictl get mx securityEvents {clientId} {networkId} | --t0 --t1 --timespan --perPage --startingAfter --endingBefore --sortOrder | List the security events for a client. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
 GET  | dhcpSubnets | merakictl get mx dhcpSubnets {serial} | | Return the DHCP subnet information for an appliance.
 GET  | performance | merakictl get mx performance {serial} | | Return the performance score for a single MX. Only primary MX devices supported. If no data is available, a 204 error code is returned.
 GET  | networkSecurityEvents | merakictl get mx networkSecurityEvents {networkId} | --t0 --t1 --timespan --perPage --startingAfter --endingBefore --sortOrder | List the security events for a network.
 GET  | organizationSecurityEvents | merakictl get mx organizationSecurityEvents {organizationId} | --t0 --t1 --timespan --perPage --startingAfter --endingBefore --sortOrder | List the security events for an organization.
 GET  | uplinkStatuses | merakictl get mx uplinkStatuses {organizationId} | --perPage --startingAfter --endingBefore | List the uplink status of every Meraki MX and Z series appliances in the organization.
 GET  | vpnStats | merakictl get mx vpnStats {organizationId} | --perPage --startingAfter --endingBefore {networkIds} --t0 --t1 --timespan | Show VPN history stat for networks in an organization.
 GET  | vpnStatuses | merakictl get mx vpnstatuses {organizationId} | --perPage --startingAfter --endingBefore {networkIds} | Show VPN status for networks in an organization.


## Camera MV  
 
 All MV level API calls. 
 
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
 GET  | qualityAndRetention | merakictl get mv qualityAndRetention {serial} | | Returns quality and retention settings for the given camera.
 GET  | qualityRetentionProfiles | merakictl get mv qualityRetentionProfiles {networkId} | | List the quality retention profiles for this network.
 GET  | qualityRetentionProfile | merakictl get mv qualityRetentionProfile {qualityRetentionProfileId} {networkId} | | Retrieve a single quality retention profile.
 GET  | schedules | merakictl get mv schedules  {networkId} | | Returns a list of all camera recording schedules.
 GET  | objectDetectionModels | merakictl get mv objectDetectionModelss {serial} | | Returns the MV Sense object detection model list for the given camera.
 GET  | sense | merakictl get mv sense {serial} | | Returns sense settings for a given camera.
 GET  | videoSettings | merakictl get mv videoSettings {serial} | | Returns video settings for the given camera.
 GET  | videoLink | merakictl get mv videoLink {serial} | --timestamp | Returns video link to the specified camera. If a timestamp is supplied, it links to that timestamp.
 GET  | liveAnalytics | merakictl get mv liveAnalytics {serial} | | Returns live state from camera of analytics zones.
 GET  | analyticsOverview | merakictl get mv analyticsOverview {serial} | --t0 --t1 --timespan --objectType | Returns an overview of aggregate analytics data for a timespan.
 GET  | recentAnalytics | merakictl get mv recentAnalytics {serial} | --objectType | Returns most recent record for analytics zones.
 GET  | analyticsZonesHistory | merakictl get mv analyticsZonesHistory {zoneId} {serial} | --t0 --t1 --timespan --resolution --objectType | Return historical records for analytic zones.
 GET  | analyticsZones | merakictl get mv analyticsZones {serial} | | Returns all configured analytic zones for this camera.
 
 
## Cellular Gateway mg  
  
  All MG level API calls.
  
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
 GET  | connectivityMonitor | merakictl get mg connectivityMonitor {networkId} | | Return the connectivity testing destinations for an MG network.
 GET  | dhcp | merakictl get mg dhcp {networkId} | | List common DHCP settings of MGs.
 GET  | lan | merakictl get mg lan {serial} | | Show the LAN Settings of a MG.
 GET  | portForwardingRules | merakictl get mg portForwardingRules {serial} | | Returns the port forwarding rules for a single MG.
 GET  | subnetPool | merakictl get mg subnetPool {networkId} | | Return the subnet pool and mask configured for MGs in the network.
 GET  | uplink | merakictl get mg uplink {networkId} | | Returns the uplink settings for your MG network.
 


## Switch MS  
 
 All MS level API calls. 
 
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
  GET  | accessControlLists | merakictl get ms accesscontrollists {networkId} | | Return the access control lists for a MS network.
  GET  | accessPolicies | merakictl get ms accessPolicies {networkId} | | List the access policies for a switch network. Only returns access policies with 'my RADIUS server' as authentication method.
  GET  | accessPolicy | merakictl get ms accessPolicy {accessPolicyNumber}  {networkId}| | Return a specific access policy for a switch network.
  GET  | switchPortsProfiles | merakictl get ms switchPortsProfiles {configTemplateId} {profileId} {organizationId} | | Return all the ports of a switch profile.
  GET  | switchPortProfile | merakictl get ms switchPortProfile {configTemplateId} {profileId} {portId} {organizationId} | | Return a switch profile port.
  GET  | switchProfiles | merakictl get ms switchProfiles {configTemplateId} {organizationId} | | List the switch profiles for your switch template configuration.
  GET  | dhcpServerPolicy | merakictl get ms dhcpServerPolicy {networkId} | | Return the DHCP server policy.
  GET  | dscp | merakictl get ms dscp {networkId} | | Return the DSCP to CoS mappings.
  GET  | linkAggregations | merakictl get ms linkAggregations {networkId} | | List link aggregation groups.
  GET  | mtu | merakictl get ms mtu {networkId} | | Return the MTU configuration.
  GET  | portSchedules | merakictl get ms portSchedules {networkId} | | List switch port schedules.
  GET  | switchPorts | merakictl get ms switchPorts {serial} | | List the switch ports for a switch.
  GET  | switchPort| merakictl get ms switchPort {portId} {serial}  | | Return a switch port.
  GET  | qoSRuleOrder | merakictl get ms qoSRuleOrder {networkId} | | Return the quality of service rule IDs by order in which they will be processed by the switch.
  GET  | qosRules | merakictl get ms qosRules {networkId} | | List quality of service rules.
  GET  | qosRule | merakictl get ms qosRule {qosRuleId} {networkId} | | Return a quality of service rule.
  GET  | dhcp | merakictl get ms dhcp interfaceId {serial} | | Return a layer 3 interface DHCP configuration for a switch.
  GET  | l3Interfaces | merakictl get ms l3Interfaces {serial} | | List layer 3 interfaces for a switch.
  GET  | l3Interface | merakictl get ms l3Interface {interfaceId} {serial} | | Return a layer 3 interface for a switch.
  GET  | rendezvousPoints | merakictl get ms rendezvousPoints {networkId} | | List multicast rendezvous points.
  GET  | rendezvousPoint | merakictl get ms rendezvousPoint {rendezvousPointId} {networkId} | | Return a multicast rendezvous point.
  GET  | multicast | merakictl get ms multicast {networkId} | | Return multicast settings for a network.
  GET  | ospf | merakictl get ms ospf {networkId} | | Return layer 3 OSPF routing configuration.
  GET  | staticRoutes | merakictl get ms staticRoutes {serial} | | List layer 3 static routes for a switch.
  GET  | staticRoute | merakictl get ms staticRoute {staticRouteId} {serial} | | Return a layer 3 static route for a switch.
  GET  | settings | merakictl get ms settings {networkId} | | Returns the switch network settings.
  GET  | stackDHCP | merakictl get ms stackDHCP {switchStackId} {interfaceId} {networkId} | | Return a layer 3 interface DHCP configuration for a switch stack.
  GET  | stackL3Interfaces | merakictl get ms stackL3Interfaces {switchStackId} {networkId}  | | List layer 3 interfaces for a switch stack.
  GET  | stackL3Interface | merakictl get ms stackL3Interface {switchStackId} {interfaceId} {networkId} | | Return a layer 3 interface from a switch stack.
  GET  | stackStaticRoutes | merakictl get ms stackStaticRoutes {switchStackId} {networkId} | | List layer 3 static routes for a switch stack.
  GET  | stackStaticRoute | merakictl get ms stackStaticRoute {switchStackId} {staticRouteId} {networkId}  | | Return a layer 3 static route for a switch stack.
  GET  | switchStacks | merakictl get ms switchStacks {networkId} | | List the switch stacks in a network.
  GET  | switchStack | merakictl get ms switchStack {switchStackId} {networkId} | | Show a switch stack.
  GET  | stormControl | merakictl get ms stormControl {networkId} | | Return the storm control configuration for a switch network.
  GET  | stp | merakictl get ms stp {networkId} | | storm control.
  GET  | warmspare | merakictl get ms warmspare {serial} | | Return warm spare configuration for a switch.
  GET  | packets | merakictl get ms packets {serial} | --t0 --timespan | Return the packet counters for all the ports of a switch.
  GET  | portsStatuses | merakictl get ms portsStatuses {serial} | --t0 --timespan  | Return the status for all the ports of a switch.




## Wireless MR
 
 All MR level API calls. 
 
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
  GET  | alternateManagementInterface | merakictl get mr alternateManagementInterface {networkId} | | Return alternate management interface and devices with IP assigned.
  GET  | bluetoothNetworkSettings | merakictl get mr bluetoothNetworkSettings {networkId} | | Return the Bluetooth settings for a network. Bluetooth settings must be enabled on the network.
  GET  | bluetoothDeviceSettings | merakictl get mr bluetoothDeviceSettings {serial} | | Return the bluetooth settings for a wireless device.
  GET  | radioSettings | merakictl get mr radioSettings {serial} | | Return the radio settings of a device.
  GET  | rfProfiles | merakictl get mr rfProfiles {networkId} | --includeTemplateProfiles | List the non-basic RF profiles for this network.
  GET  | rfProfile | merakictl get mr rfProfile {rfProfileId} {networkId} | | Return a RF profile.
  GET  | wirelessSettings | merakictl get mr wirelessSettings {networkId} | | Return the wireless settings for a network.
  GET  | l3FirewallRules | merakictl get mr l3FirewallRules {number} {networkId} | | Return the L3 firewall rules for an SSID on an MR network.
  GET  | l7FirewallRules | merakictl get mr l7FirewallRules {number} {networkId} | | Return the L7 firewall rules for an SSID on an MR network.
  GET  | identityPSKs | merakictl get mr identityPSKs {number} {networkId} | | List all Identity PSKs in a wireless network.
  GET  | identityPSK | merakictl get mr identityPSK {number} {identityPskId} {networkId} | | Return an Identity PSK.
  GET  | splashSettings | merakictl get mr splashSettings {number} {networkId}| | Display the splash page settings for the given SSID.
  GET  | trafficShapingRules | merakictl get mr trafficShapingRules {number} {networkId} | | Display the traffic shaping settings for a SSID on an MR network.
  GET  | ssids | merakictl get mr ssids {networkId} | | List the MR SSIDs in a network.
  GET  | ssid | merakictl get mr ssid {number} {networkId}| | Return a single MR SSID.
  GET  | airMarshal | merakictl get mr airmarshal {networkId} | --t0 --timespan | List Air Marshal scan results from a network.
  GET  | channelUtilizationHistory | merakictl get mr channelUtilizationHistory {networkId} | --t0 --t1 --timespan --resolution --autoResolution --clientId --deviceSerial --apTag --band | Return AP channel utilization over time for a device or network client.
  GET  | clientCountHistory | merakictl get mr clientCountHistory {networkId} | --t0 --t1 --timespan --resolution --autoResolution --clientId --deviceSerial --apTag --band --ssid | Return wireless client counts over time for a network, device, or network client.
  GET  | connectionStat | merakictl get mr connectionStat {networkId} | --t0 --t1 --timespan --band --ssid --vlan --apTag | Aggregated connectivity info for a given client on this network. Clients are identified by their MAC.
  GET  | connectionStats  | merakictl get mr connectionStats  {networkId} | --t0 --t1 --timespan --band --ssid --vlan --apTag  |  Aggregated connectivity info for this network, grouped by clients.
  GET  | connectivityEvents | merakictl get mr connectivityEvents {clientId} {networkId} | --perPage --startingAfter --endingBefore --t0 --t1 --timespan --types --includedSeverities --band --ssidNumber --deviceSerial | List the wireless connectivity events for a client within a network in the timespan.
  GET  | latencyHistory | merakictl get mr latencyHistory {clientId} {networkId} | --t0 --t1 --timespan --resolution | Return the latency history for a client. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP. The latency data is from a sample of 2% of packets and is grouped into 4 traffic categories: background, best effort, video, voice. Within these categories the sampled packet counters are bucketed by latency in milliseconds..
  GET  | latencyStat | merakictl get mr latencyStat {clientId} {networkId}| --t0 --t1 --timespan --band --ssid --vlan --apTag --fields | Aggregated latency info for a given client on this network. Clients are identified by their MAC.
  GET  | latencyStats | merakictl get mr latencyStats {networkId} | --t0 --t1 --timespan --band --ssid --vlan --apTag --fields | Aggregated latency info for this network, grouped by clients.
  GET  | deviceConnectionStats | merakictl get mr deviceConnectionStats {serial} | --t0 --t1 --timespan --band --ssid --vlan --apTag | Aggregated connectivity info for a given AP on this network.
  GET  | networkConnectionStats | merakictl get mr networkConnectionStats {networkId} | --t0 --t1 --timespan --band --ssid --vlan --apTag | Aggregated connectivity info for this network.
  GET  | dataRateHistory | merakictl get mr dataRateHistory {networkId} | --t0 --t1 --timespan --resolution --autoResolution --clientId --deviceSerial --apTag --band --ssid | Return PHY data rates over time for a network, device, or network client.
  GET  | connectionStats | merakictl get mr connectionStats {networkId} |--t0 --t1 --timespan --band --ssid --vlan --apTag | Aggregated connectivity info for this network, grouped by node.
  GET  | deviceLatencyStats | merakictl get mr deviceLatencyStats {networkId} | --t0 --t1 --timespan --band --ssid --vlan --apTag --fields | Aggregated latency info for this network, grouped by node.
  GET  | failedConnections | merakictl get mr failedConnections {networkId} | --t0 --t1 --timespan --band --ssid --vlan --apTag {serial} {clientId} | List of all failed client connection events on this network in a given time range.
  GET  | latencyHistoryAverage | merakictl get mr latencyHistoryAverage {networkId} | --t0 --t1 --timespan --resolution --autoResolution --clientId --deviceSerial --apTag --band --ssid --accessCategory | Return average wireless latency over time for a network, device, or network client.
  GET  | aggregatedLatencies | merakictl get mr aggregatedLatencies {serial} | --t0 --t1 --timespan --band --ssid --vlan --apTag --fields | Aggregated latency info for a given AP on this network.
  GET  | networkLatencyStats | merakictl get mr networkLatencyStats {networkId} | --t0 --t1 --timespan --band --ssid --vlan --apTag --fields | Aggregated latency info for this network.
  GET  | meshStatuses | merakictl get mr meshstatuses {networkId} | --perPage --startingAfter --endingBefore | List wireless mesh statuses for repeaters.
  GET  | signalQualityHistory | merakictl get mr signalQualityHistory {networkId} | --t0 --t1 --timespan --resolution --autoResolution --clientId --deviceSerial --apTag --band --ssid | Return signal quality (SNR/RSSI) over time for a device or network client.
  GET  | status | merakictl get mr status {serial} | | Return the SSID statuses of an access point.
  GET  | usageHistory | merakictl get mr usageHistory {networkId} |  --t0 --t1 --timespan --resolution --autoResolution --clientId --deviceSerial --apTag --band --ssid | Return AP usage over time for a device or network client.



## SM  
 
 All SM level API calls. 
 
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
  GET  | apnsCert | merakictl get sm apnsCert {organizationId} | | Get the organization's APNS certificate.
  GET  | bypassActivationLockAttempts | merakictl get sm bypassActivationLockAttempts {attemptId} {networkId} | | Bypass activation lock attempt status.
  GET  | certs | merakictl get sm certs {deviceId} {networkId} | | List the certs on a device.
  GET  | deviceProfiles | merakictl get sm deviceProfiles {deviceId} {networkId} | | Get the profiles associated with a device.
  GET  | networkAdapters | merakictl get sm networkAdapters {deviceId} {networkId} | | List the network adapters of a device.
  GET  | restrictions | merakictl get sm restrictions {deviceId} {networkId} | | List the restrictions on a device.
  GET  | securityCenters | merakictl get sm securityCenters {deviceId} {networkId} | | List the security centers on a device.
  GET  | deviceSoftware | merakictl get sm deviceSoftware {deviceId} {networkId} | | Get a list of softwares associated with a device.
  GET  | wlanLists | merakictl get sm wlanLists {deviceId} {networkId} | | List the saved SSID names on a device.
  GET  | devices | merakictl get sm devices {networkId} | --fields --wifiMacs --serials --ids --scope --perPage --startingAfter --endingBefore | List the devices enrolled in an SM network with various specified fields and filters.
  GET  | profiles | merakictl get sm profiles {networkId} | | List all profiles in a network.
  GET  | targetGroups | merakictl get sm targetGroups {networkId} | --withDetails | List the target groups in this network.
  GET  | targetGroup | merakictl get sm targetGroup {targetGroupId} {networkId} | --withDetails | Return a target group.
  GET  | userSoftware | merakictl get sm userSoftware {clientId} {networkId} | | Get a list of softwares associated with a user.
  GET  | users | merakictl get sm users {networkId} | --ids --usernames --emails --scope | List the owners in an SM network with various specified fields and filters.
  GET  | vppAccount | merakictl get sm vppAccount {vppAccountId} {organizationId} | | Get a hash containing the unparsed token of the VPP account with the given ID.
  GET  | vppAccounts | merakictl get sm vppAccounts {organizationId} | | List the VPP accounts in the organization.
  GET  | cellularUsageHistory | merakictl get sm cellularUsageHistory {deviceId} {networkId} | | Return the client's daily cellular data usage history. Usage data is in kilobytes.
  GET  | connectivity | merakictl get sm connectivity {deviceId} {networkId} | --perPage --startingAfter --endingBefore | Returns historical connectivity data (whether a device is regularly checking in to Dashboard).
  GET  | desktopLogs | merakictl get sm desktopLogs {deviceId} {networkId} | --perPage --startingAfter --endingBefore | Return historical records of various Systems Manager network connection details for desktop devices.
  GET  | deviceCommandLogs | merakictl get sm deviceCommandLogs {deviceId} {networkId} | --perPage --startingAfter --endingBefore | Return historical records of commands sent to Systems Manager devices. Note that this will include the name of the Dashboard user who initiated the command if it was generated by a Dashboard admin rather than the automatic behavior of the system; you may wish to filter this out of any reports.
  GET  | performanceHistory | merakictl get sm performanceHistory {deviceId} {networkId} | --perPage --startingAfter --endingBefore | Return historical records of various Systems Manager client metrics for desktop devices.


## Insight  
 
 All Insight level API calls. 
 
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
  GET  | mmonitoredMediaServers | merakictl get insight monitoredMediaServers {organizationId} | | List the monitored media servers for this organization. Only valid for organizations with Meraki Insight.
  GET  | monitoredMediaServer | merakictl get insight monitoredMediaServer {monitoredMediaServerId} {organizationId} | | Return a monitored media server for this organization. Only valid for organizations with Meraki Insight.
  
  

## Filters

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
perPage | | 
startingAfter | | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
endingBefore | | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.


The filters `perPage`, `startingAfter`, and `endingBefore` are part of the Meraki Dashboard APIs pagination strategy. 
They should be considered overrides as Merakictl is designed to automatically handle pagination. 
To learn more see [RFC5988 Web Linking](https://tools.ietf.org/html/rfc5988).

