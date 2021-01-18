# Merakictl CLI Reference 

*Main Documentation: [English](https://github.com/ddexterpark/merakictl/README.md)*
*CLI Reference Overview: [English](https://github.com/ddexterpark/merakictl/tree/master/meraki/README.md)*


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

