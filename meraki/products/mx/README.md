# Merakictl CLI Reference 

*Main Documentation: [English](https://github.com/ddexterpark/merakictl/README.md)*
*CLI Reference Overview: [English](https://github.com/ddexterpark/merakictl/meraki/README.md)*

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

